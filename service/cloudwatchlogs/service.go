package cloudwatchlogs

import (
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/internal/protocol/jsonrpc"
	"github.com/awslabs/aws-sdk-go/internal/signer/v4"
)

// CloudWatchLogs is a client for Amazon CloudWatch Logs.
type CloudWatchLogs struct {
	*aws.Service
}

// New returns a new CloudWatchLogs client.
func New(config *aws.Config) *CloudWatchLogs {
	if config == nil {
		config = &aws.Config{}
	}

	service := &aws.Service{
		Config:       aws.DefaultConfig.Merge(config),
		ServiceName:  "logs",
		APIVersion:   "2014-03-28",
		JSONVersion:  "1.1",
		TargetPrefix: "Logs_20140328",
	}
	service.Initialize()

	// Handlers
	service.Handlers.Sign.PushBack(v4.Sign)
	service.Handlers.Build.PushBack(jsonrpc.Build)
	service.Handlers.Unmarshal.PushBack(jsonrpc.Unmarshal)
	service.Handlers.UnmarshalMeta.PushBack(jsonrpc.UnmarshalMeta)
	service.Handlers.UnmarshalError.PushBack(jsonrpc.UnmarshalError)

	return &CloudWatchLogs{service}
}
