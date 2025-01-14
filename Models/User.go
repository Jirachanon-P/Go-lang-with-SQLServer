package Models

import (
	"first-api/Config"
)

func GetAllUsers(users *[]User) (err error) {
	if err = Config.DB.Where("isDelete = ?", false).Find(users).Error; err != nil {
		return err
	}
	return err
}

func CreatUser(user *User) (err error) {
	if err = Config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserById(user *User, id int) (err error) {
	if err = Config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User, id int) (err error) {
	if err = Config.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user User) (err error) {
	user.IsDelete = true
	if err = Config.DB.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func Login(user *User, username string) (err error) {
	if err = Config.DB.Where("username = ?", username).First(user).Error; err != nil {
		return err
	}
	return nil
}
