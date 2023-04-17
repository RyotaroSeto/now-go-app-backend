package infrastructure

// type UserRepository struct {
// 	db *DB
// }

// var _ domain.UserRepository = new(UserRepository)

// func NewUserRepository() *UserRepository {
// 	return &UserRepository{GetDB()}
// }

// func (u *UserRepository) conn(ctx context.Context) *gorm.DB {
// 	tx, ok := ctx.Value(txKey).(*gorm.DB)
// 	if ok && tx != nil {
// 		return tx
// 	}

// 	return u.db.Session(&gorm.Session{})
// }

// type Users struct {
// 	ID       int    `json:"id" gorm:"primaryKey"`
// 	UserName string `json:"user_name" gorm:"user_name"`
// 	Password string `json:"password" gorm:"password"`
// 	Email    string `json:"email" gorm:"email"`
// }

// type UsersDatail struct {
// 	ID          int       `json:"id" gorm:"primaryKey"`
// 	UserID      int       `json:"user_id" gorm:"user_id"`
// 	DateOfBirth time.Time `json:"date_of_birth" gorm:"date_of_birth"`
// 	Gender      string    `json:"gender" gorm:"gender"`
// 	Residence   string    `json:"residence" gorm:"residence"`
// 	Occupation  string    `json:"occupation" gorm:"occupation"`
// 	Height      int       `json:"height" gorm:"height"`
// 	Weight      int       `json:"weight" gorm:"weight"`
// }

// func (u *UserRepository) GetProfile(ctx context.Context, uID domain.UserID) error {
// 	// u.db.Model(&Users{}).Preload("UsersDatail")
// 	us := Users{}
// 	err := u.db.Model(&us)
// 	if err != nil {
// 		log.Println("失敗")
// 	}
// 	log.Println(us)
// 	return nil
// }
