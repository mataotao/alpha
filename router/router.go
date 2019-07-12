package router

import (
	"alpha/handler/admin/permission"
	"alpha/handler/admin/role"
	"alpha/handler/sd"
	"alpha/router/middleware"
	"alpha/router/middleware/limiter"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(mw...)
	pprof.Register(g)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	admin := g.Group("/admin/")
	admin.Use(middleware.RequestId())
	admin.Use(limiter.TBIP())
	{
		//新增权限
		admin.POST("permission", permission.Create)
		//修改权限
		admin.PUT("permission/:id", permission.Update)
		//删除权限
		admin.DELETE("permission/:id", permission.Delete)
		//删除权限
		admin.GET("permission/:id", permission.Get)
		//列表
		admin.GET("permission", permission.List)

		//新增角色
		admin.POST("role", role.Create)
		//删除角色
		admin.DELETE("role/:id", role.Delete)
		//修改角色
		admin.PUT("role/:id", role.Update)
		//获取角色详情
		admin.GET("role/:id", role.Get)
	}
	// The health check handlers
	svcd := g.Group("/sd")
	svcd.Use(limiter.TBIP())
	{
		svcd.GET("/health", sd.HealthCheck)
		svcd.GET("/disk", sd.DiskCheck)
		svcd.GET("/cpu", sd.CPUCheck)
		svcd.GET("/ram", sd.RAMCheck)
	}

	return g
}
