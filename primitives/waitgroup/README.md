# Примитив синхронизации: WaitGroup

## Что такое WaitGroup?

WaitGroup — это примитив синхронизации, который позволяет основной горутине ждать завершения набора запущенных горутин. Основной идеей является наличие счетчика, который увеличивается при запуске горутины (через метод `Add`) и уменьшается при её завершении (через метод `Done`). Метод `Wait` блокирует выполнение до тех пор, пока счетчик не достигнет нуля.

## Как работает реализация на Go?

В реализации WaitGroup:
- **Счетчик:** Используется атомарная переменная `counter` типа `int64` для отслеживания количества активных задач.
- **Методы Add/Done:** Метод `Add(delta int64)` атомарно изменяет значение счетчика. Метод `Done()` уменьшает счетчик на единицу, вызывая `Add(-1)`.
- **Ожидание:** Метод `Wait()` периодически проверяет значение счетчика с помощью `atomic.LoadInt64`. Если значение больше нуля, происходит активное ожидание с уступкой процессорного времени (`runtime.Gosched()`).

В данной реализации ожидание реализовано через «busy-wait», что является упрощённым решением для демонстрационных целей. В стандартной библиотеке Go для WaitGroup используется более эффективная блокировка, основанная на сигналах и системных вызовах, позволяющая избежать активного ожидания.

## Пример использования

```go
package main

import (
	"fmt"
	"runtime"
	"time"

	"primitives/waitgroup"
)

func main() {
	var wg waitgroup.WaitGroup
	const numGoroutines = 10

	// Устанавливаем счетчик равным количеству запускаемых горутин.
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(i int) {
			// Симуляция работы
			time.Sleep(10 * time.Millisecond)
			fmt.Printf("Горутина %d завершена\n", i)
			wg.Done() // Уменьшаем счетчик
		}(i)
	}

	// Ожидаем, пока все горутины не завершатся.
	wg.Wait()
	fmt.Println("Все горутины завершены")
	
	// Дополнительный вызов Gosched для корректного завершения, если имеются ожидающие задачи.
	runtime.Gosched()
}
