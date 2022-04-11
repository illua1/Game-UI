package UI

import (
  "fmt"
  "errors"
	"github.com/hajimehoshi/ebiten/v2"
  sort "github.com/illua1/go-helpful/Sort"
)

type RenderingImage struct {
  image *ebiten.Image
}

func(ri RenderingImage)Render(context ScreenContext) error {
  fmt.Println("V-1")
  if ri.image != nil {
    x, y := ri.image.Size()
    context.Op.GeoM = ebiten.GeoM{}
    context.Op.GeoM.Scale(float64(context.SelfRectangle.Dx())/float64(x), float64(context.SelfRectangle.Dy())/float64(y))
    context.Op.GeoM.Translate(
      float64(sort.MinF(context.SelfRectangle.Min.X,context.SelfRectangle.Max.X)),
      float64(sort.MinF(context.SelfRectangle.Min.Y,context.SelfRectangle.Max.Y)),
    )
    context.Screen.DrawImage(ri.image, context.Op)
    return nil
  }else{
    return errors.New(fmt.Sprint("RenderingImage: Nil Image"))
  }
}

func NewImage(img *ebiten.Image)Rendering{
  return RenderingImage{img}
}