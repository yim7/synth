package ui

import (
	"fmt"
	"log"
	"sync"

	"github.com/veandco/go-sdl2/ttf"
)

var DefaultFont *Font

func init() {
	if err := ttf.Init(); err != nil {
		log.Fatal(err)
	}

	if font, err := LoadFont("fonts/wqy-microhei.ttc", 17); err != nil {
		log.Fatal(err)
	} else {
		DefaultFont = font
	}
}

var (
	fonts map[string]*Font
	mutex sync.Mutex
)

type Font struct {
	*ttf.Font
	name string
	size int
}

// font loader, cache by name and size
func LoadFont(name string, size int) (*Font, error) {
	key := fmt.Sprintf("%s-%d", name, size)
	mutex.Lock()
	defer mutex.Unlock()

	if fonts == nil {
		fonts = make(map[string]*Font)
	}

	if font, ok := fonts[key]; ok {
		return font, nil
	}

	raw, err := ttf.OpenFont(name, size)
	if err != nil {
		return nil, err
	}
	font := &Font{raw, name, size}
	fonts[key] = font

	return font, err
}
