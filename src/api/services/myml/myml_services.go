package myml

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/myml/src/api/domain/myml"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"io/ioutil"
	"net/http"
)

/*func GetUsersParams(id string) myml.User {

	urlUser := "https://api.mercadolibre.com/users/" + id
	response, err := http.Get(urlUser)
	user := myml.User{}
	if err, ok := err.(*url.Error); ok {
		fmt.Printf("Ocurrió un error al traer los datos de usuario %s\n", err)
	} else {
		//fmt.Println(f)
		data, _ := ioutil.ReadAll(response.Body)
		err := json.Unmarshal([]byte(data), &user)
		if err != nil {
			panic(err)
		}
		//return user
	}
	return user
}*/

/*func GetUserFromAPI(userID int64) (*myml.User, *apierrors.ApiError) {

	urlUser := "https://api.mercadolibre.com/users/"
	final := fmt.Sprintf("%s%d", urlUser, userID)

	if userID == 0 {
		return nil, &apierrors.ApiError{
			Message: "userId id empty",
			Status: http.StatusBadRequest,
		}
	}

	response, err := http.Get(final)
	user := myml.User{}
	if err != nil {
		return nil, &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	// esto se hace para que la variable err no esté viva al pedo en la ejecucion, ya que solo se usa en el IF
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return nil, &apierrors.ApiError{
			Message: err.Error(),
			Status: http.StatusInternalServerError,
		}
	}

	return &user, nil
}*/

func GetUserFromAPI(userID int64) (*myml.User, *apierrors.ApiError) {

	if userID == 0 {
		return nil, &apierrors.ApiError{
			Message: "userId id empty",
			Status:  http.StatusBadRequest,
		}
	}

	user := &myml.User{
		ID: int(userID),
	}

	err := user.Get()

	if err != nil {
		return nil, &apierrors.ApiError{
			Message: "Error en el GET",
			Status:  http.StatusInternalServerError,
		}
	}

	return user, nil
}

func GetGeneralInfo(userID int64) (*myml.User, *apierrors.ApiError) {

	urlUser := "https://api.mercadolibre.com/users/"
	final := fmt.Sprintf("%s%d", urlUser, userID)

	if userID == 0 {
		return nil, &apierrors.ApiError{
			Message: "userId id empty",
			Status:  http.StatusBadRequest,
		}
	}

	response, err := http.Get(final)
	user := myml.User{}
	if err != nil {
		return nil, &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	// esto se hace para que la variable err no esté viva al pedo en la ejecucion, ya que solo se usa en el IF
	if err := json.Unmarshal([]byte(data), &user); err != nil {
		return nil, &apierrors.ApiError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &user, nil
}
