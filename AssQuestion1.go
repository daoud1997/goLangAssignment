// AssQuestion1
// Daoud Hamadneh, 8252115
package main

import (
	"errors"
	"fmt"
)

const DistanceToMontreal = 200. //As given in km
const DistanceToToronto = 400.  //As given in km

type Trip struct {
	destination string
	weight      float32
	deadline    int
}
type Truck struct {
	vehicle     string
	name        string
	destination string
	speed       float32
	capacity    float32
	load        float32
	success     bool
}
type Pickup struct {
	truck     Truck
	isPrivate bool
}

type TrainCar struct {
	truck   Truck
	railway string
}

func NewTruck(name string) Truck {
	var truck Truck
	truck.vehicle = "Truck"
	truck.name = name
	truck.destination = ""
	truck.speed = 40
	truck.capacity = 10
	truck.load = 0
	truck.success = false
	return truck
}

func NewPickup(name string) Pickup {
	var pickup Pickup
	pickup.truck.vehicle = "Pickup"
	pickup.truck.name = name
	pickup.truck.destination = ""
	pickup.truck.speed = 60
	pickup.truck.capacity = 2
	pickup.truck.load = 0
	pickup.isPrivate = true
	return pickup
}
func NewTrainCar(name string) TrainCar {
	var traincar TrainCar
	traincar.truck.vehicle = "TrainCar"
	traincar.truck.name = name
	traincar.truck.destination = ""
	traincar.truck.speed = 30
	traincar.truck.capacity = 30
	traincar.truck.load = 0
	traincar.railway = "CNR"
	return traincar
}

func NewTorontoTrip(weight float32, deadline int) *Trip {
	var trip Trip
	trip.destination = "Toronto"
	trip.deadline = deadline
	trip.weight = weight
	return &trip
}
func NewMontrealTrip(weight float32, deadline int) *Trip {
	var trip Trip
	trip.destination = "Montreal"
	trip.deadline = deadline
	trip.weight = weight
	return &trip
}

type Transporter interface {
	addLoad(trip Trip) error
	print()
}

func (truck *Truck) addLoad(trip Trip) error {

	var success = false

	if truck.destination == "" {
		truck.destination = trip.destination
		success = true
	}

	if truck.destination != trip.destination {
		return errors.New("Error: other destination")
	}

	if truck.capacity < trip.weight {
		if truck.load == 0 {
			truck.destination = ""
		}
		return errors.New("Error: Out of Capacity")
	}
	var goodTime bool

	if trip.destination == "Toronto" {
		//time :=
		goodTime = (int(DistanceToToronto/truck.speed) <= trip.deadline)
	}
	if trip.destination == "Montreal" {
		//timeNeeded :=
		goodTime = int((DistanceToMontreal / truck.speed)) <= trip.deadline
	}

	if !goodTime {
		if success {
			truck.destination = ""
		}
		return errors.New("Error: Not on time")
	}
	truck.load = truck.load + trip.weight
	truck.capacity = truck.capacity - truck.load
	return nil
}

func (truck *Truck) print() {

	fmt.Printf("%s to %s with %f tons \n", truck.name, truck.destination, truck.load)

}

func (pickup *Pickup) print() {

	fmt.Printf("%s to %s with %f tons (Private: %t)\n", pickup.truck.name, pickup.truck.destination, pickup.truck.load, pickup.isPrivate)

}

func (traincar *TrainCar) print() {

	fmt.Printf("%s to %s with %f tons (%s)\n", traincar.truck.name, traincar.truck.destination, traincar.truck.load, traincar.railway)

}

func (truck *Pickup) addLoad(trip Trip) error {
	// ONLY DID IT SO THAT IT CAN IMPLEMENT THE INTERFACE, TRANSPORTER
	// THEREFORE, EMBEDDED TYPE IS APPLICABLE
	return nil

}

func (truck *TrainCar) addLoad(trip Trip) error {
	// ONLY DID IT SO THAT IT CAN IMPLEMENT THE INTERFACE, TRANSPORTER
	// THEREFORE, EMBEDDED TYPE IS APPLICABLE
	return nil

}

func main() {
	defer fmt.Println("THANK YOU FOR YOUR PATIENCE")
	fmt.Println("HELLO MY FRIEND!!")
	truck_A := NewTruck("Truck A")
	truck_B := NewTruck("Truck B")
	pick_A := NewPickup("PickUp A")
	pick_B := NewPickup("Pickup B")
	pick_C := NewPickup("Pickup C")
	traincar_A := NewTrainCar("TrainCar A")
	transportation := []Transporter{&truck_A, &truck_B, &pick_A.truck, &pick_B.truck, &pick_C.truck, &traincar_A.truck}
	transPrinting := []Transporter{&pick_A, &pick_B, &pick_C, &traincar_A}
	var location string
	var weight float32
	var time int
	for {
		var trip *Trip
		var IsToronto bool
		var IsMontreal bool

		fmt.Print("Destination : (T)oronto, (M)ontreal, else exit ? ")
		fmt.Scanln(&location)
		FirstLetter := location[0:1]
		//fmt.Println(FirstLetter)
		if FirstLetter == "T" || FirstLetter == "t" {
			IsToronto = true
			//fmt.Println("Toronto")

		} else if FirstLetter == "M" || FirstLetter == "m" {
			IsMontreal = true
			//fmt.Println("Montreal")
		} else {
			fmt.Println(" Not going to Montreal nor to Toronto, BYE!!!!!")
			break
		}

		fmt.Print("Weight: ")
		fmt.Scanf("%v \n", &weight)
		fmt.Print("Deadline (In Hours): ")
		fmt.Scanf("%d \n", &time)

		if IsToronto {
			trip = NewTorontoTrip(weight, time)
		}
		if IsMontreal {
			trip = NewMontrealTrip(weight, time)
		}

		for _, transporters := range transportation {

			ill := transporters.addLoad(*trip)
			if ill == nil {
				break
			}
			fmt.Println(ill)
		}
	}
	for i := 0; i < 2; i++ {
		transportation[i].print()
	}

	for _, transporters := range transPrinting {
		transporters.print()
	}

}
