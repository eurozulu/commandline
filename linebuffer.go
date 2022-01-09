package commandline

import (
	"fmt"
	"io"
	"log"
	"strings"
)

type lineBuffer struct {
	cursorOffset int
	value        []byte
	Out          io.Writer
}

func (lo lineBuffer) String() string {
	return string(lo.value)
}

func (lo lineBuffer) Value() []byte {
	return lo.value
}

func (lo *lineBuffer) Clear() {
	lo.value = nil
	lo.cursorOffset = 0
}

func (lo *lineBuffer) Insert(v []byte) {
	lo.SetValue(lo.valueBeforeCursor(), v, lo.valueAfterCursor())
}

func (lo *lineBuffer) SetValue(v ...[]byte) {
	lo.clearOutput()
	lo.value = nil
	if len(v) > 0 {
		for _, vv := range v {
			lo.value = append(lo.value, vv...)
		}
	}
	lo.writeOutput()
}

func (lo *lineBuffer) Delete() {
	v := lo.valueBeforeCursor()
	if len(v) == 0 {
		return
	}
	lo.SetValue(v[:len(v)-1], lo.valueAfterCursor())
}

func (lo *lineBuffer) CursorPosition() int {
	return len(lo.value) - lo.cursorOffset
}

func (lo *lineBuffer) CursorLeft() {
	if lo.cursorOffset < len(lo.value) {
		lo.clearOutput()
		lo.cursorOffset++
		lo.writeOutput()
	}
}
func (lo *lineBuffer) CursorRight() {
	if lo.cursorOffset > 0 {
		lo.clearOutput()
		lo.cursorOffset--
		lo.writeOutput()
	}
}
func (lo *lineBuffer) CursorHome() {
	lo.clearOutput()
	lo.cursorOffset = len(lo.value)
	lo.writeOutput()
}
func (lo *lineBuffer) CursorEnd() {
	lo.clearOutput()
	lo.cursorOffset = 0
	lo.writeOutput()
}

func (lo *lineBuffer) writeOutput() {
	bs := strings.Repeat(BACKSPACE.String(), lo.cursorOffset)
	if _, err := fmt.Fprint(lo.Out, lo.String(), bs); err != nil {
		log.Println(err)
	}
}

func (lo *lineBuffer) clearOutput() {
	if len(lo.value) == 0 {
		return
	}
	bs := strings.Repeat(BACKSPACE.String(), len(lo.value))
	ws := strings.Repeat(" ", len(lo.value))
	tail := strings.Repeat(" ", len(lo.valueAfterCursor()))
	if _, err := fmt.Fprint(lo.Out, tail, bs, ws, bs); err != nil {
		log.Println(err)
	}
}

func (lo *lineBuffer) valueBeforeCursor() []byte {
	return lo.value[:lo.CursorPosition()]
}

func (lo *lineBuffer) valueAfterCursor() []byte {
	if lo.cursorOffset == 0 {
		return nil
	}
	return lo.value[lo.CursorPosition():]
}
