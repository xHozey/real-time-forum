package services

import (
	middleware "forum/server/internal/middleWare"
)

type ServiceLayer struct {
	ServiceDB middleware.MiddleWareLayer
}
