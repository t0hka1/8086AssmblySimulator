package main

import (
	"8086AssmblySimulator/Definer"
)

var CPU Definer.Register

func main() {
	Definer.InitReg()
	CPU.ShowReg()
	//src.MM[0] = 0x1
	//src.MM[1] = 0x2
	//src.ReadMemoryWord(0)
}
