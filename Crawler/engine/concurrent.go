package engine

type ConcurrentEngine struct {
	Scheduler		Scheduler
	WorkerCount 	int
	ItemChan 		chan interface{} // 保存 item 用的通道
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run (seeds ...Request)  {

	out := make(chan  ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		creatWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)
	}

	for _, r := range seeds {
		 e.Scheduler.Submit(r)
	}


	for  {
		result := <- out
		for _,item := range result.Items {
			go func() { e.ItemChan <- item }() // 打印换成保存 , 使用 goroutine
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}

}

func creatWorker(in chan Request,out chan ParseResult, ready ReadyNotifier)  {

	go func() {
		for {
			// tell scheduler im ready
			ready.WorkerReady(in)
			request := <- in
			result,err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
