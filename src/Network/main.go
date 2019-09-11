package main

import (
	"Network/work"
	"log"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var names = []string{
	"name1",
	"name2",
	"name3",
	"name4",
	"name5",

}

type namePrinter struct {
	name string
}

func (m *namePrinter)Task()  {
	log.Println(m.name)
	time.Sleep(time.Second * 10)
}

func main(){
	work.Start()
}


func testWork(){
	runtime.GOMAXPROCS(3)

	p := work.New(3)
	start := time.Now()

	var wg sync.WaitGroup
	wg.Add(100 * len(names))


	for i:= 0; i < 100; i++{
		for _,name := range names{
			np := namePrinter{
				name:strconv.Itoa(i) + ":" + name,
			}

			go func() {
				p.Run(&np)
				wg.Done()
			}()
		}
	}
	wg.Wait()

	p.Shutdown()

	end := time.Now()

	log.Println(end.Sub(start))
}