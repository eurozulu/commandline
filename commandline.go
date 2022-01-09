package commandline

import (
	"bufio"
	"bytes"
	"fmt"
	"golang.org/x/term"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type CommandLine struct {
	KeyFunctions Sequences
	history      *linehistory
	buffer       *lineBuffer
}

func (cli *CommandLine) ReadCommand() (string, error) {
	// switch stdin into 'raw' mode
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	defer func(state *term.State) {
		if err := term.Restore(int(os.Stdin.Fd()), oldState); err != nil {
			fmt.Println(err)
		}
	}(oldState)
	defer cli.buffer.Clear()

	keys := make(Sequence, 6)
	for {
		l, err := os.Stdin.Read(keys)
		if err != nil {
			return "", err
		}
		if l == 0 {
			break
		}
		if err := cli.insertKeys(keys[:l]); err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
	}
	s := cli.buffer.String()
	cli.history.Add(s)
	return s, nil
}

func (cli *CommandLine) LoadHistory(p string) error {
	by, err := ioutil.ReadFile(os.ExpandEnv(p))
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	cli.ClearHistory()
	if len(by) > 0 {
		s := bufio.NewScanner(bytes.NewBuffer(by))
		for s.Scan() {
			if s.Text() != "" {
				*cli.history = append((*cli.history), s.Text())
			}
		}
	}
	return nil
}

func (cli *CommandLine) SaveHistory(p string) error {
	s := strings.Join(*cli.history, "\n")
	return ioutil.WriteFile(os.ExpandEnv(p), []byte(s), 0600)
}

func (cli *CommandLine) ClearHistory() {
	cli.history = &linehistory{}
}

func (cli *CommandLine) insertKeys(keys Sequence) error {
	fn, ok := cli.KeyFunctions[keys.String()]
	if ok {
		return fn(keys)
	}
	cli.buffer.Insert(cleanBytes(keys))
	return nil
}

func (cli *CommandLine) cancelHandler(keys Sequence) error {
	return CancelledErr
}

func (cli *CommandLine) returnHandler(keys Sequence) error {
	return io.EOF
}

func (cli *CommandLine) escHandler(keys Sequence) error {
	cli.history.Add(cli.buffer.String())
	cli.buffer.SetValue(nil)
	return nil
}

func (cli *CommandLine) delHandler(keys Sequence) error {
	cli.buffer.Delete()
	return nil
}

func (cli *CommandLine) horizontalArrowHandler(keys Sequence) error {
	switch keys.String() {
	case ARROW_LEFT.String():
		cli.buffer.CursorLeft()
	case ARROW_RIGHT.String():
		cli.buffer.CursorRight()
	default:
		return fmt.Errorf("unknown horizontal arrow sequence %v", keys)
	}
	return nil
}

func (cli *CommandLine) verticalArrowHandler(keys Sequence) error {
	v := cli.buffer.String()
	switch keys.String() {
	case ARROW_UP.String():
		v = cli.history.Previous(v)
	case ARROW_DOWN.String():
		v = cli.history.Next(v)
	default:
		return fmt.Errorf("unknown vertical arrow sequence %v", keys)
	}
	cli.buffer.SetValue([]byte(v))
	return nil
}

// cleanBytes removes any non printable chars (ascii < 32/SPACE)
func cleanBytes(by []byte) []byte {
	var index int
	for _, b := range by {
		if b < SPACE {
			continue
		}
		by[index] = b
		index++
	}
	return by[:index]
}

func NewCommandLine() *CommandLine {
	cli := &CommandLine{
		history: &linehistory{},
		buffer:  &lineBuffer{Out: os.Stdout},
	}
	cli.KeyFunctions = Sequences{
		CANCEL.String():          cli.cancelHandler,
		CARRIAGE_RETURN.String(): cli.returnHandler,
		LINE_FEED.String():       cli.returnHandler,
		CR_LF.String():           cli.returnHandler,
		ESCAPE.String():          cli.escHandler,
		DELETE.String():          cli.delHandler,
		ARROW_LEFT.String():      cli.horizontalArrowHandler,
		ARROW_RIGHT.String():     cli.horizontalArrowHandler,
		ARROW_UP.String():        cli.verticalArrowHandler,
		ARROW_DOWN.String():      cli.verticalArrowHandler,
	}
	return cli
}
