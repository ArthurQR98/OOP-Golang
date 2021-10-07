package invoiceitem

type Item struct {
	id      uint64
	product string
	value   float64
}

func New(id uint64, product string, value float64) Item {
	return Item{id, product, value}
}

// getter de value
// func (i Item) Value() float64 {
// 	return i.value
// }

type Items []Item

func NewItems(items ...Item) Items {
	var is Items
	for _, item := range items {
		is = append(is, item)
	}
	return is
}

func (is Items) Total() float64 {
	var total float64
	for _, item := range is {
		total += item.value
	}
	return total
}
