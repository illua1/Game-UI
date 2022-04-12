package UI

import (
  "image"
  "image/color"
  "golang.org/x/image/font"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type RenderingText struct{
  Lines []TextToRender
}

type TextToRender struct{
  Location image.Point
  Font font.Face
  Str string
  Color color.Color
}

func NewText(lines ...TextToRender)RenderingText{
  return RenderingText{lines}
}

func(rt RenderingText)Render(context ScreenContext)error{
  for i := range rt.Lines{
    f := context.SelfRectangle.Min.Add(rt.Lines[i].Location)
    text.Draw(context.Screen, rt.Lines[i].Str, rt.Lines[i].Font, f.X, f.Y, rt.Lines[i].Color)
  }
  return nil
}