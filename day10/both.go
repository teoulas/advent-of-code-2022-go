package main

import (
	"fmt"
	"strconv"
	"strings"
)

const exInput = `addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop`

const input = `addx 1
addx 4
addx 21
addx -20
addx 4
noop
noop
addx 5
addx 3
noop
addx 2
addx 1
noop
noop
addx 4
noop
noop
noop
addx 3
addx 5
addx 2
addx 1
noop
addx -37
addx 22
addx -4
addx -14
addx 2
addx 5
addx 3
addx -2
addx 2
addx 5
addx 2
addx -15
addx 32
addx -14
addx 5
addx 2
addx 3
noop
addx -13
addx -2
addx 18
addx -36
noop
addx 11
addx -7
noop
noop
addx 6
addx 22
addx -21
addx 3
addx 2
addx 4
noop
noop
noop
addx 5
addx -16
addx 17
addx 2
addx 5
addx -11
addx 15
addx -15
addx -24
noop
noop
addx 7
addx 2
addx -6
addx 9
noop
addx 5
noop
addx -3
addx 4
addx 2
noop
noop
addx 7
noop
noop
noop
addx 5
addx -28
addx 29
noop
addx 3
addx -7
addx -29
noop
addx 7
addx -2
addx 2
addx 5
addx 2
addx -3
addx 4
addx 5
addx 2
addx 8
addx -30
addx 25
addx 7
noop
noop
addx 3
addx -2
addx 2
addx -10
addx -24
addx 2
noop
noop
addx 2
noop
addx 3
addx 2
noop
addx 3
addx 2
addx 5
addx 2
noop
addx 1
noop
addx 2
addx 8
noop
noop
addx -1
addx -9
addx 14
noop
addx 1
noop
noop`

var opCycles = map[string]int{
	"addx": 2,
	"noop": 1,
}

type executor struct {
	cmd             string
	arg             int
	regX            int
	cyclesRemaining int
}

func NewExec() *executor {
	return &executor{regX: 1}
}

func (e *executor) Done() bool {
	return e.cyclesRemaining == 0
}

func (e *executor) LoadCmd(line string) {
	var err error
	split := strings.Split(line, " ")
	e.cmd = split[0]
	e.cyclesRemaining = opCycles[e.cmd]
	if len(split) == 1 {
		return
	}
	e.arg, err = strconv.Atoi(split[1])
	if err != nil {
		panic(err)
	}
}

func (e *executor) EndCycle() {
	e.cyclesRemaining--
	if e.cyclesRemaining > 0 {
		return
	}
	switch e.cmd {
	case "addx":
		e.regX += e.arg
	}
}

func main() {
	lines := strings.Split(input, "\n")
	crt := make([][]rune, 6)
	for i := range crt {
		crt[i] = make([]rune, 40)
	}
	sum := 0
	ex := NewExec()
	pc := -1
	for c := 1; c <= 240; c++ {
		if ex.Done() {
			pc++
			ex.LoadCmd(lines[pc])
		}
		if c == 20 || c == 60 || c == 100 || c == 140 || c == 180 || c == 220 {
			sum = sum + ex.regX*c
			// fmt.Printf("Cycle %d: x is %d - signal: %d\n", c, ex.regX, ex.regX*c)
		}
		crtRow := (c - 1) / 40
		crtPos := (c - 1) % 40
		// fmt.Printf("row %d - pos %d\n", crtRow, crtPos)
		draw := ' '
		if ex.regX >= crtPos-1 && ex.regX <= crtPos+1 {
			draw = '\u2588'
		}
		crt[crtRow][crtPos] = draw
		ex.EndCycle()
	}
	fmt.Printf("Sum: %d\n", sum)
	for _, row := range crt {
		fmt.Println(string(row))
	}
}
