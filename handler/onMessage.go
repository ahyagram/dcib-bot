package handler

import (
    "fmt"
    "strconv"
    "strings"
    
    "github.com/mymmrac/telego"
    
    _"gobot/command"
    "gobot/database"
    "gobot/utils"
)


// var COMMAND = map[string]func(*telego.Bot, telego.Message, string) {
//     "cleanchannel": command.CleanChannel,
// }


func OnMessage(bot *telego.Bot, message telego.Message) {
    if message.LeftChatMember != nil {
        return
    } else if message.NewChatMembers != nil {
        unmute(bot, message)
        return
    }
    
    var (
        // argument    string
        // command     string
        // parser      *utils.Parser
    )
    
    clean(bot, message)
    /*parser = utils.NewParser(message.Text)
    command = parser.Command()
    argument = parser.Argument()
    
    i, e := COMMAND[command]
    if e == false {
        return
    }
    
    i(bot, message, argument)*/
    return
}


// Private function
func clean(bot *telego.Bot, message telego.Message) {
    var (
        chatId       int64
        chatIdTlg    telego.ChatID
        chatTitle    string
        groupLink    string
        messageId    int
        text         string
        tunnel       chan bool
        msgSent      *telego.Message
        userId       int64
        userName     string
    )
    
    chatId = message.Chat.ID
    chatIdTlg = telego.ChatID {ID: chatId}
    chatTitle = message.Chat.Title
    userId = message.From.ID
    userName = message.From.FirstName
    messageId = message.MessageID
    tunnel = make(chan bool)
    groupLink = database.Get("joinLink.json", strconv.Itoa(int(chatId)))
    groupLink = strings.Replace(groupLink,"+", "%2B", 1)
    
    go utils.IsMember(bot, chatIdTlg, userId, tunnel)
    if (<- tunnel) == true {
        return
    }
    
    text = fmt.Sprintf("<a href='tg://user?id=%s'>%s</a>", strconv.Itoa(int(userId)), userName)
    text += ", silahkan masuk grup untuk bisa berdiskusi.\nTautan: "
    text += fmt.Sprintf("<a href='%s'>%s</a>", groupLink, chatTitle)
    
    msgSent = sendMessage(bot, chatIdTlg, messageId, text, nil)
    restrictChatMember(bot, chatIdTlg, userId, telego.ChatPermissions {
        CanSendMessages: utils.BoolPointer(false),
    })
    
    mId, _ := strconv.Atoi(database.Get(
        "lastJoin.json",
        strconv.Itoa(int(chatId)),
    ))
    
    deleteMessage(bot, chatIdTlg, messageId)
    deleteMessage(bot, chatIdTlg, mId)
    database.Update(
        "lastJoin.json",
        strconv.Itoa(int(chatId)),
        strconv.Itoa(msgSent.MessageID),
    )
}


func unmute(bot *telego.Bot, message telego.Message) {
    var (
        chatIdTlg    telego.ChatID
        userId       int64
    )
    
    chatIdTlg = telego.ChatID {ID: message.Chat.ID}
    userId = message.From.ID
    restrictChatMember(bot, chatIdTlg, userId, telego.ChatPermissions {
        CanSendMessages: utils.BoolPointer(true),
        CanSendAudios: utils.BoolPointer(true),
        CanSendDocuments: utils.BoolPointer(true),
        CanSendPhotos: utils.BoolPointer(true),
        CanSendVideos: utils.BoolPointer(true),
        CanSendVideoNotes: utils.BoolPointer(true),
        CanSendVoiceNotes: utils.BoolPointer(true),
        CanSendPolls: utils.BoolPointer(true),
        CanSendOtherMessages: utils.BoolPointer(true),
        CanAddWebPagePreviews: utils.BoolPointer(true),
        CanChangeInfo: utils.BoolPointer(true),
        CanInviteUsers: utils.BoolPointer(true),
        CanPinMessages: utils.BoolPointer(true),
        CanManageTopics: utils.BoolPointer(true),
    })
}


// telego.Bot function which called using go keyword
func deleteMessage(
    bot          *telego.Bot,
    chatId       telego.ChatID,
    messageId    int,
) {
    bot.DeleteMessage(&telego.DeleteMessageParams {
        ChatID   : chatId,
        MessageID: messageId,
    })
}


func restrictChatMember(
    bot            *telego.Bot,
    chatId         telego.ChatID,
    userId         int64,
    permissions    telego.ChatPermissions,
) {
    bot.RestrictChatMember(&telego.RestrictChatMemberParams {
        ChatID     : chatId,
        UserID     : userId,
        Permissions: permissions,
    })
}


func sendMessage(
    bot          *telego.Bot,
    chatId       telego.ChatID,
    messageId    int,
    text         string,
    tunnel       chan *telego.Message,
) *telego.Message {
    var message *telego.Message
    message, _ = bot.SendMessage(
        &telego.SendMessageParams {
            ChatID   : chatId,
            Text     : text,
            ParseMode: "HTML",
            ReplyToMessageID: messageId,
        },
    )
    
    if tunnel != nil {
        tunnel <- message
        return nil
    }
    
    return message
}