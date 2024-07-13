package main

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/fthanb/web-pplbo/config"
	controller "github.com/fthanb/web-pplbo/controllers"
)

func main() {
	db, err := config.DBConn()
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/dashboard", controller.Index)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/profil", controller.NewMahasiswa(db))
	http.HandleFunc("/dosen", controller.NewDosen(db))
	http.HandleFunc("/proposal", controller.Upload(db))
	http.HandleFunc("/status", controller.Status(db))
	http.HandleFunc("/edit", controller.EditAndUpdate(db))

	fmt.Println("GAS")
	http.ListenAndServe(":8000", nil)
}
