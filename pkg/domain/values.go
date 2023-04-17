package domain

type User struct {
	ID       UserID
	UserName UserName
	Password Password
	Email    Email
}

func NewUser(uID UserID, un UserName, p Password, e Email) *User {
	return &User{
		ID:       uID,
		UserName: un,
		Password: p,
		Email:    e,
	}
}

type UserID int

func (u UserID) Num() int {
	return int(u)
}

type UserName string

func (u UserName) String() string {
	return string(u)
}

type Password string

func (p Password) String() string {
	return string(p)
}

type Email string

func (e Email) String() string {
	return string(e)
}
