package controllers

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func NewMahasiswa(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			nama := r.FormValue("nama")
			nim := r.FormValue("nim")
			alamat := r.FormValue("alamat")
			no_telp := r.FormValue("no")

			_, err := db.Exec("INSERT INTO biodata (nama, nim, alamat, no_telp) VALUES (?, ?, ?, ?)", nama, nim, alamat, no_telp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
			return

		} else if r.Method == http.MethodGet {
			fp := filepath.Join("views", "profil.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if err := tmpl.Execute(w, nil); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
	}
}
