package imgui

const (
	// WindowFlagsNone default = 0
	WindowFlagsNone = 0
	// WindowFlagsNoTitleBar disables title-bar.
	WindowFlagsNoTitleBar = 1 << 0
	// WindowFlagsNoResize disables user resizing with the lower-right grip.
	WindowFlagsNoResize = 1 << 1
	// WindowFlagsNoMove disables user moving the window.
	WindowFlagsNoMove = 1 << 2
	// WindowFlagsNoScrollbar disables scrollbars. Window can still scroll with mouse or programmatically.
	WindowFlagsNoScrollbar = 1 << 3
	// WindowFlagsNoScrollWithMouse disables user vertically scrolling with mouse wheel. On child window, mouse wheel
	// will be forwarded to the parent unless NoScrollbar is also set.
	WindowFlagsNoScrollWithMouse = 1 << 4
	// WindowFlagsNoCollapse disables user collapsing window by double-clicking on it.
	WindowFlagsNoCollapse = 1 << 5
	// WindowFlagsAlwaysAutoResize resizes every window to its content every frame.
	WindowFlagsAlwaysAutoResize = 1 << 6
	// WindowFlagsNoBackground disables drawing background color (WindowBg, etc.) and outside border. Similar as using
	// SetNextWindowBgAlpha(0.0f).
	WindowFlagsNoBackground = 1 << 7
	// WindowFlagsNoSavedSettings will never load/save settings in .ini file.
	WindowFlagsNoSavedSettings = 1 << 8
	// WindowFlagsNoMouseInputs disables catching mouse, hovering test with pass through.
	WindowFlagsNoMouseInputs = 1 << 9
	// WindowFlagsMenuBar has a menu-bar.
	WindowFlagsMenuBar = 1 << 10
	// WindowFlagsHorizontalScrollbar allows horizontal scrollbar to appear (off by default). You may use
	// SetNextWindowContentSize(ImVec2(width,0.0f)); prior to calling Begin() to specify width. Read code in imgui_demo
	// in the "Horizontal Scrolling" section.
	WindowFlagsHorizontalScrollbar = 1 << 11
	// WindowFlagsNoFocusOnAppearing disables taking focus when transitioning from hidden to visible state.
	WindowFlagsNoFocusOnAppearing = 1 << 12
	// WindowFlagsNoBringToFrontOnFocus disables bringing window to front when taking focus. e.g. clicking on it or
	// programmatically giving it focus.
	WindowFlagsNoBringToFrontOnFocus = 1 << 13
	// WindowFlagsAlwaysVerticalScrollbar always shows vertical scrollbar, even if ContentSize.y < Size.y .
	WindowFlagsAlwaysVerticalScrollbar = 1 << 14
	// WindowFlagsAlwaysHorizontalScrollbar always shows horizontal scrollbar, even if ContentSize.x < Size.x .
	WindowFlagsAlwaysHorizontalScrollbar = 1 << 15
	// WindowFlagsAlwaysUseWindowPadding ensures child windows without border uses style.WindowPadding (ignored by
	// default for non-bordered child windows, because more convenient).
	WindowFlagsAlwaysUseWindowPadding = 1 << 16
	// WindowFlagsNoNavInputs has no gamepad/keyboard navigation within the window.
	WindowFlagsNoNavInputs = 1 << 18
	// WindowFlagsNoNavFocus has no focusing toward this window with gamepad/keyboard navigation
	// (e.g. skipped by CTRL+TAB)
	WindowFlagsNoNavFocus = 1 << 19
	// WindowFlagsUnsavedDocument appends '*' to title without affecting the ID, as a convenience to avoid using the
	// ### operator. When used in a tab/docking context, tab is selected on closure and closure is deferred by one
	// frame to allow code to cancel the closure (with a confirmation popup, etc.) without flicker.
	WindowFlagsUnsavedDocument = 1 << 20

	// WindowFlagsNoNav combines WindowFlagsNoNavInputs and WindowFlagsNoNavFocus.
	WindowFlagsNoNav = WindowFlagsNoNavInputs | WindowFlagsNoNavFocus
	// WindowFlagsNoDecoration combines WindowFlagsNoTitleBar, WindowFlagsNoResize, WindowFlagsNoScrollbar and
	// WindowFlagsNoCollapse.
	WindowFlagsNoDecoration = WindowFlagsNoTitleBar | WindowFlagsNoResize | WindowFlagsNoScrollbar | WindowFlagsNoCollapse
	// WindowFlagsNoInputs combines WindowFlagsNoMouseInputs, WindowFlagsNoNavInputs and WindowFlagsNoNavFocus.
	WindowFlagsNoInputs = WindowFlagsNoMouseInputs | WindowFlagsNoNavInputs | WindowFlagsNoNavFocus
)
