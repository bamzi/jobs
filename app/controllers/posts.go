package controllers

import "github.com/robfig/revel"

type Posts struct {
	*revel.Controller
}

func (c Posts) Index() revel.Result {
	return c.Render()
}

func (c Posts) New() revel.Result {
	return c.Render()
}

func (c Posts) Create() revel.Result {
	return c.Render()
}
