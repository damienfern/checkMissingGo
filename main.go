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
	/*
		TODO : 2 ways to do it :
					* get seasons one by one and then do stuff
					* get all files recursively and do stuffs with regexs
	*/
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

	for _, oneSeason := range allSeasons {
		missingEpisodes = append(missingEpisodes, oneSeason.CheckMissingEpisodes(&series)...)
	}

	if len(missingEpisodes) > 0 {
		fmt.Println("Missing episodes are :")
		for _, value := range missingEpisodes {
			fmt.Println("* S0" + strconv.Itoa(value.AiredSeason) + "E" + strconv.Itoa(value.AiredEpisodeNumber))
		}
	} else {
		fmt.Println("No missing episodes")
	}

}
