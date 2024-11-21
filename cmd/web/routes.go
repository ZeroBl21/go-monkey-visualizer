package main

import (
	"fmt"
	"net/http"

	"github.com/ZeroBl21/go-monkey-visualizer/internal/repl"
	"github.com/ZeroBl21/go-monkey-visualizer/ui"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.FS(ui.Files))
	mux.Handle("GET /static/", fileServer)

	// HTML Routes
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/tutorial", app.tutorial)

	// Api Routes
	mux.HandleFunc("POST /api/lexer", app.lexerMonkey)

	mux.HandleFunc("POST /api/pratt", app.parserPratt)

	mux.HandleFunc("POST /api/evaluator", app.evaluateMonkey)

	mux.HandleFunc("POST /api/bytecode", app.bytecodeMonkey)

	return mux
}

// Handlers

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "home.html", data)
}

func (app *application) tutorial(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)

	app.render(w, http.StatusOK, "tutorial.html", data)
}

func (app *application) lexerMonkey(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Input string `json:"input"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := newValidator()

	if v.Check(input.Input != "", "input", "must be provided"); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	replInstance := repl.New()
	result := replInstance.ParseTokens(input.Input)

	err := app.writeJSON(w, http.StatusOK, envelope{"result": result}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) parserPratt(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Input string `json:"input"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := newValidator()

	if v.Check(input.Input != "", "input", "must be provided"); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	replInstance := repl.New()
	result := replInstance.ParseAST(input.Input)

	if len(result.Errors) != 0 {
		err := app.writeJSON(w, http.StatusOK, envelope{"result": result.Errors}, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"result": result.Program}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) evaluateMonkey(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Input string `json:"input"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := newValidator()

	if v.Check(input.Input != "", "input", "must be provided"); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	replInstance := repl.New()
	result := replInstance.EvaluateLine(input.Input)

	if len(result.Errors) != 0 {
		err := app.writeJSON(w, http.StatusOK, envelope{"result": result.Errors}, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	err := app.writeJSON(w, http.StatusOK, envelope{"result": result.Evaluate}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) bytecodeMonkey(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Input string `json:"input"`
	}

	if err := app.readJSON(w, r, &input); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := newValidator()

	if v.Check(input.Input != "", "input", "must be provided"); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	replInstance := repl.New()
	result, err := replInstance.CompileToBytecode(input.Input)
	if err != nil {
		fmt.Println(err)
		err := app.writeJSON(w, http.StatusOK, envelope{"result": err.Error()}, nil)
		if err != nil {
			app.serverErrorResponse(w, r, err)
		}

		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"result": result}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
