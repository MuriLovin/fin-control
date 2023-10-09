package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/MuriLovin/fin-control/internal/database"
	"github.com/gorilla/mux"
)

type User struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDTO struct {
	Id         uint32    `json:"id"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Created_at time.Time `json:"created_at"`
}

func SaveUser(writer http.ResponseWriter, request *http.Request) {
	var user User
	json.NewDecoder(request.Body).Decode(&user)
	db := database.GetInstance()

	if db.Error != nil {
		log.Default().Fatal(db.Error)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)

		response := ErrorHandlerResponse{
			Code:    "INIT_DB_FAIL",
			Message: "Internal server error: failed to connect with database",
		}

		json.NewEncoder(writer).Encode(response)
		return
	}

	_, err := db.Inst.Exec(
		"INSERT INTO users (name, username, password) VALUES (?, ?, ?)",
		user.Name,
		user.Username,
		user.Password,
	)

	if err != nil {
		log.Default().Println(err)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)

		response := ErrorHandlerResponse{
			Code:    "EXEC_DB_FAIL",
			Message: "Internal server error: fail to execute the resource",
		}

		json.NewEncoder(writer).Encode(response)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	response := SimpleHandlerResponse{
		Code:    "EXEC_DB_SUCCESS",
		Message: "User created successful",
	}

	json.NewEncoder(writer).Encode(response)
}

func FindUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["userId"]

	db := database.GetInstance()

	if db.Error != nil {
		log.Default().Println(db.Error)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)

		response := ErrorHandlerResponse{
			Code:    "EXEC_DB_FAIL",
			Message: "Internal server error: fail to execute the resource",
		}

		json.NewEncoder(writer).Encode(response)
		return
	}

	var user UserDTO
	query := `SELECT id, name, username, created_at FROM users WHERE id = ?`
	err := db.Inst.QueryRow(query, userId).Scan(&user.Id, &user.Name, &user.Username, &user.Created_at)

	if err != nil {
		log.Default().Println(err)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		response := SimpleHandlerResponse{
			Code:    "EXEC_DB_SUCCESS",
			Message: "User not found",
		}

		json.NewEncoder(writer).Encode(response)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	response := HandlerResponse[UserDTO]{
		Code: "EXEC_DB_SUCCESS",
		Data: user,
	}

	json.NewEncoder(writer).Encode(response)
}
