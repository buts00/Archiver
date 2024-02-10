package commands

import (
	"errors"
	"github.com/buts00/Archiver/internal/lab/vlc"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}

var ErrEmptyPath = errors.New("path to file isn't specified")

const packedExt = "vlc"

func init() {
	packCmd.AddCommand(vlcCmd)
}

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		log.Fatal(ErrEmptyPath)
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	packed := vlc.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func packedFileName(path string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(path)) + "." + packedExt
}
