package main

import "fmt"


const uSixteenBitMax float64 = 65535
const KMHMultiple float64 = 1.60934

type car struct {
	GasPedal uint16
	BrakePedal uint16
	SteeringWheel int16
	TopSpeedKMH float64
}

// Get a copy, receiver type
func (c car) kmh() float64 {
	return float64(c.GasPedal) * (c.TopSpeedKMH/uSixteenBitMax)
}

// Get a copy, receiver type
func (c car) mph() float64 {
	return float64(c.GasPedal) * (c.TopSpeedKMH/KMHMultiple/uSixteenBitMax)
}
// Modify struct itself via pointer type
func (c *car) newTopSpeed (newSpeed float64){
	c.TopSpeedKMH = newSpeed
}

func main(){
	aCar := car{
		GasPedal: 22341,
		BrakePedal : 0,
		SteeringWheel: 12345,
		TopSpeedKMH: 225.0,
	}

	fmt.Println(aCar.kmh())
	aCar.newTopSpeed(500)
	fmt.Println(aCar.kmh())


}