package model

type Cart struct {
	InfoProduct []InfoProduct
	IDUser      string
	Quantity    int
	FinalValue  float64
	ID          string
}
type InfoProduct struct {
	IDProd       string
	QuantityProd int
}
