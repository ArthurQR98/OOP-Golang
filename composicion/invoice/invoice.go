package invoice

import (
	"github.com/Arthur98/poo/composicion/customer"
	"github.com/Arthur98/poo/composicion/invoiceitem"
)

type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer // realizando la composicion de 1 a 1
	// items   []invoiceitem.Item // realizando la composicion de 1 a muchos
	items invoiceitem.Items
}

// contry, city string, client customer.Customer, items []invoiceitem.Item
func New(contry, city string, client customer.Customer, items invoiceitem.Items) Invoice {
	return Invoice{
		country: contry,
		city:    city,
		client:  client,
		items:   items,
	}
}

// SetTotal is the setter of Invoice.total
func (i *Invoice) SetTotal() {
	// for _, item := range i.items {
	// 	i.total += item.Value()
	// }
	i.total = i.items.Total()
}
