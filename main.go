package main

import (
	"fmt"
	"net/http"

	authcontroller "github.com/fthanb/web-pplbo/controllers"
)

func main() {
	http.HandleFunc("/dashboard", authcontroller.Index)
	http.HandleFunc("/proposal", authcontroller.Up)
	http.HandleFunc("/skripsi", authcontroller.UpSc)
	http.HandleFunc("/login", authcontroller.Login)
	http.HandleFunc("/logout", authcontroller.Logout)
	http.HandleFunc("/register", authcontroller.Register)
	fmt.Println("GAS")
	http.ListenAndServe(":8000", nil)
}
