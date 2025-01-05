package handlers

import "forum/server/internal/services"

type HandlerLayer struct {
	HandlerDB services.ServiceLayer
}
