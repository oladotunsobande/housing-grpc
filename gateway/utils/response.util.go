package utils

import (
	"encoding/json"
	"log"
)

// result defines api response interface
type result struct {
	Success bool        `json:"success"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

// SetResponse converts result struct to map (JSON equivalent)
func (_result result) SetResponse() map[string]interface{} {
	jsonObject, _ := json.Marshal(_result)

	vals := make(map[string]interface{})
	err := json.Unmarshal(jsonObject, &vals)
	if err != nil {
		log.Fatal(err)
	}

	return vals
}

// SuccessResult returns data for successful api call
func SuccessResult(data interface{}) map[string]interface{} {
	var _resultPointer = result{
		Success: true,
		Data:    data,
	}

	return _resultPointer.SetResponse()
}

// ErrorMessage returns error from failed api call
func ErrorMessage(err string) map[string]interface{} {
	var _resultPointer = result{
		Success: false,
		Error:   err,
	}

	return _resultPointer.SetResponse()
}
