package myml

type General struct {
	ID       int       `json:"id"`
	Category *Category `json:"category"`
	Currency *Currency `json:"currency"`
}
