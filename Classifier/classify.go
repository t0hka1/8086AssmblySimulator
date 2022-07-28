package Classifier

import (
	"8086AssmblySimulator/Definer"
	"8086AssmblySimulator/Parser"
	"strconv"
)

func Classify(op *Parser.Op) {
	if op.Opcode == "mov" { // mov 操作码
		Definer.Src.Reg = &Definer.CPU
		if op.Dst == "ax" {
			Definer.Dst.Reg1 = &Definer.CPU.Ax
			tmp, _ := strconv.Atoi(op.Src)
			Definer.Src.IMM = uint16(tmp)
			Definer.Mov(&Definer.Src, &Definer.Dst, 0)
		}
	}
}
