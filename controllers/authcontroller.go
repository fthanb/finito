package controllers

import (
	"errors"
	"html/template"
	"net/http"

	"github.com/fthanb/web-pplbo/config"
	"github.com/fthanb/web-pplbo/entities"
	"github.com/fthanb/web-pplbo/libraries"
	"github.com/fthanb/web-pplbo/models"

	"golang.org/x/crypto/bcrypt"
)

type UserInput struct {
	Id       int
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
				"id":   session.Values["id"],
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
			session.Values["id"] = user.Id

			session.Save(r, w)

			http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		}
	}
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := config.Store.Get(r, config.SESSION_ID)
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
