package utils

import (
	"context"
	"log"
	"sync"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/google/uuid"
)

func GetUUID() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
		return ""
	}
	return uuid.String()
}

type transactionId struct {
	value string
}

// AWS Req ID for Idempotency
var lock_aws_req_id = &sync.Mutex{}

var aws_req_id *transactionId

func NewAWSReqId(ctx context.Context) *transactionId {
	lc, _ := lambdacontext.FromContext(ctx)
	lock_aws_req_id.Lock()
	defer lock_aws_req_id.Unlock()
	aws_req_id = &transactionId{value: lc.AwsRequestID}
	return (*transactionId)(aws_req_id)
}

func GetAWSReqId() *transactionId {
	return aws_req_id
}
