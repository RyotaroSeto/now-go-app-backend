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

// func (u *User) ReloadUserID() {
// 	u.ID = NewUserID()
// }

type UserName string

type Password string

type Email string
