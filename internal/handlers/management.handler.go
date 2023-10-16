package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/MuriLovin/fin-control/internal/database"
	"github.com/gorilla/mux"
)

type ManagementDTO struct {
	Id          uint      `json:"id"`
	User_id     uint      `json:"user_id"`
	Category_id uint      `json:"category_id"`
	Kind        string    `json:"kind"`
	Amount      float32   `json:"amount"`
	Year        string    `json:"year"`
	Month       int8      `json:"month"`
	Created_at  time.Time `json:"created_at"`
}

type ManagementRequestBody struct {
	User_id     uint    `json:"user_id"`
	Category_id uint    `json:"category_id"`
	Kind        string  `json:"kind"`
	Amount      float32 `json:"amount"`
	Year        string  `json:"year"`
	Month       string  `json:"month"`
}

func SaveManagement(writer http.ResponseWriter, request *http.Request) {
	var body ManagementRequestBody
	json.NewDecoder(request.Body).Decode(&body)

	db := database.GetInstance()

	if db.Error != nil {
		log.Default().Println(db.Error)

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

	query := "INSERT INTO management (user_id, category_id, kind, amount, year, month) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := db.Inst.Exec(
		query,
		body.User_id,
		body.Category_id,
		body.Kind,
		body.Amount,
		body.Year,
		body.Month,
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
		Message: "Management created successful",
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

func FindManagement(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	managementId := vars["managementId"]

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
		return
	}

	var management ManagementDTO
	query := "SELECT * FROM management WHERE id = ?"
	err := db.Inst.QueryRow(query, managementId).Scan(
		&management.Id,
		&management.User_id,
		&management.Category_id,
		&management.Kind,
		&management.Amount,
		&management.Year,
		&management.Month,
		&management.Created_at,
	)

	if err != nil {
		responseNotFound := SimpleHandlerResponse{
			Message: "Management not found",
		}

		JsonResponse(
			&writer,
			HandlerResponse{
				Code: DbSuccessCode,
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
			Data: management,
		},
		http.StatusOK,
	)
}

func FindAllUserManagement(writer http.ResponseWriter, request *http.Request) {
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
		return
	}

	query := "SELECT * FROM management WHERE user_id = ?"
	rows, err := db.Inst.Query(query, userId)
	if err != nil {
		log.Default().Println(err)

		responseSuccess := SimpleHandlerResponse{
			Message: "Internal server error: fail to retrieve rows",
		}

		JsonResponse(
			&writer,
			HandlerResponse{
				Code: SuccessResponseCode,
				Data: responseSuccess,
			},
			http.StatusInternalServerError,
		)
		return
	}

	defer rows.Close()

	var managements []ManagementDTO
	for rows.Next() {
		var management ManagementDTO
		err := rows.Scan(
			&management.Id,
			&management.User_id,
			&management.Category_id,
			&management.Kind,
			&management.Amount,
			&management.Year,
			&management.Month,
			&management.Created_at,
		)

		if err != nil {
			log.Default().Println(err)

			responseError := ErrorHandlerResponse{
				Error:   "Internal server error",
				Message: "Fail to retrieve rows",
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

		managements = append(managements, management)
	}

	JsonResponse(
		&writer,
		HandlerResponse{
			Code: SuccessResponseCode,
			Data: managements,
		},
		http.StatusOK,
	)
}
