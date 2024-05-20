package controllers

import (
	"errors"
	"html/template"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/fthanb/web-pplbo/config"
	"github.com/fthanb/web-pplbo/entities"
	"github.com/fthanb/web-pplbo/libraries"
	"github.com/fthanb/web-pplbo/models"

	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Nama     string `validate:"required"`
	Nim      string `validate:"required"`
	Password string `validate:"required"`
}

var userModel = models.NewUserModel()
var validation = libraries.NewValidation()

func Index(w http.ResponseWriter, r *http.Request) {

	session, _ := config.Store.Get(r, config.SESSION_ID)

	if len(session.Values) == 0 {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {

		if session.Values["loggedIn"] != true {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		} else {
			data := map[string]interface{}{
				"nama": session.Values["nama"],
			}
			temp, _ := template.ParseFiles("views/index.html")
			temp.Execute(w, data)
		}
	}

}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		temp, _ := template.ParseFiles("views/login.html")
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		r.ParseForm()
		UserInput := &UserInput{
			Nama:     r.Form.Get("nama"),
			Nim:      r.Form.Get("nim"),
			Password: r.Form.Get("password"),
		}

		var user entities.User
		userModel.Where(&user, "nim", UserInput.Nim)
		var message error
		if user.Nim == "" {
			message = errors.New("NIM/Password Salah")
		} else {
			errPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(UserInput.Password))
			if errPassword != nil {
				message = errors.New("NIM/Password Salah")
			}
		}

		if message != nil {

			data := map[string]interface{}{
				"error": message,
			}

			temp, _ := template.ParseFiles("views/login.html")
			temp.Execute(w, data)
		} else {
			//create session
			session, _ := config.Store.Get(r, config.SESSION_ID)

			session.Values["loggedIn"] = true
			session.Values["nama"] = user.Nama

			session.Save(r, w)

			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
	// delete session
	session.Options.MaxAge = -1
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		temp, err := template.ParseFiles("views/register.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		temp.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user := entities.User{
			Nama:      r.Form.Get("Nama"),
			Nim:       r.Form.Get("Nim"),
			Password:  r.Form.Get("Password"),
			Cpassword: r.Form.Get("Cpassword"),
		}

		errorMessages := validation.Struct(user)
		if errorMessages != nil {
			data := map[string]interface{}{
				"validation": errorMessages,
				"user":       user,
			}
			temp, err := template.ParseFiles("views/register.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			temp.Execute(w, data)
		} else {
			hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			user.Password = string(hashPassword)

			userModel.Create(user)

			data := map[string]interface{}{
				"pesan": "Registrasi berhasil",
			}
			temp, err := template.ParseFiles("views/register.html")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			temp.Execute(w, data)
		}
	}
}

var (
	tempUp = template.Must(template.ParseFiles("views/proposal.html"))
)

func Up(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		handleUpload(w, r)
		return
	}
	tempUp.ExecuteTemplate(w, "proposal.html", nil)
}
func handleUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm((10 << 20))

	file, fileHeader, err := r.FormFile("pdf")
	if err != nil {
		http.Error(w, "Gagal mengunggah file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	filename := path.Base(fileHeader.Filename)
	dest, err := os.Create("C:/Users/Lenovo/Downloads/" + filename)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer dest.Close()

	if _, err = io.Copy(dest, file); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("PROPOSAL UPLOADED"))
}

var (
	tempSc = template.Must(template.ParseFiles("views/skripsi.html"))
)

func UpSc(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		handleUploadSc(w, r)
		return
	}
	tempSc.ExecuteTemplate(w, "skripsi.html", nil)
}
func handleUploadSc(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm((10 << 20))

	file, fileHeader, err := r.FormFile("pdf")
	if err != nil {
		http.Error(w, "Gagal mengunggah file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	filename := path.Base(fileHeader.Filename)
	dest, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer dest.Close()

	if _, err = io.Copy(dest, file); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("SKRIPSI UPLOADED"))
}

func Eprof() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fp := "views/eprof.html"
		temp, err := template.ParseFiles(fp)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = temp.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
