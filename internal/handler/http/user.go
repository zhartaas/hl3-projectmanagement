package http

import (
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"hl3-projectmanagement/internal/domain/user"
	"hl3-projectmanagement/internal/service/management"
	"hl3-projectmanagement/pkg/server/response"
	"net/http"
)

type UserHandler struct {
	service *management.Service
}

func NewUserHandler(s *management.Service) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", h.create)
	r.Get("/", h.list)
	r.Get("/search", h.search)

	r.Route("/{id}", func(r chi.Router) {
		r.Get("/", h.getById)
		r.Put("/", h.update)
		r.Delete("/", h.delete)
		r.Get("/tasks", h.userTasks)
	})

	return r
}

// @Summary	create new user
// @Tags		user
// @Accept		json
// @Produce	json
// @Param		request	body	user.Request	true	"body"
// @Success	201
// @Failure	404	{object}	response.Object
// @Router		/user [post]
func (h *UserHandler) create(w http.ResponseWriter, r *http.Request) {
	req := user.Request{}
	if err := render.Bind(r, &req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	if err := h.service.CreateUser(req); err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.Created(w, r, "success")

}

// @Summary	get all users
// @Tags		user
// @Accept		json
// @Produce	json
// @Success	200	{object}	response.Object
// @Failure	404	{object}	response.Object
// @Router		/user [get]
func (h *UserHandler) list(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.GetUsers()
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, res)

}

// @Summary	get user by id
// @Tags		user
// @Accept		json
// @produce	json
// @param		id	path		string	true	"user id"
// @Success	200	{object}	response.Object
// @failure	404	{object}	response.Object
// @router		/user/{id} [get]
func (h *UserHandler) getById(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	res, err := h.service.GetUserByID(idParam)
	if err != nil {
		response.BadRequest(w, r, err, res)
		return
	}

	response.OK(w, r, res)
}

// @Summary	update user data
// @Tags		user
// @Accept		json
// @Produce	json
// @Param		request	body	user.Request	true	"body"
// @Param		id		path	string			true	"user id"
// @Success	200
// @Failure	404	{object}	response.Object
// @Router		/user/{id} [put]
func (h *UserHandler) update(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	req := &user.Request{}

	if err := render.Bind(r, req); err != nil {
		response.BadRequest(w, r, err, req)
		return
	}

	err := h.service.UpdateUser(idParam, *req)
	if err != nil {
		response.InternalServerError(w, r, err)
	}

	response.OK(w, r, "success")
}

// @Summary	delete user
// @tags		user
// @accept		json
// @produce	json
// @param		id	path	string	true	"user id"
// @Success	200
// @failure	404	{object}	response.Object
// @router		/user/{id} [delete]
func (h *UserHandler) delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	err := h.service.DeleteUser(idParam)
	if err != nil {
		response.BadRequest(w, r, err, "incorrect id")
		return
	}

	response.OK(w, r, "success")
}

// @Summary	user tasks by id
// @tags		user
// @accept		json
// @produce	json
// @param		id	path		string	true	"user id"
// @success	200	{object}	response.Object
// @failure	404	{object}	response.Object
// @failure	500	{object}	response.Object
// @router		/user/{id}/tasks [get]
func (h *UserHandler) userTasks(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")

	task, err := h.service.GetUserTasks(idParam)
	if err != nil {
		response.InternalServerError(w, r, err)
		return
	}

	response.OK(w, r, task)
}

// @summary		search by name or email
// @description	fill only one and leave second empty
// @tags			user
// @accept			json
// @produce		json
// @param			name	query		string	false	"Optional* Name"
// @param			email	query		string	false	"Optional* Email"
// @success		200		{object}	response.Object
// @failure		404		{object}	response.Object
// @failure		500		{object}	response.Object
// @router			/user/search [get]
func (h *UserHandler) search(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")

	param := ""
	var findByName bool
	if name == "" && email != "" {
		findByName = false
		param = email
	} else if name != "" && email == "" {
		findByName = true
		param = name
	} else {
		response.BadRequest(w, r, errors.New("bad request"), fmt.Sprintf("name:%s, email:%s", name, email))
		return
	}

	user, err := h.service.SearchUser(param, findByName)
	if err != nil {
		response.BadRequest(w, r, errors.New("not found"), user)
	}

	response.OK(w, r, user)
}
