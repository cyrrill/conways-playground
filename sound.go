package main

// Based on:  https://github.com/hajimehoshi/oto/blob/main/example/main.go

import (
	"io"
	"math"
	"runtime"
	"sync"
	"time"

	"github.com/hajimehoshi/oto/v2"
)

var (
	sampleRate      = 44100
	channelNum      = 2
	bitDepthInBytes = 2

	otoContext *oto.Context
	ready chan struct{}
	players []oto.Player
)

type SineWave struct {
	freq   float64
	length int64
	pos    int64

	remaining []byte
}


func algoSound(totalDots int) {
	var numerator int
	if randomSeed > 5000 {
		numerator = randomSeed
	} else {
		numerator = 5000
	}

	var counting int
	if totalDots < 5000 {
		counting = 5000
	} else {
		counting = totalDots
	}


	multiple := (numerator / 2000) + 1
	freq := float64(counting / multiple) * math.Sin(float64(totalDots))

	if totalDots % 3 == 0 {
		playSound(freq)
	}
}

func NewSineWave(freq float64, duration time.Duration) *SineWave {
	l := int64(channelNum) * int64(bitDepthInBytes) * int64(sampleRate) * int64(duration) / int64(time.Second)
	l = l / 4 * 4
	return &SineWave{
		freq:   freq,
		length: l,
	}
}

func (s *SineWave) Read(buf []byte) (int, error) {
	if len(s.remaining) > 0 {
		n := copy(buf, s.remaining)
		copy(s.remaining, s.remaining[n:])
		s.remaining = s.remaining[:len(s.remaining)-n]
		return n, nil
	}

	if s.pos == s.length {
		return 0, io.EOF
	}

	eof := false
	if s.pos+int64(len(buf)) > s.length {
		buf = buf[:s.length-s.pos]
		eof = true
	}

	var origBuf []byte
	if len(buf)%4 > 0 {
		origBuf = buf
		buf = make([]byte, len(origBuf)+4-len(origBuf)%4)
	}

	length := float64(sampleRate) / float64(s.freq)

	num := (bitDepthInBytes) * (channelNum)
	p := s.pos / int64(num)
	switch bitDepthInBytes {
	case 1:
		for i := 0; i < len(buf)/num; i++ {
			const max = 127
			b := int(math.Sin(2*math.Pi*float64(p)/length) * 0.3 * max)
			for ch := 0; ch < channelNum; ch++ {
				buf[num*i+ch] = byte(b + 128)
			}
			p++
		}
	case 2:
		for i := 0; i < len(buf)/num; i++ {
			const max = 32767
			b := int16(math.Sin(2*math.Pi*float64(p)/length) * 0.3 * max)
			for ch := 0; ch < channelNum; ch++ {
				buf[num*i+2*ch] = byte(b)
				buf[num*i+1+2*ch] = byte(b >> 8)
			}
			p++
		}
	}

	s.pos += int64(len(buf))

	n := len(buf)
	if origBuf != nil {
		n = copy(origBuf, buf)
		s.remaining = buf[n:]
	}

	if eof {
		return n, io.EOF
	}
	return n, nil
}

func play(context *oto.Context, freq float64, duration time.Duration) oto.Player {
	p := context.NewPlayer(NewSineWave(freq, duration))
	p.Play()
	return p
}

func initSound() {
	otoContext, ready, err = oto.NewContext(sampleRate, channelNum, bitDepthInBytes)
}

func playSound(freq float64) error {
	if err != nil {
		return err
	}
	<-ready

	var m sync.Mutex
	go func() {
		p := play(otoContext, freq, time.Duration(1000) * time.Millisecond)
		m.Lock()
		players = append(players, p)
		m.Unlock()
	}()

	// Pin the players not to GC the players.
	runtime.KeepAlive(players)

	return nil
}
