package main

import (
	"errors"
	"fmt"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"time"

	"github.com/AllenDang/cimgui-go/imgui"
	"github.com/AllenDang/cimgui-go/imnodes"
	"github.com/AllenDang/cimgui-go/implot"
	g "github.com/AllenDang/giu"
)

var (
	fallback        = &g.ReflectiveBoundTexture{}
	dynamicImage    = &g.StatefulReflectiveBoundTexture{}
	imageScaleX     = float32(1.0)
	imageScaleY     = float32(1.0)
	linkedScale     = true
	dynamicImageUrl = "https://www.pngfind.com/pngs/b/465-4652097_gopher-png.png"
	showWindow      = true
	wnd             *g.MasterWindow
	footerLabel     string
	inputRootfs     = "."
)

func TextHeader(str string, col color.RGBA) *g.StyleSetter {
	return g.Style().SetColor(g.StyleColorText, col).To(
		g.Align(g.AlignCenter).To(
			g.Label(str),
		),
	)
}

func CanLoadHeader() *g.AlignmentSetter {
	return g.Align(g.AlignCenter).To(

		g.Row(
			g.InputText(&dynamicImageUrl),
			g.Button("Load Image from URL").OnClick(func() {
				err := dynamicImage.ResetState()
				if err == nil {
					dynamicImage.SetSurfaceFromURL(dynamicImageUrl, time.Second*5, false)
				}
			}),
		))
}

func HeaderWidget() g.Widget {
	if dynamicImage.GetState() == g.SsLoading {
		return TextHeader("Image Is Currently loading...", color.RGBA{0x80, 0x80, 0xFF, 255})
	}
	return CanLoadHeader()
}

func FooterWidget(label string) g.Widget {
	return g.Row(TextHeader(label, color.RGBA{0xFF, 0xFF, 0xFF, 255}))
}

func ShouldReturnImage() g.Widget {
	if dynamicImage.GetState() != g.SsSuccess {
		return fallback.ToImageWidget().Size(-1, -1)
	}
	return g.Custom(func() { dynamicImage.ToImageWidget().Scale(imageScaleX, imageScaleY).Build() })
}

func ShouldReturnPanel() g.Widget {

	return g.Custom(func() {
		imgui.SeparatorText("Image Scale")
		if imgui.Button("Reset##Scaling") {
			imageScaleX = 1.0
			imageScaleY = 1.0
		}
		imgui.SameLine()
		imgui.Checkbox("Linked##Scaling", &linkedScale)
		if linkedScale {
			imgui.SliderFloat("scale XY##Scaling", &imageScaleX, 0.1, 4.0)
			imageScaleY = imageScaleX
		} else {
			imgui.SliderFloat("scale X##Scaling", &imageScaleX, 0.1, 4.0)
			imgui.SliderFloat("scale Y##Scaling", &imageScaleY, 0.1, 4.0)
		}
		imgui.SeparatorText("FileSystem URLS")
		if dynamicImage.GetState() == g.SsLoading {
			imgui.Text("Unavailable while loading image...")
		} else {
			imgui.Text("Loading URLS Works with file:/// scheme too.")
			imgui.Text("By default, root is executable working dir")
			imgui.Text("-> Try loading this in the url bar:")
			imgui.Text("file:///files/sonic.png ->")
			if imgui.Button("or CLICK HERE") {
				inputRootfs = "."
				dynamicImage.SetFSRoot(inputRootfs)
				dynamicImageUrl = "file:///files/sonic.png"
				err := dynamicImage.ResetState()
				if err == nil {
					dynamicImage.SetSurfaceFromURL(dynamicImageUrl, time.Second*5, false)
				}
				linkedScale = true
				imageScaleX = 0.356
				imageScaleY = 0.356
			}
			imgui.Separator()
			imgui.Text("Set rootFS to / for full filesystem access")
			rootfs := dynamicImage.GetFSRoot()
			imgui.Text(fmt.Sprintf("Current ROOTFS: %s", rootfs))
			g.InputText(&inputRootfs).Build()
			imgui.SameLine()
			if imgui.Button("SET rootfs") {
				dynamicImage.SetFSRoot(inputRootfs)
			}
		}
	})

}

func loop() {

	if !showWindow {
		wnd.SetShouldClose(true)
	}

	g.PushColorWindowBg(color.RGBA{30, 30, 30, 255})
	g.Window("Async Images").IsOpen(&showWindow).Pos(10, 30).Size(1280, 720).Flags(g.WindowFlagsNoResize).Layout(
		HeaderWidget(),
		g.Separator(),
		g.Row(
			g.Child().Size(400, 625).Layout(
				ShouldReturnPanel(),
			),
			g.Child().Size(-1, 625).Layout(
				ShouldReturnImage(),
			)),
		g.Separator(),
		FooterWidget(footerLabel),
	)
	g.PopStyleColor()
}

func noOSDecoratedWindowsConfig() g.MasterWindowFlags {
	imgui.CreateContext()
	implot.PlotCreateContext()
	imnodes.ImNodesCreateContext()
	io := imgui.CurrentIO()
	io.SetConfigViewportsNoAutoMerge(true)
	io.SetConfigViewportsNoDefaultParent(true)
	io.SetConfigWindowsMoveFromTitleBarOnly(true)
	return g.MasterWindowFlagsHidden | g.MasterWindowFlagsTransparent | g.MasterWindowFlagsFrameless
}

func initDynamicImage() error {

	// SurfaceURL works from files:// too !
	// Note : the "root" of the scheme is willingly the Executable / working directory
	if err := fallback.SetSurfaceFromURL("file:///files/fallback.png", time.Second*5, false); err != nil {
		return err
	}

	dynamicImage.OnReset(func() {
		log.Println("DynamicImage was reset !")
	}).OnLoading(func() {
		log.Println("DynamicImage Started Loading a new surface...")
	}).OnFailure(func(e error) {
		if !errors.Is(e, g.ErrNeedReset) {
			footerLabel = fmt.Sprintf("DynamicImage failed loading with error: %v", e)
			log.Printf("DynamicImage failed loading with error: %+v\n", e)
		}
	}).OnSuccess(func() {
		footerLabel = "DynamicImage has sucessfully loaded new surface !"
		log.Println("DynamicImage has sucessfully loaded new surface !")
	})

	return nil
}

func main() {

	// This prepare creating a fully imgui window with no native decoration.
	// Flags are to be used with NewMasterWindow.
	// Should NOT use SingleLayoutWindow !
	mwFlags := noOSDecoratedWindowsConfig()
	if err := initDynamicImage(); err != nil {
		log.Fatalf("Error in DynamicImage initialization: %v", err)
	}

	wnd = g.NewMasterWindow("Load Image", 1280, 720, mwFlags)
	wnd.SetTargetFPS(60)
	wnd.Run(loop)
}
