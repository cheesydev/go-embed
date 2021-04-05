package storage

import (
	"embed"
	"errors"
)

//go:embed files/*
var files embed.FS

var ErrNotFound = errors.New("File not found")

func Find(filename string) (string, error) {
	content, err := files.ReadFile("files/" + filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
