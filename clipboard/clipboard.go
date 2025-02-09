package clipboard

import (
	"errors"
	"io"
	"os/exec"
)

func NewClipboard(cmd ClipCmd) (clip Clipboard, err error) {
	clip.cmd, err = cmd.get()
	if err != nil {
		return clip, err
	}

	clip.stdin, err = clip.cmd.StdinPipe()
	if err != nil {
		return clip, err
	}

	if err := clip.cmd.Start(); err != nil {
		return clip, err
	}

	return clip, nil
}

type Clipboard struct {
	cmd   *exec.Cmd
	stdin io.WriteCloser
	used  bool
}

func (c *Clipboard) CopyFrom(r io.Reader) error {
	if c.used {
		return errors.New("clipboards should not be reused")
	}
	c.used = true

	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	if _, err := c.stdin.Write(data); err != nil {
		return err
	}
	if err := c.stdin.Close(); err != nil {
		return err
	}

	return c.cmd.Wait()
}

type ClipCmd struct {
	Cmd  string
	Args []string
}

func (c ClipCmd) get() (*exec.Cmd, error) {
	if c.Cmd == "" {
		return nil, errors.New("no clipboard command provided")
	}
	return exec.Command(c.Cmd, c.Args...), nil
}
