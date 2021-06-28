package repository

import "github.com/hoashi-akane/moneygement-graphql/graph/model"

func (UserDB *UserDB) InsertUser(user *model.User, password string) *model.User {
	if UserDB.DB.Table("users").Create(&user).Error != nil {
		panic("エラ-")
		return nil
	}
	user = UserDB.GetUser(user.Email)
	var auth = model.UserAuth{UserID: user.ID, Password: password}
	if UserDB.DB.Table("user_auth").Create(&auth).Error != nil {
		return nil
	}
	return user
}

func (UserDB *UserDB) GetUser(email string) *model.User {
	var user model.User
	UserDB.DB.Table("users").Take(&user, "email=?", email)
	return &user
}


type UserDBInterface interface {
	InsertUser(user *model.User) *model.User
	GetUser(email int) *model.User
}
