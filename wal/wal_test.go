package wal

import (
	"encoding/binary"
	"fmt"
	"hash/crc32"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSave(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := os.CreateTemp("testdata", "wal_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Create a Wal instance
	w := &Wal{
		file:     tmpfile,
		crcTable: crc32.MakeTable(crc32.Castagnoli),
	}

	// Test case 1: Save empty entities
	err = w.Save(nil)
	assert.NoError(t, err)

	// Test case 2: Save non-empty entities
	entities := [][]byte{
		[]byte("entity1"),
		[]byte("entity22"),
		[]byte("entity333"),
	}
	err = w.Save(entities)
	assert.NoError(t, err)

	// Verify the contents of the file
	fileContent, err := os.ReadFile(tmpfile.Name())
	assert.NoError(t, err)

	const crcLen = 4

	// Verify the length of each record in the file
	offset := 0
	for _, entity := range entities {
		recordLength := binary.LittleEndian.Uint64(fileContent[offset : offset+8])
		assert.Equal(t, uint64(len(entity)), recordLength-crcLen)
		offset += 8 + int(recordLength)
	}

	// Verify the CRC checksum of each record in the file
	offset = 0
	for _, entity := range entities {
		recordLength := binary.LittleEndian.Uint64(fileContent[offset : offset+8])
		recordData := fileContent[offset+8 : offset+8+int(recordLength)]
		r := &record{}
		r.UnMarshal(recordData)
		recordCRC := crc32.Checksum(r.data, w.crcTable)
		assert.Equal(t, recordCRC, r.crc)
		assert.Equal(t, entity, r.data)
		offset += 8 + int(recordLength)
	}
}

func TestReadRecord(t *testing.T) {
	// Create a temporary file for testing
	tmpfile, err := os.CreateTemp("testdata", "wal_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())

	// Create a Wal instance
	w := &Wal{
		file: tmpfile,
	}

	// Write test data to the file
	testData := []byte("test data")

	buf := make([]byte, len(testData))
	binary.LittleEndian.PutUint64(buf, uint64(len(testData)))
	_, err = w.file.Write(buf)
	if err != nil {
		t.Fatal(err)
	}

	_, err = w.file.Write(testData)
	if err != nil {
		t.Fatal(err)
	}
	w.file.Sync()
	w.file.Seek(0, 0)
	fmt.Println("hahahahaa", len(testData))

	// Call the readRecord function
	record, err := w.readRecord()
	assert.NoError(t, err)
	fmt.Println("record", string(record))
	assert.Equal(t, testData, record)
}
