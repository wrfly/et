package limiter

import (
	"errors"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

// TODO: need to save these limiters in a `Close()` function

// error
var (
	ErrExceeded = errors.New("can not handle your request due to limiter")
)

const (
	day = time.Hour * 24
)

type Limiter struct {
	name string
	max  int

	m sync.Mutex
	c *cache.Cache
}

var (
	// same IP can submit x tasks per 24 hours
	IP = &Limiter{
		name: "limitIP",
		max:  10,
		c:    cache.New(day, time.Hour),
	}

	// a task can sent x emails before the resume
	Sent = &Limiter{
		name: "limitSent",
		max:  5,
		c:    cache.New(day*365, day*365),
	}

	// one address can receive x emails per 24 hours
	Recv = &Limiter{
		name: "limitRecv",
		max:  50,
		c:    cache.New(day, time.Hour),
	}
)

func (l *Limiter) Inc(k string) error {
	current, found := l.c.Get(k)
	if found {
		if current.(int) > l.max {
			return ErrExceeded
		}
	} else {
		l.Reset(k)
	}

	currentInt, _ := l.c.IncrementInt(k, 1)
	if currentInt > l.max {
		return ErrExceeded
	}

	return nil
}

func (l *Limiter) Dec(k string) {
	if _, found := l.c.Get(k); found {
		l.c.DecrementInt(k, 1)
	}
}

func (l *Limiter) Reset(k string) {
	l.c.SetDefault(k, 0)
}
