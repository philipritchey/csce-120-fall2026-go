package main

import (
	"fmt"
	"math/rand/v2"
)

type Priced interface {
	PriceCents() int
}

type Snack struct {
	Base    int
	Premium int
}

type Drink struct {
	Price int // cents
}

type CouponItem struct {
	Item     Priced
	Discount float32 // as %
}

func (s Snack) PriceCents() int {
	// Base + [0, Premium)
	return s.Base + rand.IntN(s.Premium)
}

func (d Drink) PriceCents() int {
	return d.Price
}

func (c CouponItem) PriceCents() int {
	return int(float32(c.Item.PriceCents()) * (1 - c.Discount))
}

func PrintPrice(p Priced) {
	fmt.Printf("$%d.%0.2d\n", p.PriceCents()/100, p.PriceCents()%100)
}

func main() {
	items := []Priced{
		Snack{123, 100},
		Drink{314},
		CouponItem{Drink{867}, 0.10},
	}
	for _, item := range items {
		PrintPrice(item)
	}
}
