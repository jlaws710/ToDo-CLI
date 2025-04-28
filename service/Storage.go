package service

import (
	"ToDo-CLI/model"
	"encoding/json"
	"fmt"
	"os"
)

const fileName = "tasks.json"

func Save(tasks []model.Task) error {
	file, err := os.Create(fileName)

	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)

	return encoder.Encode(tasks)
}

func Load() ([]model.Task, error) {
	file, err := os.Open(fileName)

	if err != nil {

		return []model.Task{}, nil
	}
	defer file.Close()

	var tasks []model.Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)

	return tasks, err
}
