package routerfx

import (
	"go.uber.org/fx"
	testfx "master/internal/di/routerfx/testfx"
	userfx "master/internal/di/routerfx/userfx"
)

var Module = fx.Options(
	userfx.Module,
	testfx.Module,
)
