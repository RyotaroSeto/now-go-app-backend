package domain

type Board struct {
	ID           BoardID
	UserID       UserID
	Body         Body
	UsersDetails UsersDetails
}

type BoardID int

func (b BoardID) Num() int {
	return int(b)
}

type Body string

func (b Body) String() string {
	return string(b)
}
