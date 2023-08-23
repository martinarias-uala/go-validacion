package dynamo

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ddb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/martinarias-uala/go-validacion/internal/logs"
	"github.com/martinarias-uala/go-validacion/pkg/models"
)

type DynamoClient interface {
	PutItem(context.Context, *ddb.PutItemInput, ...func(*ddb.Options)) (*ddb.PutItemOutput, error)
	ExecuteStatement(context.Context, *ddb.ExecuteStatementInput, ...func(*ddb.Options)) (*ddb.ExecuteStatementOutput, error)
}

type DynamoRepository interface {
	CreateItem(shape models.ShapeData) error
	GetShape(shapeType, nextToken string) (models.GetShapesResponse, error) //GET A BETTER NAME FOR THIS STRUCT
}

type Dynamo struct {
	client DynamoClient
}

func New() *Dynamo {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("configuration error, " + err.Error())
	}

	client := ddb.NewFromConfig(cfg)
	return &Dynamo{
		client: client,
	}
}

func (d *Dynamo) CreateItem(shape models.ShapeData) error {
	logger := logs.GetLoggerInstance()
	table := "devShapes"

	logger.Info().Msg("<CreateItem> Starting to create item in DynamoDB")

	_, err := d.client.PutItem(context.TODO(), &ddb.PutItemInput{
		TableName: aws.String(table),
		Item: map[string]types.AttributeValue{
			"id":      &types.AttributeValueMemberS{Value: shape.ID},
			"tipo":    &types.AttributeValueMemberS{Value: shape.Type},
			"a":       &types.AttributeValueMemberN{Value: fmt.Sprintf("%f", shape.A)},
			"b":       &types.AttributeValueMemberN{Value: fmt.Sprintf("%f", shape.B)},
			"creador": &types.AttributeValueMemberS{Value: shape.CreatedBy},
		},
	})
	if err != nil {
		logger.Error().Msg(fmt.Sprintf("<CreateItem> Error creating item: %s", err.Error()))
		return err
	}
	logger.Info().Msg("<CreateItem> Item created successfully")

	return nil
}
func (d *Dynamo) GetShape(shapeType string, nextToken string) (models.GetShapesResponse, error) {
	table := "devShapes"
	shapes := []models.ShapeData{}
	logger := logs.GetLoggerInstance()

	logger.Info().Msg("<GetShape> Starting to get items from DynamoDB")

	params, err := attributevalue.MarshalList([]interface{}{shapeType})
	if err != nil {
		logger.Error().Msg(fmt.Sprintf("<GetShape> Error marshaling params: %s", err.Error()))
		return models.GetShapesResponse{}, err
	}
	statement := &dynamodb.ExecuteStatementInput{
		Statement: aws.String(
			fmt.Sprintf("SELECT * FROM \"%v\" WHERE tipo=?",
				table)),
		Parameters: params,
		Limit:      aws.Int32(17),
	}

	if len(nextToken) > 1 {
		statement.NextToken = &nextToken
	}

	data, err := d.client.ExecuteStatement(context.TODO(), statement)
	if err != nil {
		logger.Error().Msg(fmt.Sprintf("<GetShape> Error database connection refused: %s", err.Error()))
		return models.GetShapesResponse{}, err
	}

	err = attributevalue.UnmarshalListOfMaps(data.Items, &shapes)

	if err != nil {
		logger.Error().Msg(fmt.Sprintf("<GetShape> Error decoding db response failed: %s", err.Error()))
		return models.GetShapesResponse{}, err
	}

	logger.Info().Msg("<GetShape> Items retrieved successfully from DynamoDB")
	return models.GetShapesResponse{
		ShapesData: shapes,
		PageToken:  data.NextToken,
	}, nil

}
