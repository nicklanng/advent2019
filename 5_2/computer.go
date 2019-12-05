package main

import (
	"errors"
	"fmt"
	"runtime/debug"
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
	BRK = 99
)

var (
	ErrInvalidInput = errors.New("invalid input")
	ErrInvalidOpCode = errors.New("invalid opcode")
)

type Computer struct {
	memory [1024]int
}

func (c *Computer) Load(program []int) {
	copy(c.memory[:], program)
}

func (c *Computer) Run() (err error) {
	defer func() {
		if err != nil {
			debug.PrintStack()
			fmt.Println(c.memory)
		}
		if r := recover(); r != nil {
			fmt.Println(r)
			debug.PrintStack()
			fmt.Println(c.memory)
		}
	}()

	var i int
	var opLength int
	for {
		opCode := c.memory[i] % 100
		modes := c.memory[i] / 100

		switch opCode {
		case ADD:
			opLength = 4

			operand1 := c.ReadOperand(i+1, c.isBitSet(modes, 0))
			operand2 := c.ReadOperand(i+2, c.isBitSet(modes, 1))

			c.WritePosition(i+3, operand1 + operand2)

			i += opLength
			continue

		case MUL:
			opLength = 4

			operand1 := c.ReadOperand(i+1, c.isBitSet(modes, 0))
			operand2 := c.ReadOperand(i+2, c.isBitSet(modes, 1))

			c.WritePosition(i+3, operand1 * operand2)

			i += opLength
			continue

		case RIN:
			opLength = 2

			fmt.Print("> ")

			var val int
			if _, err := fmt.Scanf("%d", &val); err != nil {
				return ErrInvalidInput
			}

			c.WritePosition(i+1,val)

			i += opLength
			continue

		case PRN:
			opLength = 2

			fmt.Println(c.ReadPosition(i + 1))

			i += opLength
			continue

		case JIT:
			opLength = 3

			operand1 := c.ReadOperand(i+1, c.isBitSet(modes, 0))

			if operand1 != 0 {
				i = c.ReadOperand(i+2, c.isBitSet(modes, 1))
				continue
			}

			i += opLength
			continue

		case JIF:
			opLength = 3

			operand1 := c.ReadOperand(i+1, c.isBitSet(modes, 0))

			if operand1 == 0 {
				i = c.ReadOperand(i+2, c.isBitSet(modes, 1))
				continue
			}

			i += opLength
			continue

		case LES:
			opLength = 4

			operand1 := c.ReadOperand(i+1, c.isBitSet(modes, 0))
			operand2 := c.ReadOperand(i+2, c.isBitSet(modes, 1))

			if operand1 < operand2 {
				c.WritePosition(i+3, 1)
			} else {
				c.WritePosition(i+3, 0)
			}

			i += opLength
			continue

		case EQL:
			opLength = 4

			operand1 := c.ReadOperand(i+1, c.isBitSet(modes, 0))
			operand2 := c.ReadOperand(i+2, c.isBitSet(modes, 1))

			if operand1 == operand2 {
				c.WritePosition(i+3, 1)
			} else {
				c.WritePosition(i+3, 0)
			}

			i += opLength
			continue

		case BRK:
			return nil

		default:
			return ErrInvalidOpCode
		}
	}
}

func (c *Computer) ReadOperand(i int, immediate bool) int {
	if immediate{
		return c.ReadImmediate(i)
	}

	return c.ReadPosition(i)
}

func (c *Computer) ReadImmediate(i int) int {
	return c.memory[i]
}

func (c *Computer) ReadPosition(i int) int {
	return c.ReadImmediate(c.memory[i])
}

func (c *Computer) WritePosition(pos int, val int) {
	ptr := c.memory[pos]
	c.memory[ptr] = val
}

func (c Computer) isBitSet(modes int, bit byte) bool {
	return modes&(1<<bit) == 1<<bit
}
