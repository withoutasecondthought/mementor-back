package handler

import "mementor-back/pkg/service"

type Handler struct {
	services service.Service
}

func NewHandler(service service.Service) *Handler {
	return &Handler{
		services: service,
	}
}

func InitRoutes() {

}
