package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/RBird111/ssh-clipboard/clipboard"
)

func main() {
	opts := clipboard.ServerOpts{
		Address: "localhost:8888",
		ClipCmd: clipboard.ClipCmd{
			Cmd:  "xsel",
			Args: []string{"--clipboard", "--input"},
		},
	}
	s := clipboard.NewServer(opts)
	defer s.Stop()

	cmd := exec.Command(
		"ssh", "root@localhost",
		"-p", "8022",
		"-R", "8888:localhost:8888",
	)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
