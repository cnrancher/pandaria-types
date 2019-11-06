package client

const (
	PodFileDownloadInputType               = "podFileDownloadInput"
	PodFileDownloadInputFieldContainerName = "containerName"
	PodFileDownloadInputFieldFilePath      = "filePath"
)

type PodFileDownloadInput struct {
	ContainerName string `json:"containerName,omitempty" yaml:"containerName,omitempty"`
	FilePath      string `json:"filePath,omitempty" yaml:"filePath,omitempty"`
}
