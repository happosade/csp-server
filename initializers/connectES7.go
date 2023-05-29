package initializers

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

var ES *elasticsearch.Client

func ConnectES7() {
	var err error
	ES, err = elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}
	log.Println(ES.Info())
}
