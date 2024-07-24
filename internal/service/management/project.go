package management

import (
	"github.com/google/uuid"
	"hl3-projectmanagement/internal/domain/project"
	"hl3-projectmanagement/internal/domain/task"
	"time"
)

func (s *Service) CreateProject(req project.Request) (err error) {
	startedat, _ := time.Parse(time.DateOnly, req.StartedAt)
	endat, _ := time.Parse(time.DateOnly, req.EndAt)
	project := project.Entity{
		ID:          uuid.New().String(),
		Title:       req.Title,
		Description: req.Description,
		StartDate:   startedat,
		EndDate:     endat,
		ManagerID:   req.ManagerID,
	}

	err = s.ProjectRepository.CreateProject(project)

	return
}

func (s *Service) ListProjects() (projects []project.Entity, err error) {
	projects, err = s.ProjectRepository.GetProjects()
	return
}

func (s *Service) GetProject(id string) (project project.Entity, err error) {
	project, err = s.ProjectRepository.GetProject(id)
	return
}

func (s *Service) UpdateProject(id string, req project.Request) (err error) {
	startedat, _ := time.Parse(time.DateOnly, req.StartedAt)
	endat, _ := time.Parse(time.DateOnly, req.EndAt)
	project := project.Entity{
		Title:       req.Title,
		Description: req.Description,
		StartDate:   startedat,
		EndDate:     endat,
		ManagerID:   req.ManagerID,
	}

	err = s.ProjectRepository.UpdateProject(id, project)
	return
}

func (s *Service) DeleteProject(id string) (err error) {
	err = s.ProjectRepository.DeleteProject(id)
	return
}

func (s *Service) SearchProjectTasks(id string) (tasks []task.Entity, err error) {
	tasks, err = s.SearchTask("", "", "", "", id)
	return
}

func (s *Service) SearchProject(title, managerID string) (tasks []project.Entity, err error) {
	params := map[string]string{
		"title":      title,
		"manager_id": managerID,
	}
	tasks, err = s.ProjectRepository.SearchProject(params)

	return
}
