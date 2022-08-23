package main

import (
	"finalproject/data"
	"fmt"
	"html/template"
	"net/http"
)

func (app *Config) HomePage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "home.page.gohtml", nil)
}

func (app *Config) LoginPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "login.page.gohtml", nil)
}

func (app *Config) PostLoginPage(w http.ResponseWriter, r *http.Request) {
	_ = app.Session.RenewToken(r.Context())

	// Parse Form post
	err := r.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	//Get email & password from form
	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := app.Models.User.GetByEmail(email)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Invalid Credentials")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	//Check password
	validPassword, err := user.PasswordMatches(password)
	if err != nil {
		app.Session.Put(r.Context(), "error", "Invalid Credentials")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if !validPassword {
		msg := Message{
			To:      email,
			Subject: "Failed login attempt",
			Data:    "Invalid login attempt",
		}

		app.sendMail(msg)

		app.Session.Put(r.Context(), "error", "Invalid Credentials")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	//okay. log user in
	app.Session.Put(r.Context(), "userID", user.ID)
	app.Session.Put(r.Context(), "user", user)

	app.Session.Put(r.Context(), "flash", "Successful Login !!!")
	//Redirect the user.
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (app *Config) Logout(w http.ResponseWriter, r *http.Request) {
	//Clean up session
	_ = app.Session.Destroy(r.Context())
	_ = app.Session.RenewToken(r.Context())

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *Config) RegisterPage(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "register.page.gohtml", nil)
}

func (app *Config) PostRegisterPage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.ErrorLog.Println(err)
	}

	// TODO - Validate data.

	// Create User
	u := data.User{
		Email:     r.Form.Get("email"),
		FirstName: r.Form.Get("first-name"),
		LastName:  r.Form.Get("last-name"),
		Password:  r.Form.Get("password"),
		Active:    0,
		IsAdmin:   0,
	}

	_, err = u.Insert(u)
	if err != nil {
		app.Session.Put(r.Context(), "error", "unable to create user.")
		http.Redirect(w, r, "/register", http.StatusSeeOther)
		return
	}

	//send an activation email.
	url := fmt.Sprintf("http://localhost/activate?email=%s", u.Email)
	signedUrl := GenerateTokenFromString(url)
	app.InfoLog.Println(signedUrl)
	msg := Message{
		To:       u.Email,
		Subject:  "Activate your email",
		Template: "confirmation-email",
		Data:     template.HTML(signedUrl),
	}

	app.sendMail(msg)
	app.Session.Put(r.Context(), "flash", "Confirmation email sent, Check your email")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *Config) ActivateAccount(w http.ResponseWriter, r *http.Request) {
	// Validate the URL
	url := r.RequestURI
	testUrl := fmt.Sprintf("http://localhost%s", url)
	ok := VerifyToken(testUrl)

	if !ok {
		app.Session.Put(r.Context(), "error", "Invalid Token")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Activate account.
	u, err := app.Models.User.GetByEmail(r.URL.Query().Get("email"))
	if err != nil {
		app.Session.Put(r.Context(), "error", "No User Found")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	u.Active = 1

	err = u.Update()
	if err != nil {
		app.Session.Put(r.Context(), "error", "Unable to update user")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	app.Session.Put(r.Context(), "flash", "Account activated, You can now log in.")
	http.Redirect(w, r, "/login", http.StatusSeeOther)

	// Generate an invoice.

	//send an email with attachments.

	//send an email with an invoice attached.
}

func (app *Config) SubscribeToPlan(writer http.ResponseWriter, request *http.Request) {

	// Get the ID of the plan that is choosen
	// Get the plan from Database
	// Get the user from the session
	// generate an invoice
	// send an email with invoice attached.
	// generate a manual.
	// send an email with the manual attached.

	//subscribe the user to account

	// redirect
}

func (app *Config) ChooseSubscription(w http.ResponseWriter, r *http.Request) {

	if !app.Session.Exists(r.Context(), "userID") {
		app.Session.Put(r.Context(), "warning", "You must login to see this page.")
		http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
	}

	plans, err := app.Models.Plan.GetAll()
	if err != nil {
		app.ErrorLog.Println(err)
		return
	}

	dataMap := make(map[string]interface{})
	dataMap["plans"] = plans

	app.render(w, r, "plans.page.gohtml", &TemplateData{
		DataMap: dataMap,
	})
}


