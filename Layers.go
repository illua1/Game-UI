package UI

import (
  "fmt"
  "errors"
)

type RenderingLayer struct{
  Layers []Rendering
}

func NewLayer(in ...Rendering)Rendering{
  return RenderingLayer{in}
}

func(rl RenderingLayer)Render(context ScreenContext)error{
  for i := range rl.Layers {
    if rl.Layers[i] != nil {
      err := rl.Layers[i].Render(context)
      if err != nil {
        return errors.New(fmt.Sprint("RenderingLayer->",err))
      }
    }else{
      return errors.New(fmt.Sprint("RenderingLayer: Nil Layer: ", i))
    }
  }
  return nil
}