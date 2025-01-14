package progbar

import (
	"bytes"
	"fmt"
	"github.com/mattn/go-runewidth"
	"math"
	"regexp"
	//"unicode/utf8"
)

const (
	_KiB = 1024
	_MiB = 1048576
	_GiB = 1073741824
	_TiB = 1099511627776
)

var ctrlFinder = regexp.MustCompile("\x1b\x5b[0-9]+\x6d")

// CellCount returns the number of runes in a string, regardless of what encoding it is in.
func CellCount(s string) int {
	n := runewidth.StringWidth(s)
	for _, sm := range ctrlFinder.FindAllString(s, -1) {
		n -= runewidth.StringWidth(sm)
	}
	return n
}

// StripString strips the string and returns it
func StripString(s string, w int) string {
	l := CellCount(s)
	if l <= w {
		return s
	}
	var buf = bytes.NewBuffer(make([]byte, 0, len(s)))
	StripStringToBuffer(s, w, buf)
	return buf.String()
}

// StripStringToBuffer strips the string and writes it the buffer
func StripStringToBuffer(s string, w int, buf *bytes.Buffer) {
	var seqs = ctrlFinder.FindAllStringIndex(s, -1)
mainloop:
	for i, r := range s {
		for _, seq := range seqs {
			if i >= seq[0] && i < seq[1] {
				buf.WriteRune(r)
				continue mainloop
			}
		}
		if rw := CellCount(string(r)); rw <= w {
			w -= rw
			buf.WriteRune(r)
		} else {
			break
		}
	}
	for w > 0 {
		buf.WriteByte(' ')
		w--
	}
	return
}

func round(val float64) (newVal float64) {
	roundOn := 0.5
	places := 0
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

// Convert bytes to human readable string. Like a 2 MiB, 64.2 KiB, 52 B
func formatBytes(i int64) (result string) {
	switch {
	case i >= _TiB:
		result = fmt.Sprintf("%.02f TiB", float64(i)/_TiB)
	case i >= _GiB:
		result = fmt.Sprintf("%.02f GiB", float64(i)/_GiB)
	case i >= _MiB:
		result = fmt.Sprintf("%.02f MiB", float64(i)/_MiB)
	case i >= _KiB:
		result = fmt.Sprintf("%.02f KiB", float64(i)/_KiB)
	default:
		result = fmt.Sprintf("%d B", i)
	}
	return
}
