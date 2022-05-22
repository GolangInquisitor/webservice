package http

import (
	http2 "Scoltest/internal/app/api/http"
	"Scoltest/internal/domain/order"
	"Scoltest/pkg/loger"
	"Scoltest/pkg/utils/httputil"
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	userRoute  = "/order/:uuid"
	usersRoute = "/order/:uuid"
)

type handler struct {
	logger  *loger.Logger
	service order.Service
}

func NewHandler(loger *loger.Logger, service order.Service) http2.Handler {
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

	userOrder := order.Order{
		Product: r.FormValue("product"),
	}

	if err := h.service.Create(context.TODO(), params.ByName("uuid"), &userOrder); err != nil {
		http.Error(w, "Error create order ", http.StatusInternalServerError)
		return
	}

	if err := httputil.SendAnswer(&userOrder, w); err != nil {
		h.logger.Error("Error marshaling order struct: %#v ", userOrder)
	}

	return
}
func (h *handler) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		h.logger.Error("Error parse form ", err.Error())
		http.Error(w, "Error parse put  ", http.StatusBadRequest)
		return
	}

	userOrder := order.Order{
		Product: r.FormValue("product"),
		Id:      r.FormValue("id"),
		Uuid:    r.FormValue("uuid"),
	}

	if err := h.service.Update(context.TODO(), params.ByName("uuid"), &userOrder); err != nil {
		http.Error(w, "Error update order ", http.StatusInternalServerError)
		return
	}

	if err := httputil.SendAnswer(&userOrder, w); err != nil {
		h.logger.Error("Error marshaling userOrder struct: %#v ", userOrder)
	}
	return
}
func (h *handler) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		h.logger.Error("Error parse form ", err.Error())
		http.Error(w, "Error parse delete ", http.StatusBadRequest)
		return
	}
	uuid := params.ByName("uuid")
	orderId := r.FormValue("id")
	if err := h.service.Delete(context.TODO(), uuid, orderId); err != nil {
		http.Error(w, "Error update user ", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(204)
	w.Write([]byte(fmt.Sprintf("This id api  %s", "Delete Order")))
	return
}
