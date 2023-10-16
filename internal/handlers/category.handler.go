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
	Name   string `json:"name"`
	UserId uint   `json:"user_id"`
}

func SaveCategory(writer http.ResponseWriter, request *http.Request) {
	var body CategoryRequestBody
	json.NewDecoder(request.Body).Decode(&body)

	db := database.GetInstance()

	if db.Error != nil {
		log.Default().Fatal(db.Error)

		responseError := ErrorHandlerResponse{
			Error:   "Internal server error",
			Message: "Failed to connect with database",
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

	query := "INSERT INTO categories (name, user_id) VALUES (?, ?)"
	_, err := db.Inst.Exec(query, body.Name, body.UserId)

	if err != nil {
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
		Message: "Category created successful",
	}

	JsonResponse(&writer,
		HandlerResponse{
			Code: SuccessResponseCode,
			Data: responseSuccess,
		},
		http.StatusCreated,
	)
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

		responseNotFound := SimpleHandlerResponse{
			Message: "Category not found",
		}

		JsonResponse(&writer,
			HandlerResponse{
				Code: SuccessResponseCode,
				Data: responseNotFound,
			},
			http.StatusNotFound,
		)
		return
	}

	JsonResponse(
		&writer,
		HandlerResponse{
			Code: SuccessResponseCode,
			Data: category,
		},
		http.StatusOK,
	)
}
