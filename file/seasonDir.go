package file

import (
	"fmt"
	"github.com/pioz/tvdb"
	"os"
	"regexp"
	"strconv"
)

type SeasonDir struct {
	SeasonID     int
	EpisodeFiles []*EpisodeFile
	FileInfo     os.FileInfo
	Filepath     string
}

func NewSeasonDirSeason(info os.FileInfo, rootpath string) *SeasonDir {
	regexpFilePath := regexp.MustCompile("^(Saison|Season) ([0-9]{1,2})$")
	filePathArray := regexpFilePath.FindStringSubmatch(info.Name())
	seasonID, _ := strconv.Atoi(filePathArray[2])
	seasonDir := SeasonDir{SeasonID: seasonID, FileInfo: info, Filepath: rootpath + "/" + info.Name()}
	seasonDir.fillEpisodeFiles()
	return &seasonDir
}

func (s *SeasonDir) fillEpisodeFiles() {
	listEpisodesFiles, _ := ListAllFilesOnlyInDir(s.Filepath)
	for _, oneFile := range listEpisodesFiles {
		s.EpisodeFiles = append(s.EpisodeFiles, NewEpisodeFile(oneFile, s.Filepath))
	}
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
