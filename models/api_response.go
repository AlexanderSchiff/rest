package models

// APIResponse is the response for the API when an object is not returned
type APIResponse struct {
	Code int `json:"code"`
	ResponseType string `json:"type"`
	Message string `json:"message"`
}