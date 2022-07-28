package main

import (
	"8086AssmblySimulator/Actuator"
	"8086AssmblySimulator/Definer"
)

func main() {
	Definer.InitReg()
	Definer.CPU.ShowReg()
	code := "mov ax,20"
	Actuator.Actuate(code)
}
