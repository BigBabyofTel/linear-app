package main

import (
	"errors"
	"net/http"
	"time"

	"github.com/lucabrx/wuhu/internal/data"
	"github.com/lucabrx/wuhu/internal/validator"
)

func (a *app) emailRegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := a.readJSON(w, r, &input); err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	user := &data.User{
		Name:     input.Name,
		Email:    input.Email,
		Verified: false,
	}

	if err := user.Password.Set(input.Password); err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	v := validator.New()
	if data.ValidateUser(v, user); !v.Valid() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	if err := a.DB.User.Insert(user); err != nil {
		switch {
		case errors.Is(err, data.ErrDuplicateEmail):
			v.AddError("email", "a user with this email address already exists")
			a.failedValidationResponse(w, r, v.Errors)
		default:
			a.serverErrorResponse(w, r, err)
		}
		return
	}

	if err := a.writeJSON(w, http.StatusCreated, user, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *app) emailLoginUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := a.readJSON(w, r, &input); err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	if data.ValidateEmail(v, input.Email); !v.Valid() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}
	if data.ValidatePasswordPlaintext(v, input.Password); !v.Valid() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	user, err := a.DB.User.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			a.invalidCredentialsResponse(w, r)
		default:
			a.serverErrorResponse(w, r, err)
		}
	}

	match, err := user.Password.Matches(input.Password)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	if !match {
		a.invalidCredentialsResponse(w, r)
		return
	}

	month := 30 * 24 * time.Hour
	sessionToken, err := a.DB.Token.New(user.ID, month, data.ScopeAuthentication)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	if err := a.writeJSON(w, http.StatusOK, envelope{"token": sessionToken.Plaintext}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}
}

func (a *app) logoutHandler(w http.ResponseWriter, r *http.Request) {
	session := a.contextGetUser(r)
	err := a.DB.Token.DeleteAllForUser(data.ScopeAuthentication, session.ID)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusTemporaryRedirect)

}
