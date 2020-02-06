package giu

type InputTextFlags int

const (
	// InputTextFlagsNone sets everything default.
	InputTextFlagsNone InputTextFlags = 0
	// InputTextFlagsCharsDecimal allows 0123456789.+-
	InputTextFlagsCharsDecimal InputTextFlags = 1 << 0
	// InputTextFlagsCharsHexadecimal allow 0123456789ABCDEFabcdef
	InputTextFlagsCharsHexadecimal InputTextFlags = 1 << 1
	// InputTextFlagsCharsUppercase turns a..z into A..Z.
	InputTextFlagsCharsUppercase InputTextFlags = 1 << 2
	// InputTextFlagsCharsNoBlank filters out spaces, tabs.
	InputTextFlagsCharsNoBlank InputTextFlags = 1 << 3
	// InputTextFlagsAutoSelectAll selects entire text when first taking mouse focus.
	InputTextFlagsAutoSelectAll InputTextFlags = 1 << 4
	// InputTextFlagsEnterReturnsTrue returns 'true' when Enter is pressed (as opposed to when the value was modified).
	InputTextFlagsEnterReturnsTrue InputTextFlags = 1 << 5
	// InputTextFlagsCallbackCompletion for callback on pressing TAB (for completion handling).
	InputTextFlagsCallbackCompletion InputTextFlags = 1 << 6
	// InputTextFlagsCallbackHistory for callback on pressing Up/Down arrows (for history handling).
	InputTextFlagsCallbackHistory InputTextFlags = 1 << 7
	// InputTextFlagsCallbackAlways for callback on each iteration. User code may query cursor position, modify text buffer.
	InputTextFlagsCallbackAlways InputTextFlags = 1 << 8
	// InputTextFlagsCallbackCharFilter for callback on character inputs to replace or discard them.
	// Modify 'EventChar' to replace or discard, or return 1 in callback to discard.
	InputTextFlagsCallbackCharFilter InputTextFlags = 1 << 9
	// InputTextFlagsAllowTabInput when pressing TAB to input a '\t' character into the text field.
	InputTextFlagsAllowTabInput InputTextFlags = 1 << 10
	// InputTextFlagsCtrlEnterForNewLine in multi-line mode, unfocus with Enter, add new line with Ctrl+Enter
	// (default is opposite: unfocus with Ctrl+Enter, add line with Enter).
	InputTextFlagsCtrlEnterForNewLine InputTextFlags = 1 << 11
	// InputTextFlagsNoHorizontalScroll disables following the cursor horizontally.
	InputTextFlagsNoHorizontalScroll InputTextFlags = 1 << 12
	// InputTextFlagsAlwaysInsertMode sets insert mode.
	InputTextFlagsAlwaysInsertMode InputTextFlags = 1 << 13
	// InputTextFlagsReadOnly sets read-only mode.
	InputTextFlagsReadOnly InputTextFlags = 1 << 14
	// InputTextFlagsPassword sets password mode, display all characters as '*'.
	InputTextFlagsPassword InputTextFlags = 1 << 15
	// InputTextFlagsNoUndoRedo disables undo/redo. Note that input text owns the text data while active,
	// if you want to provide your own undo/redo stack you need e.g. to call ClearActiveID().
	InputTextFlagsNoUndoRedo InputTextFlags = 1 << 16
	// InputTextFlagsCharsScientific allows 0123456789.+-*/eE (Scientific notation input).
	InputTextFlagsCharsScientific InputTextFlags = 1 << 17
	// inputTextFlagsCallbackResize for callback on buffer capacity change requests.
	// inputTextFlagsCallbackResize InputTextFlags = 1 << 18
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
