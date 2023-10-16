package handlers

import (
	"encoding/json"
	"net/http"
)

type HandlerResponse struct {
	Code string      `json:"code"`
	Data interface{} `json:"data"`
}

type SimpleHandlerResponse struct {
	Message string `json:"message"`
}

type ErrorHandlerResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

const InitDbErrorCode = "E_INIT_DB_FAIL"
const SuccessResponseCode = "S_EXEC"
const ErrorResponseCode = "E_FAIL"
const DbSuccessCode = "S_EXEC_DB"
const DbErrorCode = "E_EXEC_DB"

func JsonResponse(writer *http.ResponseWriter, body HandlerResponse, statusCode int) *http.ResponseWriter {
	(*writer).Header().Set("Content-Type", "application/json")
	(*writer).WriteHeader(statusCode)
	json.NewEncoder(*writer).Encode(body)
	return writer
}
