package controllers

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Stat struct {
	Id         string
	Nama       string
	NIM        string
	Nama_dosen string
	Judul      string
	Status     string
}

func Status(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			rows, err := db.Query(`SELECT user.id, biodata.nama, biodata.nim, dosen.nama_dosen, proposal.judul FROM biodata 
									INNER JOIN dosen ON biodata.no_reg = dosen.no_reg 
									INNER JOIN proposal ON biodata.no_reg = proposal.no_reg 
									INNER JOIN user
									WHERE biodata.no_reg = user.id;`)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer rows.Close()

			var stats []Stat
			for rows.Next() {
				var stat Stat

				err = rows.Scan(
					&stat.Id,
					&stat.Nama,
					&stat.NIM,
					&stat.Nama_dosen,
					&stat.Judul,
				)
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				stats = append(stats, stat)
			}

			fp := filepath.Join("views", "status.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			data := make(map[string]interface{})
			data["stats"] = stats

			if r.Method == "POST" {
				err := r.ParseForm()
				if err != nil {
					http.Error(w, err.Error(), http.StatusInternalServerError)
					return
				}
				id := r.FormValue("id")
				if id != "" {
					deleteFunc := Delete(db)
					err := deleteFunc(id)
					if err != nil {
						http.Error(w, err.Error(), http.StatusInternalServerError)
						return
					}
					http.Redirect(w, r, "/status", http.StatusSeeOther)
					return
				}
			}
			err = tmpl.Execute(w, data)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}

func DeleteHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		id := r.FormValue("id")
		if id == "" {
			http.Error(w, "ID parameter is missing", http.StatusBadRequest)
			return
		}

		deleteFunc := Delete(db)
		err = deleteFunc(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/status", http.StatusSeeOther)
	}
}

func Delete(db *sql.DB) func(id string) error {
	return func(id string) error {
		tx, err := db.Begin()
		if err != nil {
			return err
		}
		defer func() {
			if err != nil {
				tx.Rollback()
				return
			}
			err = tx.Commit()
			if err != nil {
				return
			}
		}()
		_, err = tx.Exec("DELETE FROM biodata WHERE no_reg = ?", id)
		if err != nil {
			return err
		}
		_, err = tx.Exec("DELETE FROM proposal WHERE no_reg = ?", id)
		if err != nil {
			return err
		}
		_, err = tx.Exec("DELETE FROM dosen WHERE no_reg = ?", id)
		if err != nil {
			return err
		}

		return nil
	}
}
