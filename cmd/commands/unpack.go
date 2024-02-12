package commands

import (
	"github.com/buts00/Archiver/internal/lib/compression"
	"github.com/buts00/Archiver/internal/lib/compression/vlc"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var unPackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "Unpack file",
	Run:   unpack,
}

func init() {
	rootCmd.AddCommand(unPackCmd)
	
	unPackCmd.Flags().StringP("method", "m", "", "compression method: vlc")
	err := unPackCmd.MarkFlagRequired("method")
	if err != nil {
		log.Fatal(err)
	}
}

const unPackedExt = "txt"

func unpack(cmd *cobra.Command, args []string) {
	var decoder compression.Decoder
	if len(args) == 0 || args[0] == "" {
		log.Fatal(ErrEmptyPath)
	}
	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	method := cmd.Flag("method").Value.String()
	switch method {
	case "vlc":
		decoder = vlc.New()
	default:
		cmd.PrintErr("unknown method")
	}

	data, err := io.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	defer func(r *os.File) {
		err = r.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(r)

	packed := decoder.Decode(data)

	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func unpackedFileName(path string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(path)) + "." + unPackedExt
}
