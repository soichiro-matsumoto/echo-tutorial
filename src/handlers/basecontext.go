package handlers

import (
	"echo-tutorial/src/config"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

// BaseContext ...
// echo.ContextをEmbed
type BaseContext struct {
	// 埋め込み
	echo.Context
}

// NewBaseContext ...
// BaseContextのコンストラクタ
func NewBaseContext(c echo.Context) *BaseContext {
	return &BaseContext{c}
}

// GormOpen ...
// GormDBを取得する
func (c *BaseContext) GormOpen() (*gorm.DB, error) {
	return config.Config.GormOpen()
}
