package gosip

type RouteGroup struct {
	engine *Engine
}

// handle 统一注册请求
func (group *RouteGroup) handle(method string, c callback) {
	group.engine.addMethod(method, c)
}

// Register 请求方法
func (group *RouteGroup) Register(c callback) {
	group.handle(RegisterMethod, c)
}

// Message 请求方法
func (group *RouteGroup) Message(c callback) {
	group.handle(MessageMethod, c)
}

// Invite 请求方法
func (group *RouteGroup) Invite(c callback) {
	group.handle(InviteMethod, c)
}

// Ack 请求方法
func (group *RouteGroup) Ack(c callback) {
	group.handle(AckMethod, c)
}

// Cancel 请求方法
func (group *RouteGroup) Cancel(c callback) {
	group.handle(CancelMethod, c)
}

// Options 请求方法
func (group *RouteGroup) Options(c callback) {
	group.handle(OptionsMethod, c)
}

// Bye 请求方法
func (group *RouteGroup) Bye(c callback) {
	group.handle(ByeMethod, c)
}
