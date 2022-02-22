package Model

import (
	"gin_demo/app/global/variable"
	"gorm.io/gorm"
	"time"
)

type (
	// UserModel interface {
	// 	Inserts(data []User) ([]uint, error)
	// 	Insert(data User) error
	// 	FindOne(map[string]interface{}) (User, error)
	// 	UpdateById(id uint, data map[string]interface{}) error
	// 	DeleteById(id uint) error
	// 	FindAllByIds(ids []uint, field string) ([]User, error)
	// 	UpdateByIds(ids []uint, data map[string]interface{}) error
	// 	FindCountByTel(tel string) int64
	// 	FindInfoByTel(tel string) (User, error)
	// }

	defaultUserModel struct {
		Db    *gorm.DB
		Table string
	}

	User struct {
		BaseModel `json:"-"`
		Name      string `form:"username" json:"user" comment:"用户名" uri:"user" xml:"user" `
		Tel       string `form:"tel" json:"tel" uri:"tel" xml:"tel" comment:"电话" binding:"required" `
		Password  string `gorm:"column:password" form:"password" json:"password" uri:"password" xml:"password"`
	}
)

func NewUserModel() *defaultUserModel {
	return &defaultUserModel{
		Db:    UseDbConn(variable.ConfigGormv2Yml.GetString("Gormv2.UseDbType")),
		Table: "`user`",
	}
}

func (d defaultUserModel) Inserts(users []User) ([]uint, error) {
	err := d.Db.Create(&users).Error
	if err != nil {
		return nil, err
	}
	var id []uint
	for _, user := range users {
		id = append(id, user.ID)
	}
	return id, nil
}

func (d defaultUserModel) Insert(user User) error {
	err := d.Db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (d defaultUserModel) FindOne(where map[string]interface{}) (User, error) {
	var user User
	err := d.Db.Where(where).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (d defaultUserModel) FindInfoByTel(tel string) (User, error) {
	// var user User
	user := User{}
	err := d.Db.Table(d.Table).Where("tel=?", tel).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (d defaultUserModel) FindCountByTel(tel string) int64 {
	var count int64
	d.Db.Table(d.Table).Where("tel=?", tel).Count(&count)
	return count
}

func (d defaultUserModel) UpdateById(id uint, data map[string]interface{}) error {
	data["updated_at"] = time.Now().Format(variable.DateFormat)
	panic("implement me")
}

func (d defaultUserModel) DeleteById(id uint) error {
	panic("implement me")
}

func (d defaultUserModel) FindAllByIds(ids []uint, field string) ([]User, error) {
	panic("implement me")
}

func (d defaultUserModel) UpdateByIds(ids []uint, data map[string]interface{}) error {
	d.Db.Table(d.Table).Where("id = ?", ids).Updates(data)
	panic("implement me")
}
