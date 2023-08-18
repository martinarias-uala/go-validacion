package dynamo

/*

import (
	"context"
	"errors"

	ddb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("Dynamo Repository", func() {
	Context("When receive error insert to dynamo", happyPath)
	Context("When receive error fail on insert", errorPath)
})

func happyPath() {
	var (
		item = models.ACTraceItem{
			Timestamp:   "2023-05-30",
			Subprocess:  models.ReadSubprocess,
			State:       models.ErrorState,
			Description: "some error",
			LoanId:      "abc-123",
			TraceId:     "test",
		}
		c = MockClient{}
		d = New()
	)

	c.On("PutItem", context.TODO(), mock.Anything, mock.Anything).Return(&ddb.PutItemOutput{}, nil)
	d.client = &c

	err := d.CreateItem(item)
	It("Should not return error", func() {
		Ω(err).To(BeNil())
	})
}

func errorPath() {
	var (
		item = models.ACTraceItem{
			Timestamp:   "2023-05-30",
			Subprocess:  models.ReadSubprocess,
			State:       models.ErrorState,
			Description: "some error",
			LoanId:      "abc-123",
			TraceId:     "test",
		}
		c = MockClient{}
		d = New()
	)

	c.On("PutItem", context.TODO(), mock.Anything, mock.Anything).Return(&ddb.PutItemOutput{}, errors.New("some error"))
	d.client = &c

	err := d.CreateItem(item)
	It("Should return error", func() {
		Ω(err).Should(HaveOccurred())
	})
}
*/
