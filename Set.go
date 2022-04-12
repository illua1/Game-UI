package UI

import (
  "image"
)
type RenderingSet struct{
  InSet []Rendering
  Direction_flag bool
}

func NewSet(flag bool, in ...Rendering)RenderingSet{
  return RenderingSet{in, flag}
}

func(rs RenderingSet)Render(context ScreenContext)error{
  for i := range rs.InSet{
    if rs.InSet[i] != nil {
      if err := rs.InSet[i].Render(context); err != nil {
        return RenderErrorLocation(err, "RenderingSet: ", i)
      }else{
        return nil
      }
      if rs.Direction_flag {
        context.SelfRectangle = context.SelfRectangle.Add(image.Point{
          context.SelfRectangle.Dx(),
          0,
        })
      }else{
        context.SelfRectangle = context.SelfRectangle.Add(image.Point{
          0,
          context.SelfRectangle.Dy(),
        })
      }
    }else{
      return NewRenderError("Nil InSet: ", i)
    }
  }
  return nil
}