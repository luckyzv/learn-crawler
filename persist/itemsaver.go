package persist

import (
	"crawler/engine"
	"github.com/olivere/elastic/v7"
	"log"
)

func ItemSaver() (chan engine.Item, error) {
	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}
	out := make(chan engine.Item)
	itemCount := 0
	go func() {
		result := <-out
		log.Printf("Item saver: got item #%d: %v", itemCount, result)
		itemCount++

		Save(client, "data_profile", result)
	}()
	return out, nil
}

func Save(client *elastic.Client, index string, item engine.Item) error {
	return nil
}
