package main

import (
	"./config/args"
	"./config/connection"
	"./tvdb"
	"fmt"
	"regexp"
)

func main() {
	c := connection.ConnectToTVDB()
	path := args.GetDirPathToSearch()
	fmt.Println(path)
	regexpFilePath := regexp.MustCompile("^/?(.+/)*(.+)$")
	filePathArray := regexpFilePath.FindStringSubmatch(path)
	seriesName := filePathArray[len(filePathArray)-1]

	series := tvdb.FindSeriesOrFail(seriesName, &c)

	err := c.GetSeriesEpisodes(&series, nil)
	if err != nil {
		panic(err)
	}

	// Print the title of the episode 4x08 (season 4, episode 8)
	fmt.Println(series.GetEpisode(4, 8).Overview)
}
