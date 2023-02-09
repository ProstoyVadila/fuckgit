package main

func (a *App) setRoutes() {
	a.router.GET("/ping", a.ping)
	a.router.GET("/questions", a.questions)
}
