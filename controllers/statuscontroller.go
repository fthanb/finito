package controllers

import (
	"database/sql"
	"html/template"
	"net/http"
	"path/filepath"
)

type Stat struct {
	Nama       string
	NIM        string
	Nama_dosen string
	Judul      string
	Status     string
}

func Status(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(`SELECT biodata.nama, biodata.nim, dosen.nama_dosen, proposal.judul FROM biodata 
								INNER JOIN dosen ON biodata.no_reg = dosen.no_reg 
								INNER JOIN proposal ON biodata.no_reg = proposal.no_reg 
								INNER JOIN user
								WHERE biodata.no_reg = user.id;`)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var stats []Stat
		for rows.Next() {
			var stat Stat

			err = rows.Scan(
				&stat.Nama,
				&stat.NIM,
				&stat.Nama_dosen,
				&stat.Judul,
			)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			stats = append(stats, stat)
		}

		fp := filepath.Join("views", "status.html")
		tmpl, err := template.ParseFiles(fp)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		data := make(map[string]any)
		data["stats"] = stats

		err = tmpl.Execute(w, data)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	}
}
