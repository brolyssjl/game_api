package handlers

import "github.com/brolyssjl/game_api/engine"

type Handler struct {
	Engine engine.Spec
}

func NewHandler(e engine.Spec) Handler {
	return Handler{Engine: e}
}
