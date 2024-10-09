package apires

import (
	"encoding/json"
	"net/http"
)

type Default struct {
	Success bool        `json:"success"`
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Details interface{} `json:"details"`
}

/*
	200: OK
*/

func ResponseSuccess(w http.ResponseWriter, message string, data interface{}, details interface{}) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&Default{
		Success: true,
		Code:    SuccessCode,
		Message: message,
		Data:    data,
		Details: details,
	})
}

func ResponseSuccessDataRetrieved(w http.ResponseWriter, data interface{}, details interface{}) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&Default{
		Success: true,
		Code:    SuccessCode,
		Message: SuccessDataRetrieved,
		Data:    data,
		Details: details,
	})
}

func ResponseSuccessDataCreated(w http.ResponseWriter, data interface{}, details interface{}) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&Default{
		Success: true,
		Code:    SuccessCode,
		Message: SuccessDataCreated,
		Data:    data,
		Details: details,
	})
}

func ResponseSuccessDataUpdated(w http.ResponseWriter, data interface{}, details interface{}) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&Default{
		Success: true,
		Code:    SuccessCode,
		Message: SuccessDataUpdated,
		Data:    data,
		Details: details,
	})
}

func ResponseSuccessDataDeleted(w http.ResponseWriter, data interface{}, details interface{}) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(&Default{
		Success: true,
		Code:    SuccessCode,
		Message: SuccessDataDeleted,
		Data:    data,
		Details: details,
	})
}

/*
	500: Internal Server Error
*/

func ResponseInternalServerError(w http.ResponseWriter, data interface{}, details interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(&Default{
		Success: true,
		Code:    ErrorCodeInternalServer,
		Message: ErrorMessageInternalServer,
		Data:    data,
		Details: details,
	})
}
