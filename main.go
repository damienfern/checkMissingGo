package main

import (
	"./config/args"
	"./config/connection"
	"./file"
	"./tvdb"
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	tvdbApi "github.com/pioz/tvdb"
	log "github.com/sirupsen/logrus"
	"regexp"
	"sort"
	"strconv"
)

func init() {
	log.SetFormatter(&nested.Formatter{
		ShowFullLevel:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func main() {
	c := connection.ConnectToTVDB()
	path, err := args.GetDirPathToSearch()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(path)
	regexpFilePath := regexp.MustCompile("^/?(.+/)*(.+)$")
	filePathArray := regexpFilePath.FindStringSubmatch(path)
	seriesName := filePathArray[len(filePathArray)-1]

	series := tvdb.FindSeriesOrFail(seriesName, &c)
	listSeasonsDir, errListSeasonDirs := file.ListAllDirOnlyInDir(path)

	var allSeasons []*file.SeasonDir
	for _, element := range listSeasonsDir {
		one := file.NewSeasonDirSeason(element, path)
		allSeasons = append(allSeasons, one)
	}
	if errListSeasonDirs != nil {
		log.Fatalln(err)
	}

	err2 := c.GetSeriesEpisodes(&series, nil)
	if err2 != nil {
		log.Fatal(err)
	}
	var missingEpisodes []*tvdbApi.Episode
	c.GetSeriesSummary(&series)
	seasonsFromTvDB := series.Summary.AiredSeasons
	var airedSeason []int

	for _, oneSeasonFromTvDB := range seasonsFromTvDB {
		converted, _ := strconv.Atoi(oneSeasonFromTvDB)
		airedSeason = append(airedSeason, converted)
	}
	sort.Ints(airedSeason)
	if airedSeason[0] == 0 {
		airedSeason = airedSeason[1:] // We don't care about special episodes
	}
	for _, oneSeason := range allSeasons {
		missingEpisodes = append(missingEpisodes, oneSeason.CheckMissingEpisodes(&series)...)
		airedSeason = Filter(airedSeason, func(i int) bool {
			return i != oneSeason.SeasonID
		})
	}

	if len(missingEpisodes) > 0 {
		fmt.Println("Missing episodes are :")
		for _, value := range missingEpisodes {
			fmt.Println("* S0" + strconv.Itoa(value.AiredSeason) + "E" + strconv.Itoa(value.AiredEpisodeNumber))
		}
		if len(airedSeason) > 0 {
			fmt.Println("Missing seasons are :")
			for _, value := range airedSeason {
				fmt.Println("* Season " + strconv.Itoa(value))
			}
		}
	} else {
		fmt.Println("No missing episodes")
	}

}

func Filter(vs []int, f func(int) bool) []int {
	vsf := make([]int, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
