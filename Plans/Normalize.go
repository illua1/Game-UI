package UI

import (
  "errors"
  "image"
	"github.com/hajimehoshi/ebiten/v2"
  sort "github.com/illua1/go-helpful/Sort"
)

type RenderingNormalize struct{
  body Rendering
}

func NewNormalize(in Rendering)Rendering{
  return RenderingNormalize{in}
}

func(rfs RenderingNormalize)Render(context ScreenContext)error{
  if rfs.body != nil {
    x, y, err := rfs.body.Size(context)
    if err != nil {
      return err
    }
    centre := context.SelfRectangle.Min.Add(context.SelfRectangle.Max).Div(2)
    
    factor := float64(x) / float64(y)
    
    if factor > 1. {
      x_ := int(float64(context.SelfRectangle.Dx())*factor/2)
      context.SelfRectangle = image.Rect(centre.X-x_, context.SelfRectangle.Min.Y, centre.X+x_, context.SelfRectangle.Max.Y)
      return rfs.body.Render(context)
    }else{
      y_ := int(float64(context.SelfRectangle.Dy())/factor/2)
      context.SelfRectangle = image.Rect(context.SelfRectangle.Min.X, centre.Y-y_, context.SelfRectangle.Max.X, centre.Y+y_)
      return rfs.body.Render(context)
    }
  }
  return errors.New("RenderingNormalize: Render: Nil body")
}