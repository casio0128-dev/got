package main

import (
	"bytes"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"io"
	"os"
	"os/exec"
	"strings"
)

var (
	isGetMergeHash bool
)

func init() {
	kingpin.Flag("--merge-hash", "").Default(FALSE).BoolVar(&isGetMergeHash)
}

func main() {
	//runAndShow(exec.Command("git", os.Args[1:]...), os.Stdout)

	getMergeCommitHash(os.Stdout, "cbf0e1da68f070a5061c27e597805ea05c09c4af", "main")
}

func getMergeCommitHash(out io.Writer, targetHash, branch string) {
	var b1, b2 bytes.Buffer
	runAndShow(exec.Command("git", "rev-list", fmt.Sprintf("%s..%s", targetHash, branch), "--ancestry-path"), &b1)
	runAndShow(exec.Command("git", "rev-list", fmt.Sprintf("%s..%s", targetHash, branch), "--first-parent"), &b2)

	//fmt.Fprintln(out, b1.String())
	//fmt.Fprintln(out, "-----------------------")
	//fmt.Fprintln(out, b2.String())

	var result string
	for _, h := range strings.Split(b1.String(), "\n") {
		if strings.Contains(b2.String(), h) {
			result = b2.String()
		}
	}
	fmt.Println("result", result)

	//git rev-list [特定したいコミットのSHA-1]..master --ancestry-path > FILE1
	//git rev-list [特定したいコミットのSHA-1]..master --first-parent > FILE2
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
