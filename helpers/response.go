package helpers

type ResponseWithData struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithoutData struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func Response(code int, params ResponseParams) any {
	var response any
	var status bool

	if code <= 299 && code >= 200 {
		status = true
	} else {
		status = false
	}

	if params.Data != nil {
		response = &ResponseWithData{
			Status:  status,
			Message: params.Message,
			Data:    params.Data,
		}
	} else {
		response = &ResponseWithoutData{
			Status:  status,
			Message: params.Message,
		}
	}

	return response
}
