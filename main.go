package main

import (
    _"encoding/json"
	_"fmt"
	"log"
	"os"

    "github.com/joho/godotenv"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	
	"gobot/handler"
)



func main() {
    var (
        bh          *th.BotHandler
        bot         *telego.Bot
        botToken    string
        err         error
        updates     <- chan telego.Update
    )
    
    err = godotenv.Load(".env")
    if err != nil {
        log.Fatal(err)
    }
    
    botToken = os.Getenv("BOT_TOKEN")
    bot, err = telego.NewBot(botToken)
    if err != nil {
        log.Fatal(err)
    }
    
    updates, _ = bot.UpdatesViaLongPolling(nil)
    bh, _ = th.NewBotHandler(bot, updates)
    defer bot.StopLongPolling()
    defer bh.Stop()
    
    bh.HandleMessage(func (bot *telego.Bot, message telego.Message) {
        handler.OnMessage(bot, message)
    })
    
    bh.Start()
}