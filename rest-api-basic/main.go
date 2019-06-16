package main

import (
	"gopkg.in/mgo.v2"
	"net/http"

	"go-best-practices/rest-api-basic/controller"

	"github.com/julienschmidt/httprouter"
)

func main() {
	customer := controller.Customer(session())

	r := httprouter.New()
	r.POST("/customer/show", customer.Show)
	r.POST("/customer/show/:id", customer.Show)
	r.POST("/customer/save", customer.Save)
	r.POST("/customer/save/:id", customer.Save)
	r.POST("/customer/drop/", customer.Drop)
	r.POST("/customer/drop/:id", customer.Drop)
	r.POST("/customer/field", customer.Field)

	http.ListenAndServe(":8080", r)
}

func session() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")

	if err != nil {
		panic(err)
	}

	return session
}
