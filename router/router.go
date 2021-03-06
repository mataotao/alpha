package router

import (
	"alpha/handler/admin/login"
	"alpha/handler/admin/permission"
	"alpha/handler/admin/role"
	"alpha/handler/admin/user"
	"alpha/handler/global/cache"
	"alpha/handler/global/upload"
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
	g.Use(middleware.RequestId())

	pprof.Register(g)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	g.POST("/login", login.In)
	g.POST("/upload", upload.Upload)

	global := g.Group("/global")
	{
		global.GET("cache/permission", cache.Permission)
	}
	admin := g.Group("/admin/")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(limiter.AdminRedisCell())
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
		//获取角色列表
		admin.GET("role", role.List)

		admin.GET("role-a", role.All)

		//新增用户
		admin.POST("user", user.Create)
		//更新
		admin.PUT("user/:id", user.Update)
		//更新密码
		admin.PUT("user-pwd/:id", user.UpdatePwd)
		//改变状态
		admin.PUT("user-status/:id", user.ChangeStatus)
		//获取用户详情
		admin.GET("user/:id", user.Get)
		//获取用户列表
		admin.GET("user", user.List)

		admin.GET("user-info", user.Information)
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
