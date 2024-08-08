package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/pflag"
)

func main() {
	input := ""
	output := ""

	fps := 10
	width := 320
	height := -1
	flags := "lanczos,split[s0][s1];[s0]palettegen[p];[s1][p]paletteuse"

	switch len(os.Args) {
	case 1:
		pflag.StringVarP(&input, "input", "i", input, "input file path")
		pflag.StringVarP(&output, "output", "o", output, "output file path")
	case 2:
		input = os.Args[1]
		output = strings.TrimSuffix(input, filepath.Ext(input)) + ".gif"
	case 3:
		input = os.Args[1]
		output = os.Args[2]
	default:
		pflag.PrintDefaults()
		os.Exit(1)
	}

	pflag.IntVarP(&fps, "fps", "", fps, "output frames per second")
	pflag.IntVarP(&width, "width", "w", width, "output width (px)")
	pflag.IntVarP(&height, "height", "h", height, "output height (px)")
	pflag.StringVarP(&flags, "flags", "f", flags, "ffmpeg flags")

	pflag.Parse()

	if input == "" || output == "" {
		pflag.PrintDefaults()
		os.Exit(1)
		return
	}

	if filepath.Ext(output) != ".gif" {
		fmt.Println("output file must be a gif")
		os.Exit(1)
	}

	cmd := exec.Command(
		"ffmpeg",
		"-i", input,
		"-vf", fmt.Sprintf(
			"fps=%d,scale=%d:%d:flags=%s",
			fps,
			width,
			height,
			flags,
		),
		"-loop", "0",
		output,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
