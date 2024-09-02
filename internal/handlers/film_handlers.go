package handlers

import (
	"html/template"
	"htmx_go/internal/database"
	"htmx_go/internal/models"
	"net/http"
)

var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func GetFilms(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT * FROM films")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var films []models.Film
	for rows.Next() {
		var film models.Film
		err := rows.Scan(&film.ID, &film.Title, &film.Director)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		films = append(films, film)
	}
	err = tmpl.Execute(w, map[string]interface{}{
		"Films": films,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
