package file

import (
	"os"
	"regexp"
	"strconv"
)

type EpisodeFile struct {
	SeasonID, EpisodeID int
	FileInfo            os.FileInfo
	Filepath            string
}

func NewEpisodeFile(fileInfo os.FileInfo, rootpath string) *EpisodeFile {
	regexpFilePath := regexp.MustCompile("S([0-9]{2})E([0-9]{2})")
	filePathArray := regexpFilePath.FindStringSubmatch(fileInfo.Name())
	seasonID, _ := strconv.Atoi(filePathArray[1])
	episodeID, _ := strconv.Atoi(filePathArray[2])
	filepath := rootpath + "/" + fileInfo.Name()
	return &EpisodeFile{SeasonID: seasonID, EpisodeID: episodeID, FileInfo: fileInfo, Filepath: filepath}
}

func isEpisodeInEpisodeFileList(vs []*EpisodeFile, f func(file *EpisodeFile) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}
