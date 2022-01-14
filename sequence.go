package commandline

import (
	"fmt"
)

const (
	CTRL_A = iota + 1
	CTRL_B
	CTRL_C
	CTRL_D
	CTRL_E
	CTRL_F
	CTRL_G
	CTRL_H
	CTRL_I
	CTRL_J
	CTRL_K
	CTRL_L
	CTRL_M
	CTRL_N
	CTRL_O
	CTRL_P
	CTRL_Q
	CTRL_R
	CTRL_S
	CTRL_T
	CTRL_U
	CTRL_V
	CTRL_W
	CTRL_X
	CTRL_Y
	CTRL_Z
	ESC
)
const SPACE = 32
const DEL = 127
const LSQ_BRACKET = 91

type Sequence []byte
type Sequences map[string]SequenceFunc
type SequenceFunc func(keys Sequence) error

func (s Sequence) String() string {
	return string(s)
}

var CancelledErr = fmt.Errorf("cancelled")

// sequences
var CANCEL = Sequence{CTRL_C}
var BACKSPACE = Sequence{CTRL_H}
var LINE_FEED = Sequence{CTRL_J}
var CARRIAGE_RETURN = Sequence{CTRL_M}
var CR_LF = Sequence{CTRL_M, CTRL_J}
var ESCAPE = Sequence{ESC}
var ARROW_UP = Sequence{ESC, LSQ_BRACKET, 65}
var ARROW_DOWN = Sequence{ESC, LSQ_BRACKET, 66}
var ARROW_RIGHT = Sequence{ESC, LSQ_BRACKET, 67}
var ARROW_LEFT = Sequence{ESC, LSQ_BRACKET, 68}
var DELETE = Sequence{DEL}
