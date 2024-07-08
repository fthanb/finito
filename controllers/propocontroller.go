package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

func Upload(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			fp := filepath.Join("views", "proposal.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if err := tmpl.Execute(w, nil); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		if r.Method == http.MethodPost {
			if err := r.ParseMultipartForm(10 << 20); err != nil {
				http.Error(w, "Unable to parse form", http.StatusBadRequest)
				return
			}

			file, handler, err := r.FormFile("myFile")
			if err != nil {
				http.Error(w, "Error retrieving the file", http.StatusBadRequest)
				return
			}
			defer file.Close()

			noReg := r.FormValue("no_reg")

			fmt.Printf("Uploaded File: %+v\n", handler.Filename)
			fmt.Printf("File Size: %+v\n", handler.Size)
			fmt.Printf("MIME Header: %+v\n", handler.Header)

			fileBytes, err := ioutil.ReadAll(file)
			if err != nil {
				http.Error(w, "Error reading the file", http.StatusInternalServerError)
				return
			}

			_, err = db.Exec("INSERT INTO proposal (judul, konten, no_reg) VALUES (?, ?, ?)", handler.Filename, fileBytes, noReg)
			if err != nil {
				http.Error(w, "Error inserting into the database", http.StatusInternalServerError)
				log.Println(err)
				return
			}

			http.Redirect(w, r, "/dashboard", http.StatusMovedPermanently)
			return
		}

		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
