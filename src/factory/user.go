package factory

import "telegram-bot/src/model"

func CreateUser(id int64, name string, userName string, chatId int64) model.User {
	return model.User{
		ExternalId: id,
		FirstName:  name,
		UserName:   userName,
		ChatId:     chatId,
	}
}
