package interfaces

type IAPIResponse struct {
	Error   bool        `json:"error"`
	Result  interface{} `json:"result"`
	Message string      `json:"message"`
}
