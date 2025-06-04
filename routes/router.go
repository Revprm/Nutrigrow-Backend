package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func RegisterRoutes(server *gin.Engine, injector *do.Injector) {
	User(server, injector)
	BahanMakanan(server, injector)
	Stunting(server, injector)
	Berita(server, injector)
	KategoriBerita(server, injector)
	Makanan(server, injector)
}