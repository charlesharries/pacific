package main

import "net/http"

// errorResponse sends a generic JSON error response.
func (app *application) errorJSONResponse(w http.ResponseWriter, status int, message interface{}) {
	env := envelope{"error": message}

	// Write the response using the writeJSON() helper. If this happens to return an
	// error then log it, and fall back to sending the client an empty response with a
	// 500 Internal Server Error status code.
	err := app.writeJSON(w, status, env, nil)
	if err != nil {
		w.WriteHeader(500)
	}
}

// The serverErrorResponse() method will be used when our application encounters an
// unexpected problem at runtime. It logs the detailed error message, then uses the
// errorResponse() helper to send a 500 Internal Server Error status code and JSON
// response (containing a generic error message) to the client.
func (app *application) serverErrorJSONResponse(w http.ResponseWriter, err error) {
	message := "the server encountered a problem and could not process your request"
	app.errorJSONResponse(w, http.StatusInternalServerError, message)
}

// The notFoundResponse() method will be used to send a 404 Not Found status code and
// JSON response to the client.
func (app *application) notFoundJSONResponse(w http.ResponseWriter) {
	message := "the requested resource could not be found"

	app.errorJSONResponse(w, http.StatusNotFound, message)
}

// The badRequestResponse() method will be used to send a generic 400 Bad Request
// status code and JSON response to the client.
func (app *application) badRequestJSONResponse(w http.ResponseWriter, err error) {
	app.errorJSONResponse(w, http.StatusBadRequest, err.Error())
}
