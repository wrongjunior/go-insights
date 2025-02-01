// Package waitgroup предоставляет простую реализацию примитива ожидания группы горутин.
package waitgroup

import (
	"runtime"
	"sync/atomic"
)

// WaitGroup представляет примитив синхронизации для ожидания завершения группы горутин.
type WaitGroup struct {
	counter int64
}

// Add увеличивает или уменьшает значение счетчика на заданное delta.
// Если после изменения значение становится меньше нуля, происходит паника.
func (wg *WaitGroup) Add(delta int64) {
	newCount := atomic.AddInt64(&wg.counter, delta)
	if newCount < 0 {
		panic("WaitGroup: counter стал отрицательным")
	}
}

// Done уменьшает счетчик на 1. Это удобная обёртка для wg.Add(-1).
func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

// Wait блокирует выполнение до тех пор, пока счетчик не станет равен 0.
func (wg *WaitGroup) Wait() {
	// Активное ожидание с уступкой процессорного времени.
	for atomic.LoadInt64(&wg.counter) > 0 {
		runtime.Gosched()
	}
}
