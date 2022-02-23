package router

import (
	"github.com/CaoShenZhou/Blog4Go/internal/routers/public"
	"github.com/CaoShenZhou/Blog4Go/internal/routers/user"
)

var (
	Public public.PublicRouter
	User   user.UserRouter
)
