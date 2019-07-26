package main

import (
	"./config/args"
	"./config/connection"
	"fmt"
)

func main() {
	c := connection.ConnectToTVDB()
	path := args.GetDirPathToSearch()
	fmt.Println(path)
	series, err := c.BestSearch("Game of Thrones")
	if err != nil {
		panic(err)
	}

	err = c.GetSeriesEpisodes(&series, nil)
	if err != nil {
		panic(err)
	}

	// Print the title of the episode 4x08 (season 4, episode 8)
	fmt.Println(series.GetEpisode(4, 8).Overview)
}
