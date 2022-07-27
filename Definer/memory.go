package Definer

import "fmt"

// cs:ip 最大是 ffff:ffff = 0x10ffef
var (
	MM [0x10ffef]uint8
)

// 内存写
func WriteMemory(data, dst uint16) {
	MM[dst+0] = uint8(data & 0xff)
	MM[dst+1] = uint8(data >> 8 & 0xff)
	return
}

//内存读
func ReadMemoryByte(dst uint16) uint16 {
	var data uint16 = 0
	data = data | uint16(MM[dst+0])
	fmt.Printf("%#04x处的内容为%#02x\n", dst, data)
	return data
}

func ReadMemoryWord(dst uint16) uint16 {
	var data uint16 = 0
	data = uint16(MM[dst+1])<<8 | uint16(MM[dst+0])
	fmt.Printf("%#04x处的内容为%#04x\n", dst, data)
	return data
}
