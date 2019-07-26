package tvdb

import (
	"fmt"
	"github.com/pioz/tvdb"
)

func FindSeriesOrFail(name string, tvdbConnection *tvdb.Client) tvdb.Series {
	series, err := tvdbConnection.BestSearch(name)
	if err != nil {
		// The request response is a 404: this means no results have been found
		if tvdb.HaveCodeError(404, err) {
			fmt.Println("Series" + name + " not found")
		} else {
			panic(err)
		}
	}
	return series
}
