package user

// swagger:parameters
type LoginRequest struct {
	// in: body
	Body LoginReq
}

// LoginRes
//
// swagger:response LoginResponseWrapper
type LoginResponseWrapper struct {
	// in: body
	Body LoginRes
}

// swagger:route POST /login users LoginResponseWrapper
//
// 用户登录
//
// 用户登录
//
//     Responses:
//       200: LoginResponseWrapper
