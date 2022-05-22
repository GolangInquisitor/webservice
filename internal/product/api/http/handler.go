package http

import (
	http2 "Scoltest/internal/app/api/http"
	"Scoltest/internal/domain/product"
	"Scoltest/pkg/loger"
	"Scoltest/pkg/utils/httputil"
	"context"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	productRoute  = "/product"
	productsRoute = "/products/:uuid"
)

type handler struct {
	logger  *loger.Logger
	service product.Service
}

func NewHandler(loger *loger.Logger, service product.Service) http2.Handler {
	return &handler{logger: loger, service: service}
}
func (h *handler) Register(router *httprouter.Router) {
	h.logger.Infoln("Registry routs")
	router.Handle(http.MethodPost, productRoute, h.Create)
	router.Handle(http.MethodPut, productsRoute, h.Update)
	router.Handle(http.MethodDelete, productsRoute, h.Delete)

}

func (h *handler) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		h.logger.Error("Error parse form ", err.Error())
		http.Error(w, "Error parse post ", http.StatusBadRequest)
		return
	}

	prod := product.Product{
		Description: r.FormValue("description"),
		Price:       r.FormValue("price"),
		Currency:    r.FormValue("currency"),
		LeftInStock: r.FormValue("left_in_stock"),
	}

	if err := h.service.Create(context.TODO(), &prod); err != nil {
		http.Error(w, "Error create product ", http.StatusInternalServerError)
		return
	}

	if err := httputil.SendAnswer(&prod, w); err != nil {
		h.logger.Error("Error marshaling product struct: %#v ", prod)
	}
}
func (h *handler) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	if err := r.ParseForm(); err != nil {
		h.logger.Error("Error parse form ", err.Error())
		http.Error(w, "Error parse put  ", http.StatusBadRequest)
		return
	}

	prod := product.Product{
		Uuid:        params.ByName("uuid"),
		Description: r.FormValue("description"),
		Price:       r.FormValue("price"),
		Currency:    r.FormValue("currency"),
		LeftInStock: r.FormValue("left_in_stock"),
	}

	if err := h.service.Update(context.TODO(), &prod); err != nil {
		http.Error(w, "Error update product ", http.StatusInternalServerError)
		return
	}

	if err := httputil.SendAnswer(&prod, w); err != nil {
		h.logger.Error("Error marshaling product struct: %#v ", prod)
	}

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
