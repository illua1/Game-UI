package UI

import (
  "image"
  slise "github.com/illua1/go-helpful/Slise"
)

type RenderingSet struct{
  InSet []Rendering
  Direction_flag bool
  Offsets []int
}

func NewSet(flag bool, offset []int, in ...Rendering)RenderingSet{
  return RenderingSet{in, flag, offset}
}

func(rs RenderingSet)Render(context ScreenContext)error{
  for i := range rs.InSet{
    if rs.InSet[i] != nil {
      if err := rs.InSet[i].Render(context); err != nil {
        return RenderErrorLocation(err, "RenderingSet: ", i)
      }
      v, _ := slise.GetLast[int](rs.Offsets, i)
      if rs.Direction_flag {
        context.SelfRectangle = context.SelfRectangle.Add(image.Point{v, 0})
      } else {
        context.SelfRectangle = context.SelfRectangle.Add(image.Point{0, v})
      }
    }else{
      return NewRenderError("Nil InSet: ", i)
    }
  }
  return nil
}