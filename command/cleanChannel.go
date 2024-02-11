package command

import (
    "encoding/json"
    
    "github.com/mymmrac/telego"
)


func CleanChannel(
    bot         *telego.Bot,
    message     telego.Message,
    argument    string,
) {
    var (
        
    )
    
    b, _ := bot.GetChat(&telego.GetChatParams {
        ChatID: telego.ChatID {ID: message.Chat.ID},
    })
    a, _ := json.MarshalIndent(b, "", "    ")
    bot.SendMessage(&telego.SendMessageParams {
        ChatID: telego.ChatID {ID: message.Chat.ID},
        Text  : string(a),
    })
}