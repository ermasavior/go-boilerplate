package handler

const errInvalidBodyRequest = "invalid body request"

type errorItem struct {
	Message string `json:"message"`
}

type response struct {
	Data   interface{} `json:"data"`
	Errors []errorItem `json:"errors,omitempty"`
}

func newResponse(data interface{}) *response {
	return &response{
		Data:   data,
		Errors: []errorItem{},
	}
}

func newErrorResponse(message string) *response {
	return &response{
		Errors: []errorItem{
			{
				Message: message,
			},
		},
	}
}
