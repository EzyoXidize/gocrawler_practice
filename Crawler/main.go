package main

import (
	"practice/Crawler/engine"
	"practice/Crawler/zhenai/parser"
	"practice/Crawler/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 10,
	}

	e.Run(engine.Request{
		Url			: "http://www.zhenai.com/zhenghun",
		ParserFunc	: parser.ParserCityList,
	})
}

