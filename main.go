package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"io"
	"os"
	"os/exec"
)

var (
	isGetMergeHash bool
)

func init() {
	kingpin.Flag("--merge-hash", "").Default(FALSE).BoolVar(&isGetMergeHash)
}

func main() {
	runAndShow(exec.Command("git", os.Args[1:]...), os.Stdout)
}

func runAndShow(cmd *exec.Cmd, out io.Writer) {
	res, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	if _, err := fmt.Fprintln(out, string(res)); err != nil {
		panic(err)
	}
}
