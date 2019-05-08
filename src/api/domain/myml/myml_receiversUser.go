package myml

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"io/ioutil"
	"net/http"
)

func (user *User) Get() *apierrors.ApiError {

	urlUser := "https://api.mercadolibre.com/users/"
	final := fmt.Sprintf("%s%d", urlUser, user.ID)

	response, err := http.Get(final)

	if err != nil {
		return &apierrors.ApiError{
			Message: "http.Get failed",
			Status:  http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return &apierrors.ApiError{
			Message: "Unmarshall failed",
			Status:  http.StatusInternalServerError,
		}
	}

	return nil

}

/*func (user *User) Get2(ch chan ParaCanal, userID int) *apierrors.ApiError {

	urlUser := "https://api.mercadolibre.com/users/"
	final := fmt.Sprintf("%s%d", urlUser, user.ID)

	response, err := http.Get(final)

	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return nil

}
*/
