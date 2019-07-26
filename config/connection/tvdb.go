package connection

import (
	"../env"
	"github.com/pioz/tvdb"
)

func ConnectToTVDB() tvdb.Client {
	env.LoadEnv()

	c := tvdb.Client{
		Apikey:   env.FindEnvVarOrFail("apiKey"),
		Userkey:  env.FindEnvVarOrFail("userKey"),
		Username: env.FindEnvVarOrFail("username"),
		Language: env.FindEnvVarOrFail("language"),
	}

	err := c.Login()
	if err != nil {
		panic(err)
	}
	return c
}
