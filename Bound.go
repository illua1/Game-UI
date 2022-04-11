package UI

import (
  "image"
)

type RenderingBound struct{
  InBound Rendering
  Size image.Point
}

func NewBound(in Rendering, x, y int)Rendering{
  return RenderingBound{in, image.Point{x, y}}
}

func(rb RenderingBound)Render(context ScreenContext)error{
  
  if rb.InBound != nil {
    centre := Rect_Centre(context.SelfRectangle)
    Size_ := rb.Size.Div(2)
    context.SelfRectangle = image.Rectangle{
      centre.Sub(Size_),
      centre.Add(Size_),
    }
    if err := rb.InBound.Render(context); err != nil {
      RenderErrorLocation(err, "RenderingBound")
    }else{
      return nil
    }
  }
  return NewRenderError("RenderingBound: Nil Bound")
}