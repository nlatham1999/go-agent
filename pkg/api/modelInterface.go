package api

import "github.com/nlatham1999/go-agent/pkg/model"

type ModelInterface interface {
	Init()        // runs at the very beginning
	SetUp() error // sets up the model
	Go()          // runs the model

	Model() *model.Model           // returns the model
	Stats() map[string]interface{} //returns the stats of the model
	Stop() bool                    // on whether to stop the model
	Widgets() []Widget             // returns the widgets of the model
}
