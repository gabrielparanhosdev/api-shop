package types

type ProductResponse struct {
	Id          int     `json:"_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
