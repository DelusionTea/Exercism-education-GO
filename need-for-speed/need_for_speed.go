package speed

// TODO: define the 'Car' type struct
// - battery
//- batteryDrain
//- speed
//- distance
type Car struct {
	battery, batteryDrain, speed, distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	//panic("Please implement the NewCar function")
	car := Car{
		battery:      100,
		speed:        speed,
		batteryDrain: batteryDrain,
	}
	return car

}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack created a new track
func NewTrack(distance int) Track {
	//panic("Please implement the NewTrack function")
	track := Track{
		distance: distance,
	}
	return track
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(car Car) Car {
	//panic("Please implement the Drive function")
	//ar{speed: 5, batteryDrain: 2, battery: 98, distance: 5}
	if car.batteryDrain > car.battery {
		return car
	}
	car.distance += car.speed
	car.battery -= car.batteryDrain
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	//panic("Please implement the CanFinish function")
	//track -> distance
	//distance/speed -> count of iterations
	//if battery <(batteryDrain*(disptance/speed) - return false
	countOfIteration := track.distance / car.speed
	if car.battery < (car.batteryDrain * countOfIteration) {
		return false
	}
	return true
}
