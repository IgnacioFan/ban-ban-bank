package handler

type RegisterUserReq struct {
	Account  string `json:"account"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginUserReq struct {
	Account  string `json:"account"`
	Password string `json:"password"`
	IP       string `json:"ip"`
}
