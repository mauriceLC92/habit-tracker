package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/", app.healthCheck)
	// router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)
	// router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreate)
	// router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreatePost)

	standard := alice.New(app.recoverPanic, app.logRequest)
	return standard.Then(router)
}
