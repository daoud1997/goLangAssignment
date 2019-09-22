// Q3_8252115
package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

type Point struct {
	x float64
	y float64
}
type Triangle struct {
	A Point
	B Point
	C Point
}
type Stack struct {
	lockMaster sync.Mutex
	stk        []Triangle
}

var numberInBigStack int = 0
var numberInSmallStack int = 0

func newStack() (stacks Stack) {
	stacks = Stack{sync.Mutex{}, make([]Triangle, 0)}
	return stacks
}

func (stk *Stack) Push(i int, triangle []Triangle) {

	stk.lockMaster.Lock()
	stk.stk = append(stk.stk, triangle[i])
	stk.lockMaster.Unlock()
}

func triangles10000() (result [10000]Triangle) {
	rand.Seed(2120)
	for i := 0; i < 10000; i++ {
		result[i].A = Point{rand.Float64() * 100., rand.Float64() * 100.}
		result[i].B = Point{rand.Float64() * 100., rand.Float64() * 100.}
		result[i].C = Point{rand.Float64() * 100., rand.Float64() * 100.}
	}
	return
}

func (t Triangle) Perimeter() float64 {

	lengthfromAtoB := math.Sqrt((math.Pow((t.B.x-t.A.x), 2) + math.Pow((t.B.y-t.A.y), 2)))
	lengthfromBtoC := math.Sqrt((math.Pow((t.C.x-t.B.x), 2) + math.Pow((t.C.y-t.B.y), 2)))
	lengthfromAtoC := math.Sqrt((math.Pow((t.C.x-t.A.x), 2) + math.Pow((t.C.y-t.A.y), 2)))

	return (lengthfromAtoB + lengthfromBtoC + lengthfromAtoC)
}

func (t Triangle) Area() float64 {

	areaofTriangle := ((t.B.x - t.A.x) * (t.C.y - t.A.y)) - ((t.C.x-t.A.x)*(t.B.y-t.A.y))*0.5
	return areaofTriangle
}



func classifyTriangles(highRatio *Stack, lowRatio *Stack, ratioThreshold float64, triangles []Triangle) {

	ratioAboveOne := highRatio
	ratioBelowOne := lowRatio

	for i := 0; i < len(triangles); i++ {

		if (triangles[i].Perimeter())/(triangles[i].Area()) > ratioThreshold {
			ratioAboveOne.Push(i, triangles)
			numberInBigStack++
		} else {
			ratioBelowOne.Push(i, triangles)
			numberInSmallStack++
		}
	}

}

func main() {

	var ratioThreshold float64 = 1.0
	triangles := triangles10000()
	highRatio := newStack()
	lowRatio := newStack()
	var slice1 []Triangle = triangles[0:1000]
	var slice2 []Triangle = triangles[1000:2000]
	var slice3 []Triangle = triangles[2000:3000]
	var slice4 []Triangle = triangles[3000:4000]
	var slice5 []Triangle = triangles[4000:5000]
	var slice6 []Triangle = triangles[5000:6000]
	var slice7 []Triangle = triangles[6000:7000]
	var slice8 []Triangle = triangles[7000:8000]
	var slice9 []Triangle = triangles[8000:9000]
	var slice10 []Triangle = triangles[9000:10000]
	arrTriangle := []([]Triangle){slice1, slice2, slice3, slice4, slice5, slice6, slice7, slice8, slice9, slice10}

	for _, index := range arrTriangle {
		go classifyTriangles(&highRatio, &lowRatio, ratioThreshold, index)
		time.Sleep(1 * time.Second)
	}

	fmt.Println("The number of traingles in the highRatio is: ", numberInBigStack)
	fmt.Println("The number of traingles in the lowRatio is: ", numberInSmallStack)
	fmt.Println("The top most element (triangle) in the highRatio stack is: ", highRatio.stk[numberInBigStack-1])
	fmt.Println("The top most element (triangle) in the lowRatio stack is:", lowRatio.stk[numberInSmallStack-1])

}
