package Actuator

import (
	"8086AssmblySimulator/Classifier"
	"8086AssmblySimulator/Definer"
	"8086AssmblySimulator/Parser"
)

func Actuate(code string) {
	Parser.Parse(code)
	println(code)
	Classifier.Classify(&Parser.OP)
	Definer.CPU.ShowReg()
}