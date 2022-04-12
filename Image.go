package UI

import (
	"github.com/hajimehoshi/ebiten/v2"
  sort "github.com/illua1/go-helpful/Sort"
  "fmt"
)

type RenderingImage struct {
  Image *ebiten.Image
}

func NewImage(img *ebiten.Image)RenderingImage{
  return RenderingImage{img}
}

func(ri RenderingImage)Render(context ScreenContext) error {
  fmt.Println("-v2")
  if ri.Image != nil {
    x, y := ri.Image.Size()
    context.Op.GeoM = ebiten.GeoM{}
    context.Op.GeoM.Scale(float64(context.SelfRectangle.Dx())/float64(x), float64(context.SelfRectangle.Dy())/float64(y))
    context.Op.GeoM.Translate(
      float64(sort.MinF(context.SelfRectangle.Min.X,context.SelfRectangle.Max.X)),
      float64(sort.MinF(context.SelfRectangle.Min.Y,context.SelfRectangle.Max.Y)),
    )
    context.Screen.DrawImage(ri.Image, context.Op)
    return nil
  }else{
    return NewRenderError("RenderingImage: Nil Image")
  }
}