package handlers

import (
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
