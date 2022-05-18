package gw2w

// WebhookNotifierBody is the body of webhook
// notification channel
// https://github.com/grafana/grafana/blob/v8.5.3/pkg/services/alerting/notifiers/webhook.go#L94
type WebhookNotifierBody struct {
	Title       string            `json:"title"`              // title
	RuleID      int64             `json:"ruleId"`             // ruleId
	RuleName    string            `json:"ruleName"`           // ruleName
	State       string            `json:"state"`              // state
	EvalMatches []interface{}     `json:"evalMatches"`        // evalMatches
	OrgID       int64             `json:"orgId"`              // orgId
	DashboardID int64             `json:"dashboardId"`        // dashboardId
	PanelID     int64             `json:"panelId"`            // panelId
	Tags        map[string]string `json:"tags"`               // tags
	RuleURL     string            `json:"ruleUrl,omitempty"`  // ruleUrl,omitempty
	ImageURL    string            `json:"imageUrl,omitempty"` // imageUrl,omitempty
	Message     string            `json:"message,omitempty"`  // message,omitempty
}

// Trans to WeChatBOTNewsBody
func (b WebhookNotifierBody) Trans(chatID string, visibleToUser string) WeChatBOTNewsBody {
	var cID *string
	var user *string
	if chatID != "" {
		cID = &chatID
	}
	if visibleToUser != "" {
		user = &visibleToUser
	}

	return WeChatBOTNewsBody{
		ChatID:        cID,
		MSGType:       "news",
		VisibleToUser: user,
		News: WeChatBOTNews{
			Articles: []WeChatBOTNewsArticle{
				{
					Title:       b.Title,
					Description: b.Message,
					Url:         b.RuleURL,
					PicURL:      b.ImageURL,
				},
			},
		},
	}
}
