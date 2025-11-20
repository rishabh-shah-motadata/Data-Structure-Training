package day9

import "fmt"

func (c *cart) total(discounts map[string]float64) {
	total := 0.0
	cur := c.head

	for cur != nil {
		itemTotal := cur.price * float64(cur.quantity)

		if d, ok := discounts[cur.id]; ok {
			itemTotal = itemTotal - (itemTotal * d)
		}

		total += itemTotal
		cur = cur.next
	}

	fmt.Printf("\nTotal Payable: %.2f\n", total)
}

func CartSystem() {
	cart := &cart{}
	discounts := newDiscount()
	wishlist := newWishlist()

	cart.add("101", "Shoes", 2000, 1)
	cart.add("102", "T-Shirt", 800, 2)
	cart.add("106", "Belt", 50, 1)
	cart.add("103", "Watch", 3500, 1)

	fmt.Println("Cart Items:")
	cart.display()
	cart.total(discounts.items)

	fmt.Println("\nRemoving item 106...")
	cart.remove("106")

	fmt.Println("\nUpdated Cart:")
	cart.display()
	cart.total(discounts.items)

	discounts.add("102", 0.10)
	discounts.add("104", 0.3)
	discounts.add("101", 0.05)

	fmt.Println("\nDiscount Items:")
	discounts.display()
	cart.total(discounts.items)

	fmt.Println("\nRemoving discount for item 101...")
	discounts.remove("101")
	cart.total(discounts.items)

	fmt.Println("\nUpdated Discount Items:")
	discounts.display()
	cart.total(discounts.items)

	wishlist.add("201", "Sunglasses")
	wishlist.add("202", "Jacket")
	wishlist.add("203", "Wallet")

	fmt.Println("\nWishlist:")
	wishlist.display()

	fmt.Println("\nRemoving item 202 from wishlist...")
	wishlist.remove("202")

	fmt.Println("\nUpdated Wishlist:")
	wishlist.display()
}
