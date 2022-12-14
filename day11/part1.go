package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const exInput = `Monkey 0:
Starting items: 79, 98
Operation: new = old * 19
Test: divisible by 23
  If true: throw to monkey 2
  If false: throw to monkey 3

Monkey 1:
Starting items: 54, 65, 75, 74
Operation: new = old + 6
Test: divisible by 19
  If true: throw to monkey 2
  If false: throw to monkey 0

Monkey 2:
Starting items: 79, 60, 97
Operation: new = old * old
Test: divisible by 13
  If true: throw to monkey 1
  If false: throw to monkey 3

Monkey 3:
Starting items: 74
Operation: new = old + 3
Test: divisible by 17
  If true: throw to monkey 0
  If false: throw to monkey 1`

const input = `Monkey 0:
Starting items: 56, 52, 58, 96, 70, 75, 72
Operation: new = old * 17
Test: divisible by 11
  If true: throw to monkey 2
  If false: throw to monkey 3

Monkey 1:
Starting items: 75, 58, 86, 80, 55, 81
Operation: new = old + 7
Test: divisible by 3
  If true: throw to monkey 6
  If false: throw to monkey 5

Monkey 2:
Starting items: 73, 68, 73, 90
Operation: new = old * old
Test: divisible by 5
  If true: throw to monkey 1
  If false: throw to monkey 7

Monkey 3:
Starting items: 72, 89, 55, 51, 59
Operation: new = old + 1
Test: divisible by 7
  If true: throw to monkey 2
  If false: throw to monkey 7

Monkey 4:
Starting items: 76, 76, 91
Operation: new = old * 3
Test: divisible by 19
  If true: throw to monkey 0
  If false: throw to monkey 3

Monkey 5:
Starting items: 88
Operation: new = old + 4
Test: divisible by 2
  If true: throw to monkey 6
  If false: throw to monkey 4

Monkey 6:
Starting items: 64, 63, 56, 50, 77, 55, 55, 86
Operation: new = old + 8
Test: divisible by 13
  If true: throw to monkey 4
  If false: throw to monkey 0

Monkey 7:
Starting items: 79, 58
Operation: new = old + 6
Test: divisible by 17
  If true: throw to monkey 1
  If false: throw to monkey 5`

type Monkey struct {
	ID          int
	Items       []int
	Operation   func(int) int
	Test        func(int) bool
	ThrowTo     map[bool]int
	NumInspects int
}

func NewMonkey() Monkey {
	return Monkey{
		ThrowTo: make(map[bool]int),
	}
}

func (m *Monkey) Inspect() (int, int) {
	if len(m.Items) == 0 {
		return 0, -1
	}
	m.NumInspects++
	item := m.Items[0]
	m.Items = m.Items[1:]
	fmt.Printf("  Inspects item with worry lvl %d\n", item)
	item = m.Operation(item)
	fmt.Printf("    After operation: %d\n", item)
	item = item / 3
	fmt.Printf("    After dividing by 3: %d\n", item)
	tested := m.Test(item)
	newMonkey := m.ThrowTo[tested]
	fmt.Printf("    Tested: %t - item thrown to monkey %d\n", tested, newMonkey)
	return item, newMonkey
}

func main() {
	monkeys := parseInput(input)
	numMonkeys := len(monkeys)
	for r := 1; r <= 20; r++ {
		fmt.Printf("== Round %d ==\n", r)
		for i := 0; i < numMonkeys; i++ {
			fmt.Printf("Monkey %d\n", i)
			for len(monkeys[i].Items) > 0 {
				item, nextID := monkeys[i].Inspect()
				if nextID == -1 {
					continue
				}
				nextM := monkeys[nextID]
				nextM.Items = append(nextM.Items, item)
				monkeys[nextID] = nextM
			}
		}
	}
	for i, m := range monkeys {
		fmt.Printf("Monkey %d inspected %d items\n", i, m.NumInspects)
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].NumInspects > monkeys[j].NumInspects
	})
	fmt.Printf("Monkey business: %d x %d = %d\n",
		monkeys[0].NumInspects,
		monkeys[1].NumInspects,
		monkeys[0].NumInspects*monkeys[1].NumInspects)
}

func parseInput(input string) []Monkey {
	lines := strings.Split(input, "\n")
	monkeys := make([]Monkey, 0)

	m := NewMonkey()
	for _, line := range lines {
		if line == "" {
			monkeys = append(monkeys, m)
			m = NewMonkey()
			continue
		}
		if strings.HasPrefix(line, "Monkey") {
			fmt.Sscanf(line, "Monkey %d:", &m.ID)
			continue
		}
		if strings.HasPrefix(line, "Starting items") {
			matches := regexp.MustCompile("[0-9]+").FindAllString(line, -1)
			for _, match := range matches {
				n, err := strconv.Atoi(match)
				if err != nil {
					panic(err)
				}
				m.Items = append(m.Items, n)
			}
			continue
		}
		if strings.HasPrefix(line, "Operation") {
			var operator string
			var operand int
			fmt.Sscanf(line, "Operation: new = old %s %d", &operator, &operand)
			m.Operation = func(old int) int {
				if operand == 0 { // old * old
					return old * old
				}
				switch operator {
				case "+":
					return old + operand
				case "*":
					return old * operand
				}
				return 0
			}
		}
		if strings.HasPrefix(line, "Test:") {
			var operand int
			fmt.Sscanf(line, "Test: divisible by %d", &operand)
			m.Test = func(n int) bool {
				return n%operand == 0
			}
		}
		if strings.Contains(line, "throw to monkey") {
			var b bool
			var id int
			fmt.Sscanf(strings.TrimSpace(line), "If %t: throw to monkey %d", &b, &id)
			m.ThrowTo[b] = id
		}
	}
	if monkeys[len(monkeys)-1].ID != m.ID {
		monkeys = append(monkeys, m)
	}

	return monkeys
}
