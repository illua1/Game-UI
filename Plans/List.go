package UI

import (
  "errors"
  "image"
	"github.com/hajimehoshi/ebiten/v2"
  sort "github.com/illua1/go-helpful/Sort"
)

type Rendering interface {
  Render(context ScreenContext) error
  Size(context ScreenContext)(int ,int, error)
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








type RenderingImage struct {
  body *ebiten.Image
}

func(r RenderingImage)Render(context ScreenContext) error {
  if r.body != nil {
    x, y := r.body.Size()
    context.Op.GeoM = ebiten.GeoM{}
    context.Op.GeoM.Scale(float64(context.SelfRectangle.Dx())/float64(x), float64(context.SelfRectangle.Dy())/float64(y))
    context.Op.GeoM.Translate(
      float64(sort.MinF(context.SelfRectangle.Min.X,context.SelfRectangle.Max.X)),
      float64(sort.MinF(context.SelfRectangle.Min.Y,context.SelfRectangle.Max.Y)),
    )
    context.Screen.DrawImage(r.body, context.Op)
    return nil
  }else{
    return errors.New("Nil Image")
  }
}

func(r RenderingImage)Size(context ScreenContext) (int, int, error) {
  if r.body != nil {
    x, y := r.body.Size()
    return x, y, nil
  }
  return 0,0, errors.New("RenderingImage: Size: Nil body")
}

func NewImage(img *ebiten.Image)Rendering{
  return RenderingImage{img}
}









type RenderingFullSize struct{
  body Rendering
}

func NewFullSize(in Rendering)Rendering{
  return RenderingFullSize{in}
}

func(rfs RenderingFullSize)Render(context ScreenContext)error{
  if rfs.body != nil {
    return errors.New("RenderingFullSize: Render: Nil body")
  }
  context.SelfRectangle = context.DomainRectangle
  return rfs.body.Render(context)
  
}

func(rfs RenderingFullSize)Size(context ScreenContext)(int, int, error){
  if rfs.body != nil {
    x, y, err := rfs.body.Size(context)
    return x, y, err
  }
  return 0, 0, errors.New("RenderingFullSize: Size: Nil body")
}









type RenderingAtMaxSize struct{
  body Rendering
}

func NewMaxSize(in Rendering)Rendering{
  return RenderingAtMaxSize{in}
}

func(rfs RenderingAtMaxSize)Render(context ScreenContext)error{
  if rfs.body != nil {
    max := sort.MaxF(context.DomainRectangle.Dx(), context.DomainRectangle.Dy())/2
    centre := context.SelfRectangle.Min.Add(context.SelfRectangle.Max).Div(2)
    context.SelfRectangle = image.Rectangle{
        centre.Sub(
          image.Point{max, max},
        ),
        centre.Add(
          image.Point{max, max},
        ),
      }
    return rfs.body.Render(context)
  }
  return errors.New("RenderingAtMaxSize: Render: Nil body")
}

func(rfs RenderingAtMaxSize)Size(context ScreenContext)(int, int, error){
  if rfs.body != nil {
    x, y, err := rfs.body.Size(context)
    return x, y, err
  }
  return 0, 0, errors.New("RenderingFullSize: Size: Nil body")
}









type RenderingAtMinSize struct{
  body Rendering
}

func NewMinSize(in Rendering)Rendering{
  return RenderingAtMinSize{in}
}

func(rfs RenderingAtMinSize)Render(context ScreenContext)error{
  if rfs.body != nil {
    min := sort.MinF(context.DomainRectangle.Dx(), context.DomainRectangle.Dy())/2
    centre := context.SelfRectangle.Min.Add(context.SelfRectangle.Max).Div(2)
    context.SelfRectangle = image.Rectangle{
        centre.Sub(
          image.Point{min, min},
        ),
        centre.Add(
          image.Point{min, min},
        ),
      }
    return rfs.body.Render(context)
  }
  return errors.New("RenderingAtMinSize: Render: Nil body")
}

func(rfs RenderingAtMinSize)Size(context ScreenContext)(int, int, error){
  if rfs.body != nil {
    x, y, err := rfs.body.Size(context)
    return x, y, err
  }
  return 0, 0, errors.New("RenderingFullSize: Size: Nil body")
}







type RenderingToCentre struct{
  body Rendering
}

func NewToCentre(in Rendering)Rendering{
  return RenderingToCentre{in}
}

func(rfs RenderingToCentre)Render(context ScreenContext)error{
  if rfs.body != nil {
    x, y := context.SelfRectangle.Dx()/2, context.SelfRectangle.Dy()/2
    centre := context.DomainRectangle.Min.Add(context.DomainRectangle.Max).Div(2)
    context.SelfRectangle = image.Rectangle{
        centre.Sub(
          image.Point{x, y},
        ),
        centre.Add(
          image.Point{x, y},
        ),
      }
    return rfs.body.Render(context)
  }
  return errors.New("RenderingToCentre: Render: Nil body")
}

func(rfs RenderingToCentre)Size(context ScreenContext)(int, int, error){
  if rfs.body != nil {
    x, y, err := rfs.body.Size(context)
    return x, y, err
  }
  return 0, 0, errors.New("RenderingToCentre: Size: Nil body")
}







type RenderingToFrame struct{
  a, b Rendering
  os bool
  factor float64
}

func NewFrame(first, seccond Rendering, os bool, factor float64)Rendering{
  return RenderingToFrame{first, seccond, os, factor}
}

func(rfs RenderingToFrame)Render(context ScreenContext)(err error){
  orig_DomainRectangle := context.DomainRectangle
  if rfs.a != nil {
    if rfs.os{
      min_x := sort.MinF(
        orig_DomainRectangle.Min.X, 
        orig_DomainRectangle.Max.X,
      )
      centre_x := min_x + int(float64(orig_DomainRectangle.Dx())*rfs.factor)
      min_y := sort.MinF(
        orig_DomainRectangle.Min.Y, 
        orig_DomainRectangle.Max.Y,
      )
      max_y := sort.MaxF(
        orig_DomainRectangle.Min.Y, 
        orig_DomainRectangle.Max.Y,
      )
      context.DomainRectangle = image.Rect(min_x, min_y, centre_x, max_y)
      context.SelfRectangle = context.DomainRectangle
    }else{
      
    }
    err = rfs.a.Render(context)
  }
  if err != nil {
    return err
  }
  if rfs.b != nil {
    if rfs.os{
      max_x := sort.MaxF(
        orig_DomainRectangle.Min.X, 
        orig_DomainRectangle.Max.X,
      )
      centre_x := max_x - int(float64(orig_DomainRectangle.Dx())*(1.-rfs.factor))
      min_y := sort.MinF(
        orig_DomainRectangle.Min.Y, 
        orig_DomainRectangle.Max.Y,
      )
      max_y := sort.MaxF(
        orig_DomainRectangle.Min.Y, 
        orig_DomainRectangle.Max.Y,
      )
      context.DomainRectangle = image.Rect(max_x, min_y, centre_x, max_y)
      context.SelfRectangle = context.DomainRectangle
    }else{
      
    }
    err = rfs.b.Render(context)
  }
  return err
}

func(rfs RenderingToFrame)Size(context ScreenContext)(int, int, error){
  return context.DomainRectangle.Dx(), context.DomainRectangle.Dy(), nil
}







type RenderingToLayer struct{
  body []Rendering
}

func NewToLayer(in ...Rendering)Rendering{
  return RenderingToLayer{in}
}

func(rfs RenderingToLayer)Render(context ScreenContext)error{
  for i := range rfs.body {
    err := rfs.body[i].Render(context)
    if err != nil {
      return err
    }
  }
  return nil
}

func(rfs RenderingToLayer)Size(context ScreenContext)(int, int, error){
  var x_max, y_max int
  for i := range rfs.body {
    x, y, err := rfs.body[i].Size(context)
    if err != nil {
      return 0, 0, err
    }
    x_max = sort.MaxF(x_max, x)
    y_max = sort.MaxF(y_max, y)
  }
  return x_max, y_max, nil
}





type RenderingNormalize struct{
  body Rendering
}

func NewNormalize(in Rendering)Rendering{
  return RenderingNormalize{in}
}

func(rfs RenderingNormalize)Render(context ScreenContext)error{
  if rfs.body != nil {
    x, y, err := rfs.body.Size(context)
    if err != nil {
      return err
    }
    centre := context.SelfRectangle.Min.Add(context.SelfRectangle.Max).Div(2)
    
    factor := float64(x) / float64(y)
    
    if factor > 1. {
      x_ := int(float64(context.SelfRectangle.Dx())*factor/2)
      context.SelfRectangle = image.Rect(centre.X-x_, context.SelfRectangle.Min.Y, centre.X+x_, context.SelfRectangle.Max.Y)
      return rfs.body.Render(context)
    }else{
      y_ := int(float64(context.SelfRectangle.Dy())/factor/2)
      context.SelfRectangle = image.Rect(context.SelfRectangle.Min.X, centre.Y-y_, context.SelfRectangle.Max.X, centre.Y+y_)
      return rfs.body.Render(context)
    }
  }
  return errors.New("RenderingNormalize: Render: Nil body")
}

func(rfs RenderingNormalize)Size(context ScreenContext)(int, int, error){
  if rfs.body != nil {
    x, y, err := rfs.body.Size(context)
    if err != nil {
      return 0, 0, err
    }
    centre := context.SelfRectangle.Min.Add(context.SelfRectangle.Max).Div(2)
    
    factor := float64(x) / float64(y)
    
    if factor > 1. {
      x_ := int(float64(context.SelfRectangle.Dx())*factor/2)
      return centre.X-x_+centre.X+x_, context.SelfRectangle.Min.Y+context.SelfRectangle.Max.Y, nil
    }else{
      y_ := int(float64(context.SelfRectangle.Dy())/factor/2)
      return context.SelfRectangle.Min.X+context.SelfRectangle.Max.X, centre.Y-y_+centre.Y+y_, nil
    }
  }
  return 0, 0, errors.New("RenderingNormalize: Size: Nil body")
}







type RenderingToBound struct{
  body Rendering
  Size_ image.Point
}

func NewBound(in Rendering, x, y int)Rendering{
  return RenderingToBound{in, image.Point{x, y}}
}

func(rfs RenderingToBound)Render(context ScreenContext)error{
  if rfs.body != nil {
    context.SelfRectangle = image.Rectangle{
      image.Point{
        0,
        0,
      },
      image.Point{
        sort.MinF(
          rfs.Size_.X,
          context.SelfRectangle.Dx(),
        ),
        sort.MinF(
          rfs.Size_.Y,
          context.SelfRectangle.Dy(),
        ),
      },
    }
    return rfs.body.Render(context)
  }
  return errors.New("RenderingToBound: Render: Nil body")
}

func(rfs RenderingToBound)Size(context ScreenContext)(int, int, error){
  return rfs.Size_.X, rfs.Size_.Y , nil
}




type toBorder int8

const(
  ToCentre toBorder = 0
  ToLeftBorder toBorder = 1
  ToRightBorder toBorder = 2
  ToTopBorder toBorder = 1
  ToBottomBorder toBorder = 2
)


type RenderingToBorder struct{
  body Rendering
  lr, tb toBorder
}

func NewToBorder(in Rendering, lr, tb toBorder)Rendering{
  return RenderingToBorder{in, lr, tb}
}

func(rfs RenderingToBorder)Render(context ScreenContext)error{
  var x,y int = 0, 0
  switch rfs.lr {
    case ToLeftBorder : {
      x = sort.MinF(context.DomainRectangle.Min.X, context.DomainRectangle.Max.X) - sort.MinF(context.SelfRectangle.Min.X, context.SelfRectangle.Max.X)
    }
    case ToRightBorder : {
      x = sort.MaxF(context.DomainRectangle.Min.X, context.DomainRectangle.Max.X) - sort.MaxF(context.SelfRectangle.Min.X, context.SelfRectangle.Max.X)
    }
    default : {
      x = sort.MinF(context.SelfRectangle.Min.X, context.SelfRectangle.Max.X) - sort.MinF(context.DomainRectangle.Min.X, context.DomainRectangle.Max.X)
      x += sort.MaxF(context.SelfRectangle.Min.X, context.SelfRectangle.Max.X) - sort.MaxF(context.DomainRectangle.Min.X, context.DomainRectangle.Max.X)
      x /= 2
    }
  }
  switch rfs.tb {
    case ToTopBorder : {
      y = sort.MinF(context.DomainRectangle.Min.Y, context.DomainRectangle.Max.Y) - sort.MinF(context.SelfRectangle.Min.Y, context.SelfRectangle.Max.Y)}
    case ToBottomBorder : {
      y = sort.MaxF(context.SelfRectangle.Min.Y, context.SelfRectangle.Max.Y) - sort.MaxF(context.DomainRectangle.Min.Y, context.DomainRectangle.Max.Y)
    }
    default : {
      y = sort.MinF(context.SelfRectangle.Min.Y, context.SelfRectangle.Max.Y) - sort.MinF(context.DomainRectangle.Min.Y, context.DomainRectangle.Max.Y)
      y += sort.MaxF(context.SelfRectangle.Min.Y, context.SelfRectangle.Max.Y) - sort.MaxF(context.DomainRectangle.Min.Y, context.DomainRectangle.Max.Y)
      y /= 2
    }
  }
  context.SelfRectangle = context.SelfRectangle.Add(image.Point{x, y})
  return rfs.body.Render(context)
  
}

func(rfs RenderingToBorder)Size(context ScreenContext)(int, int, error){
  if rfs.body != nil {
    x, y, err := rfs.body.Size(context)
    return x, y, err
  }
  return 0, 0, errors.New("RenderingToBorder: Size: Nil body")
}