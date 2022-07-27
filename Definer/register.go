package Definer

import "fmt"

var CPU Register

// 定义寄存器结构体
type reg struct {
	Ah uint8
	Al uint8
}

// si、di 寄存器不可分
type Register struct {
	Ax  reg    //累加寄存器
	Bx  reg    //基地址寄存器
	Cx  reg    //计数器寄存器
	Dx  reg    //数据寄存器
	Si  uint16 //源变址寄存器
	Di  uint16 //源目标寄存器
	Sp  reg    //栈顶寄存器
	Bp  reg    //栈底寄存器
	Ip  reg    //指令指针寄存器
	Cs  reg    //代码段寄存器
	Ss  reg    //栈段寄存器
	Ds  reg    //数据段寄存器
	Es  reg    //附加数据段寄存器
	Psw reg    //标志寄存器
}

func CR1(data uint16, regPointer *reg) {
	regPointer.Al = uint8(data & 0xff)
	regPointer.Ah = uint8(data >> 8 & 0xff)
}

func CR2(data uint16, regPointer *uint16) {
	*regPointer = data
}

func InitReg() {
	CR1(0, &CPU.Ax)
	CR1(0, &CPU.Bx)
	CR1(0, &CPU.Cx)
	CR1(0, &CPU.Dx)
	CR2(0, &CPU.Si)
	CR2(0, &CPU.Di)
	CR1(0, &CPU.Sp)
	CR1(0, &CPU.Bp)
	CR1(0, &CPU.Ip)
	CR1(0, &CPU.Cs)
	CR1(0, &CPU.Ss)
	CR1(0, &CPU.Ds)
	CR1(0, &CPU.Es)
	CR1(0, &CPU.Psw) //暂留异议
}

func calReg(regPointer *reg) uint16 {
	return uint16(regPointer.Ah)<<8 | uint16(regPointer.Al)
}

func (r Register) ShowReg() {
	fmt.Println("===========当前cpu寄存器的状态============")
	fmt.Printf("Ax=%#04x   Bx=%#04x   Cx=%#04x   Dx=%#04x  Sp=%#04x  Bp=%#04x  Si=%#04x  Di=%#04x\nDs=%#04x   Es=%#04x   Ss=%#04x   Cs=%#04x  Ip=%#04x\n",
		calReg(&r.Ax), calReg(&r.Bx), calReg(&r.Cx), calReg(&r.Dx), calReg(&r.Sp), calReg(&r.Bp), r.Si, r.Di, calReg(&r.Ds), calReg(&r.Es), calReg(&r.Ss), calReg(&r.Cs), calReg(&r.Ip))
	fmt.Println("=======================================")
}
