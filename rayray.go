package main

// TODO
// Duplicate in fix point and
//
import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
	"time"
)

type Vec struct {
	x, y, z float64
}

func (a Vec) Subtract(b Vec) Vec {
	return Vec{a.x - b.x, a.y - b.y, a.z - b.z}
}

func (a Vec) Dot(b Vec) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

type Ray struct {
	o, d Vec
}

type Sphere struct {
	center Vec
	radius float64
}

func (s Sphere) Intersect(ray Ray, t *float64) bool {
	oc := ray.o.Subtract(s.center)
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

type Color struct {
	r, g, b, a float64
}

func createFrame() {
	w := 800
	h := 500
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	// pixels := [500][800]Color{}
	// white := color.White
	sphere := Sphere{Vec{float64(w) / 2, float64(h) / 2, 50}, 20.0}
	tm := time.Now()
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {

			ray := Ray{Vec{float64(x), float64(y), 0}, Vec{0, 0, 1}}
			t := 20000.0
			if sphere.Intersect(ray, &t) {
				// Color the pixel
				// pixels[y][x] = white
				img.Set(x, y, color.Black)
			}
		}
	}
	fmt.Println(time.Since(tm))
	file, _ := os.Create("out.png")
	png.Encode(file, img)
}

func main() {
	// fmt.Println(Vec{-6, 8, 0}.Dot(Vec{5, 12, 0}))
	createFrame()
}
