package result

type Result struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var m = map[string]string{}

func GetSuccess(msg string) *Result {
	return &Result{
		"0",
		msg,
		nil,
	}
}

func GetSimpleSuccess() *Result {
	return GetSuccess("success")
}

func GetFail(msg string) *Result {
	return &Result{
		"-1",
		msg,
		nil,
	}
}

func GetSimpleFail() *Result {
	return GetFail("fail")
}
