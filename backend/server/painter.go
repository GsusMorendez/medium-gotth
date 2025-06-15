package server

import (
	"context"
	"log"
	"net/http"
	"version1-medium-gotth/backend/model"
	"version1-medium-gotth/frontend/component"
)

type Painter struct {
}

func NewPainter() *Painter {
	return &Painter{}
}

func (p *Painter) HelloWorld(w http.ResponseWriter, ctx context.Context) {
	if err := component.HelloWorl("Gophers").Render(ctx, w); err != nil {
		log.Println(err)
	}
}

func (p *Painter) MainPage(w http.ResponseWriter, ctx context.Context, data model.PageData) {
	if err := component.CompletePage(data).Render(ctx, w); err != nil {
		log.Println(err)
	}
}
