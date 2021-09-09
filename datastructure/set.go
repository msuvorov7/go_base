package main

import (
	"fmt"
)

type Set struct {
	arr []interface{}
}


func (s *Set) add(item interface{}) bool {
	if _, exist := s.contains(item); !exist {
		s.arr = append(s.arr, item)
		return true
	}
	return false
}

func (s *Set) contains(item interface{}) (int, bool) {
	for k, i := range s.arr {
		if i == item {
			return k, true
		}
	}
	return -1, false
}

func (lhs *Set) intersect(rhs *Set) Set {
	result := Set {
		arr: make([]interface{}, 0),
	}
	for _, i := range lhs.arr {
		if _, exist := rhs.contains(i); exist {
			result.add(i)
		}
	}
	return result
}

func (lhs *Set) union(rhs *Set) Set {
	result := Set {
		arr: make([]interface{}, 0),
	}
	for _, i := range lhs.arr {
		result.add(i)
	}
	for _, i := range rhs.arr {
		result.add(i)
	}
	return result
}

func (lhs *Set) difference(rhs *Set) Set {
	result := Set {
		arr: make([]interface{}, 0),
	}
	for _, i := range lhs.arr {
		if _, exist := rhs.contains(i); !exist {
			result.add(i)
		}
	}
	return result
}

func (lhs *Set) symmetricDifference(rhs *Set) Set {
	first := lhs.union(rhs)
	second := lhs.intersect(rhs)
	return first.difference(&second)
}

func (lhs *Set) isSubset(rhs *Set) bool {
	for _, i := range lhs.arr {
		if _, exist := rhs.contains(i); !exist {
			return false
		}
	}
	return true
}

func (s *Set) remove(item interface{}) bool {
	if i, exist := s.contains(item); exist {
		s.arr[i] = s.arr[s.getLen()-1]
		s.arr[s.getLen()-1] = nil
		s.arr = s.arr[:s.getLen()-1]
		return true
	}

	return false
}

func (s *Set) print() {
	for _, i := range s.arr {
		fmt.Print(i, " ")
	}
	fmt.Println()
}

func (s *Set) getLen() int {
	return len(s.arr)
}

func main() {
	lhs := Set {
		arr: make([]interface{}, 0),
	}
	rhs := Set {
		arr: make([]interface{}, 0),
	}
	lhs.add(1)
	lhs.add(2)
	lhs.add(3)
	lhs.add(4)

	rhs.add(3)
	rhs.add(4)
	rhs.add(5)
	rhs.add(6)

	inter := lhs.intersect(&rhs)
	inter.print()

	inter = lhs.union(&rhs)
	inter.print()

	inter = lhs.difference(&rhs)
	inter.print()

	inter = lhs.symmetricDifference(&rhs)
	inter.print()

	fmt.Println(lhs.isSubset(&rhs))
	fmt.Println(lhs.remove(3))
	lhs.print()

}