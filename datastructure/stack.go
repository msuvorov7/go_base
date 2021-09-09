package main

import (
	"fmt"
)

var (
	brakets = map[rune]rune{
		')': '(',
		'}': '{',
	}
)

type Stack struct {
	arr []interface{}
}

func (s *Stack) push(item interface{}) {
	s.arr = append(s.arr, item)
}

func (s *Stack) pop() interface{} {
	var result interface{}
	if len(s.arr) > 0 {
		result = s.arr[len(s.arr)-1]
		s.arr = s.arr[:len(s.arr)-1]
		return result
	}
	return nil

}

func (s *Stack) top() interface{} {
	if len(s.arr) > 0 {
		return s.arr[len(s.arr)-1]
	}
	return nil
}

func (s *Stack) sort() Stack {
	/*
	Снимаем с верхушки исходного стека элемент x.
	Находим позицию, в которую его можно добавить во временный стек, такую, чтобы не нарушался инвариант.
	Эта позиция - либо единственно возможная, если временный стек пуст, либо сразу после элемента, 
	который больше x, ибо тогда ниже в стеке окажутся элементы меньшие или равные (напоминаю, временный 
	стек всегда отсортирован). Для того, чтобы x поместить в эту позицию, мы последовательно вынимаем из
	стека все, что нам мешает, и кладем в исходный стек (позже мы все равно это отсортируем).
	Помещаем x во временный стек. Временный стек остается отсортирован после этого шага.
	Повторяем до тех пор, пока в исходном стеке не осталось элементов.
	*/
	helper := Stack{
		arr: make([]interface{}, 0),
	}
	for s.top() != nil {
		val := s.pop()
		for (helper.top() != nil) && (helper.top().(int) > val.(int)) {
			s.push(helper.pop())
		}

		helper.push(val)
	}

	return helper
}

func checkBrakets(line string) bool {
	/*
	Теперь пройдитесь по строке выражения exp.
	Если текущий символ является стартовой скобкой ( '(' или '{' или '[' )), поместите его в стек.
	Если текущий символ является закрывающей скобкой ( ')' или '}' или ']' ), тогда извлекается из стека, 
	и если извлеченный символ является совпадающей стартовой скобкой, то продолжаем. В противном случае круглые скобки не сбалансированы.
	После полного обхода, если в стеке остается какая-то начальная скобка, то «не сбалансировано»
	*/
	s := Stack{
		arr: make([]interface{}, 0),
	}
	for _, i := range line {
	
		if (i == '(') || (i == '{') {
			s.push(i)
			continue
		} 
		open, ok := brakets[i]
		
		if ok {
			
			if open == s.top() {
				s.pop()
			} else {
				return false
			}
		}
	}
	if s.top() != nil {
		return false
	}
	
	return true
}

func main() {
	/*
	s := Stack{
		arr: make([]interface{}, 0),
	}
	s.push(2)
	s.push(3)
	s.push(1)
	s.push(6)
	s = s.sort()
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
	*/
	var testLine = "{1+(()3)}()k"
	//fmt.Printf("%c\n", brakets['}'])
	fmt.Println(checkBrakets(testLine))
}
