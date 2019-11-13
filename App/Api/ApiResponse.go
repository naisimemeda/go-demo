package Api

type ApiData struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Success(result interface{}) ApiData {
	return ApiData{201, result}
}

func Failed(error string, code int) Error {
	return Error{code, error}
}
