package giu

import "github.com/AllenDang/imgui-go"

// InputTextFlags represents input text flags.
type InputTextFlags int

// input text flags.
const (
	// InputTextFlagsNone sets everything default.
	InputTextFlagsNone InputTextFlags = imgui.InputTextFlagsNone
	// InputTextFlagsCharsDecimal allows 0123456789.+-.
	InputTextFlagsCharsDecimal InputTextFlags = imgui.InputTextFlagsCharsDecimal
	// InputTextFlagsCharsHexadecimal allow 0123456789ABCDEFabcdef.
	InputTextFlagsCharsHexadecimal InputTextFlags = imgui.InputTextFlagsCharsHexadecimal
	// InputTextFlagsCharsUppercase turns a..z into A..Z.
	InputTextFlagsCharsUppercase InputTextFlags = imgui.InputTextFlagsCharsUppercase
	// InputTextFlagsCharsNoBlank filters out spaces, tabs.
	InputTextFlagsCharsNoBlank InputTextFlags = imgui.InputTextFlagsCharsNoBlank
	// InputTextFlagsAutoSelectAll selects entire text when first taking mouse focus.
	InputTextFlagsAutoSelectAll InputTextFlags = imgui.InputTextFlagsAutoSelectAll
	// InputTextFlagsEnterReturnsTrue returns 'true' when Enter is pressed (as opposed to when the value was modified).
	InputTextFlagsEnterReturnsTrue InputTextFlags = imgui.InputTextFlagsEnterReturnsTrue
	// InputTextFlagsCallbackCompletion for callback on pressing TAB (for completion handling).
	InputTextFlagsCallbackCompletion InputTextFlags = imgui.InputTextFlagsCallbackCompletion
	// InputTextFlagsCallbackHistory for callback on pressing Up/Down arrows (for history handling).
	InputTextFlagsCallbackHistory InputTextFlags = imgui.InputTextFlagsCallbackHistory
	// InputTextFlagsCallbackAlways for callback on each iteration. User code may query cursor position, modify text buffer.
	InputTextFlagsCallbackAlways InputTextFlags = imgui.InputTextFlagsCallbackAlways
	// InputTextFlagsCallbackCharFilter for callback on character inputs to replace or discard them.
	// Modify 'EventChar' to replace or discard, or return 1 in callback to discard.
	InputTextFlagsCallbackCharFilter InputTextFlags = imgui.InputTextFlagsCallbackCharFilter
	// InputTextFlagsAllowTabInput when pressing TAB to input a '\t' character into the text field.
	InputTextFlagsAllowTabInput InputTextFlags = imgui.InputTextFlagsAllowTabInput
	// InputTextFlagsCtrlEnterForNewLine in multi-line mode, unfocus with Enter, add new line with Ctrl+Enter
	// (default is opposite: unfocus with Ctrl+Enter, add line with Enter).
	InputTextFlagsCtrlEnterForNewLine InputTextFlags = imgui.InputTextFlagsCtrlEnterForNewLine
	// InputTextFlagsNoHorizontalScroll disables following the cursor horizontally.
	InputTextFlagsNoHorizontalScroll InputTextFlags = imgui.InputTextFlagsNoHorizontalScroll
	// InputTextFlagsAlwaysInsertMode sets insert mode.
	InputTextFlagsAlwaysInsertMode InputTextFlags = imgui.InputTextFlagsAlwaysInsertMode
	// InputTextFlagsReadOnly sets read-only mode.
	InputTextFlagsReadOnly InputTextFlags = imgui.InputTextFlagsReadOnly
	// InputTextFlagsPassword sets password mode, display all characters as '*'.
	InputTextFlagsPassword InputTextFlags = imgui.InputTextFlagsPassword
	// InputTextFlagsNoUndoRedo disables undo/redo. Note that input text owns the text data while active,
	// if you want to provide your own undo/redo stack you need e.g. to call ClearActiveID().
	InputTextFlagsNoUndoRedo InputTextFlags = imgui.InputTextFlagsNoUndoRedo
	// InputTextFlagsCharsScientific allows 0123456789.+-*/eE (Scientific notation input).
	InputTextFlagsCharsScientific InputTextFlags = imgui.InputTextFlagsCharsScientific
)

// WindowFlags represents a window flags (see (*WindowWidget).Flags.
type WindowFlags int

// window flags.
const (
	// WindowFlagsNone default = 0.
	WindowFlagsNone WindowFlags = imgui.WindowFlagsNone
	// WindowFlagsNoTitleBar disables title-bar.
	WindowFlagsNoTitleBar WindowFlags = imgui.WindowFlagsNoTitleBar
	// WindowFlagsNoResize disables user resizing with the lower-right grip.
	WindowFlagsNoResize WindowFlags = imgui.WindowFlagsNoResize
	// WindowFlagsNoMove disables user moving the window.
	WindowFlagsNoMove WindowFlags = imgui.WindowFlagsNoMove
	// WindowFlagsNoScrollbar disables scrollbars. Window can still scroll with mouse or programmatically.
	WindowFlagsNoScrollbar WindowFlags = imgui.WindowFlagsNoScrollbar
	// WindowFlagsNoScrollWithMouse disables user vertically scrolling with mouse wheel. On child window, mouse wheel
	// will be forwarded to the parent unless NoScrollbar is also set.
	WindowFlagsNoScrollWithMouse WindowFlags = imgui.WindowFlagsNoScrollWithMouse
	// WindowFlagsNoCollapse disables user collapsing window by double-clicking on it.
	WindowFlagsNoCollapse WindowFlags = imgui.WindowFlagsNoCollapse
	// WindowFlagsAlwaysAutoResize resizes every window to its content every frame.
	WindowFlagsAlwaysAutoResize WindowFlags = imgui.WindowFlagsAlwaysAutoResize
	// WindowFlagsNoBackground disables drawing background color (WindowBg, etc.) and outside border. Similar as using
	// SetNextWindowBgAlpha(0.0f).
	WindowFlagsNoBackground WindowFlags = imgui.WindowFlagsNoBackground
	// WindowFlagsNoSavedSettings will never load/save settings in .ini file.
	WindowFlagsNoSavedSettings WindowFlags = imgui.WindowFlagsNoSavedSettings
	// WindowFlagsNoMouseInputs disables catching mouse, hovering test with pass through.
	WindowFlagsNoMouseInputs WindowFlags = imgui.WindowFlagsNoMouseInputs
	// WindowFlagsMenuBar has a menu-bar.
	WindowFlagsMenuBar WindowFlags = imgui.WindowFlagsMenuBar
	// WindowFlagsHorizontalScrollbar allows horizontal scrollbar to appear (off by default). You may use
	// SetNextWindowContentSize(ImVec2(width,0.0f)); prior to calling Begin() to specify width. Read code in imgui_demo
	// in the "Horizontal Scrolling" section.
	WindowFlagsHorizontalScrollbar WindowFlags = imgui.WindowFlagsHorizontalScrollbar
	// WindowFlagsNoFocusOnAppearing disables taking focus when transitioning from hidden to visible state.
	WindowFlagsNoFocusOnAppearing WindowFlags = imgui.WindowFlagsNoFocusOnAppearing
	// WindowFlagsNoBringToFrontOnFocus disables bringing window to front when taking focus. e.g. clicking on it or
	// programmatically giving it focus.
	WindowFlagsNoBringToFrontOnFocus WindowFlags = imgui.WindowFlagsNoBringToFrontOnFocus
	// WindowFlagsAlwaysVerticalScrollbar always shows vertical scrollbar, even if ContentSize.y < Size.y .
	WindowFlagsAlwaysVerticalScrollbar WindowFlags = imgui.WindowFlagsAlwaysVerticalScrollbar
	// WindowFlagsAlwaysHorizontalScrollbar always shows horizontal scrollbar, even if ContentSize.x < Size.x .
	WindowFlagsAlwaysHorizontalScrollbar WindowFlags = imgui.WindowFlagsAlwaysHorizontalScrollbar
	// WindowFlagsAlwaysUseWindowPadding ensures child windows without border uses style.WindowPadding (ignored by
	// default for non-bordered child windows, because more convenient).
	WindowFlagsAlwaysUseWindowPadding WindowFlags = imgui.WindowFlagsAlwaysUseWindowPadding
	// WindowFlagsNoNavInputs has no gamepad/keyboard navigation within the window.
	WindowFlagsNoNavInputs WindowFlags = imgui.WindowFlagsNoNavInputs
	// WindowFlagsNoNavFocus has no focusing toward this window with gamepad/keyboard navigation
	// (e.g. skipped by CTRL+TAB).
	WindowFlagsNoNavFocus WindowFlags = imgui.WindowFlagsNoNavFocus
	// WindowFlagsUnsavedDocument appends '*' to title without affecting the ID, as a convenience to avoid using the
	// ### operator. When used in a tab/docking context, tab is selected on closure and closure is deferred by one
	// frame to allow code to cancel the closure (with a confirmation popup, etc.) without flicker.
	WindowFlagsUnsavedDocument WindowFlags = imgui.WindowFlagsUnsavedDocument

	// WindowFlagsNoNav combines WindowFlagsNoNavInputs and WindowFlagsNoNavFocus.
	WindowFlagsNoNav WindowFlags = imgui.WindowFlagsNoNav
	// WindowFlagsNoDecoration combines WindowFlagsNoTitleBar, WindowFlagsNoResize, WindowFlagsNoScrollbar and
	// WindowFlagsNoCollapse.
	WindowFlagsNoDecoration WindowFlags = imgui.WindowFlagsNoDecoration
	// WindowFlagsNoInputs combines WindowFlagsNoMouseInputs, WindowFlagsNoNavInputs and WindowFlagsNoNavFocus.
	WindowFlagsNoInputs WindowFlags = imgui.WindowFlagsNoInputs
)

// ComboFlags represents imgui.ComboFlags.
type ComboFlags int

// combo flags list.
const (
	// ComboFlagsNone default = 0.
	ComboFlagsNone ComboFlags = imgui.ComboFlagsNone
	// ComboFlagsPopupAlignLeft aligns the popup toward the left by default.
	ComboFlagsPopupAlignLeft ComboFlags = imgui.ComboFlagsPopupAlignLeft
	// ComboFlagsHeightSmall has max ~4 items visible.
	// Tip: If you want your combo popup to be a specific size you can use SetNextWindowSizeConstraints() prior to calling BeginCombo().
	ComboFlagsHeightSmall ComboFlags = imgui.ComboFlagsHeightSmall
	// ComboFlagsHeightRegular has max ~8 items visible (default).
	ComboFlagsHeightRegular ComboFlags = imgui.ComboFlagsHeightRegular
	// ComboFlagsHeightLarge has max ~20 items visible.
	ComboFlagsHeightLarge ComboFlags = imgui.ComboFlagsHeightLarge
	// ComboFlagsHeightLargest has as many fitting items as possible.
	ComboFlagsHeightLargest ComboFlags = imgui.ComboFlagsHeightLargest
	// ComboFlagsNoArrowButton displays on the preview box without the square arrow button.
	ComboFlagsNoArrowButton ComboFlags = imgui.ComboFlagsNoArrowButton
	// ComboFlagsNoPreview displays only a square arrow button.
	ComboFlagsNoPreview ComboFlags = imgui.ComboFlagsNoPreview
)

// SelectableFlags represents imgui.SelectableFlags.
type SelectableFlags int

// selectable flags list.
const (
	// SelectableFlagsNone default = 0.
	SelectableFlagsNone SelectableFlags = imgui.SelectableFlagsNone
	// SelectableFlagsDontClosePopups makes clicking the selectable not close any parent popup windows.
	SelectableFlagsDontClosePopups SelectableFlags = imgui.SelectableFlagsDontClosePopups
	// SelectableFlagsSpanAllColumns allows the selectable frame to span all columns (text will still fit in current column).
	SelectableFlagsSpanAllColumns SelectableFlags = imgui.SelectableFlagsSpanAllColumns
	// SelectableFlagsAllowDoubleClick generates press events on double clicks too.
	SelectableFlagsAllowDoubleClick SelectableFlags = imgui.SelectableFlagsAllowDoubleClick
	// SelectableFlagsDisabled disallows selection and displays text in a greyed out color.
	SelectableFlagsDisabled SelectableFlags = imgui.SelectableFlagsDisabled
)

// TabItemFlags represents tab item flags.
type TabItemFlags int

// tab item flags list.
const (
	// TabItemFlagsNone default = 0.
	TabItemFlagsNone TabItemFlags = imgui.TabItemFlagsNone
	// TabItemFlagsUnsavedDocument Append '*' to title without affecting the ID, as a convenience to avoid using the
	// ### operator. Also: tab is selected on closure and closure is deferred by one frame to allow code to undo it
	// without flicker.
	TabItemFlagsUnsavedDocument TabItemFlags = imgui.TabItemFlagsUnsavedDocument
	// TabItemFlagsSetSelected Trigger flag to programmatically make the tab selected when calling BeginTabItem().
	TabItemFlagsSetSelected TabItemFlags = imgui.TabItemFlagsSetSelected
	// TabItemFlagsNoCloseWithMiddleMouseButton  Disable behavior of closing tabs (that are submitted with
	// p_open != NULL) with middle mouse button. You can still repro this behavior on user's side with if
	// (IsItemHovered() && IsMouseClicked(2)) *p_open = false.
	TabItemFlagsNoCloseWithMiddleMouseButton TabItemFlags = imgui.TabItemFlagsNoCloseWithMiddleMouseButton
	// TabItemFlagsNoPushID Don't call PushID(tab->ID)/PopID() on BeginTabItem()/EndTabItem().
	TabItemFlagsNoPushID TabItemFlags = imgui.TabItemFlagsNoPushID
)

// TabBarFlags represents imgui.TabBarFlags.
type TabBarFlags int

// tab bar flags list.
const (
	// TabBarFlagsNone default = 0.
	TabBarFlagsNone TabBarFlags = imgui.TabBarFlagsNone
	// TabBarFlagsReorderable Allow manually dragging tabs to re-order them + New tabs are appended at the end of list.
	TabBarFlagsReorderable TabBarFlags = imgui.TabBarFlagsReorderable
	// TabBarFlagsAutoSelectNewTabs Automatically select new tabs when they appear.
	TabBarFlagsAutoSelectNewTabs TabBarFlags = imgui.TabBarFlagsAutoSelectNewTabs
	// TabBarFlagsTabListPopupButton Disable buttons to open the tab list popup.
	TabBarFlagsTabListPopupButton TabBarFlags = imgui.TabBarFlagsTabListPopupButton
	// TabBarFlagsNoCloseWithMiddleMouseButton Disable behavior of closing tabs (that are submitted with p_open != NULL)
	// with middle mouse button. You can still repro this behavior on user's side with if
	// (IsItemHovered() && IsMouseClicked(2)) *p_open = false.
	TabBarFlagsNoCloseWithMiddleMouseButton TabBarFlags = imgui.TabBarFlagsNoCloseWithMiddleMouseButton
	// TabBarFlagsNoTabListScrollingButtons Disable scrolling buttons (apply when fitting policy is
	// TabBarFlagsFittingPolicyScroll).
	TabBarFlagsNoTabListScrollingButtons TabBarFlags = imgui.TabBarFlagsNoTabListScrollingButtons
	// TabBarFlagsNoTooltip Disable tooltips when hovering a tab.
	TabBarFlagsNoTooltip TabBarFlags = imgui.TabBarFlagsNoTooltip
	// TabBarFlagsFittingPolicyResizeDown Resize tabs when they don't fit.
	TabBarFlagsFittingPolicyResizeDown TabBarFlags = imgui.TabBarFlagsFittingPolicyResizeDown
	// TabBarFlagsFittingPolicyScroll Add scroll buttons when tabs don't fit.
	TabBarFlagsFittingPolicyScroll TabBarFlags = imgui.TabBarFlagsFittingPolicyScroll
	// TabBarFlagsFittingPolicyMask combines
	// TabBarFlagsFittingPolicyResizeDown and TabBarFlagsFittingPolicyScroll.
	TabBarFlagsFittingPolicyMask TabBarFlags = imgui.TabBarFlagsFittingPolicyMask
	// TabBarFlagsFittingPolicyDefault alias for TabBarFlagsFittingPolicyResizeDown.
	TabBarFlagsFittingPolicyDefault TabBarFlags = imgui.TabBarFlagsFittingPolicyDefault
)

// TreeNodeFlags represents tree node widget flags.
type TreeNodeFlags int

// tree node flags list.
const (
	// TreeNodeFlagsNone default = 0.
	TreeNodeFlagsNone TreeNodeFlags = imgui.TreeNodeFlagsNone
	// TreeNodeFlagsSelected draws as selected.
	TreeNodeFlagsSelected TreeNodeFlags = imgui.TreeNodeFlagsSelected
	// TreeNodeFlagsFramed draws full colored frame (e.g. for CollapsingHeader).
	TreeNodeFlagsFramed TreeNodeFlags = imgui.TreeNodeFlagsFramed
	// TreeNodeFlagsAllowItemOverlap hit testing to allow subsequent widgets to overlap this one.
	TreeNodeFlagsAllowItemOverlap TreeNodeFlags = imgui.TreeNodeFlagsAllowItemOverlap
	// TreeNodeFlagsNoTreePushOnOpen doesn't do a TreePush() when open
	// (e.g. for CollapsingHeader) = no extra indent nor pushing on ID stack.
	TreeNodeFlagsNoTreePushOnOpen TreeNodeFlags = imgui.TreeNodeFlagsNoTreePushOnOpen
	// TreeNodeFlagsNoAutoOpenOnLog doesn't automatically and temporarily open node when Logging is active
	// (by default logging will automatically open tree nodes).
	TreeNodeFlagsNoAutoOpenOnLog TreeNodeFlags = imgui.TreeNodeFlagsNoAutoOpenOnLog
	// TreeNodeFlagsDefaultOpen defaults node to be open.
	TreeNodeFlagsDefaultOpen TreeNodeFlags = imgui.TreeNodeFlagsDefaultOpen
	// TreeNodeFlagsOpenOnDoubleClick needs double-click to open node.
	TreeNodeFlagsOpenOnDoubleClick TreeNodeFlags = imgui.TreeNodeFlagsOpenOnDoubleClick
	// TreeNodeFlagsOpenOnArrow opens only when clicking on the arrow part.
	// If TreeNodeFlagsOpenOnDoubleClick is also set, single-click arrow or double-click all box to open.
	TreeNodeFlagsOpenOnArrow TreeNodeFlags = imgui.TreeNodeFlagsOpenOnArrow
	// TreeNodeFlagsLeaf allows no collapsing, no arrow (use as a convenience for leaf nodes).
	TreeNodeFlagsLeaf TreeNodeFlags = imgui.TreeNodeFlagsLeaf
	// TreeNodeFlagsBullet displays a bullet instead of an arrow.
	TreeNodeFlagsBullet TreeNodeFlags = imgui.TreeNodeFlagsBullet
	// TreeNodeFlagsFramePadding uses FramePadding (even for an unframed text node) to
	// vertically align text baseline to regular widget height. Equivalent to calling AlignTextToFramePadding().
	TreeNodeFlagsFramePadding TreeNodeFlags = imgui.TreeNodeFlagsFramePadding
	// TreeNodeFlagsSpanAvailWidth extends hit box to the right-most edge, even if not framed.
	// This is not the default in order to allow adding other items on the same line.
	// In the future we may refactor the hit system to be front-to-back, allowing natural overlaps
	// and then this can become the default.
	TreeNodeFlagsSpanAvailWidth TreeNodeFlags = imgui.TreeNodeFlagsSpanAvailWidth
	// TreeNodeFlagsSpanFullWidth extends hit box to the left-most and right-most edges (bypass the indented area).
	TreeNodeFlagsSpanFullWidth TreeNodeFlags = imgui.TreeNodeFlagsSpanFullWidth
	// TreeNodeFlagsNavLeftJumpsBackHere (WIP) Nav: left direction may move to this TreeNode() from any of its child
	// (items submitted between TreeNode and TreePop).
	TreeNodeFlagsNavLeftJumpsBackHere TreeNodeFlags = imgui.TreeNodeFlagsNavLeftJumpsBackHere
	// TreeNodeFlagsCollapsingHeader combines TreeNodeFlagsFramed and TreeNodeFlagsNoAutoOpenOnLog.
	TreeNodeFlagsCollapsingHeader TreeNodeFlags = imgui.TreeNodeFlagsCollapsingHeader
)

// FocusedFlags represents imgui.FocusedFlags.
type FocusedFlags int

// focused flags list.
const (
	FocusedFlagsNone             = imgui.FocusedFlagsNone
	FocusedFlagsChildWindows     = imgui.FocusedFlagsChildWindows     // Return true if any children of the window is focused
	FocusedFlagsRootWindow       = imgui.FocusedFlagsRootWindow       // Test from root window (top most parent of the current hierarchy)
	FocusedFlagsAnyWindow        = imgui.FocusedFlagsAnyWindow        // Return true if any window is focused. Important: If you are trying to tell how to dispatch your low-level inputs do NOT use this. Use 'io.WantCaptureMouse' instead! Please read the FAQ!
	FocusedFlagsNoPopupHierarchy = imgui.FocusedFlagsNoPopupHierarchy // Do not consider popup hierarchy (do not treat popup emitter as parent of popup) (when used with ChildWindows or RootWindow)
	// FocusedFlagsDockHierarchy               = 1 << 4   // Consider docking hierarchy (treat dockspace host as parent of docked window) (when used with ChildWindows or RootWindow).
	FocusedFlagsRootAndChildWindows = imgui.FocusedFlagsRootAndChildWindows
)

// HoveredFlags represents a hovered flags.
type HoveredFlags int

// hovered flags list.
const (
	// HoveredFlagsNone Return true if directly over the item/window, not obstructed by another window,
	// not obstructed by an active popup or modal blocking inputs under them.
	HoveredFlagsNone HoveredFlags = imgui.HoveredFlagsNone
	// HoveredFlagsChildWindows IsWindowHovered() only: Return true if any children of the window is hovered.
	HoveredFlagsChildWindows HoveredFlags = imgui.HoveredFlagsChildWindows
	// HoveredFlagsRootWindow IsWindowHovered() only: Test from root window (top most parent of the current hierarchy).
	HoveredFlagsRootWindow HoveredFlags = imgui.HoveredFlagsRootWindow
	// HoveredFlagsAnyWindow IsWindowHovered() only: Return true if any window is hovered.
	HoveredFlagsAnyWindow HoveredFlags = imgui.HoveredFlagsAnyWindow
	// HoveredFlagsAllowWhenBlockedByPopup Return true even if a popup window is normally blocking access to this item/window.
	HoveredFlagsAllowWhenBlockedByPopup HoveredFlags = imgui.HoveredFlagsAllowWhenBlockedByPopup
	// HoveredFlagsAllowWhenBlockedByActiveItem Return true even if an active item is blocking access to this item/window.
	// Useful for Drag and Drop patterns.
	HoveredFlagsAllowWhenBlockedByActiveItem HoveredFlags = imgui.HoveredFlagsAllowWhenBlockedByActiveItem
	// HoveredFlagsAllowWhenOverlapped Return true even if the position is overlapped by another window.
	HoveredFlagsAllowWhenOverlapped HoveredFlags = imgui.HoveredFlagsAllowWhenOverlapped
	// HoveredFlagsAllowWhenDisabled Return true even if the item is disabled.
	HoveredFlagsAllowWhenDisabled HoveredFlags = imgui.HoveredFlagsAllowWhenDisabled
)

// ColorEditFlags for ColorEdit3V(), etc.
type ColorEditFlags int

// list of color edit flags.
const (
	// ColorEditFlagsNone default = 0.
	ColorEditFlagsNone ColorEditFlags = imgui.ColorEditFlagsNone
	// ColorEditFlagsNoAlpha ignores Alpha component (read 3 components from the input pointer).
	ColorEditFlagsNoAlpha ColorEditFlags = imgui.ColorEditFlagsNoAlpha
	// ColorEditFlagsNoPicker disables picker when clicking on colored square.
	ColorEditFlagsNoPicker ColorEditFlags = imgui.ColorEditFlagsNoPicker
	// ColorEditFlagsNoOptions disables toggling options menu when right-clicking on inputs/small preview.
	ColorEditFlagsNoOptions ColorEditFlags = imgui.ColorEditFlagsNoOptions
	// ColorEditFlagsNoSmallPreview disables colored square preview next to the inputs. (e.g. to show only the inputs).
	ColorEditFlagsNoSmallPreview ColorEditFlags = imgui.ColorEditFlagsNoSmallPreview
	// ColorEditFlagsNoInputs disables inputs sliders/text widgets (e.g. to show only the small preview colored square).
	ColorEditFlagsNoInputs ColorEditFlags = imgui.ColorEditFlagsNoInputs
	// ColorEditFlagsNoTooltip disables tooltip when hovering the preview.
	ColorEditFlagsNoTooltip ColorEditFlags = imgui.ColorEditFlagsNoTooltip
	// ColorEditFlagsNoLabel disables display of inline text label (the label is still forwarded to the tooltip and picker).
	ColorEditFlagsNoLabel ColorEditFlags = imgui.ColorEditFlagsNoLabel
	// ColorEditFlagsNoDragDrop disables drag and drop target. ColorButton: disable drag and drop source.
	ColorEditFlagsNoDragDrop ColorEditFlags = imgui.ColorEditFlagsNoDragDrop

	// User Options (right-click on widget to change some of them). You can set application defaults using SetColorEditOptions().
	// The idea is that you probably don't want to override them in most of your calls, let the user choose and/or call SetColorEditOptions()
	// during startup.

	// ColorEditFlagsAlphaBar shows vertical alpha bar/gradient in picker.
	ColorEditFlagsAlphaBar ColorEditFlags = imgui.ColorEditFlagsAlphaBar
	// ColorEditFlagsAlphaPreview displays preview as a transparent color over a checkerboard, instead of opaque.
	ColorEditFlagsAlphaPreview ColorEditFlags = imgui.ColorEditFlagsAlphaPreview
	// ColorEditFlagsAlphaPreviewHalf displays half opaque / half checkerboard, instead of opaque.
	ColorEditFlagsAlphaPreviewHalf ColorEditFlags = imgui.ColorEditFlagsAlphaPreviewHalf
	// ColorEditFlagsHDR = (WIP) surrently only disable 0.0f..1.0f limits in RGBA edition (note: you probably want to use
	// ImGuiColorEditFlags_Float flag as well).
	ColorEditFlagsHDR ColorEditFlags = imgui.ColorEditFlagsHDR
	// ColorEditFlagsRGB sets the format as RGB.
	ColorEditFlagsRGB ColorEditFlags = imgui.ColorEditFlagsRGB
	// ColorEditFlagsHSV sets the format as HSV.
	ColorEditFlagsHSV ColorEditFlags = imgui.ColorEditFlagsHSV
	// ColorEditFlagsHEX sets the format as HEX.
	ColorEditFlagsHEX ColorEditFlags = imgui.ColorEditFlagsHEX
	// ColorEditFlagsUint8 _display_ values formatted as 0..255.
	ColorEditFlagsUint8 ColorEditFlags = imgui.ColorEditFlagsUint8
	// ColorEditFlagsFloat _display_ values formatted as 0.0f..1.0f floats instead of 0..255 integers. No round-trip of value via integers.
	ColorEditFlagsFloat ColorEditFlags = imgui.ColorEditFlagsFloat
)

// TableFlags represents table flags.
type TableFlags int

// Table flags enum:.
const (
	TableFlagsNone                       TableFlags = TableFlags(imgui.TableFlags_None)
	TableFlagsResizable                  TableFlags = TableFlags(imgui.TableFlags_Resizable)
	TableFlagsReorderable                TableFlags = TableFlags(imgui.TableFlags_Reorderable)
	TableFlagsHideable                   TableFlags = TableFlags(imgui.TableFlags_Hideable)
	TableFlagsSortable                   TableFlags = TableFlags(imgui.TableFlags_Sortable)
	TableFlagsNoSavedSettings            TableFlags = TableFlags(imgui.TableFlags_NoSavedSettings)
	TableFlagsContextMenuInBody          TableFlags = TableFlags(imgui.TableFlags_ContextMenuInBody)
	TableFlagsRowBg                      TableFlags = TableFlags(imgui.TableFlags_RowBg)
	TableFlagsBordersInnerH              TableFlags = TableFlags(imgui.TableFlags_BordersInnerH)
	TableFlagsBordersOuterH              TableFlags = TableFlags(imgui.TableFlags_BordersOuterH)
	TableFlagsBordersInnerV              TableFlags = TableFlags(imgui.TableFlags_BordersInnerV)
	TableFlagsBordersOuterV              TableFlags = TableFlags(imgui.TableFlags_BordersOuterV)
	TableFlagsBordersH                   TableFlags = TableFlags(imgui.TableFlags_BordersH)
	TableFlagsBordersV                   TableFlags = TableFlags(imgui.TableFlags_BordersV)
	TableFlagsBordersInner               TableFlags = TableFlags(imgui.TableFlags_BordersInner)
	TableFlagsBordersOuter               TableFlags = TableFlags(imgui.TableFlags_BordersOuter)
	TableFlagsBorders                    TableFlags = TableFlags(imgui.TableFlags_Borders)
	TableFlagsNoBordersInBody            TableFlags = TableFlags(imgui.TableFlags_NoBordersInBody)
	TableFlagsNoBordersInBodyUntilResize TableFlags = TableFlags(imgui.TableFlags_NoBordersInBodyUntilResizeTableFlags)
	TableFlagsSizingFixedFit             TableFlags = TableFlags(imgui.TableFlags_SizingFixedFit)
	TableFlagsSizingFixedSame            TableFlags = TableFlags(imgui.TableFlags_SizingFixedSame)
	TableFlagsSizingStretchProp          TableFlags = TableFlags(imgui.TableFlags_SizingStretchProp)
	TableFlagsSizingStretchSame          TableFlags = TableFlags(imgui.TableFlags_SizingStretchSame)
	TableFlagsNoHostExtendX              TableFlags = TableFlags(imgui.TableFlags_NoHostExtendX)
	TableFlagsNoHostExtendY              TableFlags = TableFlags(imgui.TableFlags_NoHostExtendY)
	TableFlagsNoKeepColumnsVisible       TableFlags = TableFlags(imgui.TableFlags_NoKeepColumnsVisible)
	TableFlagsPreciseWidths              TableFlags = TableFlags(imgui.TableFlags_PreciseWidths)
	TableFlagsNoClip                     TableFlags = TableFlags(imgui.TableFlags_NoClip)
	TableFlagsPadOuterX                  TableFlags = TableFlags(imgui.TableFlags_PadOuterX)
	TableFlagsNoPadOuterX                TableFlags = TableFlags(imgui.TableFlags_NoPadOuterX)
	TableFlagsNoPadInnerX                TableFlags = TableFlags(imgui.TableFlags_NoPadInnerX)
	TableFlagsScrollX                    TableFlags = TableFlags(imgui.TableFlags_ScrollX)
	TableFlagsScrollY                    TableFlags = TableFlags(imgui.TableFlags_ScrollY)
	TableFlagsSortMulti                  TableFlags = TableFlags(imgui.TableFlags_SortMulti)
	TableFlagsSortTristate               TableFlags = TableFlags(imgui.TableFlags_SortTristate)
	TableFlagsSizingMask                 TableFlags = TableFlags(imgui.TableFlags_SizingMask_)
)

// TableRowFlags represents table row flags.
type TableRowFlags int

// table row flags:.
const (
	TableRowFlagsNone TableRowFlags = TableRowFlags(imgui.TableRowFlags_None)
	// Identify header row (set default background color + width of its contents accounted different for auto column width).
	TableRowFlagsHeaders TableRowFlags = TableRowFlags(imgui.TableRowFlags_Headers)
)

// TableColumnFlags represents a flags for table column (see (*TableColumnWidget).Flags()).
type TableColumnFlags int

// table column flags list.
const (
	// Input configuration flags.
	TableColumnFlagsNone                 TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_None)
	TableColumnFlagsDefaultHide          TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_DefaultHide)
	TableColumnFlagsDefaultSort          TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_DefaultSort)
	TableColumnFlagsWidthStretch         TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_WidthStretch)
	TableColumnFlagsWidthFixed           TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_WidthFixed)
	TableColumnFlagsNoResize             TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_NoResize)
	TableColumnFlagsNoReorder            TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_NoReorder)
	TableColumnFlagsNoHide               TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_NoHide)
	TableColumnFlagsNoClip               TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_NoClip)
	TableColumnFlagsNoSort               TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_NoSort)
	TableColumnFlagsNoSortAscending      TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_NoSortAscending)
	TableColumnFlagsNoSortDescending     TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_NoSortDescending)
	TableColumnFlagsNoHeaderWidth        TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_NoHeaderWidth)
	TableColumnFlagsPreferSortAscending  TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_PreferSortAscending)
	TableColumnFlagsPreferSortDescending TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_PreferSortDescending)
	TableColumnFlagsIndentEnable         TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_IndentEnable)
	TableColumnFlagsIndentDisable        TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_IndentDisable)

	// Output status flags read-only via TableGetColumnFlags().
	TableColumnFlagsIsEnabled TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_IsEnabled)
	TableColumnFlagsIsVisible TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_IsVisible)
	TableColumnFlagsIsSorted  TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_IsSorted)
	TableColumnFlagsIsHovered TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_IsHovered)

	// [Internal] Combinations and masks.
	TableColumnFlagsWidthMask      TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_WidthMask_)
	TableColumnFlagsIndentMask     TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_IndentMask_)
	TableColumnFlagsStatusMask     TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_StatusMask_)
	TableColumnFlagsNoDirectResize TableColumnFlags = TableColumnFlags(imgui.TableColumnFlags_NoDirectResize_)
)

// SliderFlags represents imgui.SliderFlags
// TODO: Hard-reffer to these constants.
type SliderFlags int

// slider flags.
const (
	SliderFlagsNone SliderFlags = 0
	// Clamp value to min/max bounds when input manually with CTRL+Click. By default CTRL+Click allows going out of bounds.
	SliderFlagsAlwaysClamp SliderFlags = 1 << 4
	// Make the widget logarithmic (linear otherwise). Consider using ImGuiSliderFlagsNoRoundToFormat with this if using
	// a format-string with small amount of digits.
	SliderFlagsLogarithmic SliderFlags = 1 << 5
	// Disable rounding underlying value to match precision of the display format string (e.g. %.3f values are rounded to those 3 digits).
	SliderFlagsNoRoundToFormat SliderFlags = 1 << 6
	// Disable CTRL+Click or Enter key allowing to input text directly into the widget.
	SliderFlagsNoInput SliderFlags = 1 << 7
	// [Internal] We treat using those bits as being potentially a 'float power' argument from the previous API that has got miscast
	// to this enum, and will trigger an assert if needed.
	SliderFlagsInvalidMask SliderFlags = 0x7000000F
)

// PlotFlags represents imgui.ImPlotFlags.
type PlotFlags int

// plot flags.
const (
	PlotFlagsNone        = PlotFlags(imgui.ImPlotFlags_None)
	PlotFlagsNoTitle     = PlotFlags(imgui.ImPlotFlags_NoTitle)
	PlotFlagsNoLegend    = PlotFlags(imgui.ImPlotFlags_NoLegend)
	PlotFlagsNoMenus     = PlotFlags(imgui.ImPlotFlags_NoMenus)
	PlotFlagsNoBoxSelect = PlotFlags(imgui.ImPlotFlags_NoBoxSelect)
	PlotFlagsNoMousePos  = PlotFlags(imgui.ImPlotFlags_NoMousePos)
	PlotFlagsNoHighlight = PlotFlags(imgui.ImPlotFlags_NoHighlight)
	PlotFlagsNoChild     = PlotFlags(imgui.ImPlotFlags_NoChild)
	PlotFlagsEqual       = PlotFlags(imgui.ImPlotFlags_Equal)
	PlotFlagsYAxis2      = PlotFlags(imgui.ImPlotFlags_YAxis2)
	PlotFlagsYAxis3      = PlotFlags(imgui.ImPlotFlags_YAxis3)
	PlotFlagsQuery       = PlotFlags(imgui.ImPlotFlags_Query)
	PlotFlagsCrosshairs  = PlotFlags(imgui.ImPlotFlags_Crosshairs)
	PlotFlagsAntiAliased = PlotFlags(imgui.ImPlotFlags_AntiAliased)
	PlotFlagsCanvasOnly  = PlotFlags(imgui.ImPlotFlags_CanvasOnly)
)

// PlotAxisFlags represents imgui.ImPlotAxisFlags.
type PlotAxisFlags int

// plot axis flags.
const (
	PlotAxisFlagsNone          PlotAxisFlags = PlotAxisFlags(imgui.ImPlotAxisFlags_None)
	PlotAxisFlagsNoLabel       PlotAxisFlags = PlotAxisFlags(imgui.ImPlotAxisFlags_NoLabel)
	PlotAxisFlagsNoGridLines   PlotAxisFlags = PlotAxisFlags(imgui.ImPlotAxisFlags_NoGridLines)
	PlotAxisFlagsNoTickMarks   PlotAxisFlags = PlotAxisFlags(imgui.ImPlotAxisFlags_NoTickMarks)
	PlotAxisFlagsNoTickLabels  PlotAxisFlags = PlotAxisFlags(imgui.ImPlotAxisFlags_NoTickLabels)
	PlotAxisFlagsForeground    PlotAxisFlags = PlotAxisFlags(imgui.ImPlotAxisFlags_Foreground)
	PlotAxisFlagsLogScale      PlotAxisFlags = PlotAxisFlags(imgui.ImPlotAxisFlags_LogScale)
	PlotAxisFlagsTime          PlotAxisFlags = PlotAxisFlags(imgui.ImPlotAxisFlags_Time)
	PlotAxisFlagsInvert        PlotAxisFlags = PlotAxisFlags(imgui.ImPlotAxisFlags_Invert)
	PlotAxisFlagsNoInitialFit  PlotAxisFlags = PlotAxisFlags(imgui.ImPlotAxisFlags_NoInitialFit)
	PlotAxisFlagsAutoFit       PlotAxisFlags = PlotAxisFlags(imgui.ImPlotAxisFlags_AutoFit)
	PlotAxisFlagsRangeFit      PlotAxisFlags = PlotAxisFlags(imgui.ImPlotAxisFlags_RangeFit)
	PlotAxisFlagsLockMin       PlotAxisFlags = PlotAxisFlags(imgui.ImPlotAxisFlags_LockMin)
	PlotAxisFlagsLockMax       PlotAxisFlags = PlotAxisFlags(imgui.ImPlotAxisFlags_LockMax)
	PlotAxisFlagsLock          PlotAxisFlags = PlotAxisFlags(imgui.ImPlotAxisFlags_Lock)
	PlotAxisFlagsNoDecorations PlotAxisFlags = PlotAxisFlags(imgui.ImPlotAxisFlags_NoDecorations)
)
