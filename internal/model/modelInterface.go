package model

type ModelInterface interface {
	Init()         // runs at the very beginning
	SetUp()        // sets up the model
	Go()           // runs the model
	Model() *Model // returns the model
}
