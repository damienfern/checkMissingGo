package main

import (
	"./config/env"
	"fmt"
	"github.com/pioz/tvdb"
)

func main() {
	fmt.Println("Hello, Arch!")

	env.LoadEnv()

	c := tvdb.Client{
		Apikey:   env.FindEnvVarOrFail("apiKey"),
		Userkey:  env.FindEnvVarOrFail("userKey"),
		Username: env.FindEnvVarOrFail("username"),
	}

	err := c.Login()
	if err != nil {
		panic(err)
	}

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
