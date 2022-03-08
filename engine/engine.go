package engine

import (
	"github.com/brolyssjl/game_api/models"
)

type Engine struct {
	DB models.Spec
}

func NewEngine(db models.Spec) Spec {
	return &Engine{
		DB: db,
	}
}
