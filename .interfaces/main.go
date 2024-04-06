package main

import "fmt"

type Yeller interface {
	Yell()
}

type CanYell struct {
}

func (c *CanYell) Yell() {
	fmt.Println("YELLING!")
}

// Interface compliance to get a compile time check
var _ Yeller = (*CanYell)(nil)

func yell(yeller Yeller) {
	yeller.Yell()
}

func main() {
	// here we define a simple interface, `Yeller` - something which can 'yell'.
	// Go naming convention prefers the `er` based suffiing to interfaces.
	// To enforce the interface implementation on our concrete type we also
	// use an interface compliance technique
	// Now, its much easier to decouple and apply the open closed principle
	yeller := CanYell{}
	yell(&yeller)

	// An example of how interfaces can be used to allow user code to modify
	// core behaviour without relying on internal refactoring, decoupling the
	// the code.  (One of many benefits to interfaces)
	half := &HalfDiscounter{}
	none := &NoDiscounter{}
	fmt.Println(ApplyDiscount(100, half))
	fmt.Println(ApplyDiscount(100, none))
}

// An interface that can apply discounts
type Discounter interface {
	Discount(amount int) int
}

// An implementation of the Discounter interface
type HalfDiscounter struct{}

func (h *HalfDiscounter) Discount(amount int) int {
	return amount / 2
}

var _ Discounter = (*HalfDiscounter)(nil)

type NoDiscounter struct{}

func (n *NoDiscounter) Discount(amount int) int {
	return amount
}

var _ Discounter = (*NoDiscounter)(nil)

// ApplyDiscount returns the discounted value.  It accepts
// any concrete instance of Discounter as a strategy.
func ApplyDiscount(amount int, discountable Discounter) int {
	return discountable.Discount(amount)
}
