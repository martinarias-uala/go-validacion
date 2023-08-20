package models

import (
	"math"
)

type ShapeMetadata struct {
	ID        string  `dynamodbav:"id" json:"id"`
	Type      string  `dynamodbav:"tipo" json:"type"`
	CreatedBy string  `dynamodbav:"creador" json:"created_by"`
	Area      float64 `json:"area"`
}

// This struct is used to insert item to dynamoDB
type ShapeData struct {
	A float64 `dynamodbav:"a" json:"a"`
	B float64 `dynamodbav:"b" json:"b"`
	ShapeMetadata
}

type IShape interface {
	CalculateArea() float64
	ToGenericShape(data ShapeMetadata) ShapeData
}

type Rectangle struct {
	Length float64
	Width  float64
	ShapeMetadata
}

type Ellipse struct {
	SemiMajorAxis float64
	SemiMinorAxis float64
	ShapeMetadata
}

type Triangle struct {
	Base   float64
	Height float64
	ShapeMetadata
}

func (r Rectangle) CalculateArea() float64 {
	return r.Length * r.Width
}

func (r Rectangle) ToGenericShape(data ShapeMetadata) ShapeData {
	return ShapeData{
		A:             r.Width,
		B:             r.Length,
		ShapeMetadata: data,
	}
}

func (e Ellipse) CalculateArea() float64 {
	return math.Pi * e.SemiMajorAxis * e.SemiMinorAxis
}

func (e Ellipse) ToGenericShape(data ShapeMetadata) ShapeData {
	return ShapeData{
		A:             e.SemiMajorAxis,
		B:             e.SemiMinorAxis,
		ShapeMetadata: data,
	}
}

func (t Triangle) CalculateArea() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) ToGenericShape(data ShapeMetadata) ShapeData {
	return ShapeData{
		A:             t.Base,
		B:             t.Height,
		ShapeMetadata: data,
	}
}
