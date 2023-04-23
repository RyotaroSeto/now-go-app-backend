package domain

type Like struct {
	UserID      UserID
	LikedUserID UserID
	Status      Status
	MessageBody MessageBody
}

type MessageBody string

func (m MessageBody) String() string {
	return string(m)
}
