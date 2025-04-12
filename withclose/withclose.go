package main

import (
	"fmt"
	"io"
	"os"
)

func WithClose(closer io.Closer, fn func()) {
	defer func() {
		closer.Close()
		fmt.Printf("close %s\n", closer.(*os.File).Name())
	}()
	fn()
}

func ReadFile(paths []string) error {
	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		WithClose(file, func() {
			var content []byte
			content, err = io.ReadAll(file)
			if err != nil {
				return
			}
			fmt.Printf("%s content: %s\n", file.Name(), content)
		})
		if err != nil {
			return err
		}
	}
	return nil
}
