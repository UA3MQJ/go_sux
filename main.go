package main

import (
    "fmt"
	"time"
	"runtime"
)

func green_process() {
	fmt.Println("start gorutine")
	i := 0
	for {
		i = i + 1
	}
}

func getGOMAXPROCS() int {
    return runtime.GOMAXPROCS(0)
}

func main() {
	fmt.Printf("GOMAXPROCS is %d\n", getGOMAXPROCS())
	fmt.Println("Start")

	go green_process()

	time.Sleep(1 * time.Second)
	fmt.Println("Finish")
}