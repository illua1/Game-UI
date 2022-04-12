package UI

import (
  "image/color"
  "github.com/hajimehoshi/ebiten/v2"
  value "github.com/illua1/go-helpful"
)

type RenderingFill struct{
  Color color.Color
}

func NewFill[V value.Values](r, g, b, a V)RenderingFill{
  return RenderingFill{
    color.RGBA{
      uint8(int8(r)),
      uint8(int8(g)),
      uint8(int8(b)),
      uint8(int8(a)),
    },
  }
}

func(rf RenderingFill)Render(context ScreenContext)error{
  //ebiten.NewImageFromImage(context.Screen.SubImage(context.SelfRectangle)).Fill(rf.Color)
  context.Screen.SubImage(context.SelfRectangle).(*ebiten.Image).Fill(rf.Color)
  return nil
}