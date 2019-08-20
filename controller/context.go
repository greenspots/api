package controller

import (
	"github.com/greenspots/api/elastic"
)

type (
	AppContext struct {
		EsApi *elastic.ESApi
	}
)

func ProvideAppContext() *AppContext {
	return &AppContext{elastic.NewESApi()}
}

func ProvideDataInitContext(ctx *AppContext) *DataInitContext {
	return &DataInitContext{*ctx}
}

func ProvideEventContext(ctx *AppContext) *EventContext {
	return &EventContext{*ctx}
}
