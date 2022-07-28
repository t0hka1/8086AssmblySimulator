package main

import (
	"8086AssmblySimulator/Definer"
	"8086AssmblySimulator/Parser"
)

var CPU Definer.Register

func main() {
	Definer.InitReg()
	CPU.ShowReg()
	//src.MM[0] = 0x1
	//src.MM[1] = 0x2
	//src.ReadMemoryWord(0)
	//  可以将mov ax,18 解析成 下面16-19 行可以封装起来 放到解析器去做
	//Definer.Src.Reg = &CPU
	//Definer.Src.Reg1 = &CPU.Ax
	//Definer.Src.Reg2 = &CPU.Si
	//Definer.Dst.IMM = 0xff11
	//Definer.Mov(&Definer.Src, &Definer.Dst, 0)
	//CPU.ShowReg()
	code := "mov ax,18"
	Parser.Parse(code)
}
