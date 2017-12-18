package models

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Age      int     `json:"age"`
	Weight   float32 `json:"weight"`
	Verified bool    `json:"verified"`
	Token    string  `json:"token"`
}

func AuthorizeUser(user *User, authorization string) error {
	err = db.Where("token = ? ", authorization).Find(&user).Error
	return err
}

func RegisterUser(user *User) (*User, error) {
	err = db.Where("email = ? ", user.Email).Find(&user).Error
	if err == nil {
		return nil, nil
	}
	err = db.Create(user).Error
	return user, err
}

func LoginUser(user *User) error {
	err = db.Where("email = ? ", user.Email).Find(&user).Error
	return err
}

func LogoutUser(user *User, authorization string) error {
	err = db.Model(user).Update("token", user.Token).Where("token", authorization).Error
	return err
}

func UpdateToken(user *User) (*User, error) {
	err = db.Model(user).Update("token", user.Token).Error
	return user, err
}

func GetUsers() ([]*User, error) {
	users := []*User{}
	err = db.Find(&users).Error
	return users, err
}

func GetUser(id int) (*User, error) {
	user := &User{}
	err = db.Where("id = ?", id).First(&user).Error
	return user, err
}

func CreateUser(user *User) (*User, error) {
	err = db.Create(user).Error
	return user, err
}

func UpdateUser(user *User) (*User, error) {
	err = db.Save(user).Error
	return user, err
}

func DeleteUser(email string) error {
	err = db.Where("email = ?", email).Delete(&User{}).Error
	return err
}
