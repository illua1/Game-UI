package UI

import (
  "image"
  "fmt"
)

type RenderingFrame struct{
  Frames [2]Rendering
  os bool
  factor float64
}

func NewFrame(first, seccond Rendering, os bool, factor float64)Rendering{
  return RenderingFrame{[2]Rendering{first, seccond}, os, factor}
}

func(rf RenderingFrame)Render(context ScreenContext)(err error){
  orig_DomainRectangle := context.DomainRectangle
  
  var (
    factors_x [2]float64
    factors_y [2]float64
    
    p1, p2 image.Point
  )
  if rf.os {
    factors_x = [2]float64{rf.factor, 1-rf.factor}
    factors_y = [2]float64{1.0, 0.0}
  }else{
    factors_x = [2]float64{1.0, 0.0}
    factors_y = [2]float64{rf.factor, 1-rf.factor}
  }
  
  for i := range rf.Frames {
    if rf.Frames[i] != nil {
      
      if i == 0 {
        p1 = orig_DomainRectangle.Min
        p2 = image.Point{
          Lerp(orig_DomainRectangle.Min.X, orig_DomainRectangle.Max.X, factors_x[i]),
          Lerp(orig_DomainRectangle.Min.Y, orig_DomainRectangle.Max.Y, factors_y[i]),
        }
      }else{
        p1 = image.Point{
          Lerp(orig_DomainRectangle.Min.X, orig_DomainRectangle.Max.X, factors_x[i]),
          Lerp(orig_DomainRectangle.Min.Y, orig_DomainRectangle.Max.Y, factors_y[i]),
        }
        p2 = orig_DomainRectangle.Max
      }
      
      context.DomainRectangle = image.Rectangle{p1,p2}
      
      err := rf.Frames[i].Render(context)
      
      if err != nil {
        return fmt.Errorf("NewFrame: ", i,"-> ", err)
      }
      
    }else{
      return fmt.Errorf("NewFrame: Nil Frame: ", i)
    }
  }
  
  return err
}