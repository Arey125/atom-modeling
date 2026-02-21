package field

import (
	"fmt"
	"math"
	"math/cmplx"
)

type Field struct {
	values [][]complex128
	prev   [][]complex128
	buf    [][]complex128
	config Config
}

const Dx = float64(0.1)
const Dt = float64(0.3)

const Dxc = complex128(Dx)
const Dtc = complex128(Dt)

const RDx2 = complex128(float64(1) / Dx / Dx)

const L = 200

func initValues() [][]complex128 {
	values := make([][]complex128, L)
	for i := range values {
		values[i] = make([]complex128, L)
	}
	return values
}

func getInitialCond() [][]complex128 {
	values := initValues()
	for i := range values {
		for j := range values[i] {
			dist := float64((i-L/2)*(i-L/2) + (j-L/2-L/8)*(j-L/2-L/8))
			values[i][j] = complex(math.Exp(-float64(dist)/2000), 0) *
				cmplx.Exp(complex(0, float64(i)*0.01))
		}
	}
	return values
}

func New(config Config) Field {
	values := getInitialCond()
	prev := getInitialCond()
	buf := initValues()
	f := Field{
		values: values,
		prev:   prev,
		buf:    buf,
		config: config,
	}
	f.normalize()
	return f
}

func (f *Field) get(i, j int) complex128 {
	if i < 0 || j < 0 {
		return 0
	}
	if i >= L || j >= L {
		return 0
	}
	return f.values[i][j]
}

func (f *Field) Laplacian5p(i, j int) complex128 {
	return RDx2 * (f.get(i+1, j) +
		f.get(i, j+1) +
		f.get(i-1, j) +
		f.get(i, j-1) -
		4*f.get(i, j))
}

func (f *Field) Laplacian9p(i, j int) complex128 {
	return RDx2 * (f.get(i-1, j+1) + 4*f.get(i, j+1) + f.get(i+1, j+1) +
		4*f.get(i-1, j) - 20*f.get(i, j) + 4*f.get(i+1, j) +
		f.get(i-1, j-1) + 4*f.get(i, j-1) + f.get(i+1, j-1)) / 6
}

func getPotential(i, j int) complex128 {
	dist := math.Sqrt(float64((i-L/2)*(i-L/2) + (j-L/2)*(j-L/2)))
	if dist < 3 {
		dist = 3
	}
	return -complex128(1.) / complex(dist, 0)
}

func (f *Field) Next() {
	next := f.buf

	a := complex(0, Dt/2/f.config.Mass)
	b := complex(0, Dt/f.config.ReducedPlancksConstant)

	for i := range f.values {
		for j := range f.values[i] {
			v := a*f.Laplacian9p(i, j) -
				b*f.get(i, j)*getPotential(i, j) +
				f.prev[i][j]

			next[i][j] = v
		}
	}

	f.buf = f.prev
	f.prev = f.values
	f.values = next
	f.normalize()
}

func squareOfAbs(c complex128) float64 {
	x := real(c)
	y := imag(c)
	return x*x + y*y
}

func (f *Field) normalize() {
	c := Dx * Dx
	sum := float64(0)
	for i := range f.values {
		for j := range f.values[i] {
			sum += squareOfAbs(f.values[i][j]) * c
		}
	}

	if sum > 1.0001 {
		fmt.Println(sum)
	}
	m := 1 / math.Sqrt(sum)

	for i := range f.values {
		for j := range f.values[i] {
			f.values[i][j] *= complex(m, 0)
		}
	}
}
