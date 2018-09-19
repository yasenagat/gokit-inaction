package user

const (
	SVRNAME = "yuser"
)


type User struct {
	Username string `json:"username"`
	Id       string `json:"id"`
}
