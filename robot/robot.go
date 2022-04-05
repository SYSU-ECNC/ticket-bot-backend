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

// NewChat id是工单id，预计后期可以改成ticket
func NewChat(ID string, unionId string, netid string) {
	ctx := context.Background()
	cli := RunRobot()
	ID = "#" + ID + ":新工单"
	chatname := &ID
	resp, _, _ := cli.Chat.CreateChat(ctx, &lark.CreateChatReq{ //群聊名需要以指针形式传入
		Name: chatname,
	})
	//fmt.Println(*resp, err)
	fmt.Println(fmt.Sprintf("%+v", *resp))

	idtype := "union_id"
	idlist := []string{unionId}
	resp1, _, err1 := cli.Chat.AddChatMember(ctx, &lark.AddChatMemberReq{
		MemberIDType: (*lark.IDType)(&idtype),
		ChatID:       resp.ChatID,
		IDList:       idlist,
	})
	fmt.Println(resp1, err1)

	//resp2, _, _ := cli.Message.Send().ToChatID(resp.ChatID).SendText(ctx, "信息")
	//fmt.Println(fmt.Sprintf("%+v", *resp2))

	data := &lark.MessageContentPostAll{ //@创建者
		ZhCn: &lark.MessageContentPost{
			Title: ID,
			Content: [][]lark.MessageContentPostItem{
				{
					lark.MessageContentPostText{Text: "你创建了一个新的工单"},
				},
				{
					lark.MessageContentPostAt{UserID: unionId},
				},
			},
		},
		JaJp: nil,
		EnUs: nil,
	}
	_, _, _ = cli.Message.Send().ToChatID(resp.ChatID).SendPost(ctx, data.String())

	texttag := lark.MessageContentCardObjectTextType("lark_md")
	card := &lark.MessageContentCard{
		Header: &lark.MessageContentCardHeader{
			Template: "blue",
			Title: &lark.MessageContentCardObjectText{
				Tag:     "plain_text",
				Content: "📪 有一条新工单需要处理！",
				Lines:   5,
			},
		},
		Config: &lark.MessageContentCardConfig{
			EnableForward: false,
			UpdateMulti:   true,
		},
		Modules: []lark.MessageContentCardModule{
			lark.MessageContentCardModuleDIV{
				Text: nil,
				Fields: []*lark.MessageContentCardObjectField{
					{
						IsShort: true,
						Text: &lark.MessageContentCardObjectText{
							Tag: "lark_md",
							//"is_short": true,
							//		"text": {
							//		"content": "**👤 提交人：**\n<at email=test@email.com></at>",
							//		"tag": "lark_md"
							Content: "创建者：<at id=" + netid + "><at>", //这里@指定用户需要netid
							Lines:   5,
						},
					}, {
						IsShort: true,
						Text: &lark.MessageContentCardObjectText{
							Tag:     "lark_md",
							Content: "创建时间：2022/4/5 19:00",
							Lines:   5,
						},
					}, {
						IsShort: true,
						Text: &lark.MessageContentCardObjectText{
							Tag:     "lark_md",
							Content: "地点：东实验楼A栋404",
							Lines:   5,
						},
					}, {
						IsShort: true,
						Text: &lark.MessageContentCardObjectText{
							Tag:     "lark_md",
							Content: "联系电话：84036866",
							Lines:   5,
						},
					}, {
						IsShort: true,
						Text: &lark.MessageContentCardObjectText{
							Tag:     "lark_md",
							Content: "联系人：网络中心",
							Lines:   5,
						},
					}, {
						IsShort: false,
						Text: &lark.MessageContentCardObjectText{
							Tag:     texttag,
							Content: "问题描述：写不完了",
							Lines:   0,
						},
					},
				},
				Extra: nil,
			},
			lark.MessageContentCardModuleAction{
				Actions: []lark.MessageContentCardElement{
					lark.MessageContentCardElementButton{
						Text: &lark.MessageContentCardObjectText{
							Tag:     "lark_md",
							Content: "指派人员",
							Lines:   5,
						},
						URL:      "",
						MultiURL: nil,
						Type:     "primary",
						Value:    nil,
						Confirm:  nil,
					}, lark.MessageContentCardElementButton{
						Text: &lark.MessageContentCardObjectText{
							Tag:     "lark_md",
							Content: "结束工单",
							Lines:   5,
						},
						URL:      "",
						MultiURL: nil,
						Type:     "",
						Value:    nil,
						Confirm:  nil,
					},
				},
				Layout: "bisected",
			},
		},
	}
	resp5, _, err5 := cli.Message.Send().ToChatID(resp.ChatID).SendCard(ctx, card.String())
	fmt.Println(resp5, err5)
	FinishTicket(resp.ChatID)
}

// FinishTicket 结束工单的卡片 (还需要完善，后期可能需要受派者的netid）
func FinishTicket(chatid string) {
	ctx := context.Background()
	cli := RunRobot()
	card := &lark.MessageContentCard{
		Header: &lark.MessageContentCardHeader{
			Template: "green",
			Title: &lark.MessageContentCardObjectText{
				Tag:     "plain_text",
				Content: "📅工单已结束",
				Lines:   5,
			},
		},
		Config: &lark.MessageContentCardConfig{
			EnableForward: false,
			UpdateMulti:   false,
		},
		Modules: []lark.MessageContentCardModule{
			lark.MessageContentCardModuleDIV{
				Text: &lark.MessageContentCardObjectText{
					Tag:     "lark_md",
					Content: "受派者：<at id=all></at>", //需要netid，这里先用所有人代替
					Lines:   5,
				},
				Fields: nil,
				Extra:  nil,
			},
		},
	}
	resp5, _, err5 := cli.Message.Send().ToChatID(chatid).SendCard(ctx, card.String())
	fmt.Println(resp5, err5)
}
