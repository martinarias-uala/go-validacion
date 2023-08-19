package handler

/* import (
	"context"

	"github.com/aws/aws-lambda-go/lambdacontext"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var (
	mapstate = models.MapState{
		TraceId:       "ABC123",
		BatchNumber:   "123",
		BatchSaleDate: "YYYYMMDD",
		Step:          "mockStep",
		Portfolio: models.LoanPortfolio{
			LoanId:    "LID123",
			BranchKey: "123ABC",
		},
	}
	trace_item = models.TraceItem{
		Timestamp:   "YYYY-MM-DD HH:ii:ss",
		Subprocess:  "mockStep",
		State:       "MockState",
		Description: "mock",
		LoanId:      "LID123",
		TraceId:     "ABC123",
	}
	response = models.DynamicAPICallResponse{
		MapState:  mapstate,
		TraceItem: trace_item,
	}
	ct = &lambdacontext.LambdaContext{
		AwsRequestID:       "awsRequestId1234",
		InvokedFunctionArn: "arn:aws:lambda:xxx",
		Identity:           lambdacontext.CognitoIdentity{},
		ClientContext:      lambdacontext.ClientContext{},
	}
	ctx = lambdacontext.NewContext(context.TODO(), ct)
)

var _ = Describe("Handler", func() {
	Context("When receive event handle it successfully", happyPath)
})

func happyPath() {

	It("Should handle event", func() {
		p := processor.MockProcessor{}
		p.On("Process", mapstate).Return(response)
		h := New(&p)
		res, err := h.Handle(ctx, mapstate)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(res).Should(Equal(response))
	})
} */
