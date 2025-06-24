package main

import (
	"fmt"
	"time"
)

// order struct

type customer struct {
	name  string
	phone string
}

// composition
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

// func newOrder(id string, amount float32, status string) *order {
// 	// initial setup goes here...
// 	myOrder := order{
// 		id:     id,
// 		amount: amount,
// 		status: status,
// 	}

// 	return &myOrder
// }

// // receiver type
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// func (o order) getAmount() float32 {
// 	return o.amount
// }

func main() {
	// newCustomer := customer{
	// 	name:  "john",
	// 	phone: "1234567890",
	// }
	newOrder := order{
		id:     "1",
		amount: 30,
		status: "received",
		customer: customer{
			name:  "john",
			phone: "1234567890",
		},
	}

	newOrder.customer.name = "robin"
	fmt.Println(newOrder)

	// language := struct {
	// 	name   string
	// 	isGood bool
	// }{"golang", true}

	// fmt.Println(language)

	// myOrder := newOrder("1", 30.50, "received")
	// fmt.Println(myOrder.amount)
	// if you don't set any field, default value is zero value
	// int => 0, float => 0, string "", bool => false
	// myOrder := order{
	// 	id:     "1",
	// 	amount: 50.00,
	// 	status: "received",
	// }
	// myOrder.changeStatus("confirmed")
	// fmt.Println(myOrder)
	// myOrder.createdAt = time.Now()
	// fmt.Println(myOrder.status)

	// myOrder2 := order{
	// 	id:        "2",
	// 	amount:    100,
	// 	status:    "delivered",
	// 	createdAt: time.Now(),
	// }

	// myOrder.status = "paid"

	// fmt.Println("Order struct", myOrder2)
	// fmt.Println("Order struct", myOrder)
}

// package main  

// import (
// 	"fmt"
// 	"time"
// )

// // Customer struct represents a customer with name and phone details
// type customer struct {
// 	name  string
// 	phone string
// }

// // Order struct represents an order with embedded customer details
// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time // nanosecond precision
// 	customer
// }

// // newOrder creates and returns a new order with initialized fields
// func newOrder(id string, amount float32, status string, cust customer) *order {
// 	return &order{
// 		id:        id,
// 		amount:    amount,
// 		status:    status,
// 		createdAt: time.Now(),
// 		customer:  cust,
// 	}
// }

// // changeStatus updates the status of an order (receiver is a pointer for modification)
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// // getAmount returns the amount of an order (receiver is a value for read-only access)
// func (o order) getAmount() float32 {
// 	return o.amount
// }

// func main() {
// 	// Create a new customer
// 	newCustomer := customer{
// 		name:  "John",
// 		phone: "1234567890",
// 	}

// 	// Create a new order using the constructor function
// 	order1 := newOrder("1", 30.50, "received", newCustomer)
// 	fmt.Println(order1)
// 	fmt.Println(order1.status)

// 	// Update customer name for order1
// 	order1.customer.name = "Robin"
// 	fmt.Println(order1)

// 	// Create an anonymous struct for a language
// 	language := struct {
// 		name   string
// 		isGood bool
// 	}{"Golang", true}
// 	fmt.Println(language)

// 	// Create another order using struct literal
// 	order2 := order{
// 		id:        "2",
// 		amount:    50.00,
// 		status:    "received",
// 		createdAt: time.Now(),
// 		customer: customer{
// 			name:  "Alice",
// 			phone: "0987654321",
// 		},
// 	}
// 	fmt.Println(order2)

// 	// Change status of order2
// 	order2.changeStatus("confirmed")
// 	fmt.Println(order2.status)

// 	// Create a third order
// 	order3 := newOrder("3", 100.00, "delivered", customer{name: "Bob", phone: "5555555555"})
// 	fmt.Println(order3)

// 	// Update status of order2
// 	order2.status = "paid"
// 	fmt.Println(order2)
// }
