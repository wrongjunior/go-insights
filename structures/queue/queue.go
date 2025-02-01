// Package queue предоставляет реализацию очереди на основе среза.
package queue

// Queue представляет собой очередь, реализованную на основе среза.
// Элементы очереди могут быть любого типа (interface{}).
type Queue struct {
	elements []interface{}
}

// New создаёт и возвращает новую пустую очередь.
func New() *Queue {
	return &Queue{
		elements: make([]interface{}, 0),
	}
}

// Enqueue добавляет элемент в конец очереди.
func (q *Queue) Enqueue(item interface{}) {
	q.elements = append(q.elements, item)
}

// Dequeue удаляет и возвращает первый элемент очереди.
// Если очередь пуста, возвращает (nil, false).
func (q *Queue) Dequeue() (interface{}, bool) {
	if len(q.elements) == 0 {
		return nil, false
	}
	// Извлекаем первый элемент очереди.
	item := q.elements[0]
	// Обновляем срез, исключая извлечённый элемент.
	q.elements = q.elements[1:]
	return item, true
}

// Peek возвращает первый элемент очереди без его удаления.
// Если очередь пуста, возвращает (nil, false).
func (q *Queue) Peek() (interface{}, bool) {
	if len(q.elements) == 0 {
		return nil, false
	}
	return q.elements[0], true
}

// IsEmpty проверяет, пуста ли очередь.
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

// Size возвращает текущее количество элементов в очереди.
func (q *Queue) Size() int {
	return len(q.elements)
}
