package dao

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/go-webauthn/webauthn/webauthn"
	"gorm.io/gorm"
	"time"
)

var (
	ErrDuplicateUsername = errors.New("用户名冲突")
)

type UserDAO interface {
	Insert(ctx context.Context, u User) error
	UpdateById(ctx context.Context, entity User) error
	FindByUsername(ctx context.Context, username string) (User, error)
}

type GORMUserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) UserDAO {
	return &GORMUserDAO{
		db: db,
	}
}

func (dao *GORMUserDAO) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli()
	u.Ctime = now
	u.Utime = now
	err := dao.db.WithContext(ctx).Create(&u).Error
	if me, ok := err.(*mysql.MySQLError); ok {
		const duplicateErr uint16 = 1062
		if me.Number == duplicateErr {
			return ErrDuplicateUsername
		}
	}
	return err
}

func (dao *GORMUserDAO) UpdateById(ctx context.Context, entity User) error {
	return dao.db.WithContext(ctx).Model(&entity).Where("id = ?", entity.Id).
		Updates(map[string]any{
			"utime":       time.Now().UnixMilli(),
			"credentials": entity.Credentials,
		}).Error
}

type User struct {
	Id int64 `gorm:"primaryKey,autoIncrement"`
	// 代表这是一个可以为 NULL 的列
	//Email    *string
	Email    sql.NullString `gorm:"unique"`
	Password string
	Username string `gorm:"type=varchar(128),unique"`
	Nickname string `gorm:"type=varchar(128)"`
	// YYYY-MM-DD
	Birthday int64
	// 代表这是一个可以为 NULL 的列
	Phone       sql.NullString `gorm:"unique"`
	Ctime       int64
	Utime       int64
	Credentials ColumnCredentials `json:"credentials,omitempty" gorm:"type:VARCHAR(4096)"`
}

type ColumnCredentials []webauthn.Credential

func (c ColumnCredentials) Value() (driver.Value, error) {
	str, err := json.Marshal(c)
	return driver.Value(str), err
}

// Scan when the DB driver reads from the DB
func (c *ColumnCredentials) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("CredentialColumn value is not []byte")
	}
	var credential []webauthn.Credential
	err := json.Unmarshal(b, &credential)
	*c = credential

	return err
}

func (dao *GORMUserDAO) FindByUsername(ctx context.Context, username string) (User, error) {
	var u User
	err := dao.db.WithContext(ctx).Where("username=?", username).First(&u).Error
	return u, err
}
