package initializers

import "github.com/elastic/go-elasticsearch/v8"

var ES8 *elasticsearch.Client

func ConnectES8() {
	var err error
	ES8, err = elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}
}
