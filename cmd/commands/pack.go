package commands

import (
	"errors"
	"github.com/buts00/Archiver/internal/compression"
	"github.com/buts00/Archiver/internal/compression/methods"
	"github.com/buts00/Archiver/internal/compression/methods/table/shannon-fano"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack file",
	Run:   pack,
}

func init() {
	rootCmd.AddCommand(packCmd)
	packCmd.Flags().StringP("method", "m", "", "compression method: sf")
	err := packCmd.MarkFlagRequired("method")
	if err != nil {
		log.Fatal(err)
	}
}

var ErrEmptyPath = errors.New("path to file isn't specified")

const packedExt = "sf"

func pack(cmd *cobra.Command, args []string) {
	var encoder compression.Encoder
	if len(args) == 0 || args[0] == "" {
		log.Fatal(ErrEmptyPath)
	}
	filePath := args[0]
	method := cmd.Flag("method").Value.String()
	switch method {
	case "shannon-fano":
		encoder = methods.New(shannon_fano.NewGenerator())
	default:
		cmd.PrintErr("unknown method")
	}
	r, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer func(r *os.File) {
		err = r.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(r)

	data, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	packed := encoder.Encode(string(data))

	err = os.WriteFile(packedFileName(filePath), packed, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func packedFileName(path string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(path)) + "." + packedExt
}
