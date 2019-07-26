package args

import "os"

func GetDirPathToSearch() string {
	arg := getArgByIndex(1)
	return arg
}

func getArgByIndex(index int) string {
	arg := os.Args[index]
	return arg
}

func getAllArgs() []string {
	args := os.Args[1:]
	return args
}
