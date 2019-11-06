package client

const (
	PodFileDownloadOutputType             = "podFileDownloadOutput"
	PodFileDownloadOutputFieldFileContent = "fileContent"
)

type PodFileDownloadOutput struct {
	FileContent string `json:"fileContent,omitempty" yaml:"fileContent,omitempty"`
}
