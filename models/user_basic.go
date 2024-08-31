package models

import (
	"fmt"
	"ginchat/utils"
	"log"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string `valid:"matches(^1[3-9]{1}\\d{9}$)"`
	Email         string `valid:"email"`
	Avatar        string //头像
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     time.Time `gorm:"default:NULL"`
	HeartbeatTime time.Time `gorm:"default:NULL"`
	LoginOutTime  time.Time `gorm:"column:login_out_time;default:NULL" json:"login_out_time"`
	IsLogout      bool
	DeviceInfo    string
}

func FindUserByNameAndPwd(name string, password string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ? and pass_word=?", name, password).First(&user)

	//token加密
	str := fmt.Sprintf("%d", time.Now().Unix())
	temp := utils.MD5Encode(str)
	utils.DB.Model(&user).Where("id = ?", user.ID).Update("identity", temp)
	return user
}

func GetUserList() []*UserBasic {
	data := make([]*UserBasic, 10)
	utils.DB.Find(&data)
	for _, v := range data {
		log.Println(v)
	}
	return data
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("Name = ?", name).First(&user)
	return user
}

func FindUserByPhone(phone string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("Phone = ?", phone).First(&user)
}

func FindUserByEmail(email string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("Email = ?", email).First(&user)
}

func CreateUser(user UserBasic) *gorm.DB {
	return utils.DB.Create(&user)
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user UserBasic) *gorm.DB {
	// return utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, Password: user.Password})
	return utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, PassWord: user.PassWord, Phone: user.Phone, Email: user.Email, Avatar: user.Avatar})
}
