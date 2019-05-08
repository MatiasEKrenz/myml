package myml

import "github.com/mercadolibre/myml/src/api/utils/apierrors"

type General struct {
	Category *Category           `json:"categories"`
	Currency *Currency           `json:"currency"`
	Errores  *apierrors.ApiError `json:"error"`
}
