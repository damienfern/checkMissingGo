package file

import (
	"fmt"
	"github.com/pioz/tvdb"
	"os"
)

type SeasonDir struct {
	SeasonID     int
	EpisodeFiles []*EpisodeFile
	FileInfo     os.FileInfo
	Filepath     string
}

func NewSeasonDir(seasonID int, episodeFiles []*EpisodeFile, fileInfo os.FileInfo, rootpath string) *SeasonDir {
	return &SeasonDir{SeasonID: seasonID, EpisodeFiles: episodeFiles, FileInfo: fileInfo, Filepath: rootpath + "/" + fileInfo.Name()}
}
func NewSeasonDirSeasonV2(seasonID int, episodeFiles []*EpisodeFile) *SeasonDir {
	return &SeasonDir{SeasonID: seasonID, EpisodeFiles: episodeFiles}
}

func (s SeasonDir) toString() {
	fmt.Println(s.SeasonID, s.Filepath)
}

func (s SeasonDir) CheckMissingEpisodes(series *tvdb.Series) []*tvdb.Episode {
	var result []*tvdb.Episode
	dbEpisodes := series.GetSeasonEpisodes(s.SeasonID)

	for _, element := range dbEpisodes {
		isPresent := isEpisodeInEpisodeFileList(s.EpisodeFiles, func(file *EpisodeFile) bool {
			return file.SeasonID == element.AiredSeason && file.EpisodeID == element.AiredEpisodeNumber
		})
		if !isPresent {
			result = append(result, element)
		}
	}

	return result
}
