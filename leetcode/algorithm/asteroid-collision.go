package main

import "math"

type Stack struct {
	data    []int
	current int
}

func (s *Stack) IsEmpty() bool {
	return s.current == 0
}

func (s *Stack) Push(num int) {
	if s.current == len(s.data) {
		s.data = append(s.data, num)
	} else {
		s.data[s.current] = num
		s.current++
	}
}

func (s *Stack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		s.current--
		return s.data[s.current], true
	}
}

func asteroidCollision(asteroids []int) []int {
	neg := []int{}
	stack := Stack{data: make([]int, len(asteroids)), current: 0}
	for _, asteroid := range asteroids {
		if asteroid > 0 && stack.IsEmpty() {
			stack.Push(asteroid)
			continue
		}
		for !stack.IsEmpty() {
			if asteroid > 0 {
				stack.Push(asteroid)
				break
			}

			a, _ := stack.Pop()
			aAbs := math.Abs(float64(a))
			atAbs := math.Abs(float64(asteroid))

			if aAbs == atAbs {
				asteroid = 0
				break
			}
			if aAbs > atAbs {
				asteroid = 0
				stack.Push(a)
				break
			}
		}
		if stack.IsEmpty() && asteroid < 0 {
			neg = append(neg, asteroid)
		}
	}
	return append(neg, stack.data[:stack.current]...)
}

func main() {
	asteroidCollision([]int{-2, -1, 1, 2})
}
