package main

import (
	"fmt"
	"strings"
)

const example = `Sensor at x=2, y=18: closest beacon is at x=-2, y=15
Sensor at x=9, y=16: closest beacon is at x=10, y=16
Sensor at x=13, y=2: closest beacon is at x=15, y=3
Sensor at x=12, y=14: closest beacon is at x=10, y=16
Sensor at x=10, y=20: closest beacon is at x=10, y=16
Sensor at x=14, y=17: closest beacon is at x=10, y=16
Sensor at x=8, y=7: closest beacon is at x=2, y=10
Sensor at x=2, y=0: closest beacon is at x=2, y=10
Sensor at x=0, y=11: closest beacon is at x=2, y=10
Sensor at x=20, y=14: closest beacon is at x=25, y=17
Sensor at x=17, y=20: closest beacon is at x=21, y=22
Sensor at x=16, y=7: closest beacon is at x=15, y=3
Sensor at x=14, y=3: closest beacon is at x=15, y=3
Sensor at x=20, y=1: closest beacon is at x=15, y=3
`

const input = `Sensor at x=3923513, y=2770279: closest beacon is at x=3866712, y=2438950
Sensor at x=675683, y=3223762: closest beacon is at x=-224297, y=2997209
Sensor at x=129453, y=2652332: closest beacon is at x=92656, y=2629486
Sensor at x=3906125, y=2154618: closest beacon is at x=3866712, y=2438950
Sensor at x=65723, y=902062: closest beacon is at x=92656, y=2629486
Sensor at x=3137156, y=2876347: closest beacon is at x=2907507, y=3100765
Sensor at x=32848, y=2676435: closest beacon is at x=92656, y=2629486
Sensor at x=3272472, y=3445147: closest beacon is at x=2907507, y=3100765
Sensor at x=2926008, y=128948: closest beacon is at x=3089364, y=-501737
Sensor at x=2975, y=2769838: closest beacon is at x=92656, y=2629486
Sensor at x=3540455, y=2469135: closest beacon is at x=3866712, y=2438950
Sensor at x=3674809, y=2062166: closest beacon is at x=3719980, y=2000000
Sensor at x=3693706, y=2027384: closest beacon is at x=3719980, y=2000000
Sensor at x=3869683, y=2291983: closest beacon is at x=3866712, y=2438950
Sensor at x=2666499, y=2796436: closest beacon is at x=2650643, y=2489479
Sensor at x=492, y=2601991: closest beacon is at x=92656, y=2629486
Sensor at x=2710282, y=3892347: closest beacon is at x=2907507, y=3100765
Sensor at x=28974, y=3971342: closest beacon is at x=-224297, y=2997209
Sensor at x=3990214, y=2399722: closest beacon is at x=3866712, y=2438950
Sensor at x=3853352, y=1009020: closest beacon is at x=3719980, y=2000000
Sensor at x=1231833, y=3999338: closest beacon is at x=1313797, y=4674300
Sensor at x=2083669, y=875035: closest beacon is at x=1369276, y=-160751
Sensor at x=1317274, y=2146819: closest beacon is at x=2650643, y=2489479
Sensor at x=3712875, y=2018770: closest beacon is at x=3719980, y=2000000
Sensor at x=963055, y=23644: closest beacon is at x=1369276, y=-160751
Sensor at x=3671967, y=64054: closest beacon is at x=3089364, y=-501737
Sensor at x=3109065, y=2222392: closest beacon is at x=2650643, y=2489479
Sensor at x=3218890, y=1517419: closest beacon is at x=3719980, y=2000000
Sensor at x=3856777, y=3987650: closest beacon is at x=4166706, y=3171774
Sensor at x=1912696, y=3392788: closest beacon is at x=2907507, y=3100765
Sensor at x=3597620, y=3100104: closest beacon is at x=4166706, y=3171774
`

type Point struct {
	X, Y int
}

type SBPair struct {
	Sensor Point
	Beacon Point
	MDist  int
}

func ManhattanDistance(a, b Point) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

type Pairs []SBPair

type Range struct {
	Start int
	End   int
}

type RangeToInt map[Range]int

func (ri RangeToInt) findInt(p int) int {
	for r, i := range ri {
		if p >= r.Start && p <= r.End {
			return i
		}
	}
	return -1
}

func (p Pairs) DistressCoords(maxC int) (int, int) {
	for _, pair := range p {
		dist1 := ManhattanDistance(pair.Sensor, pair.Beacon) + 1
		for j := 0; j < dist1; j++ {
			pt2r := Point{
				X: pair.Sensor.X + j,
				Y: pair.Sensor.Y - dist1 + j,
			}
			pr2b := Point{
				X: pair.Sensor.X + dist1 - j,
				Y: pair.Sensor.Y + j,
			}
			pb2l := Point{
				X: pair.Sensor.X - j,
				Y: pair.Sensor.Y + dist1 - j,
			}
			pl2t := Point{
				X: pair.Sensor.X - dist1 + j,
				Y: pair.Sensor.Y - j,
			}
			if !p.AnyCovers(pt2r, maxC) {
				return pt2r.X, pt2r.Y
			}
			if !p.AnyCovers(pr2b, maxC) {
				return pr2b.X, pr2b.Y
			}
			if !p.AnyCovers(pb2l, maxC) {
				return pb2l.X, pb2l.Y
			}
			if !p.AnyCovers(pl2t, maxC) {
				return pl2t.X, pl2t.Y
			}
		}
	}
	return 0, 0
}

func (p Pairs) AnyCovers(pt Point, maxC int) bool {
	if pt.X < 0 || pt.Y < 0 || pt.X > maxC || pt.Y > maxC {
		return true // out of bounds
	}
	for _, pair := range p {
		if ManhattanDistance(pair.Sensor, pt) <= pair.MDist {
			return true
		}
	}
	return false
}

func Freq(x, y int) int {
	return x*4_000_000 + y
}

func main() {
	area := parseInput(example)
	x, y := area.DistressCoords(20)
	f := Freq(x, y)
	fmt.Printf("Example point: %d x %d - freq.: %d\n", x, y, f)

	area = parseInput(input)
	x, y = area.DistressCoords(4_000_000)
	f = Freq(x, y)
	fmt.Printf("Input point: %d x %d - freq.: %d\n", x, y, f)
}

func parseInput(input string) Pairs {
	lines := strings.Split(input, "\n")
	var sx, sy, bx, by int
	pairs := make([]SBPair, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		s := Point{sx, sy}
		b := Point{bx, by}
		pairs = append(pairs, SBPair{
			Sensor: s,
			Beacon: b,
			MDist:  ManhattanDistance(s, b),
		})

	}
	return pairs
}

func minMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}
