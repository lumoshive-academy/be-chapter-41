package model

type Shipping struct {
	Base
	Name  string  `json:"name" gorm:"type:varchar(100);not null"`
	Price float64 `json:"price" gorm:"type:decimal(10,2);not null"` // Sesuaikan skala decimal
}

func ShippingSeed() []Shipping {
	return []Shipping{
		{Name: "JNE", Price: 20000.00},
		{Name: "JNT", Price: 50000.00},
		{Name: "Ninja", Price: 100000.00},
	}
}
