package utils

import (
    "encoding/json"
    "github.com/mymmrac/telego"
)


func IsMember(
    bot       *telego.Bot,
    chatId    telego.ChatID,
    userId    int64,
    tunnel    chan bool,
) {
    var (
        errors        error
        getMember     telego.ChatMember
        members       [3]string
        restricted    telego.ChatMemberRestricted
        status        string
        user          telego.User
        userByte      []byte
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
    if status == "restricted" {
        restricted = telego.ChatMemberRestricted {}
        userByte, _ = json.Marshal(getMember)
        json.Unmarshal(userByte, &restricted)
        
        if restricted.IsMember == true {
            tunnel <- true
            return
        }
    }
    
    members = [3]string {"creator", "administrator", "member"}
    for i := 0; i < len(members); i++ {
        if members[i] == status {
            tunnel <- true
            return
        }
    }
    
    tunnel <- false
    return
}