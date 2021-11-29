package handlers

import "github.com/DanielDiaz18/rappi-cut/backend/data"

type ServerHandler struct {
	r *data.UsersRepo
	p *data.ParnerRepo
	o *data.OrderRepo
}

func NewServerHandler(r *data.UsersRepo, p *data.ParnerRepo, o *data.OrderRepo) *ServerHandler {
	return &ServerHandler{r, p, o}
}
