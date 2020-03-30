package client

const (
	QueryGlobalMetricInputType          = "queryGlobalMetricInput"
	QueryGlobalMetricInputFieldExpr     = "expr"
	QueryGlobalMetricInputFieldFrom     = "from"
	QueryGlobalMetricInputFieldInterval = "interval"
	QueryGlobalMetricInputFieldTo       = "to"
)

type QueryGlobalMetricInput struct {
	Expr     string `json:"expr,omitempty" yaml:"expr,omitempty"`
	From     string `json:"from,omitempty" yaml:"from,omitempty"`
	Interval string `json:"interval,omitempty" yaml:"interval,omitempty"`
	To       string `json:"to,omitempty" yaml:"to,omitempty"`
}
