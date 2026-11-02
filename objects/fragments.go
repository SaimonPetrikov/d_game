package objects

// import (
// 	"container/list"
// 	_ "image/png"
// 	"math"
// 	"math/rand"

// 	"github.com/hajimehoshi/ebiten/v2"
// )

// //todo interface
// type sprite struct {
// 	speed int
// 	count    int
// 	maxCount int
// 	dir      float64

// 	img   *ebiten.Image
// 	scale float64
// 	angle float64
// 	alpha float32
// 	ox float64
// 	oy float64
// }

// func (s *sprite) update() {
// 	// s.speed = s.speed + 2
// 	if s.count == 0 {
// 		return
// 	}
// 	s.count = s.count - 1
// }

// func (s *sprite) draw(screen *ebiten.Image) {
// 	if s.count == 0 {
// 		return
// 	}

// 	x := math.Cos(s.dir) * float64(s.maxCount-(s.count)) + float64(s.speed)
// 	y := math.Sin(s.dir) * float64(s.maxCount-(s.count)) + float64(s.speed)

// 	op := &ebiten.DrawImageOptions{}

// 	sx, sy := s.img.Bounds().Dx(), s.img.Bounds().Dy()
// 	op.GeoM.Translate(-float64(sx)/2, -float64(sy)/2)
// 	op.GeoM.Rotate(s.angle)
// 	op.GeoM.Scale(s.scale, s.scale)
// 	op.GeoM.Translate(x, y)
// 	op.GeoM.Translate(s.ox, s.oy)
// 	op.GeoM.Translate(sceneState.Cam.X, sceneState.Cam.Y)

// 	rate := float32(s.count) / float32(s.maxCount)
// 	var alpha float32
// 	if rate < 0.2 {
// 		alpha = rate / 0.2
// 	} else if rate > 0.8 {
// 		alpha = (1 - rate) / 0.2
// 	} else {
// 		alpha = 1
// 	}
// 	alpha *= s.alpha
// 	op.ColorScale.ScaleAlpha(alpha)

// 	screen.DrawImage(s.img, op)
// }

// func newSprite(img *ebiten.Image, ox, oy, angle float64) *sprite {
// 	c := rand.Intn(10) + 10
// 	dir := rand.Float64() - rand.Float64() + angle
// 	a := rand.Float64() - rand.Float64() + angle
// 	s := rand.Float64()*0.1 + 0.4
// 	return &sprite{
// 		img: img,

// 		maxCount: c,
// 		count:    c,
// 		dir:      dir,

// 		angle: a,
// 		scale: s,
// 		alpha: 0.8,
// 		ox: ox,
// 		oy: oy,
// 	}
// }

// type Fragments struct {
// 	Sprites *list.List
// 	FragmentImage *ebiten.Image
// 	Count int
// 	Ox float64
// 	Oy float64
// 	Angle float64
// }


// func NewFragments(img *ebiten.Image, count int, ox, oy, angle float64) *Fragments {
// 	return &Fragments{
// 		FragmentImage: img,
// 		Count: count,
// 		Ox: ox,
// 		Oy: oy,
// 		Angle: angle,
// 	}
// }

// func (f *Fragments) Init(s *scene) {
// 	s.AddGraphics(f)
// }

// func (g *Fragments) Draw(screen *ebiten.Image) {
// 	if g.Sprites != nil {
// 		for e := g.Sprites.Front(); e != nil; e = e.Next() {
// 			s := e.Value.(*sprite)
// 			s.draw(screen)
// 		}
// 	}
// }

// func (g *Fragments) Update(delta float64) {
// 	if g.Sprites == nil {
// 		g.Sprites = list.New()
// 	}

// 	if g.Sprites.Len() < g.Count {
// 		g.Sprites.PushBack(newSprite(g.FragmentImage, g.Ox, g.Oy, g.Angle))
// 	}

// 	for e := g.Sprites.Front(); e != nil; e = e.Next() {
// 		s := e.Value.(*sprite)
// 		s.update()
// 	}
// }

// func (g *Fragments) IsDisposed() bool {
// 	return false
// }