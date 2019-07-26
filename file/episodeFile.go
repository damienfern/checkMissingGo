package file

import (
	"os"
)

type EpisodeFile struct {
	SeasonID, EpisodeID int
	FileInfo            os.FileInfo
	Filepath            string
}

func NewEpisodeFile(fileInfo os.FileInfo, rootpath string) *EpisodeFile {
	// TODO : Regex things
	seasonID := 1
	episodeID := 1
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
