package client

const (
	NotificationType                      = "notification"
	NotificationFieldAlibabaCloudConfig   = "alibabaCloudConfig"
	NotificationFieldDingtalkConfig       = "dingtalkConfig"
	NotificationFieldMessage              = "message"
	NotificationFieldMicrosoftTeamsConfig = "microsoftTeamsConfig"
	NotificationFieldPagerdutyConfig      = "pagerdutyConfig"
	NotificationFieldSMTPConfig           = "smtpConfig"
	NotificationFieldSlackConfig          = "slackConfig"
	NotificationFieldWebhookConfig        = "webhookConfig"
	NotificationFieldWechatConfig         = "wechatConfig"
)

type Notification struct {
	AlibabaCloudConfig   *AlibabaCloudConfig   `json:"alibabaCloudConfig,omitempty" yaml:"alibabaCloudConfig,omitempty"`
	DingtalkConfig       *DingtalkConfig       `json:"dingtalkConfig,omitempty" yaml:"dingtalkConfig,omitempty"`
	Message              string                `json:"message,omitempty" yaml:"message,omitempty"`
	MicrosoftTeamsConfig *MicrosoftTeamsConfig `json:"microsoftTeamsConfig,omitempty" yaml:"microsoftTeamsConfig,omitempty"`
	PagerdutyConfig      *PagerdutyConfig      `json:"pagerdutyConfig,omitempty" yaml:"pagerdutyConfig,omitempty"`
	SMTPConfig           *SMTPConfig           `json:"smtpConfig,omitempty" yaml:"smtpConfig,omitempty"`
	SlackConfig          *SlackConfig          `json:"slackConfig,omitempty" yaml:"slackConfig,omitempty"`
	WebhookConfig        *WebhookConfig        `json:"webhookConfig,omitempty" yaml:"webhookConfig,omitempty"`
	WechatConfig         *WechatConfig         `json:"wechatConfig,omitempty" yaml:"wechatConfig,omitempty"`
}
