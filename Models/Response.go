package Models

type Response struct {
	Code   int    `json:"code"`   // Response Code
	Status string `json:"status"` // Response Status
	Data   any    `json:"data"`   //Response Data
}
