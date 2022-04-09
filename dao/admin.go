package dao

import (
	"errors"
	"github.com/e421083458/yqb_monitor/dto"
	"github.com/e421083458/yqb_monitor/public"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
	"time"
)

type Admin struct {
	Id        int       `json:"id" gorm:"primary_key" description:"自增主键"`
	UserName  string    `json:"user_name" gorm:"column:user_name" description:"管理员用户名"`
	Salt      string    `json:"salt" gorm:"column:salt" description:"盐"`
	Password  string    `json:"password" gorm:"column:password" description:"密码"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"创建时间"`
	IsDelete  int       `json:"is_delete" gorm:"column:is_delete" description:"是否删除"`
}

func (a *Admin) TableName() string {
	return "gateway_admin"
}

func(a *Admin) Find(c *gin.Context, tx *gorm.DB, search *Admin) (*Admin, error) {
	out := &Admin{}
	err := tx.WithContext(c).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (a *Admin) Save(c *gin.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(a).Error
}

func (a *Admin) LoginCheck(c *gin.Context, tx *gorm.DB, params *dto.AdminLoginInput) (*Admin, error) {
	adminInfo, err := a.Find(c, tx, (&Admin{UserName: params.UserName, IsDelete: 0}))
	if err != nil {
		return nil, errors.New("No username ")
	}
	saltPassword := public.GenSaltPassword(adminInfo.Salt, params.Password)
	if adminInfo.Password != saltPassword {
		return nil, errors.New("password is err, please input again")
	}
	return adminInfo, nil
}
