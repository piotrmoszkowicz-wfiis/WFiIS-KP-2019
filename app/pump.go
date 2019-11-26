package app

import "github.com/stianeikeland/go-rpio/v4"

// Structure which keeps Pump data inside
type Pump struct {
	Pin  rpio.Pin
	Type bool
}

// Initializes pump structure
func Init(pin uint8, pumpType bool) *Pump {
	pump := Pump{rpio.Pin(pin), pumpType}

	pump.Pin.Output()

	return &pump
}

// Start flow of the liquid
func (p *Pump) StartFlow() {
	if p.Type {
		p.Pin.High()
	} else {
		p.Pin.Low()
	}

}

// Ends flow of the liquid
func (p *Pump) StopFlow() {
	if p.Type {
		p.Pin.Low()
	} else {
		p.Pin.High()
	}
}
