package UI

import (
  "image"
  value "github.com/illua1/go-helpful"
)

type RenderingBorder struct {
  InBorder Rendering
  Border int
}

func NewBorder[V value.Values](in Rendering, border V)RenderingBorder{
  return RenderingBorder{in, int(border)}
}

func(rb RenderingBorder)Render(context ScreenContext) error {
  if rb.InBorder != nil {
    context.SelfRectangle = image.Rectangle{
      context.SelfRectangle.Min.Add(image.Point{rb.Border, rb.Border}),
      context.SelfRectangle.Max.Sub(image.Point{rb.Border, rb.Border}),
    }
    if err := rb.InBorder.Render(context); err != nil {
      return RenderErrorLocation(err, "RenderingBorder")
    }else{
      return nil
    }
  }else{
    return NewRenderError("Nil InBorder")
  }
}