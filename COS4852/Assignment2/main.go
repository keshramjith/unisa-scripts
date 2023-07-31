package main

import (
	"fmt"
	"math/rand"
)

func main() {
	question1b()
}

func question1b() {
	x1 := make([]float32, 8)
	x1[0] = 1
	x1[1] = 6
	x1[2] = 5
	x1[3] = 8
	x1[4] = 2
	x1[5] = 1
	x1[6] = 6
	x1[7] = 2

	x2 := make([]float32, 8)
	x2[0] = 1
	x2[1] = 6
	x2[2] = -1
	x2[3] = 2
	x2[4] = 4
	x2[5] = -2
	x2[6] = -5
	x2[7] = -5

	t := make([]float32, 8)
	t[0] = 1
	t[1] = 1
	t[2] = 1
	t[3] = 1
	t[4] = -1
	t[5] = -1
	t[6] = -1
	t[7] = -1

	w := make([]float32, 3)
	w[0] = rand.Float32()
	w[1] = rand.Float32()
	w[2] = rand.Float32()

	const learning_rate = 0.1

	x := make([][]float32, 2)
	x[0] = x1
	x[1] = x2

	i := 0
	var delta_w0 float32
	var delta_w1 float32
	var delta_w2 float32
	for {
		i := i % 8
		o := w[0] + x[0][i]*w[1] + x[1][i]*w[2]
		delta_w0_new := learning_rate * (t[i] - o)
		delta_w1_new := learning_rate * (t[i] - o) * x[0][i]
		delta_w2_new := learning_rate * (t[i] - o) * x[1][i]
		if delta_w1_new == 0 && delta_w2_new == 0 && delta_w0_new == 0 {
			break
		}
		delta_w0 = delta_w0_new
		delta_w1 = delta_w1_new
		delta_w2 = delta_w2_new

		w[0] = w[0] + delta_w0
		w[1] = w[1] + delta_w1
		w[2] = w[2] + delta_w2
		i++
	}

	fmt.Println(w)
}
