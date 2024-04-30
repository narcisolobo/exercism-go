package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	amount, exists := units[unit]
	if !exists {
		return false
	}
	_, exists = bill[item]
	if exists {
		bill[item] += amount
	} else {
		bill[item] = amount
	}
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	result := true
	current, exists := bill[item]
	if !exists {
		result = false
	}
	amount, exists := units[unit]
	if !exists {
		result = false
	}
	adjusted := current - amount

	switch {
	case adjusted < 0:
		result = false
	case adjusted == 0:
		delete(bill, item)
	case adjusted > 0:
		bill[item] = adjusted
	}
	return result
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, exists := bill[item]
	if !exists {
		return 0, false
	}
	return qty, true
}
