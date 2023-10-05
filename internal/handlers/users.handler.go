package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/MuriLovin/fin-control/internal/database"
	"github.com/gorilla/mux"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func AllUsers(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello, World! Thanks for your request\n")
}

func SaveUser(writer http.ResponseWriter, request *http.Request) {
	var user User
	json.NewDecoder(request.Body).Decode(&user)
	db := database.InitDB()

	if db.Error != nil {
		log.Default().Fatal(db.Error)
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Internal server error: failed to connect with database")
		return
	}

	_, err := db.Inst.Exec(
		"INSERT INTO users (name, password, created_at) VALUES (?, ?, ?)",
		user.Name,
		user.Password,
		time.Now(),
	)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Error to save user")
		return
	}

	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, "User created successfully")
}

func FriendHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["friend"]

	user := User{name, "here i am"}
	fmt.Fprintf(writer, "Hello, %s: your password is %s\n", user.Name, user.Password)
}
