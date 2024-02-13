package methods

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"github.com/buts00/Archiver/internal/compression/methods/table"
	"log"
	"strings"
)

type EncoderDecoder struct {
	tblGenerator table.Generator
}

func New(tblG table.Generator) EncoderDecoder {
	return EncoderDecoder{
		tblGenerator: tblG,
	}
}

func (ed EncoderDecoder) Encode(data string) []byte {
	tbl := ed.tblGenerator.NewTable(data)
	encoded := EncodeToBin(data, tbl)
	return buildEncodedFile(tbl, encoded)
}

func (ed EncoderDecoder) Decode(encodedData []byte) string {
	tbl, data := parseFile(encodedData)
	return tbl.Decode(data)
}

func parseFile(data []byte) (table.EncodingTable, string) {
	const tableSizeBytesCount = 4
	tableSizeBinary, data := data[:tableSizeBytesCount], data[tableSizeBytesCount:]
	dataSizeBinary, data := data[:tableSizeBytesCount], data[tableSizeBytesCount:]
	tableSize := binary.BigEndian.Uint32(tableSizeBinary)
	dataSize := binary.BigEndian.Uint32(dataSizeBinary)

	tblBinary, data := data[:tableSize], data[tableSize:]
	tbl := decodeTable(tblBinary)

	body := NewBinChunks(data).Join()

	return tbl, body[:dataSize]

}

func buildEncodedFile(tbl table.EncodingTable, data string) []byte {
	encodedTable := encodeTable(tbl)
	var buff bytes.Buffer
	buff.Write(encodeInt(len(encodedTable)))
	buff.Write(encodeInt(len(data)))
	buff.Write(encodedTable)
	buff.Write(SplitByChunks(data, chunkSize).Binary())
	return buff.Bytes()

}

func encodeInt(num int) []byte {
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(num))
	return res
}

func decodeTable(tblBinary []byte) table.EncodingTable {
	var tbl table.EncodingTable
	r := bytes.NewReader(tblBinary)
	if err := gob.NewDecoder(r).Decode(&tbl); err != nil {
		log.Fatal(err)
	}
	return tbl
}

func encodeTable(tbl table.EncodingTable) []byte {
	var tableBuffer bytes.Buffer

	if err := gob.NewEncoder(&tableBuffer).Encode(tbl); err != nil {
		log.Fatal(err)
	}

	return tableBuffer.Bytes()
}

func EncodeToBin(data string, table table.EncodingTable) string {
	builder := strings.Builder{}
	for _, ch := range data {
		builder.WriteString(bin(ch, table))
	}
	return builder.String()
}

func bin(ch rune, table table.EncodingTable) string {

	val, ok := table[ch]
	if !ok {
		panic("unknown symbol " + string(ch))
	}
	return val
}
