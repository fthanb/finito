package entities

type User struct {
	Id        int64
	Nama      string `validate:"required" label:"Nama"`
	Nim       string `validate:"required,gte=14,isunique=user-nim"`
	Password  string `validate:"required,gte=6"`
	Cpassword string `validate:"required,eqfield=Password" label:"Konfirmasi Password"`
}
