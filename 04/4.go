package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала
и выводят в stdout. Необходима возможность выбора количества воркеров при старте.

Программа должна завершаться по нажатию Ctrl+C.
Выбрать и обосновать способ завершения работы всех воркеров.
*/

type Consumer struct {
	ingestChan chan int
	jobsChan   chan int
}

func (c Consumer) workerFunc(wg *sync.WaitGroup, index int) {
	defer wg.Done()

	fmt.Printf("Воркер %d начинает работу\n", index)
	for eventIndex := range c.jobsChan {
		// симулируем работу
		fmt.Printf("Воркер %d начал работу %d\n", index, eventIndex)
		time.Sleep(time.Millisecond * time.Duration(1000+rand.Intn(2000)))
		fmt.Printf("Воркер %d завершил работу %d\n", index, eventIndex)
	}
}

// callbackFunc вызвается каждый раз когда Producer отправляет нам событие.
func (c Consumer) callbackFunc(event int) {
	c.ingestChan <- event
}

// startConsumer выступает прокси между ingestChan и jobsChan + здесь реализуется graceful shutdown
func (c Consumer) startConsumer(ctx context.Context) {
	for {
		select {
		case job := <-c.ingestChan:
			c.jobsChan <- job
		case <-ctx.Done():
			fmt.Println("Consumer получил сигнал закрытия, закрываем jobsChan")
			close(c.jobsChan)
			fmt.Println("Consumer закрыл jobsChan")
			return
		}
	}
}

// Producer - симулирует внешнюю библиотеку которая, которая генерирует события
type Producer struct {
	callbackFunc func(event int)
}

func (p Producer) start() {
	eventIndex := 0
	for {
		p.callbackFunc(eventIndex)
		eventIndex++
		time.Sleep(time.Millisecond * 250)
	}
}

func main() {
	var workerPoolSize int
	fmt.Print("Введите число воркеров: ")
	fmt.Scanln(&workerPoolSize)

	consumer := Consumer{
		ingestChan: make(chan int, 1),
		jobsChan:   make(chan int, workerPoolSize),
	}

	// Симулируем работу внешней бибилиотеки, которая генерирует работу для воркеров
	producer := Producer{callbackFunc: consumer.callbackFunc}
	go producer.start()

	// Настройка контекста для graceful shutdown и waitgroup
	ctx, cancelFunc := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	// Запускаем consumer. пишем данные в канал jobschan
	go consumer.startConsumer(ctx)

	// Запускаем воркеров
	wg.Add(workerPoolSize)
	for i := 0; i < workerPoolSize; i++ {
		go consumer.workerFunc(wg, i)
	}

	// Handle sigterm and await termChan signal
	// Ловим Ctlr+C и ждем сигнал в termChan
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	<-termChan // Лочимся до тех пор пока не прийдет сигнал в канал

	// Обработка Ctrl+c
	fmt.Println("******\nShutdown signal received\n******")
	cancelFunc() // Закрываем контекст
	wg.Wait()    // Ждем пока воркеры сделают всю оставшуюся в канале работу

	fmt.Println("All workers done, shutdown!")
}
