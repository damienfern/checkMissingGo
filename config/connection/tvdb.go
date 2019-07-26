package connection

import (
	"../env"
	"github.com/pioz/tvdb"
	log "github.com/sirupsen/logrus"
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
		log.Fatal(err)
	}
	return c
}
