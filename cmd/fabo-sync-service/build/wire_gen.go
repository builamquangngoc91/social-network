// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package build

import (
	"context"
	"o.o/backend/cmd/fabo-sync-service/config"
	"o.o/backend/com/fabo/main/fbmessaging"
	"o.o/backend/com/fabo/main/fbpage"
	"o.o/backend/com/fabo/main/fbuser"
	"o.o/backend/com/fabo/pkg/fbclient"
	redis2 "o.o/backend/com/fabo/pkg/redis"
	"o.o/backend/com/fabo/pkg/sync"
	"o.o/backend/com/main"
	"o.o/backend/com/shopping/customering/query"
	"o.o/backend/pkg/common/apifw/health"
	"o.o/backend/pkg/common/bus"
	"o.o/backend/pkg/common/redis"
)

// Injectors from wire.go:

func Build(ctx context.Context, cfg config.Config) (Output, func(), error) {
	redisRedis := cfg.Redis
	store := redis.Connect(redisRedis)
	service := health.New(store)
	v := BuildServers(cfg, service)
	databases := cfg.Databases
	mainDB, err := com.BuildDatabaseMain(databases)
	if err != nil {
		return Output{}, nil, err
	}
	appConfig := cfg.FacebookApp
	fbClient := fbclient.New(appConfig)
	busBus := bus.New()
	fbExternalMessagingAggregate := fbmessaging.NewFbExternalMessagingAggregate(mainDB, busBus, fbClient)
	commandBus := fbmessaging.FbExternalMessagingAggregateMessageBus(fbExternalMessagingAggregate)
	faboRedis := redis2.NewFaboRedis(store)
	fbMessagingQuery := fbmessaging.NewFbMessagingQuery(mainDB, faboRedis)
	queryBus := fbmessaging.FbMessagingQueryMessageBus(fbMessagingQuery)
	fbPageUtil := fbpage.NewFbPageUtil(store)
	fbExternalPageAggregate := fbpage.NewFbPageAggregate(mainDB, fbPageUtil, busBus)
	fbpagingCommandBus := fbpage.FbExternalPageAggregateMessageBus(fbExternalPageAggregate)
	customerQuery := query.NewCustomerQuery(mainDB)
	customeringQueryBus := query.CustomerQueryMessageBus(customerQuery)
	fbUserAggregate := fbuser.NewFbUserAggregate(mainDB, fbpagingCommandBus, customeringQueryBus)
	fbuseringCommandBus := fbuser.FbUserAggregateMessageBus(fbUserAggregate)
	fbUserQuery := fbuser.NewFbUserQuery(mainDB, customeringQueryBus)
	fbuseringQueryBus := fbuser.FbUserQueryMessageBus(fbUserQuery)
	syncConfig := cfg.SyncConfig
	synchronizer := sync.New(mainDB, fbClient, commandBus, queryBus, fbuseringCommandBus, fbuseringQueryBus, faboRedis, syncConfig)
	fbPageQuery := fbpage.NewFbPageQuery(mainDB, fbPageUtil)
	fbpagingQueryBus := fbpage.FbPageQueryMessageBus(fbPageQuery)
	processManager := fbmessaging.NewProcessManager(busBus, queryBus, commandBus, fbpagingQueryBus, fbuseringQueryBus, fbuseringCommandBus, faboRedis)
	output := Output{
		Servers:      v,
		Sync:         synchronizer,
		Health:       service,
		_fbmessaging: processManager,
	}
	return output, func() {
	}, nil
}