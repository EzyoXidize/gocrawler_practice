package main

import (
	"practice/Crawler/engine"
	"practice/Crawler/scheduler"
	"practice/Crawler/zhenai/parser"
	"practice/Crawler/persist"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler  	: &scheduler.QueuedScheduler{},
		WorkerCount	: 100,
		ItemChan	: persist.ItemSaver(),
	}

	//e.Run(engine.Request{
	//	Url			: "http://www.zhenai.com/zhenghun",
	//	ParserFunc	: parser.ParserCityList,

	e.Run(engine.Request{
		Url			: 	"http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc	: 	parser.ParseCity,
	})
}

