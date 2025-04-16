package domain

type Product struct {
	ID       int     `json:"id" gorm:"primaryKey"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
}
