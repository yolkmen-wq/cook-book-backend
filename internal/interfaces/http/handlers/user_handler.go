package handlers

import (
	"cook-book-backend/internal/domain/services"
	"cook-book-backend/internal/infrastructure/logger"
)

type UserHandler struct {
	*BaseHandler
	userService services.UserService
}

func NewUserHandler(userService services.UserService, logger logger.Logger) *UserHandler {
	return &UserHandler{
		BaseHandler: NewBaseHandler(logger),
		userService: userService,
	}
}

// // Login 用户登录
// func (h *UserHandler) Login(c *gin.Context) {
//     var req dto.LoginRequest

//     if err := h.BindAndValidate(c, &req); err != nil {
//         h.HandleError(c, err, "bind_login_params")
//         return
//     }

//     // 获取客户端信息
//     clientInfo := h.extractClientInfo(c)

//     loginResp, err := h.userService.Login(&req, clientInfo)
//     if err != nil {
//         h.HandleError(c, err, "user_login")
//         return
//     }

//     h.Success(c, loginResp)
// }

// // AdminLogin 管理员登录
// func (h *UserHandler) AdminLogin(c *gin.Context) {
//     var req dto.AdminLoginRequest

//     if err := h.BindAndValidate(c, &req); err != nil {
//         h.HandleError(c, err, "bind_admin_login_params")
//         return
//     }

//     clientInfo := h.extractClientInfo(c)

//     loginResp, err := h.userService.AdminLogin(&req, clientInfo)
//     if err != nil {
//         h.HandleError(c, err, "admin_login")
//         return
//     }

//     h.Success(c, loginResp)
// }

// // Logout 用户登出
// func (h *UserHandler) Logout(c *gin.Context) {
//     userID := h.getUserIDFromContext(c)

//     if err := h.userService.Logout(userID); err != nil {
//         h.HandleError(c, err, "user_logout")
//         return
//     }

//     h.Success(c, gin.H{"message": "登出成功"})
// }

// // RefreshToken 刷新令牌
// func (h *UserHandler) RefreshToken(c *gin.Context) {
//     var req dto.RefreshTokenRequest

//     if err := h.BindAndValidate(c, &req); err != nil {
//         h.HandleError(c, err, "bind_refresh_token_params")
//         return
//     }

//     tokenResp, err := h.userService.RefreshToken(req.RefreshToken)
//     if err != nil {
//         h.HandleError(c, err, "refresh_token")
//         return
//     }

//     h.Success(c, tokenResp)
// }

// // GetProfile 获取用户信息
// func (h *UserHandler) GetProfile(c *gin.Context) {
//     userID := h.getUserIDFromContext(c)

//     profile, err := h.userService.GetProfile(userID)
//     if err != nil {
//         h.HandleError(c, err, "get_user_profile")
//         return
//     }

//     h.Success(c, profile)
// }

// // UpdateProfile 更新用户信息
// func (h *UserHandler) UpdateProfile(c *gin.Context) {
//     userID := h.getUserIDFromContext(c)

//     var req dto.UpdateProfileRequest
//     if err := h.BindAndValidate(c, &req); err != nil {
//         h.HandleError(c, err, "bind_update_profile_params")
//         return
//     }

//     profile, err := h.userService.UpdateProfile(userID, &req)
//     if err != nil {
//         h.HandleError(c, err, "update_user_profile")
//         return
//     }

//     h.Success(c, profile)
// }

// // GetAsyncRoutes 获取异步路由
// func (h *UserHandler) GetAsyncRoutes(c *gin.Context) {
//     userID := h.getUserIDFromContext(c)

//     routes, err := h.userService.GetAsyncRoutes(userID)
//     if err != nil {
//         h.HandleError(c, err, "get_async_routes")
//         return
//     }

//     h.Success(c, routes)
// }

// // 辅助方法
// func (h *UserHandler) extractClientInfo(c *gin.Context) *dto.ClientInfo {
//     return &dto.ClientInfo{
//         IP:        c.ClientIP(),
//         UserAgent: c.GetHeader("User-Agent"),
//         // 可以添加更多客户端信息提取逻辑
//     }
// }

// func (h *UserHandler) getUserIDFromContext(c *gin.Context) int64 {
//     // 从JWT中间件设置的上下文中获取用户ID
//     if userID, exists := c.Get("userID"); exists {
//         if id, ok := userID.(int64); ok {
//             return id
//         }
//     }
//     return 0
// }
