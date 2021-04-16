package imgui

// #include "ListClipper.h"
// #include "imguiWrapper.h"
import "C"

// ListClipper is a helper to manually clip large list of items.
// If you are submitting lots of evenly spaced items and you have a random access to the list,
// you can perform coarse clipping based on visibility to save yourself from processing those items at all.
// The clipper calculates the range of visible items and advance the cursor to compensate for
// the non-visible items we have skipped.
// ImGui already clips items based on their bounds but it needs to measure text size to do so.
// Coarse clipping before submission makes this cost and your own data fetching/submission cost null.
//
// Usage
//  var clipper imgui.ListClipper
//  clipper.Begin(1000)  // we have 1000 elements, evenly spaced.
//  for clipper.Step()
//      for i := clipper.DisplayStart; i < clipper.DisplayEnd; i++
//          imgui.Text(fmt.Sprintf("line number %d", i))
//
// Step 0: the clipper let you process the first element, regardless of it being visible or not,
// so it can measure the element height (step skipped if user passed a known height at begin).
//
// Step 1: the clipper infers height from first element, calculates the actual range of elements to display,
// and positions the cursor before the first element.
//
// Step 2: dummy step only required if an explicit itemsHeight was passed to Begin() and user call Step().
// Does nothing and switch to Step 3.
//
// Step 3: the clipper validates that we have reached the expected Y position (corresponding to element DisplayEnd),
// advance the cursor to the end of the list and then returns 'false' to end the loop.
type ListClipper struct {
	StartPosY    float32
	ItemsHeight  float32
	ItemsCount   int
	StepNo       int
	DisplayStart int
	DisplayEnd   int
}

// wrapped return C struct and func for setting the values when done
func (clipper *ListClipper) wrapped() (out *C.IggListClipper, finisher func()) {
	if clipper == nil {
		return nil, func() {}
	}
	out = &C.IggListClipper{
		StartPosY:    C.float(clipper.StartPosY),
		ItemsHeight:  C.float(clipper.ItemsHeight),
		ItemsCount:   C.int(clipper.ItemsCount),
		StepNo:       C.int(clipper.StepNo),
		DisplayStart: C.int(clipper.DisplayStart),
		DisplayEnd:   C.int(clipper.DisplayEnd),
	}
	finisher = func() {
		clipper.StartPosY = float32(out.StartPosY)
		clipper.ItemsHeight = float32(out.ItemsHeight)
		clipper.ItemsCount = int(out.ItemsCount)
		clipper.StepNo = int(out.StepNo)
		clipper.DisplayStart = int(out.DisplayStart)
		clipper.DisplayEnd = int(out.DisplayEnd)
	}
	return
}

// Step must be called in a loop until it returns false.
// The DisplayStart/DisplayEnd fields will be set and you can process/draw those items.
func (clipper *ListClipper) Step() bool {
	arg, finishFunc := clipper.wrapped()
	defer finishFunc()
	return C.iggListClipperStep(arg) != 0
}

// Begin calls BeginV(itemsCount, -1.0) .
func (clipper *ListClipper) Begin(itemsCount int) {
	clipper.BeginV(itemsCount, -1.0)
}

// BeginV must be called before stepping.
// Use an itemCount of math.MaxInt if you don't know how many items you have.
// In this case the cursor won't be advanced in the final step.
//
// For itemsHeight, use -1.0 to be calculated automatically on first step.
// Otherwise pass in the distance between your items, typically
// GetTextLineHeightWithSpacing() or GetFrameHeightWithSpacing().
func (clipper *ListClipper) BeginV(itemsCount int, itemsHeight float32) {
	arg, finishFunc := clipper.wrapped()
	defer finishFunc()
	C.iggListClipperBegin(arg, C.int(itemsCount), C.float(itemsHeight))
}

// End resets the clipper. This function is automatically called on the last call of Step() that returns false.
func (clipper *ListClipper) End() {
	arg, finishFunc := clipper.wrapped()
	defer finishFunc()
	C.iggListClipperEnd(arg)
}
