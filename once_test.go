package main

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	group := sync.WaitGroup{}
	

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			
			group.Done()
			once.Do(OnlyOnce)
			
		}()
	}

	group.Wait()
	fmt.Println("Counter", counter)
}
