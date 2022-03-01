package main
impor
const Floors = 4
const Ack = 1
const nAck = 0

type ElevatorDirection int

const (
	ElevatorUp         ElevatorDirection = 1  //Elevator moving up
	ElevatorDown       ElevatorDirection = -1 //Elevator moving down
	ElevatorStationary ElevatorDirection = 0  //Elevator stationary
)

type DoorStatus int

const (
	Open   DoorStatus = 1  // Door open
	Closed DoorStatus = -1 // Door closed
	Other  DoorStatus = 0  // Door opening or closing
)

// var elevatorPosition int = 1
type ElevatorPositon int

const (
	Level_1 ElevatorPosition = 1 // Elevator is at Level 1
	Level_2 ElevatorPosition = 2 // Elevator is at Level 2
	Level_3 ElevatorPosition = 3 // Elevator is at Level 3
	Level_4 ElevatorPosition = 4 // Elevator is at Level 4
)

type Elevator struct {
	Id                  string // ID of elevator
	Direction           ElevDirection
	DoorStatus          DoorStatus
	Service_state       string // Active or inactive i.e to know if Elevator is online, working well (having power to motors, not lost connection, door stuck or stuck somewhere, e.tc)
	status              bool   // Elevator is in use or idle
	last_floor_serviced string // The last floor the elevator serviced
	running_hours       int    // How many hours elevator has been running
	destination         string // Last point of service - This tells the elevator's final stop as per the user(s) request
}

type CabinButtonPress struct {
	Destinations []ElevatorPosition
	obsk         bool // obstruction switch active or inactive - this keeps the elevator door open
	stop         int  // Emergency stop
}

type HallButtonPress struct {
	Direction ElevDirection
	Floor     ElevatorPosition
	Timestamp int // import timee object
}

// func (h *HallButtonPress) Press(){
// 	if h.Floor === Level1 {
// 		// do something special
// 		Direction  = ElevatorUp
// 	}
// }

func getElDirection() int {
	ED := ElevDirection{
		moving_Up:   0,
		moving_Down: 0,
		stopped:     0,
	}

	switch MotorDirection {
	case 1:
		ED.moving_Up = Level_1
		return 1
	case -1:
		ED.moving_Up = 2
		return 2
	case 0:
		ED.stopped = -1
		return -1
	}

}

func getDrStatus() {

}
func getElPositon(getFloor int) int {

	EP := ElevatorPositon{
		level_1: 0,
		level_2: 0,
		level_3: 0,
		level_4: 0,
	}
	switch getFloor {
	case 1:
		EP.level_1 = 1
		return 1
	case 2:
		EP.level_2 = 2
		return 2
	case 3:
		EP.level_3 = 3
		return 3
	case 4:
		EP.level_4 = 4
		return 4
	default:
		return -1
	}

}
func getElProfile() {

}
func getCbnBtnPress() {

}
func getHllBtnPress() {

}

func main() {
	// looping forever
	// detect when aswitch on a level is pressed
	//      create instance ofhall press (record direction, hall's switch was pressed..) -> store it in a list of hallpress objects
	//


	// 1. Create an elevator (provide an id)
	// 2. Instantiate all the halls switches and buttons and etc
	// 3. run elevator
	// 4. coordinate hall button presses and elevator arrival
	// 5. COordinate cabin button presses and delvery

	// var hallPresses = []HallButtonPress{}  // creates a queue of requests or orders
	// for {
	// 	var currentFloor = getCurrentLevel() // 1
	// 	var press = getCurrentPressed()// 1,2,3,4
	// 	var hallPress = HallButtonPress{
	// 		Floor: currentFloor
	// 		Direction: direction
	// 		Time: time
	// 	}
	// 	hallPresses = append(hallPresses, hallPress)

	// }
}
