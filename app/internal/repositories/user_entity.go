package repositories

type UserEntity struct {
	Id       int
	Login    string
	Password string
	Salt     string
}
