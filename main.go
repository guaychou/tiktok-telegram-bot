package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	scraper :=newScraper()

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		if update.Message.IsCommand() {
			msg:=tgbotapi.NewMessage(update.Message.Chat.ID,"")
			msgVideo := tgbotapi.NewVideoShare(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "start":
				msg.Text="I'm a babu of lordchou thanks for add me,\nlordchou LinkedIn: https://www.linkedin.com/in/kevin-jonathan-harnanta-b0745216b/ \nSend /help for more information."

			case "tiktok":
				split:=strings.Split(update.Message.Text," ")
				if len(split)!=2{
					msg.Text="Some argument is missing. \nUse /tiktok <link> to get the value."

				}else if len(split)==2 {
					copiedLink:=split[1]
					if (isTiktokUrl(copiedLink)){
						data,err:=getVideoLink(copiedLink,scraper)
						checkErr(err)
						msgVideo.FileID=data.ContentUrl
					}else{
						msg.Text="Invalid Tiktok URL"
					}
				}
			case "help":
				msg.Text="Type /tiktok: for download tiktok video"
			}
			msg.ParseMode="markdown"
			bot.Send(msg)
			bot.Send(msgVideo)

		}
	}
}

func isTiktokUrl(str string) bool {
	if (strings.Contains(str,"tiktok")){
		u, err := url.Parse(str)
		return err == nil && u.Scheme != "" && u.Host != ""
	}
	return false
}