package order

import (
	"fmt"
	orders "github.com/romaxa83/hra/proto"
	"sync"
	"time"
)

// OrderDispatcher - это процесс-демон, который создает набор обработчиков с помощью sync.waitGroup, чтобы параллельно
// обрабатывать и отправлять заказы
type OrderDispatcher struct {
	ordersCh   chan *orders.Order
	orderLimit int // maximum number of orders the pool will process concurrently
}

// NewOrderDispatcher создает новый OrderDispatcher
func NewOrderDispatcher(orderLimit int, bufferSize int) OrderDispatcher {
	return OrderDispatcher{
		ordersCh:   make(chan *orders.Order, bufferSize), // initiliaze as a buffered channel
		orderLimit: orderLimit,
	}
}

// SubmitOrder отправляет заказ на обработку
func (d OrderDispatcher) SubmitOrder(order *orders.Order) {
	go func() {
		d.ordersCh <- order
	}()
}

// Start запускает диспетчера в фоновом режиме
func (d OrderDispatcher) Start() {
	go d.processOrders()
}

// Shutdown отключает OrderDispatcher путем закрытия канала заказов
// Примечание: эта функция должна выполняться только после того, как последний заказ
// попадет в канал заказов. Отправка заказа в закрытый канал вызовет панику.
func (d OrderDispatcher) Shutdown() {
	close(d.ordersCh)
}

// processOrders обрабатывает все входящие заказы в фоновом режиме с помощью
// for-range и sync.waitGroup
func (d OrderDispatcher) processOrders() {
	limiter := make(chan struct{}, d.orderLimit)
	var wg sync.WaitGroup

	// Непрерывная обработка заказов, полученных из канала заказов
	// Этот цикл завершится после закрытия канала
	for order := range d.ordersCh {
		limiter <- struct{}{}
		wg.Add(1)

		go func(order *orders.Order) {
			// Что нужно сделать: запустить процесс выполнения, чтобы собрать заказ в пакет и отправить
			// Пока используем спящий режим и печать
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("Order (%v) has shipped \n", order)
			<-limiter
			wg.Done()
		}(order)
	}
	wg.Wait()
}

func main() {
	dispatcher := NewOrderDispatcher(3, 100)
	dispatcher.Start()
	defer dispatcher.Shutdown()

	dispatcher.SubmitOrder(&orders.Order{Items: []*orders.Item{{Description: "iPhone Screen Protector", Price: 9.99}}})
	dispatcher.SubmitOrder(&orders.Order{Items: []*orders.Item{{Description: "iPhone Case", Price: 19.99}}})
	dispatcher.SubmitOrder(&orders.Order{Items: []*orders.Item{{Description: "Pixel Case", Price: 14.99}}})
	dispatcher.SubmitOrder(&orders.Order{Items: []*orders.Item{{Description: "Bluetooth Speaker", Price: 29.99}}})
	dispatcher.SubmitOrder(&orders.Order{Items: []*orders.Item{{Description: "4K Monitor", Price: 159.99}}})
	dispatcher.SubmitOrder(&orders.Order{Items: []*orders.Item{{Description: "Inkjet Printer", Price: 79.99}}})

	time.Sleep(5 * time.Second) // just for testing
}
