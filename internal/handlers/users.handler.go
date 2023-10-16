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

		responseError := ErrorHandlerResponse{
			Error:   "Internal server error",
			Message: "Fail to connect with database",
		}

		JsonResponse(
			&writer,
			HandlerResponse{
				Code: InitDbErrorCode,
				Data: responseError,
			},
			http.StatusInternalServerError,
		)
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

		responseError := ErrorHandlerResponse{
			Error:   "Internal server error",
			Message: "Fail to execute the resource",
		}

		JsonResponse(
			&writer,
			HandlerResponse{
				Code: DbErrorCode,
				Data: responseError,
			},
			http.StatusInternalServerError,
		)
		return
	}

	responseSuccess := SimpleHandlerResponse{
		Message: "User created successful",
	}

	JsonResponse(
		&writer,
		HandlerResponse{
			Code: SuccessResponseCode,
			Data: responseSuccess,
		},
		http.StatusOK,
	)
}

func FindUser(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	userId := vars["userId"]

	db := database.GetInstance()

	if db.Error != nil {
		log.Default().Println(db.Error)

		responseError := ErrorHandlerResponse{
			Error:   "Internal server error",
			Message: "Fail to execute the resource",
		}

		JsonResponse(
			&writer,
			HandlerResponse{
				Code: DbErrorCode,
				Data: responseError,
			},
			http.StatusInternalServerError,
		)
	}

	var user UserDTO
	query := `SELECT id, name, username, created_at FROM users WHERE id = ?`
	err := db.Inst.QueryRow(query, userId).Scan(&user.Id, &user.Name, &user.Username, &user.Created_at)

	if err != nil {
		log.Default().Println(err)

		responseSuccess := SimpleHandlerResponse{
			Message: "User not found",
		}

		JsonResponse(
			&writer,
			HandlerResponse{
				Code: SuccessResponseCode,
				Data: responseSuccess,
			},
			http.StatusNotFound,
		)
		return
	}

	JsonResponse(
		&writer,
		HandlerResponse{
			Code: SuccessResponseCode,
			Data: user,
		},
		http.StatusOK,
	)
}
