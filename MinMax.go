package UI

import (
  "image"
  sort "github.com/illua1/go-helpful/Sort"
)

type RenderingMaxSize struct{
  InMax Rendering
}

func NewMaxSize(in Rendering)Rendering{
  return RenderingMaxSize{in}
}

func(rms RenderingMaxSize)Render(context ScreenContext)error{
  if rms.InMax != nil {
    min := sort.MaxF(context.DomainRectangle.Dx(), context.DomainRectangle.Dy())/2
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
      return RenderErrorLocation(err, "RenderingMaxSize")
    }else{
      return nil
    }
  }
  return NewRenderError("RenderingMaxSize:  Nil InMax")
}