// +build wireinject

package sender

import "github.com/google/wire"

var WireSet = wire.NewSet(
	New,
)
