package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
	}
}

// TODO: define the 'Track' type struct
type Track struct {
	distance int
}

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
	return car
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	// Calculate the number of times the car needs to drive to cover the track distance
	requiredDrives := track.distance / car.speed
	// Adjust for any remaining distance that cannot be covered with a full drive
	if track.distance%car.speed != 0 {
		requiredDrives++
	}

	// Calculate the total battery drain for the required drives
	totalBatteryDrain := requiredDrives * car.batteryDrain

	// Check if the car has enough battery to cover the total battery drain
	return car.battery >= totalBatteryDrain
}
