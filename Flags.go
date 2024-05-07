package giu

import imgui "github.com/AllenDang/cimgui-go"

// InputTextFlags represents input text flags.
type InputTextFlags imgui.InputTextFlags

// input text flags.
const (
	// InputTextFlagsNone sets everything default.
	InputTextFlagsNone = InputTextFlags(imgui.InputTextFlagsNone)
	// InputTextFlagsCharsDecimal allows 0123456789.+-.
	InputTextFlagsCharsDecimal = InputTextFlags(imgui.InputTextFlagsCharsDecimal)
	// InputTextFlagsCharsHexadecimal allow 0123456789ABCDEFabcdef.
	InputTextFlagsCharsHexadecimal = InputTextFlags(imgui.InputTextFlagsCharsHexadecimal)
	// InputTextFlagsCharsUppercase turns a..z into A..Z.
	InputTextFlagsCharsUppercase = InputTextFlags(imgui.InputTextFlagsCharsUppercase)
	// InputTextFlagsCharsNoBlank filters out spaces, tabs.
	InputTextFlagsCharsNoBlank = InputTextFlags(imgui.InputTextFlagsCharsNoBlank)
	// InputTextFlagsAutoSelectAll selects entire text when first taking mouse focus.
	InputTextFlagsAutoSelectAll = InputTextFlags(imgui.InputTextFlagsAutoSelectAll)
	// InputTextFlagsEnterReturnsTrue returns 'true' when Enter is pressed (as opposed to when the value was modified).
	InputTextFlagsEnterReturnsTrue = InputTextFlags(imgui.InputTextFlagsEnterReturnsTrue)
	// InputTextFlagsCallbackCompletion for callback on pressing TAB (for completion handling).
	InputTextFlagsCallbackCompletion = InputTextFlags(imgui.InputTextFlagsCallbackCompletion)
	// InputTextFlagsCallbackHistory for callback on pressing Up/Down arrows (for history handling).
	InputTextFlagsCallbackHistory = InputTextFlags(imgui.InputTextFlagsCallbackHistory)
	// InputTextFlagsCallbackAlways for callback on each iteration. User code may query cursor position, modify text buffer.
	InputTextFlagsCallbackAlways = InputTextFlags(imgui.InputTextFlagsCallbackAlways)
	// InputTextFlagsCallbackCharFilter for callback on character inputs to replace or discard them.
	// Modify 'EventChar' to replace or discard, or return 1 in callback to discard.
	InputTextFlagsCallbackCharFilter = InputTextFlags(imgui.InputTextFlagsCallbackCharFilter)
	// InputTextFlagsAllowTabInput when pressing TAB to input a '\t' character into the text field.
	InputTextFlagsAllowTabInput = InputTextFlags(imgui.InputTextFlagsAllowTabInput)
	// InputTextFlagsCtrlEnterForNewLine in multi-line mode, unfocus with Enter, add new line with Ctrl+Enter
	// (default is opposite: unfocus with Ctrl+Enter, add line with Enter).
	InputTextFlagsCtrlEnterForNewLine = InputTextFlags(imgui.InputTextFlagsCtrlEnterForNewLine)
	// InputTextFlagsNoHorizontalScroll disables following the cursor horizontally.
	InputTextFlagsNoHorizontalScroll = InputTextFlags(imgui.InputTextFlagsNoHorizontalScroll)
	// InputTextFlagsAlwaysInsertMode sets insert mode.

	// InputTextFlagsReadOnly sets read-only mode.
	InputTextFlagsReadOnly = InputTextFlags(imgui.InputTextFlagsReadOnly)
	// InputTextFlagsPassword sets password mode, display all characters as '*'.
	InputTextFlagsPassword = InputTextFlags(imgui.InputTextFlagsPassword)
	// InputTextFlagsNoUndoRedo disables undo/redo. Note that input text owns the text data while active,
	// if you want to provide your own undo/redo stack you need e.g. to call ClearActiveID().
	InputTextFlagsNoUndoRedo = InputTextFlags(imgui.InputTextFlagsNoUndoRedo)
	// InputTextFlagsCharsScientific allows 0123456789.+-*/eE (Scientific notation input).
	InputTextFlagsCharsScientific = InputTextFlags(imgui.InputTextFlagsCharsScientific)
)

// WindowFlags represents a window flags (see (*WindowWidget).Flags.
type WindowFlags imgui.GLFWWindowFlags

// window flags.
const (
	// WindowFlagsNone default = 0.
	WindowFlagsNone WindowFlags = WindowFlags(imgui.WindowFlagsNone)
	// WindowFlagsNoTitleBar disables title-bar.
	WindowFlagsNoTitleBar WindowFlags = WindowFlags(imgui.WindowFlagsNoTitleBar)
	// WindowFlagsNoResize disables user resizing with the lower-right grip.
	WindowFlagsNoResize WindowFlags = WindowFlags(imgui.WindowFlagsNoResize)
	// WindowFlagsNoMove disables user moving the window.
	WindowFlagsNoMove WindowFlags = WindowFlags(imgui.WindowFlagsNoMove)
	// WindowFlagsNoScrollbar disables scrollbars. Window can still scroll with mouse or programmatically.
	WindowFlagsNoScrollbar WindowFlags = WindowFlags(imgui.WindowFlagsNoScrollbar)
	// WindowFlagsNoScrollWithMouse disables user vertically scrolling with mouse wheel. On child window, mouse wheel
	// will be forwarded to the parent unless NoScrollbar is also set.
	WindowFlagsNoScrollWithMouse WindowFlags = WindowFlags(imgui.WindowFlagsNoScrollWithMouse)
	// WindowFlagsNoCollapse disables user collapsing window by double-clicking on it.
	WindowFlagsNoCollapse WindowFlags = WindowFlags(imgui.WindowFlagsNoCollapse)
	// WindowFlagsAlwaysAutoResize resizes every window to its content every frame.
	WindowFlagsAlwaysAutoResize WindowFlags = WindowFlags(imgui.WindowFlagsAlwaysAutoResize)
	// WindowFlagsNoBackground disables drawing background color (WindowBg, etc.) and outside border. Similar as using
	// SetNextWindowBgAlpha(0.0f).
	WindowFlagsNoBackground WindowFlags = WindowFlags(imgui.WindowFlagsNoBackground)
	// WindowFlagsNoSavedSettings will never load/save settings in .ini file.
	WindowFlagsNoSavedSettings WindowFlags = WindowFlags(imgui.WindowFlagsNoSavedSettings)
	// WindowFlagsNoMouseInputs disables catching mouse, hovering test with pass through.
	WindowFlagsNoMouseInputs WindowFlags = WindowFlags(imgui.WindowFlagsNoMouseInputs)
	// WindowFlagsMenuBar has a menu-bar.
	WindowFlagsMenuBar WindowFlags = WindowFlags(imgui.WindowFlagsMenuBar)
	// WindowFlagsHorizontalScrollbar allows horizontal scrollbar to appear (off by default). You may use
	// SetNextWindowContentSize(ImVec2(width,0.0f)); prior to calling Begin() to specify width. Read code in imgui_demo
	// in the "Horizontal Scrolling" section.
	WindowFlagsHorizontalScrollbar WindowFlags = WindowFlags(imgui.WindowFlagsHorizontalScrollbar)
	// WindowFlagsNoFocusOnAppearing disables taking focus when transitioning from hidden to visible state.
	WindowFlagsNoFocusOnAppearing WindowFlags = WindowFlags(imgui.WindowFlagsNoFocusOnAppearing)
	// WindowFlagsNoBringToFrontOnFocus disables bringing window to front when taking focus. e.g. clicking on it or
	// programmatically giving it focus.
	WindowFlagsNoBringToFrontOnFocus WindowFlags = WindowFlags(imgui.WindowFlagsNoBringToFrontOnFocus)
	// WindowFlagsAlwaysVerticalScrollbar always shows vertical scrollbar, even if ContentSize.y < Size.y .
	WindowFlagsAlwaysVerticalScrollbar WindowFlags = WindowFlags(imgui.WindowFlagsAlwaysVerticalScrollbar)
	// WindowFlagsAlwaysHorizontalScrollbar always shows horizontal scrollbar, even if ContentSize.x < Size.x .
	WindowFlagsAlwaysHorizontalScrollbar WindowFlags = WindowFlags(imgui.WindowFlagsAlwaysHorizontalScrollbar)
	// WindowFlagsNoNavInputs has no gamepad/keyboard navigation within the window.
	WindowFlagsNoNavInputs WindowFlags = WindowFlags(imgui.WindowFlagsNoNavInputs)
	// WindowFlagsNoNavFocus has no focusing toward this window with gamepad/keyboard navigation
	// (e.g. skipped by CTRL+TAB).
	WindowFlagsNoNavFocus WindowFlags = WindowFlags(imgui.WindowFlagsNoNavFocus)
	// WindowFlagsUnsavedDocument appends '*' to title without affecting the ID, as a convenience to avoid using the
	// ### operator. When used in a tab/docking context, tab is selected on closure and closure is deferred by one
	// frame to allow code to cancel the closure (with a confirmation popup, etc.) without flicker.
	WindowFlagsUnsavedDocument WindowFlags = WindowFlags(imgui.WindowFlagsUnsavedDocument)

	// WindowFlagsNoNav combines WindowFlagsNoNavInputs and WindowFlagsNoNavFocus.
	WindowFlagsNoNav WindowFlags = WindowFlags(imgui.WindowFlagsNoNav)
	// WindowFlagsNoDecoration combines WindowFlagsNoTitleBar, WindowFlagsNoResize, WindowFlagsNoScrollbar and
	// WindowFlagsNoCollapse.
	WindowFlagsNoDecoration WindowFlags = WindowFlags(imgui.WindowFlagsNoDecoration)
	// WindowFlagsNoInputs combines WindowFlagsNoMouseInputs, WindowFlagsNoNavInputs and WindowFlagsNoNavFocus.
	WindowFlagsNoInputs WindowFlags = WindowFlags(imgui.WindowFlagsNoInputs)
)

// ComboFlags represents imgui.ComboFlags.
type ComboFlags imgui.ComboFlags

// combo flags list.
const (
	// ComboFlagsNone default = 0.
	ComboFlagsNone = ComboFlags(imgui.ComboFlagsNone)
	// ComboFlagsPopupAlignLeft aligns the popup toward the left by default.
	ComboFlagsPopupAlignLeft = ComboFlags(imgui.ComboFlagsPopupAlignLeft)
	// ComboFlagsHeightSmall has maxValue ~4 items visible.
	// Tip: If you want your combo popup to be a specific size you can use SetNextWindowSizeConstraints() prior to calling BeginCombo().
	ComboFlagsHeightSmall = ComboFlags(imgui.ComboFlagsHeightSmall)
	// ComboFlagsHeightRegular has maxValue ~8 items visible (default).
	ComboFlagsHeightRegular = ComboFlags(imgui.ComboFlagsHeightRegular)
	// ComboFlagsHeightLarge has maxValue ~20 items visible.
	ComboFlagsHeightLarge = ComboFlags(imgui.ComboFlagsHeightLarge)
	// ComboFlagsHeightLargest has as many fitting items as possible.
	ComboFlagsHeightLargest = ComboFlags(imgui.ComboFlagsHeightLargest)
	// ComboFlagsNoArrowButton displays on the preview box without the square arrow button.
	ComboFlagsNoArrowButton = ComboFlags(imgui.ComboFlagsNoArrowButton)
	// ComboFlagsNoPreview displays only a square arrow button.
	ComboFlagsNoPreview = ComboFlags(imgui.ComboFlagsNoPreview)
)

// SelectableFlags represents imgui.SelectableFlags.
type SelectableFlags imgui.SelectableFlags

// selectable flags list.
const (
	// SelectableFlagsNone default = 0.
	SelectableFlagsNone = SelectableFlags(imgui.SelectableFlagsNone)
	// SelectableFlagsDontClosePopups makes clicking the selectable not close any parent popup windows.
	SelectableFlagsDontClosePopups = SelectableFlags(imgui.SelectableFlagsDontClosePopups)
	// SelectableFlagsSpanAllColumns allows the selectable frame to span all columns (text will still fit in current column).
	SelectableFlagsSpanAllColumns = SelectableFlags(imgui.SelectableFlagsSpanAllColumns)
	// SelectableFlagsAllowDoubleClick generates press events on double clicks too.
	SelectableFlagsAllowDoubleClick = SelectableFlags(imgui.SelectableFlagsAllowDoubleClick)
	// SelectableFlagsDisabled disallows selection and displays text in a greyed out color.
	SelectableFlagsDisabled = SelectableFlags(imgui.SelectableFlagsDisabled)
)

// TabItemFlags represents tab item flags.
type TabItemFlags imgui.TabItemFlags

// tab item flags list.
const (
	// TabItemFlagsNone default = 0.
	TabItemFlagsNone = TabItemFlags(imgui.TabItemFlagsNone)
	// TabItemFlagsUnsavedDocument Append '*' to title without affecting the ID, as a convenience to avoid using the
	// ### operator. Also: tab is selected on closure and closure is deferred by one frame to allow code to undo it
	// without flicker.
	TabItemFlagsUnsavedDocument = TabItemFlags(imgui.TabItemFlagsUnsavedDocument)
	// TabItemFlagsSetSelected Trigger flag to programmatically make the tab selected when calling BeginTabItem().
	TabItemFlagsSetSelected = TabItemFlags(imgui.TabItemFlagsSetSelected)
	// TabItemFlagsNoCloseWithMiddleMouseButton  Disable behavior of closing tabs (that are submitted with
	// p_open != NULL) with middle mouse button. You can still repro this behavior on user's side with if
	// (IsItemHovered() && IsMouseClicked(2)) *p_open = false.
	TabItemFlagsNoCloseWithMiddleMouseButton = TabItemFlags(imgui.TabItemFlagsNoCloseWithMiddleMouseButton)
	// TabItemFlagsNoPushID Don't call PushID(tab->ID)/PopID() on BeginTabItem()/EndTabItem().

)

// TabBarFlags represents imgui.TabBarFlags.
type TabBarFlags imgui.TabBarFlags

// tab bar flags list.
const (
	// TabBarFlagsNone default = 0.
	TabBarFlagsNone = TabBarFlags(imgui.TabBarFlagsNone)
	// TabBarFlagsReorderable Allow manually dragging tabs to re-order them + New tabs are appended at the end of list.
	TabBarFlagsReorderable = TabBarFlags(imgui.TabBarFlagsReorderable)
	// TabBarFlagsAutoSelectNewTabs Automatically select new tabs when they appear.
	TabBarFlagsAutoSelectNewTabs = TabBarFlags(imgui.TabBarFlagsAutoSelectNewTabs)
	// TabBarFlagsTabListPopupButton Disable buttons to open the tab list popup.
	TabBarFlagsTabListPopupButton = TabBarFlags(imgui.TabBarFlagsTabListPopupButton)
	// TabBarFlagsNoCloseWithMiddleMouseButton Disable behavior of closing tabs (that are submitted with p_open != NULL)
	// with middle mouse button. You can still repro this behavior on user's side with if
	// (IsItemHovered() && IsMouseClicked(2)) *p_open = false.
	TabBarFlagsNoCloseWithMiddleMouseButton = TabBarFlags(imgui.TabBarFlagsNoCloseWithMiddleMouseButton)
	// TabBarFlagsNoTabListScrollingButtons Disable scrolling buttons (apply when fitting policy is
	// TabBarFlagsFittingPolicyScroll).
	TabBarFlagsNoTabListScrollingButtons = TabBarFlags(imgui.TabBarFlagsNoTabListScrollingButtons)
	// TabBarFlagsNoTooltip Disable tooltips when hovering a tab.
	TabBarFlagsNoTooltip = TabBarFlags(imgui.TabBarFlagsNoTooltip)
	// TabBarFlagsFittingPolicyResizeDown Resize tabs when they don't fit.
	TabBarFlagsFittingPolicyResizeDown = TabBarFlags(imgui.TabBarFlagsFittingPolicyResizeDown)
	// TabBarFlagsFittingPolicyScroll Add scroll buttons when tabs don't fit.
	TabBarFlagsFittingPolicyScroll = TabBarFlags(imgui.TabBarFlagsFittingPolicyScroll)
	// TabBarFlagsFittingPolicyMask combines
	// TabBarFlagsFittingPolicyResizeDown and TabBarFlagsFittingPolicyScroll.
	TabBarFlagsFittingPolicyMask = TabBarFlags(imgui.TabBarFlagsFittingPolicyMask)
	// TabBarFlagsFittingPolicyDefault alias for TabBarFlagsFittingPolicyResizeDown.
	TabBarFlagsFittingPolicyDefault = TabBarFlags(imgui.TabBarFlagsFittingPolicyDefault)
)

// TreeNodeFlags represents tree node widget flags.
type TreeNodeFlags imgui.TreeNodeFlags

// tree node flags list.
const (
	// TreeNodeFlagsNone default = 0.
	TreeNodeFlagsNone = TreeNodeFlags(imgui.TreeNodeFlagsNone)
	// TreeNodeFlagsSelected draws as selected.
	TreeNodeFlagsSelected = TreeNodeFlags(imgui.TreeNodeFlagsSelected)
	// TreeNodeFlagsFramed draws full colored frame (e.g. for CollapsingHeader).
	TreeNodeFlagsFramed = TreeNodeFlags(imgui.TreeNodeFlagsFramed)
	// TreeNodeFlagsAllowItemOverlap hit testing to allow subsequent widgets to overlap this one.
	TreeNodeFlagsAllowItemOverlap = TreeNodeFlags(imgui.TreeNodeFlagsAllowOverlap)
	// TreeNodeFlagsNoTreePushOnOpen doesn't do a TreePush() when open
	// (e.g. for CollapsingHeader) = no extra indent nor pushing on ID stack.
	TreeNodeFlagsNoTreePushOnOpen = TreeNodeFlags(imgui.TreeNodeFlagsNoTreePushOnOpen)
	// TreeNodeFlagsNoAutoOpenOnLog doesn't automatically and temporarily open node when Logging is active
	// (by default logging will automatically open tree nodes).
	TreeNodeFlagsNoAutoOpenOnLog = TreeNodeFlags(imgui.TreeNodeFlagsNoAutoOpenOnLog)
	// TreeNodeFlagsDefaultOpen defaults node to be open.
	TreeNodeFlagsDefaultOpen = TreeNodeFlags(imgui.TreeNodeFlagsDefaultOpen)
	// TreeNodeFlagsOpenOnDoubleClick needs double-click to open node.
	TreeNodeFlagsOpenOnDoubleClick = TreeNodeFlags(imgui.TreeNodeFlagsOpenOnDoubleClick)
	// TreeNodeFlagsOpenOnArrow opens only when clicking on the arrow part.
	// If TreeNodeFlagsOpenOnDoubleClick is also set, single-click arrow or double-click all box to open.
	TreeNodeFlagsOpenOnArrow = TreeNodeFlags(imgui.TreeNodeFlagsOpenOnArrow)
	// TreeNodeFlagsLeaf allows no collapsing, no arrow (use as a convenience for leaf nodes).
	TreeNodeFlagsLeaf = TreeNodeFlags(imgui.TreeNodeFlagsLeaf)
	// TreeNodeFlagsBullet displays a bullet instead of an arrow.
	TreeNodeFlagsBullet = TreeNodeFlags(imgui.TreeNodeFlagsBullet)
	// TreeNodeFlagsFramePadding uses FramePadding (even for an unframed text node) to
	// vertically align text baseline to regular widget height. Equivalent to calling AlignTextToFramePadding().
	TreeNodeFlagsFramePadding = TreeNodeFlags(imgui.TreeNodeFlagsFramePadding)
	// TreeNodeFlagsSpanAvailWidth extends hit box to the right-most edge, even if not framed.
	// This is not the default in order to allow adding other items on the same line.
	// In the future we may refactor the hit system to be front-to-back, allowing natural overlaps
	// and then this can become the default.
	TreeNodeFlagsSpanAvailWidth = TreeNodeFlags(imgui.TreeNodeFlagsSpanAvailWidth)
	// TreeNodeFlagsSpanFullWidth extends hit box to the left-most and right-most edges (bypass the indented area).
	TreeNodeFlagsSpanFullWidth = TreeNodeFlags(imgui.TreeNodeFlagsSpanFullWidth)
	// TreeNodeFlagsNavLeftJumpsBackHere (WIP) Nav: left direction may move to this TreeNode() from any of its child
	// (items submitted between TreeNode and TreePop).
	TreeNodeFlagsNavLeftJumpsBackHere = TreeNodeFlags(imgui.TreeNodeFlagsNavLeftJumpsBackHere)
	// TreeNodeFlagsCollapsingHeader combines TreeNodeFlagsFramed and TreeNodeFlagsNoAutoOpenOnLog.
	TreeNodeFlagsCollapsingHeader = TreeNodeFlags(imgui.TreeNodeFlagsCollapsingHeader)
)

// FocusedFlags represents imgui.FocusedFlags.
type FocusedFlags imgui.FocusedFlags

// focused flags list.
const (
	FocusedFlagsNone             = (imgui.FocusedFlagsNone)
	FocusedFlagsChildWindows     = (imgui.FocusedFlagsChildWindows)   // Return true if any children of the window is focused
	FocusedFlagsRootWindow       = (imgui.FocusedFlagsRootWindow)     // Test from root window (top most parent of the current hierarchy)
	FocusedFlagsAnyWindow        = (imgui.FocusedFlagsAnyWindow)      // Return true if any window is focused. Important: If you are trying to tell how to dispatch your low-level inputs do NOT use this. Use 'io.WantCaptureMouse' instead! Please read the FAQ!
	FocusedFlagsNoPopupHierarchy = imgui.FocusedFlagsNoPopupHierarchy // Do not consider popup hierarchy (do not treat popup emitter as parent of popup) (when used with ChildWindows or RootWindow)
	// FocusedFlagsDockHierarchy               = 1 << 4   // Consider docking hierarchy (treat dockspace host as parent of docked window) (when used with ChildWindows or RootWindow).
	FocusedFlagsRootAndChildWindows = imgui.FocusedFlagsRootAndChildWindows
)

// HoveredFlags represents a hovered flags.
type HoveredFlags imgui.HoveredFlags

// hovered flags list.
const (
	// HoveredFlagsNone Return true if directly over the item/window, not obstructed by another window,
	// not obstructed by an active popup or modal blocking inputs under them.
	HoveredFlagsNone = HoveredFlags(imgui.HoveredFlagsNone)
	// HoveredFlagsChildWindows IsWindowHovered() only: Return true if any children of the window is hovered.
	HoveredFlagsChildWindows = HoveredFlags(imgui.HoveredFlagsChildWindows)
	// HoveredFlagsRootWindow IsWindowHovered() only: Test from root window (top most parent of the current hierarchy).
	HoveredFlagsRootWindow = HoveredFlags(imgui.HoveredFlagsRootWindow)
	// HoveredFlagsAnyWindow IsWindowHovered() only: Return true if any window is hovered.
	HoveredFlagsAnyWindow = HoveredFlags(imgui.HoveredFlagsAnyWindow)
	// HoveredFlagsAllowWhenBlockedByPopup Return true even if a popup window is normally blocking access to this item/window.
	HoveredFlagsAllowWhenBlockedByPopup = HoveredFlags(imgui.HoveredFlagsAllowWhenBlockedByPopup)
	// HoveredFlagsAllowWhenBlockedByActiveItem Return true even if an active item is blocking access to this item/window.
	// Useful for Drag and Drop patterns.
	HoveredFlagsAllowWhenBlockedByActiveItem = HoveredFlags(imgui.HoveredFlagsAllowWhenBlockedByActiveItem)
	// HoveredFlagsAllowWhenOverlapped Return true even if the position is overlapped by another window.
	HoveredFlagsAllowWhenOverlapped = HoveredFlags(imgui.HoveredFlagsAllowWhenOverlapped)
	// HoveredFlagsAllowWhenDisabled Return true even if the item is disabled.
	HoveredFlagsAllowWhenDisabled = HoveredFlags(imgui.HoveredFlagsAllowWhenDisabled)
)

// ColorEditFlags for ColorEdit3V(), etc.
type ColorEditFlags int

// list of color edit flags.
const (
	ColorEditFlagsNone ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsNone)
	//              // ColorEdit, ColorPicker, ColorButton: ignore Alpha component (will only read 3 components from the input pointer).
	ColorEditFlagsNoAlpha ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsNoAlpha)
	//              // ColorEdit: disable picker when clicking on color square.
	ColorEditFlagsNoPicker ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsNoPicker)
	//              // ColorEdit: disable toggling options menu when right-clicking on inputs/small preview.
	ColorEditFlagsNoOptions ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsNoOptions)
	//              // ColorEdit, ColorPicker: disable color square preview next to the inputs. (e.g. to show only the inputs)
	ColorEditFlagsNoSmallPreview ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsNoSmallPreview)
	//              // ColorEdit, ColorPicker: disable inputs sliders/text widgets (e.g. to show only the small preview color square).
	ColorEditFlagsNoInputs ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsNoInputs)
	//              // ColorEdit, ColorPicker, ColorButton: disable tooltip when hovering the preview.
	ColorEditFlagsNoTooltip ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsNoTooltip)
	//              // ColorEdit, ColorPicker: disable display of inline text label (the label is still forwarded to the tooltip and picker).
	ColorEditFlagsNoLabel ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsNoLabel)
	//              // ColorPicker: disable bigger color preview on right side of the picker, use small color square preview instead.
	ColorEditFlagsNoSidePreview ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsNoSidePreview)
	//              // ColorEdit: disable drag and drop target. ColorButton: disable drag and drop source.
	ColorEditFlagsNoDragDrop ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsNoDragDrop)
	//              // ColorButton: disable border (which is enforced by default)
	ColorEditFlagsNoBorder ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsNoBorder)
	//              // ColorEdit, ColorPicker: show vertical alpha bar/gradient in picker.
	ColorEditFlagsAlphaBar ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsAlphaBar)
	//              // ColorEdit, ColorPicker, ColorButton: display preview as a transparent color over a checkerboard, instead of opaque.
	ColorEditFlagsAlphaPreview ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsAlphaPreview)
	//              // ColorEdit, ColorPicker, ColorButton: display half opaque / half checkerboard, instead of opaque.
	ColorEditFlagsAlphaPreviewHalf ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsAlphaPreviewHalf)
	//              // (WIP) ColorEdit: Currently only disable 0.0f..1.0f limits in RGBA edition (note: you probably want to use ImGuiColorEditFlags_Float flag as well).
	ColorEditFlagsHDR ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsHDR)
	// [Display]    // ColorEdit: override _display_ type among RGB/HSV/Hex. ColorPicker: select any combination using one or more of RGB/HSV/Hex.
	ColorEditFlagsDisplayRGB ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsDisplayRGB)
	// [Display]    // ".
	ColorEditFlagsDisplayHSV ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsDisplayHSV)
	// [Display]    // ".
	ColorEditFlagsDisplayHex ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsDisplayHex)
	// [DataType]   // ColorEdit, ColorPicker, ColorButton: _display_ values formatted as 0..255.
	ColorEditFlagsUint8 ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsUint8)
	// [DataType]   // ColorEdit, ColorPicker, ColorButton: _display_ values formatted as 0.0f..1.0f floats instead of 0..255 integers. No round-trip of value via integers.
	ColorEditFlagsFloat ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsFloat)
	// [Picker]     // ColorPicker: bar for Hue, rectangle for Sat/Value.
	ColorEditFlagsPickerHueBar ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsPickerHueBar)
	// [Picker]     // ColorPicker: wheel for Hue, triangle for Sat/Value.
	ColorEditFlagsPickerHueWheel ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsPickerHueWheel)
	// [Input]      // ColorEdit, ColorPicker: input and output data in RGB format.
	ColorEditFlagsInputRGB ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsInputRGB)
	// [Input]      // ColorEdit, ColorPicker: input and output data in HSV format.
	ColorEditFlagsInputHSV       ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsInputHSV)
	ColorEditFlagsDefaultOptions ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsDefaultOptions)
	ColorEditFlagsDisplayMask    ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsDisplayMask)
	ColorEditFlagsDataTypeMask   ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsDataTypeMask)
	ColorEditFlagsPickerMask     ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsPickerMask)
	ColorEditFlagsInputMask      ColorEditFlags = ColorEditFlags(imgui.ColorEditFlagsInputMask)
)

// TableFlags represents table flags.
type TableFlags imgui.TableFlags

// Table flags enum:.
const (
	TableFlagsNone                       TableFlags = TableFlags(imgui.TableFlagsNone)
	TableFlagsResizable                  TableFlags = TableFlags(imgui.TableFlagsResizable)
	TableFlagsReorderable                TableFlags = TableFlags(imgui.TableFlagsReorderable)
	TableFlagsHideable                   TableFlags = TableFlags(imgui.TableFlagsHideable)
	TableFlagsSortable                   TableFlags = TableFlags(imgui.TableFlagsSortable)
	TableFlagsNoSavedSettings            TableFlags = TableFlags(imgui.TableFlagsNoSavedSettings)
	TableFlagsContextMenuInBody          TableFlags = TableFlags(imgui.TableFlagsContextMenuInBody)
	TableFlagsRowBg                      TableFlags = TableFlags(imgui.TableFlagsRowBg)
	TableFlagsBordersInnerH              TableFlags = TableFlags(imgui.TableFlagsBordersInnerH)
	TableFlagsBordersOuterH              TableFlags = TableFlags(imgui.TableFlagsBordersOuterH)
	TableFlagsBordersInnerV              TableFlags = TableFlags(imgui.TableFlagsBordersInnerV)
	TableFlagsBordersOuterV              TableFlags = TableFlags(imgui.TableFlagsBordersOuterV)
	TableFlagsBordersH                   TableFlags = TableFlags(imgui.TableFlagsBordersH)
	TableFlagsBordersV                   TableFlags = TableFlags(imgui.TableFlagsBordersV)
	TableFlagsBordersInner               TableFlags = TableFlags(imgui.TableFlagsBordersInner)
	TableFlagsBordersOuter               TableFlags = TableFlags(imgui.TableFlagsBordersOuter)
	TableFlagsBorders                    TableFlags = TableFlags(imgui.TableFlagsBorders)
	TableFlagsNoBordersInBody            TableFlags = TableFlags(imgui.TableFlagsNoBordersInBody)
	TableFlagsNoBordersInBodyUntilResize TableFlags = TableFlags(imgui.TableFlagsNoBordersInBodyUntilResize)
	TableFlagsSizingFixedFit             TableFlags = TableFlags(imgui.TableFlagsSizingFixedFit)
	TableFlagsSizingFixedSame            TableFlags = TableFlags(imgui.TableFlagsSizingFixedSame)
	TableFlagsSizingStretchProp          TableFlags = TableFlags(imgui.TableFlagsSizingStretchProp)
	TableFlagsSizingStretchSame          TableFlags = TableFlags(imgui.TableFlagsSizingStretchSame)
	TableFlagsNoHostExtendX              TableFlags = TableFlags(imgui.TableFlagsNoHostExtendX)
	TableFlagsNoHostExtendY              TableFlags = TableFlags(imgui.TableFlagsNoHostExtendY)
	TableFlagsNoKeepColumnsVisible       TableFlags = TableFlags(imgui.TableFlagsNoKeepColumnsVisible)
	TableFlagsPreciseWidths              TableFlags = TableFlags(imgui.TableFlagsPreciseWidths)
	TableFlagsNoClip                     TableFlags = TableFlags(imgui.TableFlagsNoClip)
	TableFlagsPadOuterX                  TableFlags = TableFlags(imgui.TableFlagsPadOuterX)
	TableFlagsNoPadOuterX                TableFlags = TableFlags(imgui.TableFlagsNoPadOuterX)
	TableFlagsNoPadInnerX                TableFlags = TableFlags(imgui.TableFlagsNoPadInnerX)
	TableFlagsScrollX                    TableFlags = TableFlags(imgui.TableFlagsScrollX)
	TableFlagsScrollY                    TableFlags = TableFlags(imgui.TableFlagsScrollY)
	TableFlagsSortMulti                  TableFlags = TableFlags(imgui.TableFlagsSortMulti)
	TableFlagsSortTristate               TableFlags = TableFlags(imgui.TableFlagsSortTristate)
	TableFlagsSizingMask                 TableFlags = TableFlags(imgui.TableFlagsSizingMask)
)

// TableRowFlags represents table row flags.
type TableRowFlags imgui.TableRowFlags

// table row flags:.
const (
	TableRowFlagsNone TableRowFlags = TableRowFlags(imgui.TableRowFlagsNone)
	// Identify header row (set default background color + width of its contents accounted different for auto column width).
	TableRowFlagsHeaders TableRowFlags = TableRowFlags(imgui.TableRowFlagsHeaders)
)

// TableColumnFlags represents a flags for table column (see (*TableColumnWidget).Flags()).
type TableColumnFlags imgui.TableColumnFlags

// table column flags list.
const (
	// Input configuration flags.
	TableColumnFlagsNone                 TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsNone)
	TableColumnFlagsDefaultHide          TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsDefaultHide)
	TableColumnFlagsDefaultSort          TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsDefaultSort)
	TableColumnFlagsWidthStretch         TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsWidthStretch)
	TableColumnFlagsWidthFixed           TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsWidthFixed)
	TableColumnFlagsNoResize             TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsNoResize)
	TableColumnFlagsNoReorder            TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsNoReorder)
	TableColumnFlagsNoHide               TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsNoHide)
	TableColumnFlagsNoClip               TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsNoClip)
	TableColumnFlagsNoSort               TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsNoSort)
	TableColumnFlagsNoSortAscending      TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsNoSortAscending)
	TableColumnFlagsNoSortDescending     TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsNoSortDescending)
	TableColumnFlagsNoHeaderWidth        TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsNoHeaderWidth)
	TableColumnFlagsPreferSortAscending  TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsPreferSortAscending)
	TableColumnFlagsPreferSortDescending TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsPreferSortDescending)
	TableColumnFlagsIndentEnable         TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsIndentEnable)
	TableColumnFlagsIndentDisable        TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsIndentDisable)

	// Output status flags read-only via TableGetColumnFlags().
	TableColumnFlagsIsEnabled TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsIsEnabled)
	TableColumnFlagsIsVisible TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsIsVisible)
	TableColumnFlagsIsSorted  TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsIsSorted)
	TableColumnFlagsIsHovered TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsIsHovered)

	// [Internal] Combinations and masks.
	TableColumnFlagsWidthMask      TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsWidthMask)
	TableColumnFlagsIndentMask     TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsIndentMask)
	TableColumnFlagsStatusMask     TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsStatusMask)
	TableColumnFlagsNoDirectResize TableColumnFlags = TableColumnFlags(imgui.TableColumnFlagsNoDirectResize)
)

// SliderFlags represents imgui.SliderFlags.
type SliderFlags imgui.SliderFlags

// slider flags.
const (
	SliderFlagsNone = SliderFlags(imgui.SliderFlagsNone)
	// Clamp value to minValue/maxValue bounds when input manually with CTRL+Click. By default CTRL+Click allows going out of bounds.
	SliderFlagsAlwaysClamp = SliderFlags(imgui.SliderFlagsAlwaysClamp)
	// Make the widget logarithmic (linear otherwise). Consider using ImGuiSliderFlagsNoRoundToFormat = SliderFlags(imgui.SliderFlagsNoRoundToFormat)
	// a format-string with small amount of digits.
	SliderFlagsLogarithmic = SliderFlags(imgui.SliderFlagsLogarithmic)
	// Disable rounding underlying value to match precision of the display format string (e.g. %.3f values are rounded to those 3 digits).
	SliderFlagsNoRoundToFormat = SliderFlags(imgui.SliderFlagsNoRoundToFormat)
	// Disable CTRL+Click or Enter key allowing to input text directly into the widget.
	SliderFlagsNoInput = SliderFlags(imgui.SliderFlagsNoInput)
	// [Internal] We treat using those bits as being potentially a 'float power' argument from the previous API that has got miscast
	// to this enum, and will trigger an assert if needed.
	SliderFlagsInvalidMask = SliderFlags(imgui.SliderFlagsInvalidMask)
)

// PlotFlags represents imgui.PlotFlags.
type PlotFlags imgui.PlotFlags

// plot flags.
const (
	PlotFlagsNone        = PlotFlags(imgui.PlotFlagsNone)
	PlotFlagsNoTitle     = PlotFlags(imgui.PlotFlagsNoTitle)
	PlotFlagsNoLegend    = PlotFlags(imgui.PlotFlagsNoLegend)
	PlotFlagsNoMenus     = PlotFlags(imgui.PlotFlagsNoMenus)
	PlotFlagsNoBoxSelect = PlotFlags(imgui.PlotFlagsNoBoxSelect)
	// 	PlotFlagsNoMousePos  = PlotFlags(imgui.PlotFlagsNoMousePos)
	// 	PlotFlagsNoHighlight = PlotFlags(imgui.PlotFlagsNoHighlight)
	// PlotFlagsNoChild = PlotFlags(imgui.PlotFlagsNoChild).
	PlotFlagsEqual = PlotFlags(imgui.PlotFlagsEqual)
	// 	PlotFlagsYAxis2      = PlotFlags(imgui.PlotFlagsYAxis2)
	// 	PlotFlagsYAxis3      = PlotFlags(imgui.PlotFlagsYAxis3)
	// 	PlotFlagsQuery       = PlotFlags(imgui.PlotFlagsQuery)
	PlotFlagsCrosshairs = PlotFlags(imgui.PlotFlagsCrosshairs)
	// 	PlotFlagsAntiAliased = PlotFlags(imgui.PlotFlagsAntiAliased)
	PlotFlagsCanvasOnly = PlotFlags(imgui.PlotFlagsCanvasOnly)
)

// PlotAxisFlags represents imgui.PlotAxisFlags.
type PlotAxisFlags imgui.PlotAxisFlags

// plot axis flags.
const (
	PlotAxisFlagsNone         PlotAxisFlags = PlotAxisFlags(imgui.PlotAxisFlagsNone)
	PlotAxisFlagsNoLabel      PlotAxisFlags = PlotAxisFlags(imgui.PlotAxisFlagsNoLabel)
	PlotAxisFlagsNoGridLines  PlotAxisFlags = PlotAxisFlags(imgui.PlotAxisFlagsNoGridLines)
	PlotAxisFlagsNoTickMarks  PlotAxisFlags = PlotAxisFlags(imgui.PlotAxisFlagsNoTickMarks)
	PlotAxisFlagsNoTickLabels PlotAxisFlags = PlotAxisFlags(imgui.PlotAxisFlagsNoTickLabels)
	PlotAxisFlagsForeground   PlotAxisFlags = PlotAxisFlags(imgui.PlotAxisFlagsForeground)
	//	PlotAxisFlagsLogScale      PlotAxisFlags = PlotAxisFlags(imgui.PlotAxisFlagsLogScale)
	//	PlotAxisFlagsTime          PlotAxisFlags = PlotAxisFlags(imgui.PlotAxisFlagsTime)
	PlotAxisFlagsInvert        PlotAxisFlags = PlotAxisFlags(imgui.PlotAxisFlagsInvert)
	PlotAxisFlagsNoInitialFit  PlotAxisFlags = PlotAxisFlags(imgui.PlotAxisFlagsNoInitialFit)
	PlotAxisFlagsAutoFit       PlotAxisFlags = PlotAxisFlags(imgui.PlotAxisFlagsAutoFit)
	PlotAxisFlagsRangeFit      PlotAxisFlags = PlotAxisFlags(imgui.PlotAxisFlagsRangeFit)
	PlotAxisFlagsLockMin       PlotAxisFlags = PlotAxisFlags(imgui.PlotAxisFlagsLockMin)
	PlotAxisFlagsLockMax       PlotAxisFlags = PlotAxisFlags(imgui.PlotAxisFlagsLockMax)
	PlotAxisFlagsLock          PlotAxisFlags = PlotAxisFlags(imgui.PlotAxisFlagsLock)
	PlotAxisFlagsNoDecorations PlotAxisFlags = PlotAxisFlags(imgui.PlotAxisFlagsNoDecorations)
)
