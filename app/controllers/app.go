package controllers

import (
	"github.com/revel/revel"
	"log"
)

type App struct {
	*revel.Controller
}

func (c App) Index(chart_name, chartt0, chartt1, chartt2, chartt3, chartt4 string) revel.Result {
	c.Validation.Required(chart_name).Message("Please enter the name of your 1st chart!")
	c.Validation.Required(chartt0).Message("Pleaser enter the name of your 2nd char!t")
	c.Validation.Required(chartt1).Message("Please enter the name of your 3rd chart!")
	//c.Validation.Required(chartt2).Message("Please enter the name of your fourth chart!")

	c.Validation.MaxSize(chart_name, 15)
	c.Validation.MinSize(chart_name, 2)

	c.Validation.MaxSize(chartt0, 15)
	c.Validation.MinSize(chartt0, 2)

	c.Validation.MaxSize(chartt1, 15)
	c.Validation.MinSize(chartt1, 2)

	//c.Validation.MaxSize(chartt2, 15)
	//c.Validation.MinSize(chartt2, 2)

	c.Validation.MaxSize(chart_name, 15)
	c.Validation.MinSize(chart_name, 2)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Helm)
	}
	return c.Render(chart_name, chartt0, chartt1, chartt2, chartt3, chartt4)
}

func (c App) Helm() revel.Result {
	if err := c.Request.ParseForm(); err != nil {
		log.Println(err)
	}
	values := c.Request.Form["text"]
	for i := range values {
		log.Println(values[i])
	}
	return c.Render()
}

func (c App) Contact() revel.Result {
	return c.Render()
}
func (c App) About() revel.Result {
	return c.Render()
}
