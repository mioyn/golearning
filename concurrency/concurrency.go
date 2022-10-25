package main

import (
	"fmt"
	"sync"
	"time"
)

//main is also a goroutine, it is the main one so as soon as it exits the whole program does
// we need to use WaitGroup from sync package
func main() {
	var waitgrp sync.WaitGroup
	waitgrp.Add(2)

	go func() {
		defer waitgrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()

	go func() {
		defer waitgrp.Done()
		fmt.Println("Bob")
	}()

	waitgrp.Wait()
}
