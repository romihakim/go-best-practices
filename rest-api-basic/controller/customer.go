package controller

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"

	"go-best-practices/rest-api-basic/model"

	"github.com/julienschmidt/httprouter"
)

type CustomerController struct {
	session *mgo.Session
}

func Customer(s *mgo.Session) *CustomerController {
	return &CustomerController{s}
}

func (c CustomerController) Show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var j []byte

	if p.ByName("id") == "" {
		m := []model.Customer{}

		if err := c.session.DB("rest-basic").C("customer").Find(nil).All(&m); err != nil {
			w.WriteHeader(404)
			return
		}

		j, _ = json.Marshal(m)
	} else {
		m := model.Customer{}

		if !bson.IsObjectIdHex(p.ByName("id")) {
			w.WriteHeader(404)
			return
		}

		id := bson.ObjectIdHex(p.ByName("id"))

		if err := c.session.DB("rest-basic").C("customer").FindId(id).One(&m); err != nil {
			w.WriteHeader(404)
			return
		}

		j, _ = json.Marshal(m)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	fmt.Fprintf(w, "%s", j)
}

func (c CustomerController) Save(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	m := model.Customer{}

	json.NewDecoder(r.Body).Decode(&m)

	if p.ByName("id") == "" {
		m.Id = bson.NewObjectId()

		if err := c.session.DB("rest-basic").C("customer").Insert(m); err != nil {
			w.WriteHeader(404)
			return
		}
	} else {
		if !bson.IsObjectIdHex(p.ByName("id")) {
			w.WriteHeader(404)
			return
		}

		m.Id = bson.ObjectIdHex(p.ByName("id"))

		if err := c.session.DB("rest-basic").C("customer").UpdateId(m.Id, m); err != nil {
			w.WriteHeader(404)
			return
		}
	}

	j, _ := json.Marshal(m)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	fmt.Fprintf(w, "%s", j)
}

func (c CustomerController) Drop(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if p.ByName("id") == "" {
		info, err := c.session.DB("rest-basic").C("customer").RemoveAll(nil)

		if info.Removed == 0 {
			w.WriteHeader(404)
			return
		}

		if err != nil {
			w.WriteHeader(404)
			return
		}
	} else {
		if !bson.IsObjectIdHex(p.ByName("id")) {
			w.WriteHeader(404)
			return
		}

		id := bson.ObjectIdHex(p.ByName("id"))

		if err := c.session.DB("rest-basic").C("customer").RemoveId(id); err != nil {
			w.WriteHeader(404)
			return
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
}

func (c CustomerController) Field(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	m := model.Customer{}
	m.Name = "ROMI A. HAKIM"
	m.Email = "romi@mail.com"

	json.NewDecoder(r.Body).Decode(&m)
	j, _ := json.Marshal(m)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	fmt.Fprintf(w, "%s", j)
}
