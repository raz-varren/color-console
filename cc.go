//Package cc is a simple library for printing ANSI colored text to an io.Writer, usually stdout.
package cc

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type ANSIColor []byte

var (
	ansiColorPallet = map[string][]byte{
		"none":         []byte("\x1b[0m"),
		"black":        []byte("\x1b[0;30m"),
		"red":          []byte("\x1b[0;31m"),
		"green":        []byte("\x1b[0;32m"),
		"orange":       []byte("\x1b[0;33m"),
		"blue":         []byte("\x1b[0;34m"),
		"purple":       []byte("\x1b[0;35m"),
		"cyan":         []byte("\x1b[0;36m"),
		"light-gray":   []byte("\x1b[0;37m"),
		"dark-gray":    []byte("\x1b[1;30m"),
		"light-red":    []byte("\x1b[1;31m"),
		"light-green":  []byte("\x1b[1;32m"),
		"yellow":       []byte("\x1b[1;33m"),
		"light-blue":   []byte("\x1b[1;34m"),
		"light-purple": []byte("\x1b[1;35m"),
		"light-cyan":   []byte("\x1b[1;36m"),
		"white":        []byte("\x1b[1;37m"),
	}

	colorList = sortedColorList(ansiColorPallet)

	//Available ANSI colors
	ACNone        ANSIColor = ansiColorPallet["none"]
	ACBlack       ANSIColor = ansiColorPallet["black"]
	ACRed         ANSIColor = ansiColorPallet["red"]
	ACGreen       ANSIColor = ansiColorPallet["green"]
	ACOrange      ANSIColor = ansiColorPallet["orange"]
	ACBlue        ANSIColor = ansiColorPallet["blue"]
	ACPurple      ANSIColor = ansiColorPallet["purple"]
	ACCyan        ANSIColor = ansiColorPallet["cyan"]
	ACLightGray   ANSIColor = ansiColorPallet["light-gray"]
	ACDarkGray    ANSIColor = ansiColorPallet["dark-gray"]
	ACLightRed    ANSIColor = ansiColorPallet["light-red"]
	ACLightGreen  ANSIColor = ansiColorPallet["light-green"]
	ACYellow      ANSIColor = ansiColorPallet["yellow"]
	ACLightBlue   ANSIColor = ansiColorPallet["light-blue"]
	ACLightPurple ANSIColor = ansiColorPallet["light-purple"]
	ACLightCyan   ANSIColor = ansiColorPallet["light-cyan"]
	ACWhite       ANSIColor = ansiColorPallet["white"]

	defaultCC = NewCC(os.Stdout)
)

//PrintAvailableColors will print out all available ANSI colors to stdout in their respective colors
func PrintAvailableColors() {
	for _, v := range colorList {
		Println(ANSIColor(ansiColorPallet[v]), v)
	}
}

func Print(color ANSIColor, v ...interface{}) (int, error) {
	return defaultCC.Print(color, v...)
}

func Printf(color ANSIColor, format string, v ...interface{}) (int, error) {
	return defaultCC.Printf(color, format, v...)
}

func Println(color ANSIColor, v ...interface{}) (int, error) {
	return defaultCC.Println(color, v...)
}

func Fprint(color ANSIColor, w io.Writer, v ...interface{}) (int, error) {
	n, err := w.Write(color)
	if err != nil {
		return n, err
	}

	n, err = fmt.Fprint(w, v...)
	if err != nil {
		return n, err
	}

	return w.Write(ACNone)
}

func Fprintf(color ANSIColor, w io.Writer, format string, v ...interface{}) (int, error) {
	n, err := w.Write(color)
	if err != nil {
		return n, err
	}

	n, err = fmt.Fprintf(w, format, v...)
	if err != nil {
		return n, err
	}

	return w.Write(ACNone)
}

func Fprintln(color ANSIColor, w io.Writer, v ...interface{}) (int, error) {
	n, err := w.Write(color)
	if err != nil {
		return n, err
	}

	n, err = fmt.Fprintln(w, v...)
	if err != nil {
		return n, err
	}

	return w.Write(ACNone)
}

type CC struct {
	out io.Writer
}

//NewCC creates a new instance of CC for writing colored output to an ANSI compatible io.Writer, usually stdout
func NewCC(out io.Writer) *CC {
	return &CC{
		out: out,
	}
}

func (c *CC) Print(color ANSIColor, v ...interface{}) (int, error) {
	return Fprint(color, c.out, v...)
}

func (c *CC) Printf(color ANSIColor, format string, v ...interface{}) (int, error) {
	return Fprintf(color, c.out, format, v...)
}

func (c *CC) Println(color ANSIColor, v ...interface{}) (int, error) {
	return Fprintln(color, c.out, v...)
}

func sortedColorList(colors map[string][]byte) []string {
	var sortedList []string
	for k, _ := range colors {
		sortedList = append(sortedList, k)
	}
	sort.Strings(sortedList)
	return sortedList
}
