package args

import (
	"../../file"
	"errors"
	log "github.com/sirupsen/logrus"
	"os"
)

func GetDirPathToSearch() (string, error) {
	arg := getArgByIndex(1)
	exists, err := file.Exists(arg)
	if !exists {
		err = errors.New("dir " + arg + " not found")
	}
	return arg, err
}

func getArgByIndex(index int) string {
	if index >= len(os.Args) {
		log.Fatalln("no arg")
	}
	arg := os.Args[index]
	return arg
}

func getAllArgs() []string {
	args := os.Args[1:]
	return args
}
