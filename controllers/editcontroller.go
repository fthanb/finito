package controllers

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
)

type Biodata struct {
	No_reg  string
	Nama    string
	NIM     string
	Alamat  string
	No_telp string
}

type Dosen struct {
	No_reg     string
	Nama_dosen string
	NIP        string
}

func EditAndUpdate(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		entityType := r.URL.Query().Get("type")
		id := r.URL.Query().Get("id")
		if id == "" || entityType == "" {
			http.Error(w, "No id or type provided", http.StatusBadRequest)
			return
		}

		if r.Method == http.MethodGet {
			var tmplData interface{}
			var query string

			switch entityType {
			case "biodata":
				query = "SELECT no_reg, nama, nim, alamat, no_telp FROM biodata WHERE no_reg = ?"
				var biodata Biodata
				err := db.QueryRow(query, id).Scan(&biodata.No_reg, &biodata.Nama, &biodata.NIM, &biodata.Alamat, &biodata.No_telp)
				if err != nil {
					if err == sql.ErrNoRows {
						http.Error(w, "No record found with the provided id", http.StatusNotFound)
					} else {
						http.Error(w, err.Error(), http.StatusInternalServerError)
					}
					return
				}
				tmplData = biodata
			case "dosen":
				query = "SELECT no_reg, nama_dosen, nip FROM dosen WHERE no_reg = ?"
				var dosen Dosen
				err := db.QueryRow(query, id).Scan(&dosen.No_reg, &dosen.Nama_dosen, &dosen.NIP)
				if err != nil {
					if err == sql.ErrNoRows {
						http.Error(w, "No record found with the provided id", http.StatusNotFound)
					} else {
						http.Error(w, err.Error(), http.StatusInternalServerError)
					}
					return
				}
				tmplData = dosen
			default:
				http.Error(w, "Invalid type", http.StatusBadRequest)
				return
			}

			fp := filepath.Join("views", "edit.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			data := map[string]interface{}{
				"Entity": tmplData,
				"Type":   entityType,
			}

			err = tmpl.Execute(w, data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		} else if r.Method == http.MethodPost {
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			switch entityType {
			case "biodata":
				nama := r.FormValue("nama")
				nim := r.FormValue("nim")
				alamat := r.FormValue("alamat")
				noTelp := r.FormValue("no_telp")

				_, err := db.Exec("UPDATE biodata SET nama = ?, nim = ?, alamat = ?, no_telp = ? WHERE no_reg = ?", nama, nim, alamat, noTelp, id)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			case "dosen":
				namaDosen := r.FormValue("nama_dosen")
				nip := r.FormValue("nip")

				_, err := db.Exec("UPDATE dosen SET nama_dosen = ?, nip = ? WHERE no_reg = ?", namaDosen, nip, id)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
			default:
				http.Error(w, "Invalid type", http.StatusBadRequest)
				return
			}

			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	}
}
