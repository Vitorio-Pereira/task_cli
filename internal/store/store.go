package store

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Task struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
}

type Store struct {
	Task []Task
	Path string
}

func New() *Store {
	userDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Erro ao obter home dir: %v", err)
	}
	dir := filepath.Join(userDir, ".task")
	path := filepath.Join(dir, "task.json")
	os.Mkdir(dir, 0775)

	return &Store{
		Path: path,
	}
}

func (s *Store) Load() error {
	data, err := os.ReadFile(s.Path)

	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}
	err = json.Unmarshal(data, &s.Task)
	return err
}

func (s *Store) Save() error {
	data, err := json.Marshal(&s.Task)
	if err != nil {
		return err
	}
	err = os.WriteFile(s.Path, data, 0775)
	return err
}

func (s *Store) Add(title string) error {
	task := Task{
		Id:        len(s.Task) + 1,
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}

	s.Task = append(s.Task, task)
	return s.Save()
}

func (s *Store) List() []Task {

	return s.Task
}

func (s *Store) Done(id int) error {
	for i := range s.Task {
		if s.Task[i].Id == id {
			s.Task[i].Done = true
			return s.Save()
		}
	}

	return fmt.Errorf("task %d not found", id)
}

func (s *Store) Remove(id int) error {
	for i := range s.Task {
		if s.Task[i].Id == id {
			s.Task = append(s.Task[:i], s.Task[i+1:]...)
			return s.Save()
		}
	}
	return fmt.Errorf("task %d not found", id)

}
