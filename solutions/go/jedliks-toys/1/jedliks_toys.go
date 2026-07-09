package jedlik

import "fmt"

// Drive drives the car one time.
func (car *Car) Drive() {
	if car.battery >= car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

// DisplayDistance returns the distance driven.
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery returns battery percentage.
func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// CanFinish checks if the car can finish the track.
func (car Car) CanFinish(trackDistance int) bool {
	return trackDistance <= (car.battery/car.batteryDrain)*car.speed
}