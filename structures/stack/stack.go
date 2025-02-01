// Package stack предоставляет реализацию стека на основе среза.
package stack

// Stack представляет собой структуру данных «стек», реализованную на основе среза.
// Элементы стека могут быть любого типа (interface{}).
type Stack struct {
	elements []interface{}
}

// New создаёт и возвращает новый пустой стек.
func New() *Stack {
	return &Stack{
		elements: make([]interface{}, 0),
	}
}

// Push добавляет новый элемент в вершину стека.
func (s *Stack) Push(item interface{}) {
	s.elements = append(s.elements, item)
}

// Pop удаляет и возвращает элемент, находящийся на вершине стека.
// Если стек пуст, возвращает (nil, false).
func (s *Stack) Pop() (interface{}, bool) {
	if len(s.elements) == 0 {
		return nil, false
	}
	// Извлекаем последний элемент.
	index := len(s.elements) - 1
	item := s.elements[index]
	// Очищаем последний элемент и изменяем срез.
	s.elements = s.elements[:index]
	return item, true
}

// Peek возвращает элемент на вершине стека без его удаления.
// Если стек пуст, возвращает (nil, false).
func (s *Stack) Peek() (interface{}, bool) {
	if len(s.elements) == 0 {
		return nil, false
	}
	return s.elements[len(s.elements)-1], true
}

// IsEmpty проверяет, пуст ли стек.
func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Size возвращает текущее количество элементов в стеке.
func (s *Stack) Size() int {
	return len(s.elements)
}
