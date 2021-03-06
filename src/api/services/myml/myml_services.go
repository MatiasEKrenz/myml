package myml

import (
	"encoding/json"
	"fmt"
	"github.com/mercadolibre/myml/src/api/domain/myml"
	"github.com/mercadolibre/myml/src/api/utils/apierrors"
	"io/ioutil"
	"net/http"
	"sync"
)

const (
	urlApi  = "https://api.mercadolibre.com/"
	urlMock = "http://localhost:8081/"

	urlAUsar = urlMock
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

	user := myml.User{
		ID: int(userID),
	}

	err := user.Get()

	if err != nil {
		return nil, &apierrors.ApiError{
			Message: "Error en el GET",
			Status:  http.StatusInternalServerError,
		}
	}

	return &user, nil
}

func GetGeneralInfo(userID int64) (*myml.General, *apierrors.ApiError) {

	final := fmt.Sprintf("%susers/%d", urlAUsar, userID)

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

	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan *myml.General, 2)

	//go routines
	go func() { ch <- getCategories(user.SiteID, &wg) }()
	go func() { ch <- getCurrencies(user.CountryID, &wg) }()

	wg.Wait()
	close(ch)
	var MyMl myml.General

	//for
	for i := range ch {

		if i.Category != nil {
			MyMl.Category = i.Category
			continue
		}

		if i.Currency != nil {
			MyMl.Currency = i.Currency
			continue
		}

		if i.Errores != nil {
			return nil, &apierrors.ApiError{
				Message: "Error in reply papa",
				Status:  http.StatusInternalServerError,
			}
		}
	}

	return &MyMl, nil
}

func getCategories(siteId string, wg *sync.WaitGroup) *myml.General {

	defer wg.Done()

	final := fmt.Sprintf("%ssites/%s/categories", urlAUsar, siteId)

	fmt.Println("CLA1:" + final)

	response, err := http.Get(final)

	var category myml.Category

	if err != nil {
		return &myml.General{
			Category: nil,
			Currency: nil,
			Errores: &apierrors.ApiError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &myml.General{
			Category: nil,
			Currency: nil,
			Errores: &apierrors.ApiError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		}
	}

	// esto se hace para que la variable err no esté viva al pedo en la ejecucion, ya que solo se usa en el IF
	if err := json.Unmarshal([]byte(data), &category); err != nil {
		return &myml.General{
			Category: nil,
			Currency: nil,
			Errores: &apierrors.ApiError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		}
	}

	MyMl := myml.General{
		Category: &category,
	}

	return &MyMl
}

func getCurrencies(countryId string, wg *sync.WaitGroup) *myml.General {

	defer wg.Done()

	final := fmt.Sprintf("%sclassified_locations/countries/%s", urlAUsar, countryId)

	fmt.Println("CLA2:" + final)

	response, err := http.Get(final)

	country := myml.Country{}

	if err != nil {
		return &myml.General{
			Category: nil,
			Currency: nil,
			Errores: &apierrors.ApiError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		}
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &myml.General{
			Category: nil,
			Currency: nil,
			Errores: &apierrors.ApiError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		}
	}

	// esto se hace para que la variable err no esté viva al pedo en la ejecucion, ya que solo se usa en el IF
	if err := json.Unmarshal([]byte(data), &country); err != nil {
		return &myml.General{
			Category: nil,
			Currency: nil,
			Errores: &apierrors.ApiError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		}
	}

	final = fmt.Sprintf("%scurrencies/%s", urlAUsar, country.CurrencyID)

	fmt.Println("CLA3:" + final)

	response, err = http.Get(final)

	var currency myml.Currency

	if err != nil {
		return &myml.General{
			Category: nil,
			Currency: nil,
			Errores: &apierrors.ApiError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		}
	}

	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return &myml.General{
			Category: nil,
			Currency: nil,
			Errores: &apierrors.ApiError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		}
	}

	// esto se hace para que la variable err no esté viva al pedo en la ejecucion, ya que solo se usa en el IF
	if err := json.Unmarshal([]byte(data), &currency); err != nil {
		return &myml.General{
			Category: nil,
			Currency: nil,
			Errores: &apierrors.ApiError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			},
		}
	}

	MyMl := myml.General{
		Currency: &currency,
	}

	return &MyMl
}
