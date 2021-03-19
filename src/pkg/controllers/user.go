package controllers

import (
	u "croupier/pkg/utils"
	"encoding/json"
	"net/http"

	"croupier/pkg/user"
)

//CreateAccount !
var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &user.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := account.Create()
	u.Respond(w, resp)
}

//Authenticate !
var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &user.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		u.Respond(w, u.Message(false, "Invalid request"))
		return
	}

	resp := user.Login(account.Email, account.Login, account.Password)
	u.Respond(w, resp)
}
