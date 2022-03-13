package elevator

import (
	"Elevator-go/elevio"
	"Elevator-go/fsm"
	cf "Elevator-go/type_"
	"time"
)

func Elevator(onRequestButtonPress chan elevio.ButtonEvent, ch_onFloorArrival chan int) {
	//, ch_obstruction chan bool)

	e := fsm.Fsm_onInitElevator()

	doorTimer := time.NewTimer(time.Duration(cf.DoorOpenDuration) * time.Second)

	for {
		select {
		case buttonEvent := <-onRequestButtonPress:
			fsm.Fsm_onRequestButtonPress(e, buttonEvent.Floor, buttonEvent.Button, doorTimer)
		case floor := <-ch_onFloorArrival:
			fsm.Fsm_onFloorArrival(e, floor, doorTimer)
		case <-doorTimer.C:
			fsm.Fsm_onDoorTimeout(e, doorTimer)

			/*case obstruction := <-ch_obstruction:
			if e.Behave == cf.DoorOpen && obstruction {
				doorTimer.Reset(time.Duration(cf.DoorOpenDuration) * time.Second)
			}*/
		}
	}
}
