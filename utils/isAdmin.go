package utils

import (
    "github.com/mymmrac/telego"
)


func IsAdmin(
    bot       *telego.Bot,
    chatId    telego.ChatID,
    userId    int64,
    tunnel    chan bool,
) {
    var (
        admins       [2]string
        errors       error
        getMember    telego.ChatMember
        status       string
        user         telego.User
    )
    
    if getMember, errors = bot.GetChatMember(
        &telego.GetChatMemberParams {
            ChatID: chatId,
            UserID: userId,
        },
    ); errors != nil {
        tunnel <- false
        return
    }
    
    user = getMember.MemberUser()
    if user.Username == "GroupAnonymousBot" ||
       user.ID == 777000 {
        tunnel <- true
        return
    }
    
    status = getMember.MemberStatus()
    admins = [2]string {"creator", "administrator"}
    for i := 0; i < len(admins); i++ {
        if admins[i] == status {
            tunnel <- true
            return
        }
    }
    
    tunnel <- false
    return
}