package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/MuriLovin/fin-control/internal/database"
	"github.com/gorilla/mux"
)

type CategoryDTO struct {
	Id         uint      `json:"id"`
	User_id    uint      `json:"user_id"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
}

type CategoryRequestBody struct {
	Name string `json:"name"`
	UserId uint `json:"user_id"`
}

func SaveCategory(writer http.ResponseWriter, request *http.Request) {
	var body CategoryRequestBody
	json.NewDecoder(request.Body).Decode(&body)

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

	query := "INSERT INTO categories (name, user_id) VALUES (?, ?)"
	_, err := db.Inst.Exec(query, body.Name, body.UserId)

	if err != nil {
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
	writer.WriteHeader(http.StatusCreated)

	response := SimpleHandlerResponse{
		Code:    "EXEC_DB_SUCCESS",
		Message: "Category created successful",
	}

	json.NewEncoder(writer).Encode(response)
}

func FindCategory(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	categoryId := vars["categoryId"]

	db := database.GetInstance()

	var category CategoryDTO
	query := "SELECT * FROM categories WHERE id = ?"
	err := db.Inst.QueryRow(query, categoryId).Scan(
		&category.Id,
		&category.User_id,
		&category.Name,
		&category.Created_at,
	)

	if err != nil {
		log.Default().Println(err)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		response := SimpleHandlerResponse{
			Code:    "EXEC_DB_SUCCESS",
			Message: "Category not found",
		}

		json.NewEncoder(writer).Encode(response)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	response := HandlerResponse[CategoryDTO]{
		Code: "EXEC_DB_SUCCESS",
		Data: category,
	}

	json.NewEncoder(writer).Encode(response)
}
