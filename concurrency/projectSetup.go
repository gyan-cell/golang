package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID     int
	Status string
	mu     sync.Mutex
}

var (
	totalUpdates int
	updateMutex  sync.Mutex
)

// waitgroup is used to wait for the goroutines to complete and it coordinates various Go Routines
func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	orders := GenerateOrders(10)

	go func() {
		defer wg.Done()
		ProcessOrders(orders)
	}()

	for i := 0; i < 4; i++ {
		go func(i int) {
			defer wg.Done()
			for _, ord := range orders {

				updateOrderStatuses(ord)

			}
		}(i)
		fmt.Println("Total updates:", totalUpdates)
	}

	OrderStatusReport(orders)
	wg.Wait()
	fmt.Println("All operations have been completed.")
}

// mutex is a locking mechanism that ensures that only one go routine can access the section of the code , sues FIFO method

func ProcessOrders(orders []*Order) {
	for _, order := range orders {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		fmt.Printf("Processing order %d\n", order.ID)
	}

}
func updateOrderStatuses(orders *Order) {
	orders.mu.Lock()
	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	Status := []string{
		"Pending",
		"Processing",
		"Shipped",
		"Delivered",
	}[rand.Intn(4)]
	orders.Status = Status
	fmt.Printf("Order statuses have been updated. of %d and status is  :%s\n ", orders.ID, orders.Status)
	orders.mu.Unlock()
	currentUpdates := totalUpdates
	time.Sleep(time.Duration(5 * time.Millisecond))
	totalUpdates = currentUpdates + 1
}

func OrderStatusReport(order []*Order) {

	fmt.Println("Order Status Report")
	for _, order := range order {
		fmt.Printf("Order %d: %s\n", order.ID, order.Status)
	}
}

func GenerateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{
			ID:     i + 1,
			Status: "Pending",
		}

	}
	return orders
}
