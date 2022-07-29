package main

import (
	"crawler/crawler_distributed/rpcsupport"
	"crawler/engine"
	"testing"
)

func TestItemSaver(t *testing.T) {
	go serveRpc(":1234", "test1")
	client, err := rpcsupport.NewClient(":1234")
	if err != nil {
		panic(err)
	}

	result := ""
	item := engine.Item{}
	err = client.Call("ItemSaverService.Save", item, &result)
	if err != nil || result != "ok" {
		t.Errorf("wrong")
	}

}
