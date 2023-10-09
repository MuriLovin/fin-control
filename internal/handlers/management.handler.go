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
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)

		response := ErrorHandlerResponse{
			Code:    "INIT_DB_FAIL",
			Message: "Internal server error: failed to connect with database",
		}

		json.NewEncoder(writer).Encode(response)
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
		Message: "Management created successful",
	}

	json.NewEncoder(writer).Encode(response)
}

func FindManagement(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	managementId := vars["managementId"]

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
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		response := SimpleHandlerResponse{
			Code:    "EXEC_DB_SUCCESS",
			Message: "Management not found",
		}

		json.NewEncoder(writer).Encode(response)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	response := HandlerResponse[ManagementDTO]{
		Code: "EXEC_DB_SUCCESS",
		Data: management,
	}

	json.NewEncoder(writer).Encode(response)
}

func FindAllUserManagement(writer http.ResponseWriter, request *http.Request) {
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

	query := "SELECT * FROM management WHERE user_id = ?"
	rows, err := db.Inst.Query(query, userId)
	if err != nil {
		log.Default().Println(err)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)

		response := SimpleHandlerResponse{
			Code:    "EXEC_DB_FAIL",
			Message: "Internal server error: fail to retrieve rows",
		}

		json.NewEncoder(writer).Encode(response)
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
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Header().Add("Content-type", "application/json")

			response := SimpleHandlerResponse{
				Code:    "EXEC_DB_FAIL",
				Message: "Internal server error: fail to retrieve rows",
			}

			json.NewEncoder(writer).Encode(response)
			return
		}

		managements = append(managements, management)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	response := HandlerResponse[[]ManagementDTO]{
		Code: "EXEC_DB_SUCCESS",
		Data: managements,
	}

	json.NewEncoder(writer).Encode(response)
}
