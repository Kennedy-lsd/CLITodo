package main

import (
	"encoding/json"
	"errors"
	"os"
)

type TaskStorage interface {
	Save(tasks []Task) error
	Load() ([]Task, error)
}

type FileStorage struct {
	FilePath string
}

func (f *FileStorage) Save(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	file, err := os.Create(f.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	return err
}

func (f *FileStorage) Load() ([]Task, error) {
	file, err := os.Open(f.FilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	return tasks, err
}
