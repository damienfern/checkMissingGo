package tvdb

import (
	"github.com/pioz/tvdb"
	log "github.com/sirupsen/logrus"
)

func FindSeriesOrFail(name string, tvdbConnection *tvdb.Client) tvdb.Series {
	series, err := tvdbConnection.BestSearch(name)
	if err != nil {
		// The request response is a 404: this means no results have been found
		if tvdb.HaveCodeError(404, err) {
			log.Fatal("Series " + name + " not found in TVDB")
		} else {
			log.Fatal(err)
		}
	}
	return series
}
