package robot

import (
	"context"
	"fmt"
	"github.com/chyroc/lark"
	"ticket-bot-backend/ticket"
	"time"
)

func RunRobot() *lark.Lark {
	appId := GetAppId()
	appSecret := GetAppSecret()
	cli := lark.New(lark.WithAppCredential(appId, appSecret))
	return cli
}

// NewChat
func NewChat(thisticket ticket.Ticket, unionId string, netid string) string {
	ID := thisticket.BMCID
	ctx := context.Background()
	cli := RunRobot()
	ID = "#" + ID + ":" + thisticket.Label
	chatname := &ID
	resp, _, _ := cli.Chat.CreateChat(ctx, &lark.CreateChatReq{ //群聊名需要以指针形式传入
		Name: chatname,
	})
	//fmt.Println(*resp, err)
	fmt.Println("创建群聊", fmt.Sprintf("%+v", *resp))

	idtype := "union_id"
	idlist := []string{unionId}
	resp1, _, _ := cli.Chat.AddChatMember(ctx, &lark.AddChatMemberReq{
		MemberIDType: (*lark.IDType)(&idtype),
		ChatID:       resp.ChatID,
		IDList:       idlist,
	})
	fmt.Println("拉人", fmt.Sprintf("%+v", *resp1))

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
							Tag:     "lark_md",
							Content: "创建者：<at id=" + netid + "><at>", //这里@指定用户需要netid
							Lines:   5,
						},
					}, {
						IsShort: true,
						Text: &lark.MessageContentCardObjectText{
							Tag:     "lark_md",
							Content: "创建时间：" + time.Now().Format("2006-01-02 15:04:05"),
							Lines:   5,
						},
					}, {
						IsShort: true,
						Text: &lark.MessageContentCardObjectText{
							Tag:     "lark_md",
							Content: "地点：" + thisticket.RelatedInf.Client.Department,
							Lines:   5,
						},
					}, {
						IsShort: true,
						Text: &lark.MessageContentCardObjectText{
							Tag:     "lark_md",
							Content: "联系电话：" + thisticket.RelatedInf.Client.Phone,
							Lines:   5,
						},
					}, {
						IsShort: true,
						Text: &lark.MessageContentCardObjectText{
							Tag:     "lark_md",
							Content: "联系人：" + thisticket.RelatedInf.Client.Name,
							Lines:   5,
						},
					}, {
						IsShort: false,
						Text: &lark.MessageContentCardObjectText{
							Tag:     texttag,
							Content: "问题描述：" + thisticket.RelatedInf.Summary,
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
	resp5, _, _ := cli.Message.Send().ToChatID(resp.ChatID).SendCard(ctx, card.String())
	fmt.Println("发送工单信息卡片", fmt.Sprintf("%+v", resp5))
	return resp.ChatID
}

// FinishTicket 结束工单的卡片
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
	resp5, _, _ := cli.Message.Send().ToChatID(chatid).SendCard(ctx, card.String())
	fmt.Println("发送结束工单卡片", fmt.Sprintf("%+v", resp5))
}
