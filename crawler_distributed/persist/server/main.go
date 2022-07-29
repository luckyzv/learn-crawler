package main

import (
	"crawler/crawler_distributed/persist"
	"crawler/crawler_distributed/rpcsupport"
	"github.com/olivere/elastic/v7"
)

func main() {
	serveRpc(":1235", "data_profile")
}

func serveRpc(host, index string) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	return rpcsupport.ServeRpc(host, &persist.ItemSaverService{
		Client: client,
		Index:  index,
	})
}
