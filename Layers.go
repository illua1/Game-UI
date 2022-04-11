package UI

import (
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
      if err := rl.Layers[i].Render(context); err != nil {
        return RenderErrorLocation(err, "RenderingLayer: ", i)
      }
    }else{
      return NewRenderError("RenderingLayer: Nil Layer: ", i)
    }
  }
  return nil
}