package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

const separator = ","

type Writer interface {
	Write(objects []Object) error
}

func objectsToCSV(objects []Object) (string, error) {
	out := ""
	columns := ""
	for _, object := range objects {
		columns = ToCSVLine(object.SortedKeys())
		out = out + object.String() + "\n"
	}
	return columns + "\n" + out, nil
}

type LocalWriter struct {
	name string
	path string
}

func NewLocalWriter(name string, path string) *LocalWriter {
	return &LocalWriter{name, path}
}

func (w LocalWriter) Write(objects []Object, guests int) error {
	now := time.Now()
	fileName := fmt.Sprintf("%s/%s_%s_%d_guests.csv", w.path, w.name, now.Format("2006-01-02_15-04-05"), guests)
	s, err := objectsToCSV(objects)
	if err != nil {
		return err
	}
	data := []byte(s)
	return ioutil.WriteFile(fileName, data, 0644)
}
