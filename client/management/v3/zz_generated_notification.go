package client

const (
	NotificationType                  = "notification"
	NotificationFieldAliyunSMSConfig  = "aliyunsmsConfig"
	NotificationFieldDingtalkConfig   = "dingtalkConfig"
	NotificationFieldMSTeamsConfig    = "msteamsConfig"
	NotificationFieldMessage          = "message"
	NotificationFieldPagerdutyConfig  = "pagerdutyConfig"
	NotificationFieldSMTPConfig       = "smtpConfig"
	NotificationFieldServiceNowConfig = "servicenowConfig"
	NotificationFieldSlackConfig      = "slackConfig"
	NotificationFieldWebhookConfig    = "webhookConfig"
	NotificationFieldWechatConfig     = "wechatConfig"
)

type Notification struct {
	AliyunSMSConfig  *AliyunSMSConfig  `json:"aliyunsmsConfig,omitempty" yaml:"aliyunsmsConfig,omitempty"`
	DingtalkConfig   *DingtalkConfig   `json:"dingtalkConfig,omitempty" yaml:"dingtalkConfig,omitempty"`
	MSTeamsConfig    *MSTeamsConfig    `json:"msteamsConfig,omitempty" yaml:"msteamsConfig,omitempty"`
	Message          string            `json:"message,omitempty" yaml:"message,omitempty"`
	PagerdutyConfig  *PagerdutyConfig  `json:"pagerdutyConfig,omitempty" yaml:"pagerdutyConfig,omitempty"`
	SMTPConfig       *SMTPConfig       `json:"smtpConfig,omitempty" yaml:"smtpConfig,omitempty"`
	ServiceNowConfig *ServiceNowConfig `json:"servicenowConfig,omitempty" yaml:"servicenowConfig,omitempty"`
	SlackConfig      *SlackConfig      `json:"slackConfig,omitempty" yaml:"slackConfig,omitempty"`
	WebhookConfig    *WebhookConfig    `json:"webhookConfig,omitempty" yaml:"webhookConfig,omitempty"`
	WechatConfig     *WechatConfig     `json:"wechatConfig,omitempty" yaml:"wechatConfig,omitempty"`
}
