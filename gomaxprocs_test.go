package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	totalcpu := runtime.NumCPU()
	fmt.Println("Total Cpu :", totalcpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total Thread :", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}
	totalcpu := runtime.NumCPU()
	fmt.Println("Total Cpu :", totalcpu)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total Thread :", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine", totalGoroutine)

	group.Wait()
}
