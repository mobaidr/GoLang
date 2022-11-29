package main

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

type RequestPayload struct {
	Action string      `json:"action"`
	Auth   AuthPayload `json:"auth,omitempty"`
	Log    LogPayload  `json:"log,omitempty"`
	Mail   MailPayload `json:"mail:omitempty"`
}

type MailPayload struct {
	From string `json:"from"`
	To string `json:"to"`
	Subject string `json:"subject"`
	Message string `json:"message"`
}

type LogPayload struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type AuthPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "Hit the broker",
	}

	_ = app.writeJSON(w, http.StatusAccepted, payload)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload

	err := app.readJSON(w, r, &requestPayload)

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "auth":
		{
			app.authenticate(w, requestPayload.Auth)
		}
	case "log":
		{
			app.logItem(w, requestPayload.Log)
		}
	case "mail" :
		{
			app.sendMail(w, requestPayload.Mail)
		}
	default:
		app.errorJSON(w, errors.New("Unknown action"))
	}
}

func (app *Config) logItem(w http.ResponseWriter, payload LogPayload) {
	// Create a some json to send to Auth micro service.
	jsonData, _ := json.MarshalIndent(payload, "", "\t")

	logServiceURL := "http://logger-service/log"

	request, err := http.NewRequest("POST", logServiceURL, bytes.NewBuffer(jsonData))

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	var payLoad jsonResponse
	payLoad.Error = false
	payLoad.Message = "logger"

	app.writeJSON(w, http.StatusAccepted, payLoad)
}

func (app *Config) authenticate(w http.ResponseWriter, a AuthPayload) {
	// Create a some json to send to Auth micro service.
	jsonData, _ := json.MarshalIndent(a, "", "\t")

	//Call the service
	request, err := http.NewRequest("POST", "http://authentication-service/authenticate", bytes.NewBuffer(jsonData))

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	//make sure we get back the correct status code.
	if response.StatusCode == http.StatusUnauthorized {
		app.errorJSON(w, errors.New("Invalid Credentials"))
		return
	} else if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling auth service"))
		return
	}

	/// Create a variable to read response.Body into
	var jsonFromService jsonResponse

	// Decode the json from the auth service.
	err = json.NewDecoder(response.Body).Decode(&jsonFromService)

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.errorJSON(w, err, http.StatusUnauthorized)
		return
	}

	var payLoad jsonResponse
	payLoad.Error = false
	payLoad.Message = "Authenticated"
	payLoad.Data = jsonFromService.Data

	app.writeJSON(w, http.StatusAccepted, payLoad)
}

func (app *Config) sendMail(w http.ResponseWriter, msg MailPayload) {
	jsonData, _ := json.MarshalIndent(msg,"","\t")

	//call the mail service
	mailServiceUrl := "http://mail-service/send"

	// post to mail service...
	request, err := http.NewRequest("POST", mailServiceUrl,bytes.NewBuffer(jsonData))

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	request.Header.Set("Content-Type","application/json")

	client := http.Client{}

	response, err := client.Do(request)

	if err != nil {
		app.errorJSON(w, err)
		return
	}

	defer response.Body.Close()

	//make sure you get back the right status code.

	if response.StatusCode != http.StatusAccepted {
		app.errorJSON(w, errors.New("error calling mail service"))
		return
	}

	// send back the response
	var payload jsonResponse
	payload.Error = false
	payload.Message = "Message sent to : " + msg.To

	app.writeJSON(w, http.StatusAccepted, payload)
}