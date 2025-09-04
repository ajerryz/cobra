package storage

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	DefaultStorageDbFile = "~/.taskmanager/taskmanager"
)

type Task struct {
	Index int
	Name  string
	Desc  string
}

func (task *Task) ToString() string {
	return fmt.Sprintf("index:%d, name:%s, desc:%s", task.Index, task.Name, task.Desc)
}

type FileStorage struct {
	FilePath string
}

func NewFileStorage(filePath string) *FileStorage {
	return &FileStorage{
		FilePath: filePath,
	}
}

func (storage *FileStorage) open() (*os.File, error) {
	filePath := storage.FilePath
	dir := filepath.Dir(filePath)
	fmt.Println(dir)
	stat, _ := os.Stat(dir)
	fmt.Println(stat.Name(), stat.IsDir())
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				return nil, err
			}
		}
		return nil, err
	}
	return os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
}

func (storage *FileStorage) Append(task *Task) error {
	file, err := storage.open()
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(task.ToString())
	return err
}

var DefaultFileStorage = NewFileStorage(DefaultStorageDbFile)

func AppendTask(t *Task) error {
	return DefaultFileStorage.Append(t)
}
