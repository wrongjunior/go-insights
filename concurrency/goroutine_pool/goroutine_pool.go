// Package goroutine_pool реализует пул горутин для управления конкурентным выполнением задач.
package goroutine_pool

import (
	"sync"
)

// Pool представляет пул горутин для выполнения переданных задач.
type Pool struct {
	// tasks - канал, в который помещаются задачи для выполнения.
	tasks chan func()
	// wg - ожидание завершения всех задач.
	wg sync.WaitGroup
	// workers - количество воркеров в пуле.
	workers int
}

// NewPool создаёт новый пул горутин с заданным числом воркеров и размером буфера канала задач.
// После создания пула воркеры сразу запускаются.
func NewPool(workers int, taskQueueSize int) *Pool {
	p := &Pool{
		tasks:   make(chan func(), taskQueueSize),
		workers: workers,
	}
	p.startWorkers()
	return p
}

// startWorkers запускает фиксированное число воркеров, каждый из которых непрерывно обрабатывает задачи из канала.
func (p *Pool) startWorkers() {
	for i := 0; i < p.workers; i++ {
		go func(workerID int) {
			// Каждый воркер слушает канал tasks до его закрытия.
			for task := range p.tasks {
				// Выполнение задачи.
				task()
				// Сигнализируем о завершении одной задачи.
				p.wg.Done()
			}
		}(i)
	}
}

// Submit добавляет новую задачу в пул. Задача представлена функцией без аргументов.
// Перед добавлением происходит увеличение счётчика ожидания, что гарантирует корректное завершение.
func (p *Pool) Submit(task func()) {
	p.wg.Add(1)
	p.tasks <- task
}

// Shutdown завершает работу пула. После вызова этого метода пул не принимает новые задачи.
// Метод ждёт завершения всех ранее добавленных задач.
func (p *Pool) Shutdown() {
	// Закрываем канал, что приводит к завершению цикла в каждом воркере.
	close(p.tasks)
	// Ожидаем завершения всех задач.
	p.wg.Wait()
}
