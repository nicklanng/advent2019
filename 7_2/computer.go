package main

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strconv"
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
	ErrInvalidInput = errors.New("invalid Input")
	ErrInvalidOpCode = errors.New("invalid opcode")
)

type Computer struct {
	Input  chan string
	Output chan string
	memory [1024]int
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

	defer func() {
		if err != nil {
			fmt.Println(c.memory)
		}
		if r := recover(); r != nil {
			fmt.Println(r)
			debug.PrintStack()
			fmt.Println(c.memory)
		}
	}()

	var cursor int
	for {
		opCode := c.memory[cursor] % 100
		modes,_  := strconv.ParseUint(fmt.Sprintf("%d", c.memory[cursor] / 100), 2, 8)

		switch opCode {
		case ADD:
			operand1 := c.ReadOperand(cursor+1, c.isBitSet(modes, 0))
			operand2 := c.ReadOperand(cursor+2, c.isBitSet(modes, 1))

			c.WritePosition(cursor+3, operand1 + operand2)

			cursor += 4
			continue

		case MUL:
			operand1 := c.ReadOperand(cursor+1, c.isBitSet(modes, 0))
			operand2 := c.ReadOperand(cursor+2, c.isBitSet(modes, 1))

			c.WritePosition(cursor+3, operand1 * operand2)

			cursor += 4
			continue

		case RIN:
			val, _ := strconv.Atoi(<-c.Input)

			c.WritePosition(cursor+1,val)

			cursor += 2
			continue

		case PRN:
			c.Output <- fmt.Sprintf("%d", c.ReadPosition(cursor + 1))
			cursor += 2
			continue

		case JIT:
			operand1 := c.ReadOperand(cursor+1, c.isBitSet(modes, 0))

			if operand1 != 0 {
				cursor = c.ReadOperand(cursor+2, c.isBitSet(modes, 1))
				continue
			}

			cursor += 3
			continue

		case JIF:
			operand1 := c.ReadOperand(cursor+1, c.isBitSet(modes, 0))

			if operand1 == 0 {
				cursor = c.ReadOperand(cursor+2, c.isBitSet(modes, 1))
				continue
			}

			cursor += 3
			continue

		case LES:
			operand1 := c.ReadOperand(cursor+1, c.isBitSet(modes, 0))
			operand2 := c.ReadOperand(cursor+2, c.isBitSet(modes, 1))

			if operand1 < operand2 {
				c.WritePosition(cursor+3, 1)
			} else {
				c.WritePosition(cursor+3, 0)
			}

			cursor += 4
			continue

		case EQL:
			operand1 := c.ReadOperand(cursor+1, c.isBitSet(modes, 0))
			operand2 := c.ReadOperand(cursor+2, c.isBitSet(modes, 1))

			if operand1 == operand2 {
				c.WritePosition(cursor+3, 1)
			} else {
				c.WritePosition(cursor+3, 0)
			}

			cursor += 4
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

func (c Computer) isBitSet(modes uint64, bit byte) bool {
	return modes&(1<<bit) == 1<<bit
}
