package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	ADD = 1
	MUL = 2
	RIN = 3
	PRN = 4
	JIT = 5
	JIF = 6
	LES = 7
	EQL = 8
	REL = 9
	BRK = 99
)

var (
	ErrInvalidInput  = errors.New("invalid Input")
	ErrInvalidOpCode = errors.New("invalid opcode")
)

type Computer struct {
	Input        chan string
	Output       chan string
	memory       [65536]int
	relativeBase int
}

func (c *Computer) Load(program []int) {
	copy(c.memory[:], program)
	c.Input = make(chan string)
	c.Output = make(chan string)
}

func (c *Computer) Run() (err error) {
	defer func() {
		close(c.Input)
		close(c.Output)
	}()

	var cursor int
	for {
		opCode := c.memory[cursor] % 100
		modes := Reverse(PadMemory(fmt.Sprintf("%d", c.memory[cursor]/100), 4))

		switch opCode {
		case ADD:
			operand1 := c.ReadOperand(cursor+1, modes[0])
			operand2 := c.ReadOperand(cursor+2, modes[1])

			c.WriteValue(cursor+3, modes[2], operand1+operand2)

			cursor += 4
			continue

		case MUL:
			operand1 := c.ReadOperand(cursor+1, modes[0])
			operand2 := c.ReadOperand(cursor+2, modes[1])

			c.WriteValue(cursor+3, modes[2], operand1*operand2)

			cursor += 4
			continue

		case RIN:
			val, _ := strconv.Atoi(<-c.Input)

			c.WriteValue(cursor+1, modes[0], val)

			cursor += 2
			continue

		case PRN:
			operand1 := c.ReadOperand(cursor+1, modes[0])
			c.Output <- fmt.Sprintf("%d", operand1)

			cursor += 2
			continue

		case JIT:
			operand1 := c.ReadOperand(cursor+1, modes[0])

			if operand1 != 0 {
				cursor = c.ReadOperand(cursor+2, modes[1])
				continue
			}

			cursor += 3
			continue

		case JIF:
			operand1 := c.ReadOperand(cursor+1, modes[0])

			if operand1 == 0 {
				cursor = c.ReadOperand(cursor+2, modes[1])
				continue
			}

			cursor += 3
			continue

		case LES:
			operand1 := c.ReadOperand(cursor+1, modes[0])
			operand2 := c.ReadOperand(cursor+2, modes[1])

			if operand1 < operand2 {
				c.WriteValue(cursor+3, modes[2], 1)
			} else {
				c.WriteValue(cursor+3, modes[2], 0)
			}

			cursor += 4
			continue

		case EQL:
			operand1 := c.ReadOperand(cursor+1, modes[0])
			operand2 := c.ReadOperand(cursor+2, modes[1])

			if operand1 == operand2 {
				c.WriteValue(cursor+3, modes[2], 1)
			} else {
				c.WriteValue(cursor+3, modes[2], 0)
			}

			cursor += 4
			continue

		case REL:
			operand1 := c.ReadOperand(cursor+1, modes[0])

			c.relativeBase += operand1

			cursor += 2
			continue

		case BRK:
			fmt.Println("EXIT")
			return nil

		default:
			return ErrInvalidOpCode
		}
	}

}

func (c *Computer) ReadOperand(i int, mode byte) int {
	switch mode {
	case '0':
		return c.ReadPosition(i)
	case '1':
		return c.ReadImmediate(i)
	case '2':
		return c.ReadRelative(i)
	}
	panic("unknown operand mode")
}

func (c *Computer) ReadImmediate(i int) int {
	return c.memory[i]
}

func (c *Computer) ReadPosition(i int) int {
	return c.ReadImmediate(c.memory[i])
}

func (c *Computer) ReadRelative(i int) int {
	return c.ReadImmediate(c.relativeBase + c.memory[i])
}

func (c *Computer) WriteValue(i int, mode byte, val int) {
	switch mode {
	case '0':
		c.WritePosition(i, val)
		return
	case '1':
		c.WriteImmediate(i, val)
		return
	case '2':
		c.WriteRelative(i, val)
		return
	}
	panic("unknown operand mode " + string(mode))
}

func (c *Computer) WriteImmediate(i int, val int) {
	c.memory[i] = val
}

func (c *Computer) WritePosition(i int, val int) {
	c.WriteImmediate(c.memory[i], val)
}

func (c *Computer) WriteRelative(i int, val int) {
	c.WriteImmediate(c.relativeBase+c.memory[i], val)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func PadMemory(str string, l int) string {
	if len(str) < l {
		return strings.Repeat("0", l-len(str)) + str
	}
	return str
}
