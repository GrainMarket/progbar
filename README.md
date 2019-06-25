# Terminal progress bar for Go

[![Coverage Status](https://coveralls.io/repos/github/TheBrazenGeek/progbar/badge.svg)](https://coveralls.io/github/TheBrazenGeek/progbar)

## What the fork!?

Two reasons came together that lead me to creating this fork:
- trying to use [https://github.com/cheggaaa/pb/tree/v2](https://github.com/cheggaaa/pb/tree/v2) in a gomodules-enabled project was too much for me to handle due to the reliance on versioned gopkg.in packages; and
- version 1 of [Sergey's](https://github.com/cheggaaa) pb didn't have the things I needed for my project.

So I forked Version 2 and updated the imports. I get that I could have created a pull-request to update v2 in the original project, but again,two things prompted me to fork:
-  the lack of activity on the v2 branch - 2 years at the point I forked this; and
- [https://github.com/cheggaaa/pb/pull/146](an existing pull-request to make the changes I need) that had stalled for no good reason.

The features I needed from the original v2 are:
- based on text/template
- can take custom elements
- using colors is easy

## Installation

```
go get github.com/TheBrazenGeek/progbar
```

## Usage

```Go
package main

import (
	"github.com/TheBrazenGeek/progbar"
	"time"
)

func main() {
	simple()
	fromPreset()
	customTemplate(`Custom template: {{counters . }}`)
	customTemplate(`{{ red "With colors:" }} {{bar . | green}} {{speed . | blue }}`)
	customTemplate(`{{ red "With funcs:" }} {{ bar . "<" "-" (cycle . "↖" "↗" "↘" "↙" ) "." ">"}} {{speed . | rndcolor }}`)
	customTemplate(`{{ bar . "[<" "·····•·····" (rnd "ᗧ" "◔" "◕" "◷" ) "•" ">]"}}`)
}

func simple() {
	count := 1000
	bar := progbar.StartNew(count)
	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond * 2)
	}
	bar.Finish()
}

func fromPreset() {
	count := 1000
	//bar := pb.Default.Start(total)
	//bar := pb.Simple.Start(total)
	bar := progbar.Full.Start(count)
	defer bar.Finish()
	bar.Set("prefix", "fromPreset(): ")
	for i := 0; i < count/2; i++ {
		bar.Add(2)
		time.Sleep(time.Millisecond * 4)
	}
}

func customTemplate(tmpl string) {
	count := 1000
	bar := progbar.ProgressBarTemplate(tmpl).Start(count)
	defer bar.Finish()
	for i := 0; i < count/2; i++ {
		bar.Add(2)
		time.Sleep(time.Millisecond * 4)
	}
}

```
