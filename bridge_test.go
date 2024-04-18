package bridge

import "testing"

func TestCircularSwitchBridgeBulb(t *testing.T) {
	m := NewCircularSwitch(&Bulb{})
	m.TurnOn()
	m.TurnOff()
}

func TestCircularSwitchBridgeFan(t *testing.T) {
	m := NewCircularSwitch(&Fan{})
	m.TurnOn()
	m.TurnOff()
}

func TestSquareSwitchBridgeBulb(t *testing.T) {
	m := NewSquareSwitch(&Bulb{})
	m.TurnOn()
	m.TurnOff()
}

func TestSquareSwitchBridgeFan(t *testing.T) {
	m := NewSquareSwitch(&Fan{})
	m.TurnOn()
	m.TurnOff()
}
