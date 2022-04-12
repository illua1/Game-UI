package UI

import (
  "image"
  sort "github.com/illua1/go-helpful/Sort"
)

type toBorder uint8

const(
  ToCentre toBorder = 0
  ToLeftBorder toBorder = 1
  ToRightBorder toBorder = 2
  ToTopBorder toBorder = 1
  ToBottomBorder toBorder = 2
)


type RenderingToBorder struct{
  ToBorder Rendering
  LR_flag, TB_flag toBorder
}

func NewToBorder(in Rendering, lr, tb toBorder)RenderingToBorder{
  return RenderingToBorder{in, lr, tb}
}

func(rtb RenderingToBorder)Render(context ScreenContext)error{
  if rtb.ToBorder != nil {
    var x,y int = 0, 0
    switch rtb.LR_flag {
      case ToLeftBorder : {
        x = sort.MinF(context.DomainRectangle.Min.X, context.DomainRectangle.Max.X) - sort.MinF(context.SelfRectangle.Min.X, context.SelfRectangle.Max.X)
      }
      case ToRightBorder : {
        x = sort.MaxF(context.DomainRectangle.Min.X, context.DomainRectangle.Max.X) - sort.MaxF(context.SelfRectangle.Min.X, context.SelfRectangle.Max.X)
      }
      default : {
        x = sort.MinF(context.SelfRectangle.Min.X, context.SelfRectangle.Max.X) - sort.MinF(context.DomainRectangle.Min.X, context.DomainRectangle.Max.X)
        x += sort.MaxF(context.SelfRectangle.Min.X, context.SelfRectangle.Max.X) - sort.MaxF(context.DomainRectangle.Min.X, context.DomainRectangle.Max.X)
        x /= 2
      }
    }
    switch rtb.TB_flag {
      case ToTopBorder : {
        y = sort.MinF(context.DomainRectangle.Min.Y, context.DomainRectangle.Max.Y) - sort.MinF(context.SelfRectangle.Min.Y, context.SelfRectangle.Max.Y)
      }
      case ToBottomBorder : {
        y = sort.MaxF(context.DomainRectangle.Min.Y, context.DomainRectangle.Max.Y) - sort.MaxF(context.SelfRectangle.Min.Y, context.SelfRectangle.Max.Y)
      }
      default : {
        y = sort.MinF(context.SelfRectangle.Min.Y, context.SelfRectangle.Max.Y) - sort.MinF(context.DomainRectangle.Min.Y, context.DomainRectangle.Max.Y)
        y += sort.MaxF(context.SelfRectangle.Min.Y, context.SelfRectangle.Max.Y) - sort.MaxF(context.DomainRectangle.Min.Y, context.DomainRectangle.Max.Y)
        y /= 2
      }
    }
    context.SelfRectangle = context.SelfRectangle.Add(image.Point{x, y})
    if err := rtb.ToBorder.Render(context); err != nil {
      return RenderErrorLocation(err, "RenderingToBorder")
    }else{
      return nil
    }
  }
  return NewRenderError("RenderingToBorder:  Nil ToBorder")
}