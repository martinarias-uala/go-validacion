package models

import (
	"math"
)

type shapeMetadata struct {
	ID        string `dynamodbav:"id" json:"id"`
	Type      string `dynamodbav:"tipo" json:"type"`
	CreatedBy string `dynamodbav:"creador" json:"created_by"`
}

// This struct is used to insert item to dynamoDB
type ShapeData struct {
	A float64 `dynamodbav:"a" json:"a"`
	B float64 `dynamodbav:"b" json:"b"`
	shapeMetadata
}

type IShape interface {
	CalculateArea() float64
	toDynamoItem(shapeMetadata) ShapeData
}

type Rectangle struct {
	Length float64
	Width  float64
	shapeMetadata
}

type Ellipse struct {
	SemiMajorAxis float64
	SemiMinorAxis float64
	shapeMetadata
}

type Triangle struct {
	Base   float64
	Height float64
	shapeMetadata
}

func (r Rectangle) CalculateArea() float64 {
	return r.Length * r.Width
}

func (r Rectangle) toDynamoItem(data shapeMetadata) ShapeData {
	return ShapeData{
		A:             r.Width,
		B:             r.Length,
		shapeMetadata: data,
	}
}

func (e Ellipse) CalculateArea() float64 {
	return math.Pi * e.SemiMajorAxis * e.SemiMinorAxis
}

func (e Ellipse) toDynamoItem(data shapeMetadata) ShapeData {
	return ShapeData{
		A:             e.SemiMajorAxis,
		B:             e.SemiMinorAxis,
		shapeMetadata: data,
	}
}

func (t Triangle) CalculateArea() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) toDynamoItem(data shapeMetadata) ShapeData {
	return ShapeData{
		A:             t.Base,
		B:             t.Height,
		shapeMetadata: data,
	}
}
