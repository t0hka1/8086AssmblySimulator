package Definer

// 初步设计的指令需要(更多的后面想到再加吧)
// 数据传送指令 mov/movsx/push/pop
// 算术运算指令 add/adc/inc/sub/dec/mul/div
// 逻辑运算指令 and/or/not/xor
// 移位操作指令 sal/sar/shl/shr/rol/ror
// 麻烦的loop  栈的相关实现

//func mov() {
//
//}

// mov指令
//eg:
//mov ax,18
//mov ah,78
//mov ax,bx

var (
	Src SrcType
	Dst DstType
)

type code uint16

type SrcType struct {
	IMM  uint16
	Reg  *Register
	Reg1 *Reg
	Reg2 *uint16
	MM   uint32
}

type DstType struct {
	IMM  uint16
	Reg  *Register
	Reg1 *Reg
	Reg2 *uint16
	MM   uint32
}

// 定义操作数的类型
const (
	Mov_IMM2Reg code = 0 //立即数到寄存器
	Mov_MM2Reg  code = 1 //内存到寄存器
	Mov_Reg2Reg code = 2 //寄存器到寄存器
	Mov_Reg2MM  code = 3 //寄存器到内存
	Mov_IMM2MM  code = 4 //立即数到内存
	Mov_MM2MM   code = 5 //内存到内存
)

func Mov(src *SrcType, dst *DstType, code code) {
	switch code {
	case Mov_IMM2Reg:
		//CR1(dst.IMM, src.Reg1)
		//CR2(dst.IMM, src.Reg2)
	case Mov_MM2Reg:
	case Mov_Reg2Reg:

	case Mov_Reg2MM:
	case Mov_IMM2MM:
	case Mov_MM2MM:
	}

}
