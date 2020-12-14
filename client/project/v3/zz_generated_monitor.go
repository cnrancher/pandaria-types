package client

const (
	MonitorType          = "monitor"
	MonitorFieldInterval = "interval"
	MonitorFieldRecv     = "recv"
	MonitorFieldSend     = "send"
	MonitorFieldTimeout  = "timeout"
	MonitorFieldType     = "type"
)

type Monitor struct {
	Interval int64  `json:"interval,omitempty" yaml:"interval,omitempty"`
	Recv     string `json:"recv,omitempty" yaml:"recv,omitempty"`
	Send     string `json:"send,omitempty" yaml:"send,omitempty"`
	Timeout  int64  `json:"timeout,omitempty" yaml:"timeout,omitempty"`
	Type     string `json:"type,omitempty" yaml:"type,omitempty"`
}
