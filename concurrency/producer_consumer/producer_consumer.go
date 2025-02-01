// Package main демонстрирует реализацию паттерна продюсер-потребитель.
// В этом примере несколько продюсеров генерируют данные, которые помещаются в канал,
// а несколько потребителей извлекают эти данные для дальнейшей обработки.
package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Конфигурационные параметры примера.
const (
	numProducers   = 3             // Количество горутин-продюсеров.
	numConsumers   = 4             // Количество горутин-потребителей.
	channelBufSize = 10            // Размер буфера канала.
	runDuration    = 5 * time.Second // Время работы системы.
)

// RunExample запускает демонстрацию паттерна продюсер-потребитель.
func RunExample() {
	// Создаём буферизированный канал для передачи данных между продюсерами и потребителями.
	dataCh := make(chan int, channelBufSize)

	// Используем контекст для управления временем работы продюсеров.
	ctx, cancel := context.WithTimeout(context.Background(), runDuration)
	defer cancel()

	var prodWg sync.WaitGroup
	var consWg sync.WaitGroup

	// Запуск продюсеров.
	for i := 0; i < numProducers; i++ {
		prodWg.Add(1)
		go producer(ctx, i, dataCh, &prodWg)
	}

	// Запуск потребителей.
	for i := 0; i < numConsumers; i++ {
		consWg.Add(1)
		go consumer(i, dataCh, &consWg)
	}

	// Ожидаем завершения работы продюсеров.
	prodWg.Wait()
	// После завершения работы продюсеров закрываем канал,
	// сигнализируя потребителям, что данные больше не будут поступать.
	close(dataCh)
	// Ожидаем завершения работы потребителей.
	consWg.Wait()

	fmt.Println("Демонстрация паттерна продюсер-потребитель завершена.")
}

// producer генерирует случайные целочисленные данные и отправляет их в канал.
// Работа продюсера завершается, когда контекст ctx сигнализирует об окончании работы.
func producer(ctx context.Context, producerID int, dataCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			// Контекст завершён: прекращаем генерацию данных.
			fmt.Printf("Продюсер #%d завершает работу.\n", producerID)
			return
		default:
			// Генерируем случайное число.
			item := rand.Intn(1000)
			// Отправляем данные в канал. Если канал заполнен, горутина блокируется до освобождения места.
			dataCh <- item
			fmt.Printf("Продюсер #%d сгенерировал: %d\n", producerID, item)
			// Имитация переменной задержки производства.
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		}
	}
}

// consumer извлекает данные из канала и обрабатывает их.
func consumer(consumerID int, dataCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for item := range dataCh {
		// Симуляция обработки: произвольная задержка.
		fmt.Printf("Потребитель #%d обработал: %d\n", consumerID, item)
		time.Sleep(time.Duration(rand.Intn(150)) * time.Millisecond)
	}
	fmt.Printf("Потребитель #%d завершает работу, канал закрыт.\n", consumerID)
}

// Для удобства запуска примера через пакет main.
func main() {
	// Инициализируем генератор случайных чисел.
	rand.Seed(time.Now().UnixNano())
	RunExample()
}
