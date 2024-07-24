package management

import (
	"github.com/google/uuid"
	"hl3-projectmanagement/internal/domain/task"
	"time"
)

func (s *Service) GetUserTasks(id string) (tasks []task.Entity, err error) {
	tasks, err = s.TaskRepository.GetUserTasks(id)
	return
}

func (s *Service) CreateTask(req task.Request) (err error) {
	completionDate, _ := time.Parse(time.DateOnly, req.CompletionAt)
	task := task.Entity{
		ID:             uuid.New().String(),
		Title:          req.Title,
		Description:    req.Description,
		Priority:       req.Priority,
		Status:         req.Status,
		ResponsibleID:  req.ResponsibleID,
		ProjectID:      req.ProjectID,
		CompletionDate: completionDate,
	}

	err = s.TaskRepository.CreateTask(task)

	return
}

func (s *Service) ListTasks() (tasks []task.Entity, err error) {
	tasks, err = s.TaskRepository.ListTasks()
	return
}

func (s *Service) GetTaskByID(id string) (task task.Entity, err error) {
	task, err = s.TaskRepository.GetTask(id)
	return
}

func (s *Service) UpdateTask(id string, req task.Request) (err error) {
	completionDate, _ := time.Parse(time.DateOnly, req.CompletionAt)
	task := task.Entity{
		Title:          req.Title,
		Description:    req.Description,
		Priority:       req.Priority,
		Status:         req.Status,
		ResponsibleID:  req.ResponsibleID,
		ProjectID:      req.ProjectID,
		CompletionDate: completionDate,
	}

	err = s.TaskRepository.UpdateTask(id, task)

	return
}

func (s *Service) DeleteTask(id string) (err error) {
	err = s.TaskRepository.DeleteTask(id)
	return
}

func (h *Service) SearchTask(title, status, priority, responsibleID, projectID string) (tasks []task.Entity, err error) {
	args := map[string]string{
		"title":          title,
		"status":         status,
		"priority":       priority,
		"responsible_id": responsibleID,
		"project_id":     projectID,
	}

	tasks, err = h.TaskRepository.Search(args)

	return
}
