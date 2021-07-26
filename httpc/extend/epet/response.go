package epet

type (
	Pagination struct {
		Total       int `json:"total"`
		Count       int `json:"count"`
		PerPage     int `json:"per_page"`
		CurrentPage int `json:"current_page"`
		TotalPages  int `json:"total_pages"`
	}
	Trace struct {
		TraceId   string `json:"trace_id"`
		RequestId string `json:"request_id"`
	}
	Meta struct {
		Pagination `json:"pagination"`
		Trace      `json:"trace"`
	}
	Error struct {
		Key   string `json:"key"`
		Error string `json:"error"`
	}
	Response struct {
		StatusCode string      `json:"status_code"`
		Message    string      `json:"message"`
		Data       interface{} `json:"data"`
		Meta       Meta        `json:"meta"`
		Errors     []Error     `json:"errors"`
	}
)
