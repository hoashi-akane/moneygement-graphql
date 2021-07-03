package repository

import "github.com/hoashi-akane/moneygement-graphql/graph/model"

func (UserDB *UserDB) Login(info *model.LoginInfo) (*model.User, error) {
	var user model.User
	UserDB.DB.Table("users").
		Select("users.id, users.name, users.nickname, users.email," +
			" users.adviser_name, users.introduction, user_auth.token").
		Joins("left join user_auth on users.id = user_auth.user_id").
		Find(&user, "email = ? and user_auth.password = ?", info.Email, info.Password)

	if &user.ID != nil && user.ID != 0 && user.Token != info.Token {
		UserDB.DB.Table("user_auth").
			Where("user_id=? AND password=?", user.ID, info.Password).
			Update("token", info.Token)
	}
	return &user, nil
}

func (UserDB *UserDB) GetUser(email string) *model.User {
	var user model.User
	UserDB.DB.Table("users").Take(&user, "email=?", email)
	return &user
}

func (UserDB *UserDB) GetIdUser(id int) *model.User {
	var user model.User
	UserDB.DB.Table("users").Where("id=?", id).Find(&user)
	return &user
}

func (UserDB *UserDB) GetAdviserList(listFilter *model.AdviserListFilter) ([]*model.User, error) {
	var userList []*model.User
	UserDB.DB.Table("users").
		Select("id, adviser_name, introduction").
		Offset(listFilter.First).Limit(listFilter.Last).Find(&userList, "is_adviser=true")
	return userList, nil
}

func (UserDB *UserDB) GetAdviserMemberList(memberFilter *model.UseAdviserMemberFilter) ([]*model.AdviserMember, error) {
	var adviserList []*model.AdviserMember
	var userList []*model.User
	var ledgerList []*model.Ledger
	var userIdList []int
	UserDB.DB.Table("ledger").Select("id, user_id, name").Find(&ledgerList, "adviser_id = ?", memberFilter.UserID)
	for _, adviser := range ledgerList {
		if userIdList == nil {
			userIdList = append(userIdList, adviser.UserID)
		} else {
			for _, v := range userIdList {
				if adviser.UserID != v {
					userIdList = append(userIdList, v)
				}
			}
		}
	}

	UserDB.DB.Table("users").Find(&userList, "id IN (?)", userIdList)
	for _, ledger := range ledgerList {
		for _, user := range userList {
			if ledger.UserID == user.ID {
				var adviser = model.AdviserMember{ID: ledger.ID, UserID: user.ID, NickName: user.Nickname, LedgerName: ledger.Name}
				adviserList = append(adviserList, &adviser)
			}
		}
	}
	return adviserList, nil
}

func (UserDB *UserDB)  GetChatList(chatFilter *model.ChatFilter) ([]*model.Chat, error) {
	var chatList []*model.Chat
	UserDB.DB.Order("created_at desc").Table("chat").
		Select("chat.id, chat.ledger_id, chat.user_id, chat.comment, chat.created_at, users.nickname").
		Joins("left join users on chat.user_id = users.id").Offset(chatFilter.First).Limit(chatFilter.Last).
		Find(&chatList, "ledger_id = ?", chatFilter.LedgerID)
	return chatList, nil
}

func (UserDB *UserDB) CreateGroup(group *model.Group) (int, error) {
	err := UserDB.DB.Table("groups").Create(&group).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (UserDB *UserDB) CreateEnrollment(enrollment *model.Enrollment) (int, error) {
	err := UserDB.DB.Table("enrollment").Create(&enrollment).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (UserDB *UserDB) CreateUser(user *model.User, password string) *model.User {
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

func (UserDB *UserDB) CreateAdviser(adviser *model.NewAdviser) (int, error) {
	err := UserDB.DB.Table("users").
		Where("id=?", adviser.ID).
		Updates(map[string]interface{}{
			"name": adviser.Name, "introduction": adviser.Introduction,
			"adviser_name": adviser.AdviserName,
			"is_adviser": true}).Error
	if err != nil {
		panic("エラ-")
		return 0, err
	}
	return 1, nil
}

func (UserDB *UserDB) CreateChat(newChat *model.NewChat) (int, error) {
	var err = UserDB.DB.Table("chat").Create(&newChat).Error
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func (UserDB *UserDB) UpdateUser(user *model.UpdateUser) (int, error){
	err := UserDB.DB.Table("users").
		Where("id = ?", user.ID).Updates(map[string]interface{}{
		"name": user.Name, "nickname": user.Nickname,
		"email": user.Email, "introduction": user.Introduction,
		"adviser_name": user.AdviserName}).
		Error
	if err != nil {
		panic("エラ-")
		return 0, err
	}
	return 1, nil
}

func (UserDB *UserDB) InsertGroupUser(newGroupUser *model.NewGroupUser) (int, error) {
	var user *model.User
	db := UserDB.DB
	db.Table("users").Select("id").
		Find(&user, "email=? AND nickname=?", newGroupUser.Email, newGroupUser.NickName)
	if user.ID == 0 {
		return 0, nil
	}
	var enrollment = model.Enrollment{UserID: user.ID, GroupID: newGroupUser.GroupID}
	err := db.Table("enrollment").Create(&enrollment).Error
	if err != nil {
		panic("エラ-")
		return 0, err
	}
	return 1, nil
}

func (UserDB *UserDB)  DeleteGroup(id int) (int, error) {
	var group = model.Group{ID: id}
	err := UserDB.DB.Table("groups").Delete(&group).Error
	if err != nil {
		panic("エラ-")
		return 0, err
	}
	return 1, nil
}

type UserDBInterface interface {
	Login(info *model.LoginInfo) (*model.User, error)
	GetUser(email string) *model.User
	GetIdUser(id int) *model.User
	GetAdviserList(listFilter *model.AdviserListFilter) ([]*model.User, error)
	GetAdviserMemberList(memberFilter *model.UseAdviserMemberFilter) ([]*model.AdviserMember, error)
	GetChatList(chatFilter *model.ChatFilter) ([]*model.Chat, error)
	CreateGroup(group *model.Group) (int, error)
	CreateEnrollment(enrollment *model.Enrollment) (int, error)
	CreateUser(user *model.User, password string) *model.User
	CreateAdviser(adviser *model.NewAdviser) (int, error)
	CreateChat(newChat *model.NewChat) (int, error)
	UpdateUser(user *model.UpdateUser) (int, error)
	InsertGroupUser(newGroupUser *model.NewGroupUser) (int, error)
	DeleteGroup(id int) (int, error)
}
