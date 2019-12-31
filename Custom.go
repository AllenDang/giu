package giu

type CustomWidget struct {
	BaseWidget
	callback func()
}

func Custom(callback func()) *CustomWidget {
	return &CustomWidget{
		BaseWidget: BaseWidget{width: 0},
		callback:   callback,
	}
}

func (c *CustomWidget) Build() {
	if c.callback != nil {
		c.callback()
	}
}
