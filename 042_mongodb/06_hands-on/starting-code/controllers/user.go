package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	//"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"github.com/golang-web-dev/042_mongodb/06_hands-on/starting-code/models"
)

var Users = make(map[string]models.User) // userId -> User

type UserController struct {
	//session *mgo.Session

}

//func NewUserController(s *mgo.Session) *UserController {
//	return &UserController{s}
//}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId hex representation, otherwise return status not found
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound) // 404
		return
	}

	// ObjectIdHex returns an ObjectId from the provided hex representation.
	oid := bson.ObjectIdHex(id)

	// composite literal
	//u := models.User{}

	// Fetch user
	//if err := uc.session.DB("go-web-dev-db").C("users").FindId(oid).One(&u); err != nil {
	u, ok := Users[oid.String()];
	if !ok {
		w.WriteHeader(404)
		return
	}

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.User{}

	json.NewDecoder(r.Body).Decode(&u)

	// create bson ID
	u.Id = bson.NewObjectId()

	// store the user in mongodb
	//uc.session.DB("go-web-dev-db").C("users").Insert(u)
	Users[u.Id.String()] = u

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)
	delete(Users, oid.String())
	// Delete user
	//if err := uc.session.DB("go-web-dev-db").C("users").RemoveId(oid); err != nil {
	//	w.WriteHeader(404)
	//	return
	//}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", oid, "\n")
}
