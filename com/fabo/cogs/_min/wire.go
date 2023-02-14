// +build wireinject

package _min

import (
	"github.com/google/wire"
	"o.o/backend/com/fabo/pkg/fbclient"
	"o.o/backend/com/fabo/pkg/redis"
)

var WireSet = wire.NewSet(
	fbclient.New,
	redis.NewFaboRedis,
)

