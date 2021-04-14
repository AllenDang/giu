package giu

type InputTextFlags int

const (
	InputTextFlags_None                InputTextFlags = 0
	InputTextFlags_CharsDecimal        InputTextFlags = 1 << 0  // Allow 0123456789.+-*/
	InputTextFlags_CharsHexadecimal    InputTextFlags = 1 << 1  // Allow 0123456789ABCDEFabcdef
	InputTextFlags_CharsUppercase      InputTextFlags = 1 << 2  // Turn a..z into A..Z
	InputTextFlags_CharsNoBlank        InputTextFlags = 1 << 3  // Filter out spaces, tabs
	InputTextFlags_AutoSelectAll       InputTextFlags = 1 << 4  // Select entire text when first taking mouse focus
	InputTextFlags_EnterReturnsTrue    InputTextFlags = 1 << 5  // Return 'true' when Enter is pressed (as opposed to every time the value was modified). Consider looking at the IsItemDeactivatedAfterEdit() function.
	InputTextFlags_CallbackCompletion  InputTextFlags = 1 << 6  // Callback on pressing TAB (for completion handling)
	InputTextFlags_CallbackHistory     InputTextFlags = 1 << 7  // Callback on pressing Up/Down arrows (for history handling)
	InputTextFlags_CallbackAlways      InputTextFlags = 1 << 8  // Callback on each iteration. User code may query cursor position, modify text buffer.
	InputTextFlags_CallbackCharFilter  InputTextFlags = 1 << 9  // Callback on character inputs to replace or discard them. Modify 'EventChar' to replace or discard, or return 1 in callback to discard.
	InputTextFlags_AllowTabInput       InputTextFlags = 1 << 10 // Pressing TAB input a '\t' character into the text field
	InputTextFlags_CtrlEnterForNewLine InputTextFlags = 1 << 11 // In multi-line mode, unfocus with Enter, add new line with Ctrl+Enter (default is opposite: unfocus with Ctrl+Enter, add line with Enter).
	InputTextFlags_NoHorizontalScroll  InputTextFlags = 1 << 12 // Disable following the cursor horizontally
	InputTextFlags_AlwaysOverwrite     InputTextFlags = 1 << 13 // Overwrite mode
	InputTextFlags_ReadOnly            InputTextFlags = 1 << 14 // Read-only mode
	InputTextFlags_Password            InputTextFlags = 1 << 15 // Password mode, display all characters as '*'
	InputTextFlags_NoUndoRedo          InputTextFlags = 1 << 16 // Disable undo/redo. Note that input text owns the text data while active, if you want to provide your own undo/redo stack you need e.g. to call ClearActiveID().
	InputTextFlags_CharsScientific     InputTextFlags = 1 << 17 // Allow 0123456789.+-*/eE (Scientific notation input)
	InputTextFlags_CallbackResize      InputTextFlags = 1 << 18 // Callback on buffer capacity changes request (beyond 'buf_size' parameter value), allowing the string to grow. Notify when the string wants to be resized (for string types which hold a cache of their Size). You will be provided a new BufSize in the callback and NEED to honor it. (see misc/cpp/imgui_stdlib.h for an example of using this)
	InputTextFlags_CallbackEdit        InputTextFlags = 1 << 19 // Callback on any edit (note that InputText() already returns true on edit, the callback is useful mainly to manipulate the underlying buffer while focus is active)
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

type HoveredFlags int

const (
	// HoveredFlagsNone is the default and matches if directly over the item/window, not obstructed by another window, not obstructed by an active popup or modal blocking inputs under them.
	HoveredFlagsNone HoveredFlags = 0
	// HoveredFlagsChildWindows is for IsWindowHovered() and matches if any children of the window is hovered
	HoveredFlagsChildWindows HoveredFlags = 1 << 0
	// HoveredFlagsRootWindow is for IsWindowHovered() and tests from root window (top most parent of the current hierarchy)
	HoveredFlagsRootWindow HoveredFlags = 1 << 1
	// HoveredFlagsAnyWindow is for IsWindowHovered() and matches if any window is hovered
	HoveredFlagsAnyWindow HoveredFlags = 1 << 2
	// HoveredFlagsAllowWhenBlockedByPopup matches even if a popup window is normally blocking access to this item/window
	HoveredFlagsAllowWhenBlockedByPopup HoveredFlags = 1 << 3
	// HoveredFlagsAllowWhenBlockedByModal matches even if a modal popup window is normally blocking access to this item/window. UNIMPLEMENTED in imgui.
	//HoveredFlagsAllowWhenBlockedByModal  HoveredFlags   = 1 << 4
	// HoveredFlagsAllowWhenBlockedByActiveItem matches true even if an active item is blocking access to this item/window. Useful for Drag and Drop patterns.
	HoveredFlagsAllowWhenBlockedByActiveItem HoveredFlags = 1 << 5
	// HoveredFlagsAllowWhenOverlapped matches even if the position is obstructed or overlapped by another window
	HoveredFlagsAllowWhenOverlapped HoveredFlags = 1 << 6
	// HoveredFlagsAllowWhenDisabled matches even if the item is disabled
	HoveredFlagsAllowWhenDisabled HoveredFlags = 1 << 7
	// HoveredFlagsRectOnly combines HoveredFlagsAllowWhenBlockedByPopup, HoveredFlagsAllowWhenBlockedByActiveItem, and HoveredFlagsAllowWhenOverlapped.
	HoveredFlagsRectOnly HoveredFlags = HoveredFlagsAllowWhenBlockedByPopup | HoveredFlagsAllowWhenBlockedByActiveItem | HoveredFlagsAllowWhenOverlapped
	// HoveredFlagsRootAndChildWindows combines HoveredFlagsRootWindow and HoveredFlagsChildWindows.
	HoveredFlagsRootAndChildWindows HoveredFlags = HoveredFlagsRootWindow | HoveredFlagsChildWindows
)
