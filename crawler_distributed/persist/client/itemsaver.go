package client

import (
	"crawler/crawler_distributed/rpcsupport"
	"crawler/engine"
	"log"
)

func ItemSaver(host string) (chan engine.Item, error) {
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	itemCount := 0
	go func() {
		for {
			item := <-out
			log.Printf("Item saver: got item #%d: %v", itemCount, item)
			itemCount++

			result := ""
			// call rpc
			client.Call("ItemSaverService.Save", item, &result)
		}
	}()
	return out, nil
}
