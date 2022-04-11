package UI

import (
  "image"
  sort "github.com/illua1/go-helpful/Sort"
)

type RenderingMinSize struct{
  InMax Rendering
}

func NewMinSize(in Rendering)Rendering{
  return RenderingMinSize{in}
}

func(rms RenderingMinSize)Render(context ScreenContext)error{
  if rms.InMax != nil {
    min := sort.MinF(context.DomainRectangle.Dx(), context.DomainRectangle.Dy())/2
    centre := Rect_Centre(context.SelfRectangle)
    context.SelfRectangle = image.Rectangle{
        centre.Sub(
          image.Point{min, min},
        ),
        centre.Add(
          image.Point{min, min},
        ),
      }
    if err := rms.InMax.Render(context); err != nil {
      return RenderErrorLocation(err, "RenderingMinSize")
    }else{
      return nil
    }
  }
  return NewRenderError("RenderingMinSize:  Nil InMax")
}