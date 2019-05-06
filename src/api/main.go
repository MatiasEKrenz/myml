package main

import (
	//clacla "github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/myml/src/api/controllers/myml"
)

const (
	//aca tiene en cuenta la mayuscula para que sea publica
	//aca y en var, no se usa el ":="
	port = ":8080"
)

var(
	//aca tiene en cuenta la mayuscula para que sea publica
	// en var, se puede cambiar el valor pero no el tipo de datos

	// si creo el alias, deberia usar clacla.Default()
	router = gin.Default()
)


func main()  {
	//aca tiene en cuenta la mayuscula para que sea publica

	// pongo la funcion sin parentesis porque no tengo que ejecutar la funcion ahora, solo le paso la funcion como parametro
	//router.GET("/ping", ping.Ping)

	//router.GET("/users/:id", myml.User)

	router.GET("/users/:id", myml.GetUser)

	router.GET("/users/:id/site", myml.GetInfo)


	router.Run(port)

}



// USER_ID: 152581223
// https://api.mercadolibre.com/users/USER_ID

//https://api.mercadolibre.com/users/152581223

/*
Controllers: Validar parametros + llamar servicios
Services: Consumir la API de ML




https://github.com/emikohmann/academy-myml




user = &User{
	ID: 1234567,
	}

user.Get()
//Create waitgroup
//create channel

go categories.Get()
go currency.Get()

//read channels

//meter una estructura en el canal

return &MyMl {

	User: user,
	...
	...
	...

}


*/
