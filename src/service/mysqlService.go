package service

import (
	"fmt"
	dto2 "web-api/src/controllers/dto"
	"web-api/utils/constant"

	"github.com/anthonycorbacho/slack-webhook"
)

type TestAPI struct{}

func (slacktest *TestAPI) TestMessageForSlack(testslack dto2.TestSlackMessage) (string, error) {

	IntegrationName := testslack.IntegrationName
	HookUrl := testslack.HookUrl
	SenderName := testslack.SenderName
	ChannelName := testslack.ChannelName
	MessageTitle := testslack.MessageTitle
	TagsSentWithAlerts := testslack.TagsSentWithAlerts

	if len(IntegrationName) > 0 && len(HookUrl) > 0 && len(SenderName) > 0 && len(ChannelName) > 0 && len(MessageTitle) >= 0 && len(TagsSentWithAlerts) >= 0 {
		msg := slack.Message{
			Text:     constant.TestTextForSlack,
			Username: constant.UserNameForSlack,
			Channel:  ChannelName,
		}
		err := slack.Send(HookUrl, msg)
		if err != nil {
			fmt.Printf("failed to send message: %v\n", err)
			return "Enter the correct HookUrl", nil
		}
	} else {
		return "Enter the values", nil
	}
	return "Test Successful", nil
}