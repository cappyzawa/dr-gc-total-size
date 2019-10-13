package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
)

type CLI struct {
	Out io.Writer
	Err io.Writer
	In  io.Reader
}

var (
	r       = regexp.MustCompile("blob eligible for deletion: sha256:(.*)")
	rootDir string
)

type Blob struct {
	Dir    string
	SHA256 []byte
}

func (c *CLI) Run(args []string) int {
	flags := flag.NewFlagSet(args[0], flag.ContinueOnError)
	flags.SetOutput(c.Err)
	flags.StringVar(&rootDir, "d", "", "root directory (storage.filesystem.rootdirectory)")
	if err := flags.Parse(args[1:]); err != nil {
		return 1
	}
	scanner := bufio.NewScanner(c.In)

	blobsDir := filepath.Join(rootDir, "docker/registry/v2/blobs/sha256")
	blob := &Blob{
		Dir: blobsDir,
	}
	var totalSize int64
	for scanner.Scan() {
		line := scanner.Bytes()
		if r.Match(line) {
			fmt.Fprintf(c.Out, "%s\n", string(line))

			sha := r.FindSubmatch(line)[1]
			blob.SHA256 = sha
			size, err := blob.Find()
			if err != nil {
				return 1
			}
			totalSize += size
		} else {
			fmt.Fprintf(c.Out, "%s\n", string(line))
		}
	}

	fmt.Fprintf(c.Out, "%d bytes will be reduced\n", totalSize)
	return 0
}

func (b *Blob) Find() (int64, error) {
	shaPrefix := b.SHA256[:2]
	file := filepath.Join(b.Dir, string(shaPrefix), string(b.SHA256), "data")
	fs, err := os.Stat(file)
	if err != nil {
		return 0, err
	}
	return fs.Size(), nil
}

func main() {
	cli := &CLI{
		Out: os.Stdout,
		Err: os.Stderr,
		In:  os.Stdin,
	}
	os.Exit(cli.Run(os.Args))
}
