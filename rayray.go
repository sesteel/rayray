package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/disintegration/imaging"
)

const (
	w = 1280
	h = 720
)

func createFrame() [w][h]Color {
	pixels := [w][h]Color{}
	spheres := [1000]Sphere{}
	light := Sphere{Vec{(w / 2) - 300, 1000, 10}, 1}

	// load spheres
	for i := 0; i < len(spheres); i++ {
		x := float64(rand.Intn(w))
		y := float64(rand.Intn(h))
		z := float64(rand.Intn(100))
		spheres[i] = Sphere{Vec{x, y, z}, (100 / z) * 5}
	}

	// This demonstrates the use of wait groups to
	tm := time.Now()
	sections := 1
	sw := w / sections
	wg := sync.WaitGroup{}
	for k := 0; k < 200; k++ {
		wg.Add(sections)
		for i := 0; i < sections; i++ {
			x := sw * i
			go iterate(&wg, &light, x, 0, x+sw, h, &spheres, &pixels)
		}
		wg.Wait()
	}
	fmt.Println(sections, (time.Since(tm) / 200))

	return pixels
}

func main() {
	pixels := createFrame()

	img := image.NewNRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, color.RGBA{uint8(pixels[x][y].r), uint8(pixels[x][y].g), uint8(pixels[x][y].b), 255})
		}
	}
	img = imaging.Blur(img, 1.2)
	p, _ := os.Create("out2.png")
	png.Encode(p, img)
}

type Vec struct {
	x, y, z float64
}

func (a Vec) Add(b Vec) Vec {
	return Vec{a.x + b.x, a.y + b.y, a.z + b.z}
}

func (a Vec) Sub(b Vec) Vec {
	return Vec{a.x - b.x, a.y - b.y, a.z - b.z}
}

func (a Vec) Mul(b float64) Vec {
	return Vec{a.x * b, a.y * b, a.z * b}
}

func (a Vec) Div(b float64) Vec {
	return Vec{a.x / b, a.y / b, a.z / b}
}

func (a Vec) Dot(b Vec) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func (a Vec) Normalize() Vec {
	mg := math.Sqrt(a.x*a.x + a.y*a.y + a.z*a.z)
	return Vec{a.x / mg, a.y / mg, a.z / mg}
}

type Ray struct {
	pos, d Vec
}

type Sphere struct {
	center Vec
	radius float64
}

func (s Sphere) Intersect(ray Ray, t *float64) bool {
	oc := ray.pos.Sub(s.center)
	b := 2 * oc.Dot(ray.d)
	c := oc.Dot(oc) - s.radius*s.radius
	disc := b*b - 4*c
	if disc < 0 {
		return false
	} else {
		disc = math.Sqrt(disc)
		t0 := -b - disc
		t1 := -b + disc
		if t0 < t1 {
			*t = t0
			return true
		}
		*t = t1
		return true
	}
}

func (s Sphere) getNormal(pt Vec) Vec {
	return (s.center.Sub(pt)).Div(s.radius)
}

type Color struct {
	r, g, b float64
}

func (a Color) Mul(d float64) Color {
	return Color{a.r * d, a.g * d, a.b * d}.Normalize()
}

func (a Color) Add(b Color) Color {
	c := Color{(a.r + b.r) / 2, (a.g + b.g) / 2, (a.b + b.b) / 2}
	return c.Normalize()
}

func (a Color) Normalize() Color {
	if a.r > 255 {
		a.r = 255
	}
	if a.g > 255 {
		a.g = 255
	}
	if a.b > 255 {
		a.b = 255
	}
	if a.r < 0 {
		a.r = 0
	}
	if a.g < 0 {
		a.g = 0
	}
	if a.b < 0 {
		a.b = 0
	}
	return a
}

func iterate(wg *sync.WaitGroup,
	globalLight *Sphere,
	ax, ay, bx, by int,
	spheres *[1000]Sphere,
	pixels *[w][h]Color) {

	defer wg.Done()
	white := Color{255, 255, 255}
	green := Color{0, 255, 0}
	for y := ay; y < by; y++ {
		yf := float64(y)
		for x := ax; x < bx; x++ {
			xf := float64(x)
			ray := Ray{Vec{xf, yf, 0}, Vec{0, 0, 1}}
			t := 20000.0
			for i := 0; i < 1000; i++ {
				sphere := spheres[i]
				if sphere.Intersect(ray, &t) {
					pt := ray.pos.Add(ray.d.Mul(t))
					L := (globalLight).center.Sub(pt)
					N := sphere.getNormal(pt)
					dt := L.Normalize().Dot(N.Normalize())

					pixels[x][y] = green.Add(white.Mul(dt)).Mul(1.5)
					break
				}
			}
		}
	}
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
