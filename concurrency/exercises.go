package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// Wait group
func ex1() {

	fmt.Println("gs start\t", runtime.NumGoroutine())

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("Hello from thing one")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from thing two")
		wg.Done()
	}()

	fmt.Println("gs mid\t", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("gs end\t", runtime.NumGoroutine())
	fmt.Println("about to exit")
}

// Methods sets
type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello", p.first)
}

func saySomething(h human) {
	h.speak()
}

func ex2() {
	p := person{
		first: "Nick",
	}
	p.speak()
	saySomething(&p)
}

func ex3() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Println(incrementer)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end\t", incrementer)
}

// Mutex
func ex4() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100

	var m sync.Mutex
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)
			m.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end\t", incrementer)

}

func ex5() {
	var wg sync.WaitGroup
	var incrementor int64
	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&incrementor, 1)
			fmt.Println(atomic.LoadInt64(&incrementor))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end\t", atomic.LoadInt64(&incrementor))
}

func ex6() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
}

func exercises() {
	// ex1()
	// ex2()
	// ex3()
	// ex4()
	// ex5()
	ex6()
}
