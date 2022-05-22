package http

import (
	http2 "Scoltest/internal/app/api/http"
	"Scoltest/internal/domain/user"
	"Scoltest/pkg/loger"
	"Scoltest/pkg/utils/httputil"
	"context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	userRoute  = "/user"
	usersRoute = "/user/:uuid"
)

type handler struct {
	logger  *loger.Logger
	service user.Service
}

func NewHandler(loger *loger.Logger, service user.Service) http2.Handler {
	return &handler{logger: loger, service: service}
}
func (h *handler) Register(router *httprouter.Router) {
	h.logger.Infoln("Registry routs")
	router.Handle(http.MethodPost, userRoute, h.Create)
	router.Handle(http.MethodPut, usersRoute, h.Update)
	router.Handle(http.MethodDelete, usersRoute, h.Delete)

}

func (h *handler) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		h.logger.Error("Error parse form ", err.Error())
		http.Error(w, "Error parse post ", http.StatusBadRequest)
		return
	}
	user := user.User{
		Name:       r.FormValue("name"),
		Surname:    r.FormValue("surname"),
		MiddleName: r.FormValue("midlename"),
		Gender:     r.FormValue("gender"),
		Age:        r.FormValue("age"),
	}

	if err := h.service.Create(context.TODO(), &user); err != nil {
		http.Error(w, "Error create user ", http.StatusInternalServerError)
		return
	}

	if err := httputil.SendAnswer(&user, w); err != nil {
		h.logger.Error("Error marshaling user struct: %#v ", user)
		//	http.Error(w, "Error create response ", http.StatusInternalServerError)

	}

	return
}
func (h *handler) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		h.logger.Error("Error parse form ", err.Error())
		http.Error(w, "Error parse put ", http.StatusBadRequest)
		return
	}

	user := user.User{
		Uuid:       params.ByName("uuid"),
		Name:       r.FormValue("name"),
		Surname:    r.FormValue("surname"),
		MiddleName: r.FormValue("midlename"),
		Gender:     r.FormValue("gender"),
		Age:        r.FormValue("age"),
	}

	if err := h.service.Update(context.TODO(), &user); err != nil {
		http.Error(w, "Error update user ", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(204)
	return
}
func (h *handler) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		h.logger.Error("Error parse form ", err.Error())
		http.Error(w, "Error parse delete ", http.StatusBadRequest)
		return
	}
	uuid := params.ByName("uuid")
	if err := h.service.Delete(context.TODO(), uuid); err != nil {
		http.Error(w, "Error update user ", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(204)
	return
}
