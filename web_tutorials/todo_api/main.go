package main

import (
	"log"
	"net/http"
	"todo_api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello"))
	})

	r.Route("/todos", func(r chi.Router) {
		r.Get("/", handlers.GetTodos)          // GET /todos
		r.Post("/", handlers.CreateTodo)       // POST /todos
		r.Get("/{id}", handlers.GetTodo)       // GET /todos/{id}
		r.Put("/{id}", handlers.UpdateTodo)    // PUT /todos/{id}
		r.Delete("/{id}", handlers.DeleteTodo) // DELETE /todos/{id}
	})

	log.Println("server started on port 3000")

	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}

}
