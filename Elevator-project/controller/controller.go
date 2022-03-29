package ctrl

import (
	"Elevator-go/Elevator/elevio"
	cf "Elevator-go/Elevator/type_"
	intf "Elevator-go/network/networkInterface"
)

var No_Id string = ""

func Elevator_controller(
	ch_localOrderFromButton chan elevio.ButtonEvent,
	ch_orderFromExternalElevator chan cf.OrderToExternalElev,
	ch_orderToLocalElevator chan elevio.ButtonEvent,
	ch_orderToExternalElevator chan cf.OrderToExternalElev) {
	for {
		select {
		case localOrder := <-ch_localOrderFromButton: /* Order is from local elevator */

			switch {
			case localOrder.Button == elevio.BT_Cab: /* if the local order is cab call, send it to local elevator */

				ch_orderToLocalElevator <- localOrder /* send order to local elevator */

			case localOrder.Button == elevio.BT_HallUp || localOrder.Button == elevio.BT_HallDown:

				if intf.MasterElevatorId == cf.LocalElevId { /* if the elavetor is a master, ... */

					if intf.OnlineElevsId[0] == cf.LocalElevId { /* it either excute the order locally if it is selected by 'SelectElevator()' or ... */

						ch_orderToLocalElevator <- localOrder /* send order to local elevator */
						SelectElevator()

					} else { /* it will send it to selected elevator if the order is not for the master. */

						/* send the local hall call order to the elevator selected by 'SelectElevator()' */
						ch_orderToExternalElevator <- cf.OrderToExternalElev{Order: localOrder, Elev_Id: intf.OnlineElevsId[0]}

						/* for i := 0; i < len(cf.OnlineElevatorsState); i++ {
							if cf.OnlineElevatorsState[i].ElevatorId == intf.OnlineElevsId[0] {
								fmt.Printf("%v request: %v\n", intf.OnlineElevsId[0],
									cf.OnlineElevatorsState[i].ElevatorState.ElevRequests[localOrder.Floor][localOrder.Button])
								break
							}
						} */
						SelectElevator()
					}

				} else { /* if the elevator is not a master, it is either ... */

					if len(intf.OnlineElevsId) == 0 { /* offline or ... */

						ch_orderToLocalElevator <- localOrder /* send order to local elevator */ /* confirm if it is serviced. NOT implemented yet. */

					} else { /* a slave. */

						/* Forward order to the master without elevator Id */
						ch_orderToExternalElevator <- cf.OrderToExternalElev{Order: localOrder, Elev_Id: No_Id}
					}
				}
			}

		case externalOrder := <-ch_orderFromExternalElevator: /* order from network */

			if intf.MasterElevatorId == cf.LocalElevId && externalOrder.Elev_Id == No_Id { /* if the local elevator is master and the order is forwarded from a slave */

				/* send the hall call order forwarded by a slave to the elevator selected by 'SelectElevator()' */
				ch_orderToExternalElevator <- cf.OrderToExternalElev{Order: externalOrder.Order, Elev_Id: intf.OnlineElevsId[0]}
				SelectElevator()

			} else if externalOrder.Elev_Id == cf.LocalElevId { /* if the local elevator is a slave and the order belongs to it */

				ch_orderToLocalElevator <- externalOrder.Order /* send order to local elevator */

			}

			//fmt.Printf("From controller: %v\n", intf.OnlineElevators)
		}
	}
}

func SelectElevator() {
	a := intf.OnlineElevsId[0]
	intf.OnlineElevsId = intf.OnlineElevsId[1:]
	intf.OnlineElevsId = append(intf.OnlineElevsId, a)
}

/* confirm if it is serviced. NOT implemented yet. */
func IsOrderServiced() bool {

	return true
}
