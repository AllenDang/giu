package giu

type InputTextFlags int

const (
	InputTextFlagsNone InputTextFlags = 0
	// Allow 0123456789.+-*/
	InputTextFlagsCharsDecimal InputTextFlags = 1 << 0
	// Allow 0123456789ABCDEFabcdef
	InputTextFlagsCharsHexadecimal InputTextFlags = 1 << 1
	// Turn a..z into A..Z
	InputTextFlagsCharsUppercase InputTextFlags = 1 << 2
	// Filter out spaces, tabs
	InputTextFlagsCharsNoBlank InputTextFlags = 1 << 3
	// Select entire text when first taking mouse focus
	InputTextFlagsAutoSelectAll InputTextFlags = 1 << 4
	// Return 'true' when Enter is pressed (as opposed to every time the value was modified).
	// Consider looking at the IsItemDeactivatedAfterEdit() function.
	InputTextFlagsEnterReturnsTrue InputTextFlags = 1 << 5
	// Callback on pressing TAB (for completion handling)
	InputTextFlagsCallbackCompletion InputTextFlags = 1 << 6
	// Callback on pressing Up/Down arrows (for history handling)
	InputTextFlagsCallbackHistory InputTextFlags = 1 << 7
	// Callback on each iteration. User code may query cursor position, modify text buffer.
	InputTextFlagsCallbackAlways InputTextFlags = 1 << 8
	// Callback on character inputs to replace or discard them. Modify 'EventChar' to replace or discard, or return 1 in callback to discard.
	InputTextFlagsCallbackCharFilter InputTextFlags = 1 << 9
	// Pressing TAB input a '\t' character into the text field
	InputTextFlagsAllowTabInput InputTextFlags = 1 << 10
	// In multi-line mode, unfocus with Enter, add new line with Ctrl+Enter
	// (default is opposite: unfocus with Ctrl+Enter, add line with Enter).
	InputTextFlagsCtrlEnterForNewLine InputTextFlags = 1 << 11
	// Disable following the cursor horizontally
	InputTextFlagsNoHorizontalScroll InputTextFlags = 1 << 12
	// Overwrite mode
	InputTextFlagsAlwaysOverwrite InputTextFlags = 1 << 13
	// Read-only mode
	InputTextFlagsReadOnly InputTextFlags = 1 << 14
	// Password mode, display all characters as '*'
	InputTextFlagsPassword InputTextFlags = 1 << 15
	// Disable undo/redo. Note that input text owns the text data while active, if you want to provide your own undo/redo
	// stack you need e.g. to call ClearActiveID().
	InputTextFlagsNoUndoRedo InputTextFlags = 1 << 16
	// Allow 0123456789.+-*/eE (Scientific notation input)
	InputTextFlagsCharsScientific InputTextFlags = 1 << 17
	// Callback on buffer capacity changes request (beyond 'bufsize' parameter value), allowing the string to grow.
	// Notify when the string wants to be resized (for string types which hold a cache of their Size).
	// You will be provided a new BufSize in the callback and NEED to honor it. (see misc/cpp/imguistdlib.h for an example of using this)
	InputTextFlagsCallbackResize InputTextFlags = 1 << 18
	// Callback on any edit (note that InputText() already returns true on edit,
	// the callback is useful mainly to manipulate the underlying buffer while focus is active)
	InputTextFlagsCallbackEdit InputTextFlags = 1 << 19
)

type WindowFlags int

const (
	// WindowFlagsNone default WindowFlags = 0
	WindowFlagsNone WindowFlags = 0
	// WindowFlagsNoTitleBar disables title-bar.
	WindowFlagsNoTitleBar WindowFlags = 1 << 0
	// WindowFlagsNoResize disables user resizing with the lower-right grip.
	WindowFlagsNoResize WindowFlags = 1 << 1
	// WindowFlagsNoMove disables user moving the window.
	WindowFlagsNoMove WindowFlags = 1 << 2
	// WindowFlagsNoScrollbar disables scrollbars. Window can still scroll with mouse or programmatically.
	WindowFlagsNoScrollbar WindowFlags = 1 << 3
	// WindowFlagsNoScrollWithMouse disables user vertically scrolling with mouse wheel. On child window, mouse wheel
	// will be forwarded to the parent unless NoScrollbar is also set.
	WindowFlagsNoScrollWithMouse WindowFlags = 1 << 4
	// WindowFlagsNoCollapse disables user collapsing window by double-clicking on it.
	WindowFlagsNoCollapse WindowFlags = 1 << 5
	// WindowFlagsAlwaysAutoResize resizes every window to its content every frame.
	WindowFlagsAlwaysAutoResize WindowFlags = 1 << 6
	// WindowFlagsNoBackground disables drawing background color (WindowBg, etc.) and outside border. Similar as using
	// SetNextWindowBgAlpha(0.0f).
	WindowFlagsNoBackground WindowFlags = 1 << 7
	// WindowFlagsNoSavedSettings will never load/save settings in .ini file.
	WindowFlagsNoSavedSettings WindowFlags = 1 << 8
	// WindowFlagsNoMouseInputs disables catching mouse, hovering test with pass through.
	WindowFlagsNoMouseInputs WindowFlags = 1 << 9
	// WindowFlagsMenuBar has a menu-bar.
	WindowFlagsMenuBar WindowFlags = 1 << 10
	// WindowFlagsHorizontalScrollbar allows horizontal scrollbar to appear (off by default). You may use
	// SetNextWindowContentSize(ImVec2(width,0.0f)); prior to calling Begin() to specify width. Read code in imgui_demo
	// in the "Horizontal Scrolling" section.
	WindowFlagsHorizontalScrollbar WindowFlags = 1 << 11
	// WindowFlagsNoFocusOnAppearing disables taking focus when transitioning from hidden to visible state.
	WindowFlagsNoFocusOnAppearing WindowFlags = 1 << 12
	// WindowFlagsNoBringToFrontOnFocus disables bringing window to front when taking focus. e.g. clicking on it or
	// programmatically giving it focus.
	WindowFlagsNoBringToFrontOnFocus WindowFlags = 1 << 13
	// WindowFlagsAlwaysVerticalScrollbar always shows vertical scrollbar, even if ContentSize.y < Size.y .
	WindowFlagsAlwaysVerticalScrollbar WindowFlags = 1 << 14
	// WindowFlagsAlwaysHorizontalScrollbar always shows horizontal scrollbar, even if ContentSize.x < Size.x .
	WindowFlagsAlwaysHorizontalScrollbar WindowFlags = 1 << 15
	// WindowFlagsAlwaysUseWindowPadding ensures child windows without border uses style.WindowPadding (ignored by
	// default for non-bordered child windows, because more convenient).
	WindowFlagsAlwaysUseWindowPadding WindowFlags = 1 << 16
	// WindowFlagsNoNavInputs has no gamepad/keyboard navigation within the window.
	WindowFlagsNoNavInputs WindowFlags = 1 << 18
	// WindowFlagsNoNavFocus has no focusing toward this window with gamepad/keyboard navigation
	// (e.g. skipped by CTRL+TAB)
	WindowFlagsNoNavFocus WindowFlags = 1 << 19
	// WindowFlagsUnsavedDocument appends '*' to title without affecting the ID, as a convenience to avoid using the
	// ### operator. When used in a tab/docking context, tab is selected on closure and closure is deferred by one
	// frame to allow code to cancel the closure (with a confirmation popup, etc.) without flicker.
	WindowFlagsUnsavedDocument WindowFlags = 1 << 20

	// WindowFlagsNoNav combines WindowFlagsNoNavInputs and WindowFlagsNoNavFocus.
	WindowFlagsNoNav WindowFlags = WindowFlagsNoNavInputs | WindowFlagsNoNavFocus
	// WindowFlagsNoDecoration combines WindowFlagsNoTitleBar, WindowFlagsNoResize, WindowFlagsNoScrollbar and
	// WindowFlagsNoCollapse.
	WindowFlagsNoDecoration WindowFlags = WindowFlagsNoTitleBar | WindowFlagsNoResize | WindowFlagsNoScrollbar | WindowFlagsNoCollapse
	// WindowFlagsNoInputs combines WindowFlagsNoMouseInputs, WindowFlagsNoNavInputs and WindowFlagsNoNavFocus.
	WindowFlagsNoInputs WindowFlags = WindowFlagsNoMouseInputs | WindowFlagsNoNavInputs | WindowFlagsNoNavFocus
)

type ComboFlags int

const (
	// ComboFlagNone default ComboFlags = 0
	ComboFlagNone ComboFlags = 0
	// ComboFlagPopupAlignLeft aligns the popup toward the left by default.
	ComboFlagPopupAlignLeft ComboFlags = 1 << 0
	// ComboFlagHeightSmall has max ~4 items visible.
	// Tip: If you want your combo popup to be a specific size you can use SetNextWindowSizeConstraints() prior to calling BeginCombo().
	ComboFlagHeightSmall ComboFlags = 1 << 1
	// ComboFlagHeightRegular has max ~8 items visible (default).
	ComboFlagHeightRegular ComboFlags = 1 << 2
	// ComboFlagHeightLarge has max ~20 items visible.
	ComboFlagHeightLarge ComboFlags = 1 << 3
	// ComboFlagHeightLargest has as many fitting items as possible.
	ComboFlagHeightLargest ComboFlags = 1 << 4
	// ComboFlagNoArrowButton displays on the preview box without the square arrow button.
	ComboFlagNoArrowButton ComboFlags = 1 << 5
	// ComboFlagNoPreview displays only a square arrow button.
	ComboFlagNoPreview ComboFlags = 1 << 6
)

// SelectableFlags represents imgui.SelectableFlags
type SelectableFlags int

const (
	// SelectableFlagsNone default SelectableFlags = 0
	SelectableFlagsNone SelectableFlags = 0
	// SelectableFlagsDontClosePopups makes clicking the selectable not close any parent popup windows.
	SelectableFlagsDontClosePopups SelectableFlags = 1 << 0
	// SelectableFlagsSpanAllColumns allows the selectable frame to span all columns (text will still fit in current column).
	SelectableFlagsSpanAllColumns SelectableFlags = 1 << 1
	// SelectableFlagsAllowDoubleClick generates press events on double clicks too.
	SelectableFlagsAllowDoubleClick SelectableFlags = 1 << 2
	// SelectableFlagsDisabled disallows selection and displays text in a greyed out color.
	SelectableFlagsDisabled SelectableFlags = 1 << 3
)

// TabItemFlags represents tab item flags
type TabItemFlags int

const (
	// TabItemFlagsNone default TabItemFlags = 0
	TabItemFlagsNone TabItemFlags = 0
	// TabItemFlagsUnsavedDocument Append '*' to title without affecting the ID, as a convenience to avoid using the
	// ### operator. Also: tab is selected on closure and closure is deferred by one frame to allow code to undo it
	// without flicker.
	TabItemFlagsUnsavedDocument TabItemFlags = 1 << 0
	// TabItemFlagsSetSelected Trigger flag to programmatically make the tab selected when calling BeginTabItem()
	TabItemFlagsSetSelected TabItemFlags = 1 << 1
	// TabItemFlagsNoCloseWithMiddleMouseButton  Disable behavior of closing tabs (that are submitted with
	// p_open != NULL) with middle mouse button. You can still repro this behavior on user's side with if
	// (IsItemHovered() && IsMouseClicked(2)) *p_open TabItemFlags = false.
	TabItemFlagsNoCloseWithMiddleMouseButton TabItemFlags = 1 << 2
	// TabItemFlagsNoPushID Don't call PushID(tab->ID)/PopID() on BeginTabItem()/EndTabItem()
	TabItemFlagsNoPushID TabItemFlags = 1 << 3
)

type TabBarFlags int

const (
	// TabBarFlagsNone default TabBarFlags = 0.
	TabBarFlagsNone TabBarFlags = 0
	// TabBarFlagsReorderable Allow manually dragging tabs to re-order them + New tabs are appended at the end of list
	TabBarFlagsReorderable TabBarFlags = 1 << 0
	// TabBarFlagsAutoSelectNewTabs Automatically select new tabs when they appear
	TabBarFlagsAutoSelectNewTabs TabBarFlags = 1 << 1
	// TabBarFlagsTabListPopupButton Disable buttons to open the tab list popup
	TabBarFlagsTabListPopupButton TabBarFlags = 1 << 2
	// TabBarFlagsNoCloseWithMiddleMouseButton Disable behavior of closing tabs (that are submitted with p_open != NULL)
	// with middle mouse button. You can still repro this behavior on user's side with if
	// (IsItemHovered() && IsMouseClicked(2)) *p_open TabBarFlags = false.
	TabBarFlagsNoCloseWithMiddleMouseButton TabBarFlags = 1 << 3
	// TabBarFlagsNoTabListScrollingButtons Disable scrolling buttons (apply when fitting policy is
	// TabBarFlagsFittingPolicyScroll)
	TabBarFlagsNoTabListScrollingButtons TabBarFlags = 1 << 4
	// TabBarFlagsNoTooltip Disable tooltips when hovering a tab
	TabBarFlagsNoTooltip TabBarFlags = 1 << 5
	// TabBarFlagsFittingPolicyResizeDown Resize tabs when they don't fit
	TabBarFlagsFittingPolicyResizeDown TabBarFlags = 1 << 6
	// TabBarFlagsFittingPolicyScroll Add scroll buttons when tabs don't fit
	TabBarFlagsFittingPolicyScroll TabBarFlags = 1 << 7
	// TabBarFlagsFittingPolicyMask combines
	// TabBarFlagsFittingPolicyResizeDown and TabBarFlagsFittingPolicyScroll
	TabBarFlagsFittingPolicyMask TabBarFlags = TabBarFlagsFittingPolicyResizeDown | TabBarFlagsFittingPolicyScroll
	// TabBarFlagsFittingPolicyDefault alias for TabBarFlagsFittingPolicyResizeDown
	TabBarFlagsFittingPolicyDefault TabBarFlags = TabBarFlagsFittingPolicyResizeDown
)

// TreeNodeFlags represents tree node widget flags
type TreeNodeFlags int

const (
	// TreeNodeFlagsNone default TreeNodeFlags = 0
	TreeNodeFlagsNone TreeNodeFlags = 0
	// TreeNodeFlagsSelected draws as selected.
	TreeNodeFlagsSelected TreeNodeFlags = 1 << 0
	// TreeNodeFlagsFramed draws full colored frame (e.g. for CollapsingHeader).
	TreeNodeFlagsFramed TreeNodeFlags = 1 << 1
	// TreeNodeFlagsAllowItemOverlap hit testing to allow subsequent widgets to overlap this one.
	TreeNodeFlagsAllowItemOverlap TreeNodeFlags = 1 << 2
	// TreeNodeFlagsNoTreePushOnOpen doesn't do a TreePush() when open
	// (e.g. for CollapsingHeader) TreeNodeFlags = no extra indent nor pushing on ID stack.
	TreeNodeFlagsNoTreePushOnOpen TreeNodeFlags = 1 << 3
	// TreeNodeFlagsNoAutoOpenOnLog doesn't automatically and temporarily open node when Logging is active
	// (by default logging will automatically open tree nodes).
	TreeNodeFlagsNoAutoOpenOnLog TreeNodeFlags = 1 << 4
	// TreeNodeFlagsDefaultOpen defaults node to be open.
	TreeNodeFlagsDefaultOpen TreeNodeFlags = 1 << 5
	// TreeNodeFlagsOpenOnDoubleClick needs double-click to open node.
	TreeNodeFlagsOpenOnDoubleClick TreeNodeFlags = 1 << 6
	// TreeNodeFlagsOpenOnArrow opens only when clicking on the arrow part.
	// If TreeNodeFlagsOpenOnDoubleClick is also set, single-click arrow or double-click all box to open.
	TreeNodeFlagsOpenOnArrow TreeNodeFlags = 1 << 7
	// TreeNodeFlagsLeaf allows no collapsing, no arrow (use as a convenience for leaf nodes).
	TreeNodeFlagsLeaf TreeNodeFlags = 1 << 8
	// TreeNodeFlagsBullet displays a bullet instead of an arrow.
	TreeNodeFlagsBullet TreeNodeFlags = 1 << 9
	// TreeNodeFlagsFramePadding uses FramePadding (even for an unframed text node) to
	// vertically align text baseline to regular widget height. Equivalent to calling AlignTextToFramePadding().
	TreeNodeFlagsFramePadding TreeNodeFlags = 1 << 10
	// TreeNodeFlagsSpanAvailWidth extends hit box to the right-most edge, even if not framed.
	// This is not the default in order to allow adding other items on the same line.
	// In the future we may refactor the hit system to be front-to-back, allowing natural overlaps
	// and then this can become the default.
	TreeNodeFlagsSpanAvailWidth TreeNodeFlags = 1 << 11
	// TreeNodeFlagsSpanFullWidth extends hit box to the left-most and right-most edges (bypass the indented area).
	TreeNodeFlagsSpanFullWidth TreeNodeFlags = 1 << 12
	// TreeNodeFlagsNavLeftJumpsBackHere (WIP) Nav: left direction may move to this TreeNode() from any of its child
	// (items submitted between TreeNode and TreePop)
	TreeNodeFlagsNavLeftJumpsBackHere TreeNodeFlags = 1 << 13
	// TreeNodeFlagsCollapsingHeader combines TreeNodeFlagsFramed and TreeNodeFlagsNoAutoOpenOnLog.
	TreeNodeFlagsCollapsingHeader TreeNodeFlags = TreeNodeFlagsFramed | TreeNodeFlagsNoTreePushOnOpen | TreeNodeFlagsNoAutoOpenOnLog
)

// FocusedFlags represents imgui.FocusedFlags
type FocusedFlags int

const (
	// FocusedFlagsNone default FocusedFlags = 0
	FocusedFlagsNone FocusedFlags = 0
	// FocusedFlagsChildWindows matches if any children of the window is focused
	FocusedFlagsChildWindows FocusedFlags = 1 << 0
	// FocusedFlagsRootWindow tests from root window (top most parent of the current hierarchy)
	FocusedFlagsRootWindow FocusedFlags = 1 << 1
	// FocusedFlagsAnyWindow matches if any window is focused.
	FocusedFlagsAnyWindow FocusedFlags = 1 << 2
	// FocusedFlagsRootAndChildWindows combines FocusedFlagsRootWindow and FocusedFlagsChildWindows.
	FocusedFlagsRootAndChildWindows = FocusedFlagsRootWindow | FocusedFlagsChildWindows
)

// HoveredFlags represents a hovered flags
type HoveredFlags int

const (
	// HoveredFlagsNone is the default and matches if directly over the item/window,
	// not obstructed by another window, not obstructed by an active popup or modal blocking inputs under them.
	HoveredFlagsNone HoveredFlags = 0
	// HoveredFlagsChildWindows is for IsWindowHovered() and matches if any children of the window is hovered
	HoveredFlagsChildWindows HoveredFlags = 1 << 0
	// HoveredFlagsRootWindow is for IsWindowHovered() and tests from root window (top most parent of the current hierarchy)
	HoveredFlagsRootWindow HoveredFlags = 1 << 1
	// HoveredFlagsAnyWindow is for IsWindowHovered() and matches if any window is hovered
	HoveredFlagsAnyWindow HoveredFlags = 1 << 2
	// HoveredFlagsAllowWhenBlockedByPopup matches even if a popup window is normally blocking access to this item/window
	HoveredFlagsAllowWhenBlockedByPopup HoveredFlags = 1 << 3
	// HoveredFlagsAllowWhenBlockedByModal matches even if a modal popup window is normally blocking access to this item/window.
	// UNIMPLEMENTED in imgui.
	// HoveredFlagsAllowWhenBlockedByModal  HoveredFlags   = 1 << 4
	// HoveredFlagsAllowWhenBlockedByActiveItem matches true even if an active item is blocking access to this item/window.
	// Useful for Drag and Drop patterns.
	HoveredFlagsAllowWhenBlockedByActiveItem HoveredFlags = 1 << 5
	// HoveredFlagsAllowWhenOverlapped matches even if the position is obstructed or overlapped by another window
	HoveredFlagsAllowWhenOverlapped HoveredFlags = 1 << 6
	// HoveredFlagsAllowWhenDisabled matches even if the item is disabled
	HoveredFlagsAllowWhenDisabled HoveredFlags = 1 << 7
	// HoveredFlagsRectOnly combines HoveredFlagsAllowWhenBlockedByPopup,
	// HoveredFlagsAllowWhenBlockedByActiveItem, and HoveredFlagsAllowWhenOverlapped.
	HoveredFlagsRectOnly HoveredFlags = HoveredFlagsAllowWhenBlockedByPopup |
		HoveredFlagsAllowWhenBlockedByActiveItem | HoveredFlagsAllowWhenOverlapped
	// HoveredFlagsRootAndChildWindows combines HoveredFlagsRootWindow and HoveredFlagsChildWindows.
	HoveredFlagsRootAndChildWindows HoveredFlags = HoveredFlagsRootWindow | HoveredFlagsChildWindows
)

// ColorEditFlags for ColorEdit3V(), etc.
type ColorEditFlags int

const (
	// ColorEditFlagsNone default = 0.
	ColorEditFlagsNone ColorEditFlags = 0
	// ColorEditFlagsNoAlpha ignores Alpha component (read 3 components from the input pointer).
	ColorEditFlagsNoAlpha ColorEditFlags = 1 << 1
	// ColorEditFlagsNoPicker disables picker when clicking on colored square.
	ColorEditFlagsNoPicker ColorEditFlags = 1 << 2
	// ColorEditFlagsNoOptions disables toggling options menu when right-clicking on inputs/small preview.
	ColorEditFlagsNoOptions ColorEditFlags = 1 << 3
	// ColorEditFlagsNoSmallPreview disables colored square preview next to the inputs. (e.g. to show only the inputs).
	ColorEditFlagsNoSmallPreview ColorEditFlags = 1 << 4
	// ColorEditFlagsNoInputs disables inputs sliders/text widgets (e.g. to show only the small preview colored square).
	ColorEditFlagsNoInputs ColorEditFlags = 1 << 5
	// ColorEditFlagsNoTooltip disables tooltip when hovering the preview.
	ColorEditFlagsNoTooltip ColorEditFlags = 1 << 6
	// ColorEditFlagsNoLabel disables display of inline text label (the label is still forwarded to the tooltip and picker).
	ColorEditFlagsNoLabel ColorEditFlags = 1 << 7
	// ColorEditFlagsNoSidePreview disables bigger color preview on right side of the picker, use small colored square preview instead.
	ColorEditFlagsNoSidePreview ColorEditFlags = 1 << 8
	// ColorEditFlagsNoDragDrop disables drag and drop target. ColorButton: disable drag and drop source.
	ColorEditFlagsNoDragDrop ColorEditFlags = 1 << 9
	// ColorEditFlagsNoBorder disables border (which is enforced by default).
	ColorEditFlagsNoBorder ColorEditFlags = 1 << 10

	// User Options (right-click on widget to change some of them). You can set application defaults using SetColorEditOptions().
	// The idea is that you probably don't want to override them in most of your calls, let the user choose and/or call
	// SetColorEditOptions() during startup.

	// ColorEditFlagsAlphaBar shows vertical alpha bar/gradient in picker.
	ColorEditFlagsAlphaBar ColorEditFlags = 1 << 16
	// ColorEditFlagsAlphaPreview displays preview as a transparent color over a checkerboard, instead of opaque.
	ColorEditFlagsAlphaPreview ColorEditFlags = 1 << 17
	// ColorEditFlagsAlphaPreviewHalf displays half opaque / half checkerboard, instead of opaque.
	ColorEditFlagsAlphaPreviewHalf ColorEditFlags = 1 << 18
	// ColorEditFlagsHDR = (WIP) surrently only disable 0.0f..1.0f limits in RGBA edition.
	// Note: you probably want to use ImGuiColorEditFlags_Float flag as well.
	ColorEditFlagsHDR ColorEditFlags = 1 << 19
	// ColorEditFlagsRGB sets the format as RGB.
	ColorEditFlagsRGB ColorEditFlags = 1 << 20
	// ColorEditFlagsHSV sets the format as HSV.
	ColorEditFlagsHSV ColorEditFlags = 1 << 21
	// ColorEditFlagsHEX sets the format as HEX.
	ColorEditFlagsHEX ColorEditFlags = 1 << 22
	// ColorEditFlagsUint8 _display_ values formatted as 0..255.
	ColorEditFlagsUint8 ColorEditFlags = 1 << 23
	// ColorEditFlagsFloat _display_ values formatted as 0.0f..1.0f floats instead of 0..255 integers. No round-trip of value via integers.
	ColorEditFlagsFloat ColorEditFlags = 1 << 24

	// ColorEditFlagsPickerHueBar shows bar for Hue, rectangle for Sat/Value.
	ColorEditFlagsPickerHueBar ColorEditFlags = 1 << 25
	// ColorEditFlagsPickerHueWheel shows wheel for Hue, triangle for Sat/Value.
	ColorEditFlagsPickerHueWheel ColorEditFlags = 1 << 26
	// ColorEditFlagsInputRGB enables input and output data in RGB format.
	ColorEditFlagsInputRGB ColorEditFlags = 1 << 27
	// ColorEditFlagsInputHSV enables input and output data in HSV format.
	ColorEditFlagsInputHSV ColorEditFlags = 1 << 28
)

// TableFlags represents table flags
type TableFlags int

// Table flags enum:
const (
	// Features

	TableFlagsNone TableFlags = 0
	// Enable resizing columns.
	TableFlagsResizable TableFlags = 1 << 0
	// Enable reordering columns in header row (need calling TableSetupColumn() + TableHeadersRow() to display headers)
	TableFlagsReorderable TableFlags = 1 << 1
	// Enable hiding/disabling columns in context menu.
	TableFlagsHideable TableFlags = 1 << 2
	// Enable sorting. Call TableGetSortSpecs() to obtain sort specs. Also see TableFlagsSortMulti and TableFlagsSortTristate.
	TableFlagsSortable TableFlags = 1 << 3
	// Disable persisting columns order, width and sort settings in the .ini file.
	TableFlagsNoSavedSettings TableFlags = 1 << 4
	// Right-click on columns body/contents will display table context menu. By default it is available in TableHeadersRow().
	TableFlagsContextMenuInBody TableFlags = 1 << 5
	// Decorations

	// Set each RowBg color with ColTableRowBg or ColTableRowBgAlt
	// (equivalent of calling TableSetBgColor with TableBgFlagsRowBg0 on each row manually)
	TableFlagsRowBg TableFlags = 1 << 6
	// Draw horizontal borders between rows.
	TableFlagsBordersInnerH TableFlags = 1 << 7
	// Draw horizontal borders at the top and bottom.
	TableFlagsBordersOuterH TableFlags = 1 << 8
	// Draw vertical borders between columns.
	TableFlagsBordersInnerV TableFlags = 1 << 9
	// Draw vertical borders on the left and right sides.
	TableFlagsBordersOuterV TableFlags = 1 << 10
	// Draw horizontal borders.
	TableFlagsBordersH TableFlags = TableFlagsBordersInnerH | TableFlagsBordersOuterH
	// Draw vertical borders.
	TableFlagsBordersV TableFlags = TableFlagsBordersInnerV | TableFlagsBordersOuterV
	// Draw inner borders.
	TableFlagsBordersInner TableFlags = TableFlagsBordersInnerV | TableFlagsBordersInnerH
	// Draw outer borders.
	TableFlagsBordersOuter TableFlags = TableFlagsBordersOuterV | TableFlagsBordersOuterH
	// Draw all borders.
	TableFlagsBorders TableFlags = TableFlagsBordersInner | TableFlagsBordersOuter
	// [ALPHA] Disable vertical borders in columns Body (borders will always appears in Headers). -> May move to style
	TableFlagsNoBordersInBody TableFlags = 1 << 11
	// [ALPHA] Disable vertical borders in columns Body until hovered for resize (borders will always appears in Headers). -> May move to style
	TableFlagsNoBordersInBodyUntilResizeTableFlags TableFlags = 1 << 12
	// Sizing Policy (read above for defaults)TableFlags

	// Columns default to WidthFixed or WidthAuto (if resizable or not resizable), matching contents width.
	TableFlagsSizingFixedFit TableFlags = 1 << 13
	// Columns default to WidthFixed or WidthAuto (if resizable or not resizable), matching the maximum contents width of all columns.
	// Implicitly enable TableFlagsNoKeepColumnsVisible.
	TableFlagsSizingFixedSame TableFlags = 2 << 13
	// Columns default to WidthStretch with default weights proportional to each columns contents widths.
	TableFlagsSizingStretchProp TableFlags = 3 << 13
	// Columns default to WidthStretch with default weights all equal, unless overridden by TableSetupColumn().
	TableFlagsSizingStretchSame TableFlags = 4 << 13
	// Sizing Extra Options

	// Make outer width auto-fit to columns, overriding outersize.x value.
	// Only available when ScrollX/ScrollY are disabled and Stretch columns are not used.
	TableFlagsNoHostExtendX TableFlags = 1 << 16
	// Make outer height stop exactly at outersize.y (prevent auto-extending table past the limit).
	// Only available when ScrollX/ScrollY are disabled.
	// Data below the limit will be clipped and not visible.
	TableFlagsNoHostExtendY TableFlags = 1 << 17
	// Disable keeping column always minimally visible when ScrollX is off and table gets too small. Not recommended if columns are resizable.
	TableFlagsNoKeepColumnsVisible TableFlags = 1 << 18
	// Disable distributing remainder width to stretched columns
	// (width allocation on a 100-wide table with 3 columns: Without this flag: 33,33,34. With this flag: 33,33,33).
	// With larger number of columns, resizing will appear to be less smooth.
	TableFlagsPreciseWidths TableFlags = 1 << 19

	// Clipping

	// Disable clipping rectangle for every individual columns (reduce draw command count,
	// items will be able to overflow into other columns). Generally incompatible with TableSetupScrollFreeze().
	TableFlagsNoClip TableFlags = 1 << 20
	// Padding

	// Default if BordersOuterV is on. Enable outer-most padding. Generally desirable if you have headers.
	TableFlagsPadOuterX TableFlags = 1 << 21
	// Default if BordersOuterV is off. Disable outer-most padding.
	TableFlagsNoPadOuterX TableFlags = 1 << 22
	// Disable inner padding between columns (double inner padding if BordersOuterV is on, single inner padding if BordersOuterV is off).
	TableFlagsNoPadInnerX TableFlags = 1 << 23

	// Scrolling

	// Enable horizontal scrolling. Require 'outersize' parameter of BeginTable() to specify the container size.
	// Changes default sizing policy. Because this create a child window, ScrollY is currently generally recommended when using ScrollX.
	TableFlagsScrollX TableFlags = 1 << 24
	// Enable vertical scrolling. Require 'outersize' parameter of BeginTable() to specify the container size.
	TableFlagsScrollY TableFlags = 1 << 25
	// Sorting

	// Hold shift when clicking headers to sort on multiple column. TableGetSortSpecs() may return specs where (SpecsCount > 1).
	TableFlagsSortMulti TableFlags = 1 << 26
	// Allow no sorting, disable default sorting. TableGetSortSpecs() may return specs where (SpecsCount == 0).
	TableFlagsSortTristate TableFlags = 1 << 27

	// [Internal] Combinations and masks
	TableFlagsSizingMask TableFlags = TableFlagsSizingFixedFit | TableFlagsSizingFixedSame |
		TableFlagsSizingStretchProp | TableFlagsSizingStretchSame
)

type TableRowFlags int

const (
	TableRowFlagsNone TableRowFlags = 0
	// Identify header row (set default background color + width of its contents accounted different for auto column width)
	TableRowFlagsHeaders TableRowFlags = 1 << 0
)

type TableColumnFlags int

const (
	// Input configuration flags
	TableColumnFlagsNone TableColumnFlags = 0
	// Default as a hidden/disabled column.
	TableColumnFlagsDefaultHide TableColumnFlags = 1 << 0
	// Default as a sorting column.
	TableColumnFlagsDefaultSort TableColumnFlags = 1 << 1
	// Column will stretch. Preferable with horizontal scrolling disabled
	// (default if table sizing policy is SizingStretchSame or SizingStretchProp).
	TableColumnFlagsWidthStretch TableColumnFlags = 1 << 2
	// Column will not stretch. Preferable with horizontal scrolling enabled
	// (default if table sizing policy is SizingFixedFit and table is resizable).
	TableColumnFlagsWidthFixed TableColumnFlags = 1 << 3
	// Disable manual resizing.
	TableColumnFlagsNoResize TableColumnFlags = 1 << 4
	// Disable manual reordering this column, this will also prevent other columns from crossing over this column.
	TableColumnFlagsNoReorder TableColumnFlags = 1 << 5
	// Disable ability to hide/disable this column.
	TableColumnFlagsNoHide TableColumnFlags = 1 << 6
	// Disable clipping for this column (all NoClip columns will render in a same draw command).
	TableColumnFlagsNoClip TableColumnFlags = 1 << 7
	// Disable ability to sort on this field (even if TableFlagsSortable is set on the table).
	TableColumnFlagsNoSort TableColumnFlags = 1 << 8
	// Disable ability to sort in the ascending direction.
	TableColumnFlagsNoSortAscending TableColumnFlags = 1 << 9
	// Disable ability to sort in the descending direction.
	TableColumnFlagsNoSortDescending TableColumnFlags = 1 << 10
	// Disable header text width contribution to automatic column width.
	TableColumnFlagsNoHeaderWidth TableColumnFlags = 1 << 11
	// Make the initial sort direction Ascending when first sorting on this column (default).
	TableColumnFlagsPreferSortAscending TableColumnFlags = 1 << 12
	// Make the initial sort direction Descending when first sorting on this column.
	TableColumnFlagsPreferSortDescending TableColumnFlags = 1 << 13
	// Use current Indent value when entering cell (default for column 0).
	TableColumnFlagsIndentEnable TableColumnFlags = 1 << 14
	// Ignore current Indent value when entering cell (default for columns > 0). Indentation changes within the cell will still be honored.
	TableColumnFlagsIndentDisable TableColumnFlags = 1 << 15

	// Output status flags read-only via TableGetColumnFlags()
	// Status: is enabled == not hidden by user/api (referred to as "Hide" in DefaultHide and NoHide) flags.
	TableColumnFlagsIsEnabled TableColumnFlags = 1 << 20
	// Status: is visible == is enabled AND not clipped by scrolling.
	TableColumnFlagsIsVisible TableColumnFlags = 1 << 21
	// Status: is currently part of the sort specs
	TableColumnFlagsIsSorted TableColumnFlags = 1 << 22
	// Status: is hovered by mouse
	TableColumnFlagsIsHovered TableColumnFlags = 1 << 23

	// [Internal] Combinations and masks
	TableColumnFlagsWidthMask  TableColumnFlags = TableColumnFlagsWidthStretch | TableColumnFlagsWidthFixed
	TableColumnFlagsIndentMask TableColumnFlags = TableColumnFlagsIndentEnable | TableColumnFlagsIndentDisable
	TableColumnFlagsStatusMask TableColumnFlags = TableColumnFlagsIsEnabled |
		TableColumnFlagsIsVisible | TableColumnFlagsIsSorted | TableColumnFlagsIsHovered
	// [Internal] Disable user resizing this column directly (it may however we resized indirectly from its left edge)
	TableColumnFlagsNoDirectResize TableColumnFlags = 1 << 30
)

type SliderFlags int

const (
	SliderFlagsNone SliderFlags = 0
	// Clamp value to min/max bounds when input manually with CTRL+Click. By default CTRL+Click allows going out of bounds.
	SliderFlagsAlwaysClamp SliderFlags = 1 << 4
	// Make the widget logarithmic (linear otherwise). Consider using ImGuiSliderFlagsNoRoundToFormat with this if using
	// a format-string with small amount of digits.
	SliderFlagsLogarithmic SliderFlags = 1 << 5
	// Disable rounding underlying value to match precision of the display format string (e.g. %.3f values are rounded to those 3 digits)
	SliderFlagsNoRoundToFormat SliderFlags = 1 << 6
	// Disable CTRL+Click or Enter key allowing to input text directly into the widget
	SliderFlagsNoInput SliderFlags = 1 << 7
	// [Internal] We treat using those bits as being potentially a 'float power' argument from the previous API that has got miscast
	// to this enum, and will trigger an assert if needed.
	SliderFlagsInvalidMask SliderFlags = 0x7000000F
)

type PlotFlags int

const (
	// default
	PlotFlagsNone PlotFlags = 0
	// the plot title will not be displayed (titles are also hidden if preceded by double hashes, e.g. "##MyPlot")
	PlotFlagsNoTitle PlotFlags = 1 << 0
	// the legend will not be displayed
	PlotFlagsNoLegend PlotFlags = 1 << 1
	// the user will not be able to open context menus with right-click
	PlotFlagsNoMenus PlotFlags = 1 << 2
	// the user will not be able to box-select with right-click drag
	PlotFlagsNoBoxSelect PlotFlags = 1 << 3
	// the mouse position, in plot coordinates, will not be displayed inside of the plot
	PlotFlagsNoMousePos PlotFlags = 1 << 4
	// plot items will not be highlighted when their legend entry is hovered
	PlotFlagsNoHighlight PlotFlags = 1 << 5
	// a child window region will not be used to capture mouse scroll (can boost performance for single Gui window applications)
	PlotFlagsNoChild PlotFlags = 1 << 6
	// primary x and y axes will be constrained to have the same units/pixel (does not apply to auxiliary y-axes)
	PlotFlagsEqual PlotFlags = 1 << 7
	// enable a 2nd y-axis on the right side
	PlotFlagsYAxis2 PlotFlags = 1 << 8
	// enable a 3rd y-axis on the right side
	PlotFlagsYAxis3 PlotFlags = 1 << 9
	// the user will be able to draw query rects with middle-mouse or CTRL + right-click drag
	PlotFlagsQuery PlotFlags = 1 << 10
	// the default mouse cursor will be replaced with a crosshair when hovered
	PlotFlagsCrosshairs PlotFlags = 1 << 11
	// plot lines will be software anti-aliased (not recommended for high density plots, prefer MSAA)
	PlotFlagsAntiAliased PlotFlags = 1 << 12
	PlotFlagsCanvasOnly  PlotFlags = PlotFlagsNoTitle | PlotFlagsNoLegend | PlotFlagsNoMenus | PlotFlagsNoBoxSelect | PlotFlagsNoMousePos
)

type PlotAxisFlags int

const (
	// default
	PlotAxisFlagsNone PlotAxisFlags = 0
	// the axis label will not be displayed (axis labels also hidden if the supplied string name is NULL)
	PlotAxisFlagsNoLabel PlotAxisFlags = 1 << 0
	// the axis grid lines will not be displayed
	PlotAxisFlagsNoGridLines PlotAxisFlags = 1 << 1
	// the axis tick marks will not be displayed
	PlotAxisFlagsNoTickMarks PlotAxisFlags = 1 << 2
	// the axis tick labels will not be displayed
	PlotAxisFlagsNoTickLabels PlotAxisFlags = 1 << 3
	// a logartithmic (base 10) axis scale will be used (mutually exclusive with PlotAxisFlagsTime)
	PlotAxisFlagsLogScale PlotAxisFlags = 1 << 4
	// axis will display date/time formatted labels (mutually exclusive with PlotAxisFlagsLogScale)
	PlotAxisFlagsTime PlotAxisFlags = 1 << 5
	// the axis will be inverted
	PlotAxisFlagsInvert PlotAxisFlags = 1 << 6
	// the axis minimum value will be locked when panning/zooming
	PlotAxisFlagsLockMin PlotAxisFlags = 1 << 7
	// the axis maximum value will be locked when panning/zooming
	PlotAxisFlagsLockMax       PlotAxisFlags = 1 << 8
	PlotAxisFlagsLock          PlotAxisFlags = PlotAxisFlagsLockMin | PlotAxisFlagsLockMax
	PlotAxisFlagsNoDecorations PlotAxisFlags = PlotAxisFlagsNoLabel | PlotAxisFlagsNoGridLines |
		PlotAxisFlagsNoTickMarks | PlotAxisFlagsNoTickLabels
)
