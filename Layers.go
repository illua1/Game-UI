package UI

import (
  "fmt"
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
        return fmt.Errorf("RenderingLayer->",err)
      }
    }else{
      return fmt.Errorf("RenderingLayer: Nil Layer: ", i)
    }
  }
  return nil
}