package main

import (
	_ "image/jpeg"
	_ "image/png"

	g "github.com/AllenDang/giu"
)

var (
	spriteTexture *g.Texture
	largeTexture  *g.Texture
)

func loop() {
	g.SingleWindow("load image", g.Layout{
		g.Group(g.Layout{
			g.Label("15x20 pixel image"),
			g.Line(
				g.Group(g.Layout{
					g.Label("50%"),
					g.Image(spriteTexture, 8, 10),
				}),
				g.Group(g.Layout{
					g.Label("100%"),
					g.Image(spriteTexture, 15, 20),
				}),
				g.Group(g.Layout{
					g.Label("800%"),
					g.Image(spriteTexture, 120, 160),
				}),
			),
		}),
		g.Group(g.Layout{
			g.Label("215x140 image"),
			g.Line(
				g.Group(g.Layout{
					g.Label("50%"),
					g.Image(largeTexture, 215/2, 140/2),
				}),
				g.Group(g.Layout{
					g.Label("100%"),
					g.Image(largeTexture, 215, 140),
				}),
				g.Group(g.Layout{
					g.Label("200%"),
					g.Image(largeTexture, 215*2, 140*2),
				}),
			),
		}),
		g.Line(
			g.Button("Minify Filter Nearest", func() {
				_ = g.Context.GetRenderer().SetTextureMinFilter(g.TextureFilterNearest)
			}),
			g.Button("Minify Filter Linear", func() {
				_ = g.Context.GetRenderer().SetTextureMinFilter(g.TextureFilterLinear)
			}),
			/*g.Button("Nearest Mipmap Nearest", func() {
				g.Context.GetRenderer().SetTextureMinFilter(g.TextureFilterNearestMipmapNearest)
			}),
			g.Button("Linear Mipmap Nearest", func() {
				g.Context.GetRenderer().SetTextureMinFilter(g.TextureFilterLinearMipmapNearest)
			}),
			g.Button("Nearest Mipmap Linear", func() {
				g.Context.GetRenderer().SetTextureMinFilter(g.TextureFilterNearestMipmapLinear)
			}),
			g.Button("Linear Mipmap Linear", func() {
				g.Context.GetRenderer().SetTextureMinFilter(g.TextureFilterLinearMipmapLinear)
			}),*/
		),
		g.Line(
			g.Button("Magnify Filter Nearest", func() {
				_ = g.Context.GetRenderer().SetTextureMagFilter(g.TextureFilterNearest)
			}),
			g.Button("Magnify Filter Linear", func() {
				_ = g.Context.GetRenderer().SetTextureMagFilter(g.TextureFilterLinear)
			}),
		),
	})
}

func main() {
	wnd := g.NewMasterWindow("Texture Filtering", 800, 600, g.MasterWindowFlagsNotResizable, nil)

	spriteImg, _ := g.LoadImage("gopher-sprite.png")
	largeImg, _ := g.LoadImage("gopher.png")
	go func() {
		spriteTexture, _ = g.NewTextureFromRgba(spriteImg)
		largeTexture, _ = g.NewTextureFromRgba(largeImg)
	}()

	wnd.Main(loop)
}
