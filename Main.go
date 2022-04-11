package UI

import (
  "fmt"
  "image"
	"github.com/hajimehoshi/ebiten/v2"
  sort "github.com/illua1/go-helpful/Sort"
  values "github.com/illua1/go-helpful"
)

type Rendering interface {
  Render(context ScreenContext) error
  //Size(context ScreenContext)(int ,int, error)
}

type ScreenContext struct {
  Op *ebiten.DrawImageOptions
  
  DomainRectangle image.Rectangle
  SelfRectangle image.Rectangle
  
  Screen *ebiten.Image
}

func NewScreenContext(op *ebiten.DrawImageOptions, rect image.Rectangle, screen *ebiten.Image)ScreenContext{
  return ScreenContext{
      op, 
      rect,
      rect, 
      screen,
    }
}

func Rect_Normalize(in image.Rectangle)image.Rectangle{
  return image.Rectangle{
    image.Point{
      sort.MinF(
        in.Min.X, in.Max.X,
      ),
      sort.MinF(
        in.Min.Y, in.Max.Y,
      ),
    },
    image.Point{
      sort.MaxF(
        in.Min.X, in.Max.X,
      ),
      sort.MaxF(
        in.Min.Y, in.Max.Y,
      ),
    },
  }
}

func Rect_Centre(in image.Rectangle)image.Point{
  return in.Min.Add(in.Max).Div(2)
}

func Point_Mull[T values.Values](a image.Point, b T)image.Point{
  return image.Point{
    int(T(a.X)*b),
    int(T(a.Y)*b),
  }
}

func Rect_Mull[T values.Values](a image.Rectangle, b T)image.Rectangle{
  return image.Rectangle{
    Point_Mull[T](a.Min, b),
    Point_Mull[T](a.Max, b),
  }
}

func Rect_Mull_Centre[T values.Values](a image.Rectangle, b T)image.Rectangle{
  centre := Rect_Centre(a)
  a_ := a.Sub(centre)
  return image.Rectangle{
    Point_Mull[T](a_.Min, b),
    Point_Mull[T](a_.Max, b),
  }.Add(centre)
}

func Lerp[T, F values.Values](a, b T, f F)T{
  return a+T(F(b-a)*f)
}

func Rect_Lerp[F values.Values](r image.Rectangle, x, y F)image.Point{
  return image.Point{
    Lerp[int, F](r.Min.X, r.Max.X, x),
    Lerp[int, F](r.Min.Y, r.Max.Y, y),
  }
}

type RenderError string

func NewRenderError(text ...interface{})*RenderError{
  var v = RenderError(fmt.Sprint(text))
  return &v
}

func RenderErrorLocation(err error, i ...interface{})*RenderError{
  var v = RenderError(fmt.Sprint(i) + " -> " + err.Error())
  return &v
}

func(re *RenderError)Error()string{
  return string(*re)
}
