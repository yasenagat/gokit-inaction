package user

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
	Age      int    `json:"age"`
	Sid      string `json:"sid"`
	Balance  string `json:"balance"`
	Phone    string `json:"phone"`
}

func (u User) String() string {
	return u.Id
}

type Service interface {
	Login(username, pwd string) (User, error)
	UpdatePhone(username, phone string) error
	GetUser(username string) (User, error)
}

type UserService struct {
}

func (UserService) Login(username, pwd string) (User, error) {
	return CheckLogin(username, pwd)
}

func (UserService) UpdatePhone(username, phone string) error {
	return UpdatePhone(username, phone)
}

func (UserService) GetUser(username string) (User, error) {
	return GetUser(username)
}
