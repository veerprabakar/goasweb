package main

import (
	"fmt"
	"sync"
	"syscall/js"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		runit("goat")
		wg.Done()
	}()

	go func() {
		runit("tiger")
		wg.Done()
	}()

	wg.Wait()
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}

func add(i []js.Value) {
	fmt.Println(i)
	js.Global().Set("output", js.ValueOf(i[0].Int()+i[1].Int()))
}

func registerCallbacks() {
	js.Global().Set("add", js.NewCallback(add))
}

func runit(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
