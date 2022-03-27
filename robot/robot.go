package robot

import (
	"context"
	"fmt"
	"github.com/chyroc/lark"
)

func RunRobot() *lark.Lark {
	appId := GetAppId()
	appSecret := GetAppSecret()
	cli := lark.New(lark.WithAppCredential(appId, appSecret))
	return cli
}

// NewChat 貌似Union_id还是不知道怎么获取 （未完成）
func NewChat(ID *string) {
	ctx := context.Background()
	cli := RunRobot()
	*ID = "test"
	resp, _, _ := cli.Chat.CreateChat(ctx, &lark.CreateChatReq{
		Name: ID,
	})
	//fmt.Println(*resp, err)
	fmt.Println(fmt.Sprintf("%+v", *resp))

	idtype := "app_id"
	idlist := []string{GetAppId()}
	resp1, _, err1 := cli.Chat.AddChatMember(ctx, &lark.AddChatMemberReq{
		MemberIDType: (*lark.IDType)(&idtype),
		ChatID:       resp.ChatID,
		IDList:       idlist,
	})
	fmt.Println(resp1, err1)
}
