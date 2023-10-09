package handlers

type HandlerResponse[T any] struct {
	Code string `json:"code"`
	Data T      `json:"data"`
}

type SimpleHandlerResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ErrorHandlerResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// func JsonResponse(writer *http.ResponseWriter, statusCode int) http.ResponseWriter {
// 	(*writer).Header().Set("Content-Type", "application/json")
// 	(*writer).WriteHeader(statusCode)
// }
