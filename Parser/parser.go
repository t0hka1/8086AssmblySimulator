package Parser

import (
	"strings"
)

var (
	OP Op
)

type Op struct {
	Opcode string
	Src    string
	Dst    string
}

func Parse(code string) {
	opcode := strings.Split(code, ` `)[0]
	OP.Opcode = opcode
	//fmt.Printf("opcode:%s\n", opcode)
	dst := strings.Split(strings.Split(code, ` `)[1], `,`)[0]
	src := strings.Split(strings.Split(code, ` `)[1], `,`)[1]
	//fmt.Printf("src:%s\ndst:%s\n", src, dst)
	OP.Src = src
	OP.Dst = dst
}
