package myml

type Currency []struct {
	ID            string `json:"id"`
	Symbol        string `json:"symbol"`
	Description   string `json:"description"`
	DecimalPlaces int    `json:"decimal_places"`
}
