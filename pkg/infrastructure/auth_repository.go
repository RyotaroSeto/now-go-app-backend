package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"log"
	"now-go-kon/pkg/domain"
	"time"

	"gorm.io/gorm"
)

type Session struct {
	ID           string    `gorm:"column:id;primaryKey,omitempty"`
	UserName     string    `gorm:"column:user_name"`
	RefreshToken string    `gorm:"column:refresh_token"`
	UserAgent    string    `gorm:"column:user_agent"`
	ClientIP     string    `gorm:"column:client_ip"`
	IsBlocked    bool      `gorm:"column:is_blocked"`
	ExpiresDate  time.Time `gorm:"column:expires_date"`
	CreatedDate  time.Time `gorm:"column:created_date;autoCreateTime"`
}

func (s *Session) toEntity() *domain.Session {
	session := &domain.Session{
		SessionID:    s.ID,
		UserName:     s.UserName,
		RefreshToken: s.RefreshToken,
		UserAgent:    s.UserAgent,
		ClientIP:     s.ClientIP,
		IsBlocked:    s.IsBlocked,
		ExpiresDate:  s.ExpiresDate,
	}

	return session
}

func (u *Session) fromEntity(e *domain.Session) {
	u.ID = e.SessionID
	u.UserName = e.UserName
	u.RefreshToken = e.RefreshToken
	u.UserAgent = e.UserAgent
	u.ClientIP = e.ClientIP
	u.IsBlocked = e.IsBlocked
	u.ExpiresDate = e.ExpiresDate
}

func (m *Session) bindEntity(e *domain.Session) {
	ue := m.toEntity()
	e.SessionID = ue.SessionID
	e.UserName = ue.UserName
	e.RefreshToken = ue.RefreshToken
	e.UserAgent = ue.UserAgent
	e.ClientIP = ue.ClientIP
	e.IsBlocked = ue.IsBlocked
	e.ExpiresDate = ue.ExpiresDate
}

type AuthRepository struct {
	db *DB
}

var _ domain.AuthRepository = new(AuthRepository)

func NewAuthRepository() *AuthRepository {
	return &AuthRepository{GetDB()}
}

func (r *AuthRepository) conn(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(txKey).(*gorm.DB)
	if ok && tx != nil {
		return tx
	}

	return r.db.Session(&gorm.Session{})
}
func (r *AuthRepository) UserGet(ctx context.Context, email domain.Email) (*domain.User, error) {
	us := Users{}
	q := Users{Email: email.String()}
	res := r.conn(ctx).Where(&q).First(&us)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			msg := fmt.Sprintf("Email: %s is not found", email.String())
			return nil, errors.New(msg)
		}
		return nil, err
	}

	return us.toEntity(), nil
}

func (u *AuthRepository) SessionCreate(ctx context.Context, session *domain.Session) (*domain.Session, error) {
	var b Session
	b.fromEntity(session)

	if err := u.conn(ctx).Create(&b).Error; err != nil {
		log.Println(err)
		return nil, errors.New(err.Error())
	}

	return b.toEntity(), nil
}

func (u *AuthRepository) SessionDelete(ctx context.Context, userName domain.UserName) error {
	var s Session
	if err := u.conn(ctx).Where(&Session{UserName: userName.String()}).Delete(&s).Error; err != nil {
		log.Println(err)
		return errors.New(err.Error())
	}
	return nil
}

func (r *AuthRepository) PasswordAuth(ctx context.Context, uID domain.UserID, password domain.Password) error {
	// パスワード認証の処理は割愛
	return nil
}
