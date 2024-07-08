package controllers

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

func NewDosen(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			nama_dosen := r.FormValue("nama_d")
			nip := r.FormValue("nip")

			_, err := db.Exec("INSERT INTO dosen (nama_dosen, nip) VALUES (?, ?)", nama_dosen, nip)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
			return

		} else if r.Method == http.MethodGet {
			fp := filepath.Join("views", "dosen.html")
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
