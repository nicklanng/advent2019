package main

import (
	"errors"
	"fmt"
	"runtime/debug"
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
	BRK = 99
)

var (
	ErrSyntaxError = errors.New("syntax error")
	ErrInvalidInput = errors.New("invalid input")
	ErrInvalidOpCode = errors.New("invalid opcode")
)

type Computer struct {
	memory [1024]string
	currentInstruction string
}

func (c *Computer) Load(program []string) {
	copy(c.memory[:], program)
}

func (c *Computer) Run() (err error) {
	defer func() {
		if err != nil {
			fmt.Println(c.currentInstruction)
			debug.PrintStack()
			fmt.Println(c.memory)
		}
		if r := recover(); r != nil {
			fmt.Println(r)
			fmt.Println(c.currentInstruction)
			debug.PrintStack()
			fmt.Println(c.memory)
		}
	}()

	var i int
	for {
		memLoc := c.PadMemory(c.memory[i], 2)
		opCode, err := strconv.Atoi(memLoc[len(memLoc)-2:])
		if err != nil {
			return ErrSyntaxError
		}

		switch opCode {
		case ADD:
			length := 4
			c.currentInstruction = strings.Join(c.memory[i:i+length],",")

			modes, err := c.parseOperandModes(memLoc)
			if err != nil {
				return err
			}

			operand1, err := c.ReadOperand(i+1, c.isBitSet(modes, 0))
			if err != nil {
				return err
			}
			operand2, err := c.ReadOperand(i+2, c.isBitSet(modes, 1))
			if err != nil {
				return err
			}

			if err := c.WritePosition(i+3, fmt.Sprintf("%d", operand1+operand2)); err != nil {
				return ErrSyntaxError
			}

			i += length
			continue

		case MUL:
			length := 4
			c.currentInstruction = strings.Join(c.memory[i:i+length], ",")

			modes, err := c.parseOperandModes(memLoc)
			if err != nil {
				return err
			}

			operand1, err := c.ReadOperand(i+1, c.isBitSet(modes, 0))
			if err != nil {
				return err
			}
			operand2, err := c.ReadOperand(i+2, c.isBitSet(modes, 1))
			if err != nil {
				return err
			}

			if err := c.WritePosition(i+3, fmt.Sprintf("%d", operand1*operand2)); err != nil {
				return ErrSyntaxError
			}

			i += length
			continue

		case RIN:
			length := 2
			c.currentInstruction = strings.Join(c.memory[i:i+length], ",")

			fmt.Print("> ")

			var val int
			if _, err := fmt.Scanf("%d", &val); err != nil {
				return ErrInvalidInput
			}

			if err := c.WritePosition(i+1, fmt.Sprintf("%d", val)); err != nil {
				return ErrSyntaxError
			}

			i += length
			continue

		case PRN:
			length := 2
			c.currentInstruction = strings.Join(c.memory[i:i+length], ",")

			val, err := c.ReadPosition(i + 1)
			if err != nil {
				return err
			}
			fmt.Println(val)

			i += length
			continue

		case JIT:
			length := 3
			c.currentInstruction = strings.Join(c.memory[i:i+length], ",")

			modes, err := c.parseOperandModes(memLoc)
			if err != nil {
				return err
			}

			comp, err := c.ReadOperand(i+1, c.isBitSet(modes, 0))
			if err != nil {
				return err
			}

			if comp != 0 {
				ptr, err := c.ReadOperand(i+2, c.isBitSet(modes, 1))
				if err != nil {
					return err
				}
				i = ptr
				continue
			}

			i += length
			continue

		case JIF:
			length := 3
			c.currentInstruction = strings.Join(c.memory[i:i+length], ",")

			modes, err := c.parseOperandModes(memLoc)
			if err != nil {
				return err
			}

			comp, err := c.ReadOperand(i+1, c.isBitSet(modes, 0))
			if err != nil {
				return err
			}

			if comp == 0 {
				ptr, err := c.ReadOperand(i+2, c.isBitSet(modes, 1))
				if err != nil {
					return err
				}
				i = ptr
				continue
			}

			i += length
			continue

		case LES:
			length := 4
			c.currentInstruction = strings.Join(c.memory[i:i+length], ",")

			modes, err := c.parseOperandModes(memLoc)
			if err != nil {
				return err
			}

			operand1, err := c.ReadOperand(i+1, c.isBitSet(modes, 0))
			if err != nil {
				return err
			}
			operand2, err := c.ReadOperand(i+2, c.isBitSet(modes, 1))
			if err != nil {
				return err
			}

			if operand1 < operand2 {
				if err := c.WritePosition(i+3, "1"); err != nil {
					return ErrSyntaxError
				}
			} else {
				if err := c.WritePosition(i+3, "0"); err != nil {
					return ErrSyntaxError
				}
			}

			i += length
			continue

		case EQL:
			length := 4
			c.currentInstruction = strings.Join(c.memory[i:i+length], ",")

			modes, err := c.parseOperandModes(memLoc)
			if err != nil {
				return err
			}

			operand1, err := c.ReadOperand(i+1, c.isBitSet(modes, 0))
			if err != nil {
				return err
			}
			operand2, err := c.ReadOperand(i+2, c.isBitSet(modes, 1))
			if err != nil {
				return err
			}

			if operand1 == operand2 {
				if err := c.WritePosition(i+3, "1"); err != nil {
					return ErrSyntaxError
				}
			} else {
				if err := c.WritePosition(i+3, "0"); err != nil {
					return ErrSyntaxError
				}
			}

			i += length
			continue

		case BRK:
			return nil

		default:
			return ErrInvalidOpCode
		}
	}
}

func (c *Computer) PadMemory(str string, l int) string {
	if len(str) < l {
		return strings.Repeat("0", l - len(str))+str
	}
	return str
}

func (c *Computer) ReadOperand(i int, immediate bool) (int, error) {
	var (
		val int
		err error
	)

	if immediate{
		val, err = c.ReadImmediate(i)
	} else {
		val, err = c.ReadPosition(i)
	}

	if err != nil {
		return 0, ErrSyntaxError
	}
	return val, nil

}

func (c *Computer) ReadImmediate(i int) (int, error) {
	return strconv.Atoi(c.memory[i])
}

func (c *Computer) ReadPosition(i int) (int, error) {
	ptr, err := strconv.Atoi(c.memory[i])
	if err != nil {
		return 0, err
	}
	return c.ReadImmediate(ptr)
}

func (c *Computer) WritePosition(pos int, val string) error {
	ptr, err := strconv.Atoi(c.memory[pos])
	if err != nil {
		return err
	}
	c.memory[ptr] = val
	return nil
}

func (c *Computer) parseOperandModes(memLoc string) (byte, error) {
	subString := memLoc[:len(memLoc) - 2]
	if len(subString) == 0 {
		return 0, nil
	}

	m, err := strconv.ParseInt(subString, 2, 8)
	if err != nil {
		return 0, ErrSyntaxError
	}
	modes := byte(m)
	return modes, nil
}

func (c Computer) isBitSet(modes byte, bit byte) bool {
	return modes&(1<<bit) == 1<<bit
}
