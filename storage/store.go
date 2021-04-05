package storage

import (
	"embed"
)

//go:embed files/*
var files embed.FS

func Find(filename string) (string, error) {
	content, err := files.ReadFile("files/" + filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
