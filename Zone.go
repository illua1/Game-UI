package UI

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type RenderingZone struct{
  InZone Rendering
}

func NewZone(in Rendering)RenderingZone{
  return RenderingZone{in}
}

func(rz RenderingZone)Render(context ScreenContext)error{
  if rz.InZone != nil {
    context.Screen = context.Screen.SubImage(context.SelfRectangle).(*ebiten.Image)
    context.DomainRectangle = context.SelfRectangle
    if err := rz.InZone.Render(context); err != nil {
      return RenderErrorLocation(err, "RenderingZone")
    }else{
      return nil
    }
  }else{
    return NewRenderError("Nil InZone")
  }
  return nil
}