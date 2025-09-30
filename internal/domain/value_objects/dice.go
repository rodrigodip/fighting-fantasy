package dice

import (
	"math/rand"
	"time"
)

type Dados struct{}

func (d *Dados) RollD6() int {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	return r.Intn(5) + 1
}

func (d *Dados) Roll2D6() (d1, d2 int) {
	return d.RollD6(), d.RollD6()
}
