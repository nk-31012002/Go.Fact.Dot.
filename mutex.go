package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer wg.Done()

	p.mu.Lock()
	p.views++
	p.mu.Unlock()

}

func main() {
	myPost := post{views: 0}
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	//myPost.inc()
	wg.Wait()
	fmt.Println(myPost.views)
}
