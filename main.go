package main

import (
	"./config/args"
	"./config/connection"
	"./file"
	"./tvdb"
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"regexp"
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
		TODO : get all files recursively and do stuffs with regexs
	*/
	for _, element := range listSeasonsDir {
		fmt.Println(element.Name())
	}
	if errListSeasonDirs != nil {
		log.Fatalln(err)
	}

	error := c.GetSeriesEpisodes(&series, nil)
	if error != nil {
		log.Fatal(err)
	}

	// Print the title of the episode 4x08 (season 4, episode 8)
	fmt.Println(series.GetEpisode(4, 8).Overview)
}
