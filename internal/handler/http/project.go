package http

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"hl3-projectmanagement/internal/domain/project"
	"hl3-projectmanagement/internal/service/management"
	"hl3-projectmanagement/pkg/server/response"
	"net/http"
)

type ProjectHandler struct {
	service *management.Service
}

func NewProjectHandler(s *management.Service) *ProjectHandler {
	return &ProjectHandler{service: s}
}

func (h *ProjectHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.create)
	r.Get("/", h.list)
	r.Get("/search", h.search)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.getByID)
		r.Put("/", h.update)
		r.Delete("/", h.delete)
		r.Get("/tasks", h.getTasks)
	})

	return r
}

//	@Summary	Create a project
//	@tags		project
//	@accept		json
//	@produce	json
//	@param		request	body	project.Request	true	"body"
//	@Success	201
//	@Failure	400	{object}	response.Object
//	@Failure	500	{object}	response.Object
//	@Router		/project [post]
func (h *ProjectHandler) create(w http.ResponseWriter, r *http.Request) {
	req := project.Request{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	exists := h.service.UserExists(req.ManagerID)
	if !exists {
		response.BadRequest(w, r, errors.New("manager_id: user doesn't exist"), req)
		return
	}

	if err := h.service.CreateProject(req); err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.Created(w, r, "success")
}

//	@Summary	Get all projects
//	@tags		project
//	@accept		json
//	@produce	json
//	@Success	200	{array}		response.Object
//	@Failure	500	{object}	response.Object
//	@Router		/project [get]
func (h *ProjectHandler) list(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.ListProjects()
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

//	@Summary	Get project by id
//	@tags		project
//	@accept		json
//	@produce	json
//	@param		id	path		string	true	"project id"
//	@Success	200	{object}	response.Object
//	@Failure	404	{object}	response.Object
//	@Failure	500	{object}	response.Object
//	@Router		/project/{id} [get]
func (h *ProjectHandler) getByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	res, err := h.service.GetProject(id)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

//	@Summary	Update project by id
//	@tags		project
//	@accept		json
//	@produce	json
//	@param		id		path	string			true	"project id"
//	@param		request	body	project.Request	true	"body"
//	@Success	200
//	@Failure	400	{object}	response.Object
//	@Failure	500	{object}	response.Object
//	@Router		/project/{id} [put]
func (h *ProjectHandler) update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	req := project.Request{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	if err := h.service.UpdateProject(id, req); err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, "success")
}

//	@Summary	Delete project by id
//	@tags		project
//	@accept		json
//	@produce	json
//	@param		id	path	string	true	"project id"
//	@Success	200
//	@Failure	500	{object}	response.Object
//	@Router		/project/{id} [delete]
func (h *ProjectHandler) delete(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if err := h.service.DeleteProject(id); err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, "success")
}

//	@Summary	Get tasks by project id
//	@tags		project
//	@accept		json
//	@produce	json
//	@param		id	path		string	true	"project id"
//	@Success	200	{array}		response.Object
//	@Failure	500	{object}	response.Object
//	@Router		/project/{id}/tasks [get]
func (h *ProjectHandler) getTasks(w http.ResponseWriter, r *http.Request) {
	projectID := chi.URLParam(r, "id")
	res, err := h.service.SearchProjectTasks(projectID)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}

//	@Summary	Search project by param
//	@tags		project
//	@accept		json
//	@produce	json
//	@param		title		query		string	false	"title"
//	@param		manager_id	query		string	false	"manager_id"
//	@Success	200			{array}		response.Object
//	@Failure	500			{object}	response.Object
//	@Router		/project/search [get]
func (h *ProjectHandler) search(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	managerID := r.URL.Query().Get("manager_id")

	res, err := h.service.SearchProject(title, managerID)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)
}
