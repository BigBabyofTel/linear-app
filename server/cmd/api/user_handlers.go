package main

import (
	"net/http"

	"github.com/lucabrx/wuhu/internal/data"
	"github.com/lucabrx/wuhu/internal/validator"
)

func (a *app) getSessionUserHandler(w http.ResponseWriter, r *http.Request) {
	session := a.contextGetUser(r)

	if err := a.writeJSON(w, http.StatusOK, envelope{"user": session}, nil); err != nil {
		a.badRequestResponse(w, r, err)
	}
}

func (a *app) deleteUserAccountHandler(w http.ResponseWriter, r *http.Request) {
	user := a.contextGetUser(r)
	token := a.contextGetToken(r)
	if err := a.DB.User.Delete(user.ID); err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	if err := a.Cache.User.DelUserByToken(r.Context(), token); err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	http.Redirect(w, r, a.config.ClientUrl, http.StatusSeeOther)
}

func (a *app) updatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		CurrentPassword string `json:"currentPassword"`
		NewPassword     string `json:"newPassword"`
	}

	if err := a.readJSON(w, r, &input); err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	if data.ValidatePasswordPlaintext(v, input.CurrentPassword); !v.Valid() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	if data.ValidatePasswordPlaintext(v, input.NewPassword); !v.Valid() {
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	user := a.contextGetUser(r)

	match, err := user.Password.Matches(input.CurrentPassword)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	if !match {
		v.AddError("current_password", "password is incorrect")
		a.failedValidationResponse(w, r, v.Errors)
		return
	}

	if err := user.Password.Set(input.NewPassword); err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	if err := a.DB.User.Update(user); err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	if err := a.writeJSON(w, http.StatusOK, envelope{"message": "password updated"}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}
}
