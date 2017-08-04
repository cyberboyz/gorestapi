package models

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Age      int     `json:"age"`
	Weight   float32 `json:"weight"`
	Verified bool    `json:"verified"`
}

func GetUsers() ([]*User, error) {
	users := []*User{}
	err = db.Find(&users).Error
    return users, err
}

func GetUser(id int) (*User, error){
	user := &User{}
	err = db.Where("id = ?", id).First(&user).Error
	return user, err
}

func CreateUser(user *User) (*User, error){
	err = db.Create(user).Error
	return user, err
}

func UpdateUser(user *User) (*User, error){
	err = db.Save(user).Error
	return user, err
}

func DeleteUser(id int) error {
	err = db.Where("id = ?",id).Delete(&User{}).Error
	return err
}

