package util

import (
	"alert/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

func Sendmsg(msg,chatid string) error{
	bot,err:=tgbotapi.NewBotAPI(config.Configmanager.Config.App.BotToken)
	if err!=nil{
		return err
	}
	if chatid!=""{
		personConf:=getMsgConf(msg,chatid)
		if _,err:=bot.Send(*personConf);err !=nil{
			log.Println("send person error",err)
		}
	}
	groupConf:=getMsgConf(msg,config.Configmanager.App.ChatID)
	if _,err:=bot.Send(*groupConf);err!=nil{
		return err
	}
	return nil
}

func getMsgConf(msg,chatid string) *tgbotapi.MessageConfig {
	chatID, err := strconv.ParseInt(chatid, 10, 64)
	if err != nil {
		return nil
	}
	NewMsg:=tgbotapi.NewMessage(chatID,msg)
	NewMsg.ParseMode=tgbotapi.ModeHTML
	return &NewMsg

}