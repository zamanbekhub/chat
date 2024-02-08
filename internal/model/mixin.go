package model

import (
	"time"
)

type DeleteMixin struct {
	DeletedAt *time.Time `db:"deleted_at" gorm:"column:deleted_at"`
}

type TimestampMixin struct {
	CreatedAt time.Time `db:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `db:"updated_at" gorm:"column:updated_at"`
}

//func (u *User) BeforeUpdate(tx *gocql.Session) (err error) {
//	if u.Role == "admin" {
//		return errors.New("admin user not allowed to update")
//	}
//	return
//}
//func (m *BaseModel) BeforeInsert(db gocql.Session) error {
//	now := time.Now()
//	if m.CreatedAt.IsZero() {
//		m.CreatedAt = now
//	}
//	if m.UpdatedAt.IsZero() {
//		m.UpdatedAt = now
//	}
//	return nil
//}
//
//func (m *BaseModel) BeforeUpdate(db gocql.Session) error {
//	m.UpdatedAt = time.Now()
//	return nil
//}
