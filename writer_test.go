package ach

import (
	"bytes"
	"strings"
	"testing"
)

func TestPPDWrite(t *testing.T) {
	file := NewFile().SetHeader(mockFileHeader())
	entry := mockEntryDetail()
	entry.AddAddenda(mockAddenda())
	batch := NewBatch().SetHeader(mockBatchHeader()).AddEntryDetail(entry)
	batch.Build()
	file.AddBatch(batch)
	file.Build()
	if err := file.Build(); err != nil {
		t.Errorf("Could not build file: %v", err)
	}
	if err := file.ValidateAll(); err != nil {
		t.Errorf("Could not validate built file: %v", err)
	}

	b := &bytes.Buffer{}
	f := NewWriter(b)

	err := f.WriteAll([]*File{file})
	if err != nil {
		t.Errorf("Unexpected error: %s\n", err)
	}
	//	fmt.Print(b.String())
	r := NewReader(strings.NewReader(b.String()))
	_, err = r.Read()
	if err != nil {
		t.Errorf("Can not ach.Read generated file: %v", err)
		//fmt.Printf("%+v \n", r.File.Batches[0])
	}
	err = r.File.ValidateAll()
	if err != nil {
		t.Errorf("Could not validate entire generated file: %v", err)
		//fmt.Printf("%+v", r.File)
	}
}
