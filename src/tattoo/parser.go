package main

import (
	"time"
	"webapp"
)

type T_FLAGS struct {
	Single bool
	HasTag bool
	Home   bool
	Feed   bool

	WriterOverview bool
	WriterTags     bool
	WriterComments bool
	WriterEditor   bool
}

type T_DATA struct {
	Fn          Export
	Flags       T_FLAGS
	SiteConfig  Config
	ContextInfo webapp.ContextInfo
	SiteVars    map[string]string
	Vars        interface{}
}

func MakeData(ctx *webapp.Context, vars interface{}) T_DATA {
	config := GetConfig()
	ctx.Info.During = time.Now().Sub(ctx.Info.StartTime).Nanoseconds() / 1000.0
	ctx.Info.URL = ctx.Request.URL.Path
	data := T_DATA{
		SiteConfig:  *config,
		ContextInfo: ctx.Info,
		SiteVars:    TattooDB.VarDB.Index,
		Vars:        vars,
	}
	return data
}
