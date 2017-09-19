package handlers

import (
	"gorm-tutorial/src/services"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// ActorContext ...
type ActorContext struct {
	// 埋め込み
	*BaseContext
	Service *services.ActorService
}

// NewActorContext ...
// ActorContextのコンストラクタ
func NewActorContext(c echo.Context) *ActorContext {
	return &ActorContext{&BaseContext{c}, services.NewActorService("db_config.json")}
}

// GetActorID ...
// Pathパラメータのactor_idを取得する
func (c *ActorContext) GetActorID() int {
	actorID, err := strconv.Atoi(c.Param("actor_id"))
	if err != nil {
		panic(err)
	}
	return actorID
}

// GetRequestActor ...
// RequestActorを取得する
func (c *ActorContext) GetRequestActor() *RequestActor {
	req := RequestActor{}
	c.Bind(&req)
	return &req
}

// ActorIndex ...
// GET /api/actors
func ActorIndex() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する
		ct := NewActorContext(c)
		actors := ct.Service.GetActors()
		return ct.JSON(http.StatusOK, actors)
	}
}

// ActorShow ...
// GET /api/actors/:actor_id
func ActorShow() echo.HandlerFunc {
	return func(c echo.Context) error {
		ct := NewActorContext(c)
		actorID := ct.GetActorID()
		actor := ct.Service.GetActor(actorID)
		return ct.JSON(http.StatusOK, actor)
	}
}

// ActorCreate ...
// POST /api/actors
func ActorCreate() echo.HandlerFunc {
	return func(c echo.Context) error {
		ct := NewActorContext(c)
		req := ct.GetRequestActor()
		actor := ct.Service.CreateActor(req.FirstName, req.LastName)
		return ct.JSON(http.StatusOK, actor)
	}
}

// RequestActor ...
type RequestActor struct {
	FirstName string `validate:"required" json:"first_name" xml:"first_name" form:"first_name" query:"first_name"`
	LastName  string `validate:"required" json:"last_name" xml:"last_name" form:"last_name" query:"last_name"`
}

// ActorUpdate ...
// PUT /api/actors/:actor_id
func ActorUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {
		ct := NewActorContext(c)
		actorID := ct.GetActorID()
		req := ct.GetRequestActor()
		actor := ct.Service.UpdateActor(actorID, req.FirstName, req.LastName)
		return ct.JSON(http.StatusOK, actor)
	}
}

// ActorDelete ...
// DELETE /api/actors/:actor_id
func ActorDelete() echo.HandlerFunc {
	return func(c echo.Context) error {
		ct := NewActorContext(c)
		ct.Service.DeleteActor(ct.GetActorID())
		return ct.NoContent(http.StatusOK)
	}
}
