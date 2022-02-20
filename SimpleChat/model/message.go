package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type LoginResMes struct {
	StatusCode int    `json:"statusCode"`
	Error      string `json:"error"`
}
