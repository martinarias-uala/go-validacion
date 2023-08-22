package dynamo

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ddb "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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
	table := "devShapes"

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
		return err
	}

	return nil
}
func (d *Dynamo) GetShape(shapeType string, nextToken string) (models.GetShapesResponse, error) {
	table := "devShapes"
	shapes := []models.ShapeData{}
	log.Printf("shapetype :%s", shapeType)

	params, err := attributevalue.MarshalList([]interface{}{shapeType})
	if err != nil {
		log.Printf("<middle> <repository> <GetShape> -  Error marshaling params: %s\n", err.Error())
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
		log.Printf("<middle> <repository> <GetShape> - database connection refused, error: %v\n", err)
		return models.GetShapesResponse{}, err
	}

	err = attributevalue.UnmarshalListOfMaps(data.Items, &shapes)

	if err != nil {
		log.Printf("<middle> <repository> <GetShapes> - decoding fail, error: %v\n", err)
		return models.GetShapesResponse{}, err
	}

	return models.GetShapesResponse{
		ShapesData: shapes,
		PageToken:  data.NextToken,
	}, nil

}
