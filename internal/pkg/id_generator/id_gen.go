package IDgenerator

import (
	"strconv"
	"sync"
	"time"
)

type TimestampIDGenerator struct {
	mutex    sync.Mutex
	lastTime int64
	sequence int64
}

func NewTimestampIDGenerator() *TimestampIDGenerator {
	return &TimestampIDGenerator{
		lastTime: time.Now().UnixNano(),
	}
}

func (g *TimestampIDGenerator) NewID() string {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	currentTime := time.Now().UnixNano()

	if currentTime == g.lastTime {
		g.sequence++
	} else {
		g.sequence = 0
		g.lastTime = currentTime
	}

	newID := (currentTime << 16) | (g.sequence & 0xFFFF)

	return strconv.FormatInt(newID, 10)
}

var lastID int

func NewSimpleID() string {

	var newID int = lastID + 1
	lastID = newID

	return strconv.Itoa(newID)
}
