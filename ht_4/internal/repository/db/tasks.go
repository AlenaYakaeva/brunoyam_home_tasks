package db

import (
	tasksDomain "ToDoList/internal/domain/tasks"
	"ToDoList/internal/repository/errors"
	"context"
	"time"
)

func (s *Storage) SaveTask(task tasksDomain.Task) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var uid string
	err := s.conn.QueryRow(ctx,
		`INSERT INTO tasks (
		tid, 
		uid,	
		title, 
		description,
		status) 
		VALUES ($1, $2, $3, $4, $5) RETURNING tid`,
		task.TID,
		task.UID,
		task.Title,
		task.Description,
		task.Status,
	).Scan(&uid)
	if err != nil {
		return "", err
	}
	return task.TID, nil
}

func (s *Storage) GetTasks(uid string) ([]tasksDomain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := s.conn.Query(ctx, "SELECT tid, uid, title, description, status FROM tasks WHERE uid = $1", uid)
	if err != nil {
		return nil, err
	}

	var taskList []tasksDomain.Task
	for rows.Next() {
		var task tasksDomain.Task
		err := rows.Scan(&task.TID, &task.UID, &task.Title, &task.Description, &task.Status)
		if err != nil {
			return nil, err
		}
		taskList = append(taskList, task)
	}

	return taskList, nil

}

func (s *Storage) GetTaskByID(tid string) (tasksDomain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var task tasksDomain.Task
	err := s.conn.QueryRow(ctx, "SELECT tid, uid, title, description, status FROM tasks WHERE tid = $1", tid).Scan(&task.TID, &task.UID, &task.Title, &task.Description, &task.Status)
	//TODO ошибка не найденой задачи
	if err != nil {
		return tasksDomain.Task{}, err
	}
	return task, nil
}

func (s *Storage) UpdateTask(task tasksDomain.Task, tid string) (tasksDomain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.GetTaskByID(tid)
	if err != nil {
		return tasksDomain.Task{}, errors.ErrTaskNotFound
	}

	_, err = s.conn.Exec(ctx, "UPDATE tasks SET title=$1, description=$2, status=$3 WHERE tid=$4", task.Title, task.Description, task.Status, tid)
	if err != nil {
		return tasksDomain.Task{}, err
	}
	return task, nil
}

func (s *Storage) DeleteTask(tid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.GetTaskByID(tid)
	if err != nil {
		return errors.ErrTaskNotFound
	}

	_, err = s.conn.Exec(ctx, "DELETE from tasks WHERE tid=$1", tid)
	if err != nil {
		return err
	}
	return nil
}
