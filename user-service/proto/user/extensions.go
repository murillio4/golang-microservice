package go_micro_srv_user

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

//BeforeCreate set id to a randome generated id
func (model *User) BeforeCreate(scope *gorm.Scope) error {
	id := uuid.NewV4()

	return scope.SetColumn("Id", id.String())
}
