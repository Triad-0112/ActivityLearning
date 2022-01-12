package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

/* To implement this, you'll need to:
Return false if the given unit is not in the units map.
Otherwise add the item to the customer bill, indexed by the item name, then return true.
If the item is already present in the bill, increase its quantity by the amount that belongs to the provided unit.
*/
func AddItem(bill, units map[string]int, item, unit string) bool {
	value, checks := units[unit]
	if checks {
		bill[item] += value
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	value, checks := units[unit]
	if !checks {
		return false
	}
	if bill[item] >= value {
		bill[item] -= value
		if bill[item] == 0 {
			delete(bill, item)
		}
		return true
	}
	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, exist := bill[item]
	return quantity, exist
}
