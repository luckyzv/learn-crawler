package main

import (
	"crawler/app/zhenai/parser"
	"crawler/engine"
	"crawler/persist"
	"crawler/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 3,
		ItemChan:    persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
