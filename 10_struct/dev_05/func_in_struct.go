package main

import "fmt"

type CoffeeShop struct {
	shopName string
}

type OtherShop struct {
	shopName string
	otherFunc func(shop OtherShop)
}

func greetOther(shop OtherShop) {
	fmt.Printf("Hello in shop: %s\n", shop.shopName)
}

func greetShop(shop CoffeeShop) {
	fmt.Printf("Hello in shop: %s\n", shop.shopName)
}

func main() {
	myShop1 := CoffeeShop{
		shopName: "Tula",
	}

	greetShop(myShop1)

	myOther1 := OtherShop{
		shopName: "Moscow",
		otherFunc: greetOther,
	}

	myOther1.otherFunc(myOther1)

	myOther2 := OtherShop{
		shopName: "Sochi",
	}

	fmt.Printf("%#v\n", myOther2)
}