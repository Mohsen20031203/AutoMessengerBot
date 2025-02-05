package bot

type Bot interface {
	SendMassage(chatId int64, massage string) error
}
