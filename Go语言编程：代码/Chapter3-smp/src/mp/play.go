package mp

import (
	"fmt"
	"time"
)

type Player interface {
	Play(source string)
}

func Play(source, mtype string) {
	var p Player
	switch mtype {
	case "MP3":
		p = &MP3Player{}
	case "WAV":
		p = &WAVPlayer{}
	default:
		fmt.Println("UnSupported music type", mtype)
		return
	}
	p.Play(source)
}

type MP3Player struct {
	stat     int
	progress int
}

func (p *MP3Player) Play(source string) {
	fmt.Println("Playing Mp3 music", source)
	p.progress = 0
	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(".")
		p.progress += 10
	}
	fmt.Println("\nFinished playing", source)
}

type WAVPlayer struct {
	progress int
}

func (p *WAVPlayer) Play(source string) {
	fmt.Println("Playing WAV music", source)
	p.progress = 0
	for p.progress < 100 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(".")
		p.progress += 10
	}
	fmt.Println("\nFinished playing", source)
}
