package web

type errorBody struct {
	Error errorDetails `json:"error"`
}

type errorDetails struct {
	Code string `json:"code"`
}
