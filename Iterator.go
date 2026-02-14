package main

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func Count(n int) func(func(int) bool) {
	return func(yield func(int) bool) {
		for i := 0; i < n; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

func FindLogFiles(root string) func(func(string) bool) {
	return func(yield func(string) bool) {
		_ = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return nil
			}

			if d.IsDir() {
				return nil
			}

			if !strings.HasSuffix(path, ".log") {
				return nil
			}

			if !yield(path) {
				return fs.SkipAll
			}

			return nil

		})
	}
}
