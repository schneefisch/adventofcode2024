package _17

import (
	"adventofcode2024/challenges/util"
	"log"
	"math"
	"strconv"
	"strings"
)

type Computer struct {
	A, B, C            int // registers
	programm           []int
	instructionPointer int
	output             []string
}

func (c *Computer) operation(opcode, operator int) {
	switch opcode {
	case 0:
		c.adv(operator)
	case 1:
		c.bxl(operator)
	case 2:
		c.bst(operator)
	case 3:
		c.jnz(operator)
	case 4:
		c.bxc(operator)
	case 5:
		c.out(operator)
	case 6:
		c.bdv(operator)
	case 7:
		c.cdv(operator)
	default:
		log.Fatalf("invalid opcode: %v", opcode)
	}
}

// adv divides the value in register A by 2 to the power of the operator
func (c *Computer) adv(comboOperator int) {
	c.A = c.A / int(math.Pow(float64(2), float64(c.comboOperator(comboOperator))))
	c.step()
}

// bxl calculates the bitwise XOR of the value in register B and the operator
// stores the result back in register B
func (c *Computer) bxl(operator int) {
	c.B ^= operator
	c.step()
}

// bst calculates the value of it's combo-operand modulo 8 (keeping ony it's lowest 3 bits)
// the result is written into the B register
func (c *Computer) bst(comboOperator int) {
	operator := c.comboOperator(comboOperator)
	c.B = operator % 8
	c.step()
}

// jnz jumps by setting the instruction-pointer to the value of its literal operand
// only if the A register is not 0
func (c *Computer) jnz(operand int) {
	if c.A == 0 {
		c.step()
		return
	}
	c.instructionPointer = operand
}

// bxc calculates bitwise XOR of register B and C
// stores the result in register B
// the input is ignored
func (c *Computer) bxc(_ int) {
	c.B ^= c.C
	c.step()
}

// out appends the result of the combo-operator to the output (modulo 8)
func (c *Computer) out(comboOperator int) {
	op := c.comboOperator(comboOperator)
	output := strconv.Itoa(op % 8)
	c.output = append(c.output, output)
	c.step()
}

// bdv like adv, but the result is written into the B register (numerator still read from A)
func (c *Computer) bdv(comboOperator int) {
	c.B = c.A / int(math.Pow(float64(2), float64(c.comboOperator(comboOperator))))
	c.step()
}

// cdv like adv, but the result is written into the C register (numerator still read from A)
func (c *Computer) cdv(comboOperator int) {
	c.C = c.A / int(math.Pow(float64(2), float64(c.comboOperator(comboOperator))))
	c.step()
}

func (c *Computer) comboOperator(comboOperator int) int {
	switch comboOperator {
	case 0, 1, 2, 3:
		return comboOperator
	case 4:
		return c.A
	case 5:
		return c.B
	case 6:
		return c.C
	default:
		log.Fatalf("invalid combo operator: %v", comboOperator)
		return -1
	}
}

// step increments the instruction-pointer by 2
func (c *Computer) step() {
	c.instructionPointer += 2
}

func (c *Computer) run() string {
	for c.instructionPointer < len(c.programm) {
		opcode := c.programm[c.instructionPointer]
		operator := c.programm[c.instructionPointer+1]
		c.operation(opcode, operator)
	}
	// print output
	return strings.Join(c.output, ",")
}

func ChronospatialComputer(filename string) (string, int, error) {
	input, err := util.ReadLines(filename)
	if err != nil {
		return "", 0, err
	}
	c := &Computer{}
	parse(input, c)
	result := c.run()

	return result, 0, nil
}

func parse(input []string, c *Computer) {
	if len(input) < 5 || len(input) > 6 {
		log.Fatalf("invalid input: %v", input)
	}
	c.A = parseRegister(input[0])
	c.B = parseRegister(input[1])
	c.C = parseRegister(input[2])
	c.programm = parseProgramm(input[4])
	c.instructionPointer = 0
	c.output = make([]string, 0)
}

func parseProgramm(s string) []int {
	trimmed := strings.TrimPrefix(s, "Program: ")
	parts := strings.Split(trimmed, ",")
	programm := make([]int, len(parts))
	for i, part := range parts {
		value, err := strconv.Atoi(part)
		if err != nil {
			log.Fatalf("invalid programm value: %v", part)
		}
		programm[i] = value
	}
	return programm
}

func parseRegister(s string) int {
	parts := strings.Split(s, ": ")
	if len(parts) != 2 {
		log.Fatalf("invalid register: %v", s)
	}
	value, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatalf("invalid register value: %v", s)
	}
	return value
}
