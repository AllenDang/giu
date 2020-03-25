package imgui

type Platform interface {
	// ShouldStop is regularly called as the abort condition for the program loop.
	ShouldStop() bool
	// ProcessEvents is called once per render loop to dispatch any pending events.
	ProcessEvents()
	// DisplaySize returns the dimension of the display.
	DisplaySize() [2]float32
	// FramebufferSize returns the dimension of the framebuffer.
	FramebufferSize() [2]float32
	// NewFrame marks the begin of a render pass. It must update the imgui IO state according to user input (mouse, keyboard, ...)
	NewFrame()
	// PostRender marks the completion of one render pass. Typically this causes the display buffer to be swapped.
	PostRender()
	// Dispose
	Dispose()
	// Set size change callback
	SetSizeChangeCallback(func(width, height int))
	// Set pos change callback
	SetPosChangeCallback(func(x, y int))
	// Set drop callback
	SetDropCallback(func(names []string))
	// Force Update
	Update()
	// GetContentScale function retrieves the content scale for the specified monitor.
	GetContentScale() float32
	// Get content from system clipboard
	GetClipboard() string
	// Set content to system clipboard
	SetClipboard(content string)
}
