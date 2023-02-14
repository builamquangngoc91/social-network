// +build wireinject

package carrier

import (
	"github.com/google/wire"
)

var WireSet = wire.NewSet(
	NewShipmentManager,
)
