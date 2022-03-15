package mediator

func Test() {
	stationManager := newStationManger()
	passengerTrain := &passengerTrain{
		mediator: stationManager,
	}
	goodsTrain := &goodsTrain{
		mediator: stationManager,
	}
	passengerTrain.requestArrival()
	goodsTrain.requestArrival()
	passengerTrain.departure()
}
