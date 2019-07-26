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
	//listSeasonsDir, errListSeasonDirs := file.ListAllDirOnlyInDir(path)
	///*
	//	TODO : 2 ways to do it :
	//				* get seasons one by one and then do stuff
	//				* get all files recursively and do stuffs with regexs
	// */
	//for _, element := range listSeasonsDir {
	//	fmt.Println(element.Name())
	//}
	//if errListSeasonDirs != nil {
	//	log.Fatalln(err)
	//}

	err2 := c.GetSeriesEpisodes(&series, nil)
	if err2 != nil {
		log.Fatal(err)
	}
	listEpisodesFiles := []*file.EpisodeFile{
		file.NewEpisodeFile(1, 1),
		file.NewEpisodeFile(1, 2),
		file.NewEpisodeFile(1, 3),
		file.NewEpisodeFile(1, 6),
		file.NewEpisodeFile(1, 7),
	}
	seasonDir := file.NewSeasonDirSeasonV2(1, listEpisodesFiles)
	missingEpisodes := seasonDir.CheckMissingEpisodes(&series)
	fmt.Println(missingEpisodes)

	// Print the title of the episode 4x08 (season 4, episode 8)
}
