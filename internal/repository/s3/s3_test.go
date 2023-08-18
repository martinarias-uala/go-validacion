package s3

/*
import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/Uilobank/uilo-loan-portfolio-purchase/commons/models"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"
)

var _ = Describe("S3 Repository", func() {
	models.NewTid("test")
	models.NewLid("test")
	Context("When receive event get object content", happyPath)
	Context("When receive event fail on get object content", errorPath)
	Context("Put Object", putObjectPath)

})

func happyPath() {
	var (
		bucket  = "some bucket"
		key     = "some key"
		inp     = s3.GetObjectInput{Bucket: &bucket, Key: &key}
		content = "some content"
		out     = s3.GetObjectOutput{Body: io.NopCloser(strings.NewReader(content))}
		c       = MockClient{}
		r       = New()
		empty   = models.ACTraceItem{}
	)

	c.On("GetObject", context.TODO(), &inp, mock.Anything).Return(&out, nil)
	r.client = &c

	res, err := r.GetObjectContent(bucket, key)
	It("Should return content string", func() {
		Ω(res).To(Equal(content))
	})
	It("Should not return error", func() {
		Ω(err).To(Equal(empty))
	})
}

func errorPath() {
	var (
		bucket = "some bucket"
		key    = "some key"
		inp    = s3.GetObjectInput{Bucket: &bucket, Key: &key}
		c      = MockClient{}
		r      = New()
	)

	c.On("GetObject", context.TODO(), &inp, mock.Anything).Return(nil, errors.New("some error"))
	r.client = &c

	_, err := r.GetObjectContent(bucket, key)
	It("Should throw error", func() {
		Ω(err).ToNot(BeNil())
	})
	It("Should have structured error", func() {
		Ω(err.State).To(Equal(models.ErrorState))
	})
}

func putObjectPath() {
	Context("Should Put Successfully", func() {
		bucket := "some bucket"
		bucketKey := fmt.Sprintf("PROCESSING/COBRANZA/%s.json", "abc123")

		loanTransactions := []models.LoanPortfolio{
			{LoanId: "abc123"},
			{LoanId: "abc123"},
			{LoanId: "abc123"},
			{LoanId: "abc123"},
		}
		c := MockClient{}
		out := s3.PutObjectOutput{}

		c.On("PutObject", context.TODO(), mock.Anything, mock.Anything, mock.Anything).Return(&out, nil)
		r := S3{client: &c}

		err := r.PutObject(bucket, bucketKey, loanTransactions)
		It("Should not return error", func() {
			Ω(err).To(Equal(models.ACTraceItem{}))
		})
	})
	Context("Should Fail on PutObject and success on retry", func() {
		bucket := "some bucket"
		bucketKey := fmt.Sprintf("PROCESSING/COBRANZA/%s.json", "abc123")
		loanTransactions := []models.LoanPortfolio{
			{LoanId: "abc123"},
			{LoanId: "abc123"},
			{LoanId: "abc123"},
			{LoanId: "abc123"},
		}
		c := MockClient{}
		out := s3.PutObjectOutput{}

		c.On("PutObject", context.TODO(), mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("some error")).Once()
		c.On("PutObject", context.TODO(), mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("some error")).Once()
		c.On("PutObject", context.TODO(), mock.Anything, mock.Anything, mock.Anything).Return(&out, nil).Once()
		r := S3{client: &c}

		err := r.PutObject(bucket, bucketKey, loanTransactions)
		It("Should not return error", func() {
			Ω(err).To(Equal(models.ACTraceItem{}))
		})
		It("Should execute PutObject 3 times", func() {
			Ω(len(c.ExpectedCalls)).To(Equal(3))
		})
	})
	Context("Should Fail on PutObject", func() {
		bucket := "some bucket"
		bucketKey := fmt.Sprintf("PROCESSING/COBRANZA/%s.json", "abc123")
		loanTransactions := []models.LoanPortfolio{
			{LoanId: "abc123"},
			{LoanId: "abc123"},
			{LoanId: "abc123"},
			{LoanId: "abc123"},
		}
		c := MockClient{}

		c.On("PutObject", context.TODO(), mock.Anything, mock.Anything, mock.Anything).Return(nil, errors.New("some error"))
		r := S3{client: &c}

		err := r.PutObject(bucket, bucketKey, loanTransactions)
		It("Should throw error", func() {
			Ω(err).ToNot(BeNil())
		})
		It("Should have structured error", func() {
			Ω(err.State).To(Equal(models.ErrorState))
		})
	})
}
*/
