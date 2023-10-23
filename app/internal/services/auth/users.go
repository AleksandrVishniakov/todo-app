package auth

type RequestUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type ResponseUser struct {
	Id    int    `json:"id"`
	Login string `json:"login"`
}
