//ALEXA
//8009709 
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var semRout = make(chan int, NumRoutines)


var semDisp = make(chan int, 1)


const (
	NumRoutines = 3
	NumRequests = 1000
)


type Task struct {
	a, b float32
	disp chan float32
}


var wgRout sync.WaitGroup
var wgDisp sync.WaitGroup


func solve(tsk *Task) {
	randNum := rand.Intn(16) 
	randnumber := time.Duration(randNum)
	time.Sleep( randnumber * time.Second)
	wgRout.Add(1)
	go func(){
	addition := tsk.a + tsk.b
	DisplayServer() <- addition

}()
	wgDisp.Done()
}






func ComputeServer() chan *Task {
 
	send := make(chan *Task, NumRoutines)
	
	go func() {
		for {
			semRout <- 1
			go handleReq(<-send)
			<-semRout
		}
		
	}()

	return send
}

func handleReq(taskHandling *Task) {
	solve(taskHandling)

}

func DisplayServer() chan float32 {
	/*
	receieves the addition function from solve
	*/
	taker := make(chan float32, NumRoutines)

	
	go func() {
		for i:=0; i < len(taker); i++ {
			semDisp <- 1
			fmt.Println("------------------")
			fmt.Printf("Result %v \n",<-taker)
			wgRout.Done()
			<-semDisp
	
		}
	
	}()

	return taker

}

func main() {
	dispChan := DisplayServer()
	reqChan := ComputeServer()
	for {
		var a, b float32
		// make sure to use semDisp
		semDisp <- 1
		fmt.Print("Enter two numbers: ")
		fmt.Scanf("%f %f \n", &a, &b)
		fmt.Printf("%f %f \n", a, b)
		<-semDisp
		if a == 0 && b == 0 {
			break
		}
		// Create task and send to ComputeServer
		
		var tsk Task
		tsk.a = a
		tsk.b = b
		tsk.disp = dispChan
		wgDisp.Add(1)
		reqChan <- &tsk
		time.Sleep(1e9)

	}

	wgDisp.Wait()
	wgRout.Wait()
}
