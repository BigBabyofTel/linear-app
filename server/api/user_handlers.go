package api

import (
	"net/http"

	"github.com/lucabrx/wuhu/internal/data"
	"github.com/lucabrx/wuhu/internal/validator"
)

// getSessionUserHandler retrieves the session user details.
//
//	@Summary		Retrieve session user
//	@Description	Get details of the user associated with the current session.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Security		ApiKeyAuth
//	@Success		200	{object}	map[string]interface{}	"user":		data.User
//	@Failure		400	{object}	map[string]interface{}	"error":	"bad request"
//	@Failure		401	{object}	map[string]interface{}	"error":	"unauthorized"	//	If	authentication	fails
//	@Failure		500	{object}	map[string]interface{}	"error":	"internal server error"
//	@Router			/v1/user/ [get]
func (a *app) getSessionUserHandler(w http.ResponseWriter, r *http.Request) {
	session := a.contextGetUser(r)

	if err := a.writeJSON(w, http.StatusOK,
		map[string]interface{}{"user": session}, nil); err != nil {
		a.badRequestResponse(w, r, err)
	}
}

// deleteUserAccountHandler deletes the user account.
//
//	@Summary		Delete user account
//	@Description	Deletes the user account associated with the current session.
//	@Tags			user
//	@Security		ApiKeyAuth
//	@Success		303	{header}	string					"Redirects to client URL"
//	@Failure		400	{object}	map[string]interface{}	"error":	"bad request"
//	@Router			/v1/user/ [delete]
func (a *app) deleteUserAccountHandler(w http.ResponseWriter, r *http.Request) {
	user := a.contextGetUser(r)
	if err := a.DB.User.Delete(user.ID); err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	http.Redirect(w, r, a.config.ClientUrl, http.StatusSeeOther)
}

type UpdatePasswordInput struct {
	CurrentPassword string `json:"currentPassword"`
	NewPassword     string `json:"newPassword"`
}

// updatePasswordHandler updates the user password.
//
//	@Summary		Update user password
//	@Description	Updates the password of the user associated with the current session.
//	@Tags			user
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Param			input	body		UpdatePasswordInput		true		"Update Password Input"
//	@Success		200		{object}	map[string]interface{}	"message":	"password updated"
//	@Failure		400		{object}	map[string]interface{}	"error":	"bad request"
//	@Router			/v1/user/password/ [patch]
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

	if err := a.writeJSON(w, http.StatusOK, map[string]string{"message": "password updated"}, nil); err != nil {
		a.serverErrorResponse(w, r, err)
	}
}
