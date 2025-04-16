package domain

type Order struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	UserID    int    `json:"user_id"`
	OrderDate string `json:"order_date"`
	Status    string `json:"status"`
}
type OrderItem struct {
	ID        int     `json:"id" gorm:"primaryKey"`
	OrderID   int     `json:"order_id"`
	ProductID int     `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
