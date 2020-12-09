package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const appName = "2bitprogrammers/api_addition"
const appVersion = "2007.21a"
const appPort = "1234"

// PayloadJSON contains the variables in which we will perform the operations on
type PayloadJSON struct {
	A int `json:"a"`
	B int `json:"b"`
}

// Add will sum up the values for A+B
func (p *PayloadJSON) Add() int {
	sum := p.A + p.B
	return sum
}

// RequestInfo contains all the data about the initial request
type RequestInfo struct {
	URI     string `json:"uri"`
	Method  string `json:"method"`
	Payload string `json:"payload"`
}

// Response is what we will send back to the user
type Response struct {
	Date       time.Time   `json:"date"`
	StatusCode int         `json:"statusCode"`
	StatusText string      `json:"statusText"`
	Data       string      `json:"data"`
	Errors     string      `json:"errors"`
	Request    RequestInfo `json:"request"`
}

// returnResponse - this will return a json response to the web client
func returnResponse(w http.ResponseWriter, method string, uri string, requestPayload string, status int, statusText string, data string, errors string) {
	sResponse := Response{Date: time.Now(), StatusCode: status, StatusText: statusText, Errors: errors}
	sResponse.Data = data
	sResponse.Request.URI = uri
	sResponse.Request.Method = method
	sResponse.Request.Payload = requestPayload

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Allow cors
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	if method == "OPTIONS" {
		log.Printf("%s %s %d %s", method, uri, http.StatusOK, requestPayload)
		return
	}

	log.Printf("%s %s %d %s", method, uri, status, requestPayload)
	if errors != "" {
		log.Printf("[ERROR] %s - %s", uri, errors)
	}

	joResponse, err := json.Marshal(sResponse)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("[ERROR] Internal Server Error - Failed to parse Json.  %s", err), http.StatusInternalServerError)
		return
	}

	if status == 200 {
		w.Write(joResponse)
		w.Write([]byte("\n\n\n"))
	} else {
		http.Error(w, string(joResponse), status)
	}

}

// handleStatusGet will handle all incoming status requests
func handleStatusGet(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	uri := r.RequestURI

	responsePayload := `{ "healthy": true}`
	returnResponse(w, method, uri, "", http.StatusOK, "OK", responsePayload, "")
}

// handleAdditionPost will handle all incoming /add requests
func handleAdditionPost(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	uri := r.RequestURI
	requestPayload, err := ioutil.ReadAll(r.Body)
	if err != nil {
		returnResponse(w, method, uri, "", http.StatusBadRequest, "Bad Request", "", "Failed to retrieve payload. "+err.Error())
		return
	}

	var joRequest PayloadJSON
	err = json.Unmarshal(requestPayload, &joRequest)
	if err != nil {
		returnResponse(w, method, uri, "", http.StatusInternalServerError, "Internal Server Error", "", "Failed to parse payload json. "+err.Error())
		return
	}

	responsePayload := fmt.Sprintf(`{ "value": %d }`, joRequest.Add())
	returnResponse(w, method, uri, string(requestPayload), http.StatusOK, "OK", responsePayload, "")
}

func main() {
	fmt.Printf("%s v%s\n", appName, appVersion)
	fmt.Println("www.2BitProgrammers.com\nCopyright (C) 2020. All Rights Reserved.\n")
	log.Printf("Starting App on Port %s", appPort)

	http.HandleFunc("/status", handleStatusGet)
	http.HandleFunc("/add", handleAdditionPost)
	log.Fatal(http.ListenAndServe(":"+appPort, nil))
}
