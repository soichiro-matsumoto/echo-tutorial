package handlers

import (
	"echo-tutorial/src/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// ActorContext ...
type ActorContext struct {
	// 埋め込み
	*BaseContext
}

// NewActorContext ...
// ActorContextのコンストラクタ
func NewActorContext(c echo.Context) *ActorContext {
	return &ActorContext{&BaseContext{c}}
}

// GetActorID ...
// JSONのパラメータのactor_idを取得する
func (c *ActorContext) GetActorID() (int, error) {
	return strconv.Atoi(c.Param("actor_id"))
}

// ActorIndex ...
// GET /api/actors
func ActorIndex() echo.HandlerFunc {
	return func(c echo.Context) error { //c をいじって Request, Responseを色々する

		context := NewActorContext(c)

		db, _ := context.GormOpen()
		defer db.Close()

		actors := []models.Actor{}
		db.Find(&actors)

		return c.JSON(http.StatusOK, actors)
	}
}

// ActorShow ...
// GET /api/actors/:actor_id
func ActorShow() echo.HandlerFunc {
	return func(c echo.Context) error {

		context := NewActorContext(c)

		db, _ := context.GormOpen()
		defer db.Close()

		actor := new(models.Actor)
		actor.ActorID, _ = context.GetActorID()

		db.First(&actor)

		return c.JSON(http.StatusOK, actor)
	}
}

// ActorCreate ...
// POST /api/actors
func ActorCreate() echo.HandlerFunc {
	return func(c echo.Context) error {

		context := NewActorContext(c)

		request := new(RequestActor)
		context.Bind(request)
		err := context.Validate(request)
		if err != nil {
			return err
		}

		db, _ := context.GormOpen()
		defer db.Close()

		actor := new(models.Actor)
		actor.FirstName = request.FirstName
		actor.LastName = request.LastName

		db.Create(&actor)

		return c.JSON(http.StatusOK, actor)
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

		context := NewActorContext(c)
		request := new(RequestActor)
		context.Bind(request)

		db, _ := context.GormOpen()
		defer db.Close()

		actor := new(models.Actor)
		actor.ActorID, _ = context.GetActorID()
		db.Find(&actor)

		actor.FirstName = request.FirstName
		actor.LastName = request.LastName

		db.Save(&actor)

		return c.JSON(http.StatusOK, actor)
	}
}

// ActorDelete ...
// DELETE /api/actors/:actor_id
func ActorDelete() echo.HandlerFunc {
	return func(c echo.Context) error {

		context := c.(*ActorContext)

		db, _ := context.GormOpen()
		defer db.Close()

		actor := new(models.Actor)
		actor.ActorID, _ = context.GetActorID()
		db.Find(&actor)

		db.Delete(&actor)

		return c.JSON(http.StatusOK, actor)
	}
}
