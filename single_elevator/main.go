package main

import (
	"Elevator-go/elevator"
	"Elevator-go/elevio"
)

func main() {
	elevio.Init("localhost:48613", 4)
	/*Stop button is not working*/
	ch_onRequestButtonPress := make(chan elevio.ButtonEvent)
	ch_onFloorArrival := make(chan int)
	ch_obstruction := make(chan bool)
	ch_onStopButtonPress := make(chan bool)

	go elevio.PollFloorSensor(ch_onFloorArrival)
	go elevio.PollObstructionSwitch(ch_obstruction)
	go elevio.PollButtons(ch_onRequestButtonPress)
	go elevio.PollStopButton(ch_onStopButtonPress)

	elevator.Elevator(ch_onRequestButtonPress, ch_onFloorArrival,
		ch_onStopButtonPress, ch_obstruction)

	select {}
}
