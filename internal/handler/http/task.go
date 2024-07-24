package http

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"hl3-projectmanagement/internal/domain/task"
	"hl3-projectmanagement/internal/service/management"
	"hl3-projectmanagement/pkg/server/response"
	"net/http"
)

type TaskHandler struct {
	service *management.Service
}

func NewTaskHandler(s *management.Service) *TaskHandler {
	return &TaskHandler{s}
}

func (h *TaskHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.create)
	r.Get("/", h.list)
	r.Get("/search", h.search)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.getByID)
		r.Put("/", h.update)
		r.Delete("/", h.delete)
	})
	return r
}

//	@Summary	Create a task
//	@tags		task
//	@accept		json
//	@produce	json
//	@param		request	body	task.Request	true	"body"
//	@Success	201
//	@Failure	400	{object}	response.Object
//	@Failure	500	{object}	response.Object
//	@Router		/task [post]
func (h *TaskHandler) create(w http.ResponseWriter, r *http.Request) {
	req := task.Request{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}
	exists := h.service.UserExists(req.ResponsibleID)
	if !exists {
		response.BadRequest(w, r, errors.New("assignee_id: user doesn't exist"), req)
		return
	}

	// write check if project exists

	if err := h.service.CreateTask(req); err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.Created(w, r, "success")
}

//	@Summary	List tasks
//	@tags		task
//	@accept		json
//	@produce	json
//	@Success	200	{array}		response.Object
//	@Failure	500	{object}	response.Object
//	@Router		/task [get]
func (h *TaskHandler) list(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.ListTasks()
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

//	@Summary	Get task by id
//	@tags		task
//	@accept		json
//	@produce	json
//	@param		id	path		string	true	"task id"
//	@Success	200	{object}	response.Object
//	@Failure	404	{object}	response.Object
//	@Failure	500	{object}	response.Object
//	@Router		/task/{id} [get]
func (h *TaskHandler) getByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	res, err := h.service.GetTaskByID(id)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

//	@Summary	Update task by id
//	@tags		task
//	@accept		json
//	@produce	json
//	@param		id		path	string			true	"task id"
//	@param		request	body	task.Request	true	"body"
//	@Success	200
//	@Failure	400	{object}	response.Object
//	@Failure	500	{object}	response.Object
//	@Router		/task/{id} [put]
func (h *TaskHandler) update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	req := task.Request{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	if err := h.service.UpdateTask(id, req); err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, "success")
}

//	@Summary	Delete task by id
//	@tags		task
//	@accept		json
//	@produce	json
//	@param		id	path	string	true	"task id"
//	@Success	200
//	@Failure	500	{object}	response.Object
//	@Router		/task/{id} [delete]
func (h *TaskHandler) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.service.DeleteTask(id); err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, "success")
}

//	@Summary	Search tasks
//	@tags		task
//	@accept		json
//	@produce	json
//	@param		title			query		string	false	"title"
//	@param		status			query		string	false	"status"
//	@param		priority		query		string	false	"priority"
//	@param		responsibleID	query		string	false	"responsible_id"
//	@param		projectID		query		string	false	"project_id"
//	@Success	200				{object}	response.Object
//	@Failure	500				{object}	response.Object
//	@Router		/task/search [get]
func (h *TaskHandler) search(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	status := r.FormValue("status")
	priority := r.FormValue("priority")
	responsibleID := r.FormValue("responsibleID")
	projectID := r.FormValue("projectID")

	res, err := h.service.SearchTask(title, status, priority, responsibleID, projectID)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}
