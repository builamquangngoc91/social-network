// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package build

import (
	"context"
	"o.o/backend/cmd/fabo-event-handler/config"
	"o.o/backend/cogs/shipment/_fabo"
	"o.o/backend/com/eventhandler/handler/api"
	"o.o/backend/com/eventhandler/webhook/sender"
	"o.o/backend/com/eventhandler/webhook/storage"
	"o.o/backend/com/fabo/main/fbmessaging"
	"o.o/backend/com/fabo/main/fbpage"
	"o.o/backend/com/fabo/main/fbuser"
	"o.o/backend/com/fabo/pkg/fbclient"
	redis2 "o.o/backend/com/fabo/pkg/redis"
	"o.o/backend/com/fabo/pkg/webhook"
	"o.o/backend/com/main"
	"o.o/backend/com/main/connectioning/aggregate"
	"o.o/backend/com/main/connectioning/manager"
	query2 "o.o/backend/com/main/connectioning/query"
	"o.o/backend/com/main/identity"
	"o.o/backend/com/main/location"
	"o.o/backend/com/main/ordering"
	"o.o/backend/com/main/shipmentpricing/pricelist"
	"o.o/backend/com/main/shipmentpricing/pricelistpromotion"
	"o.o/backend/com/main/shipmentpricing/shipmentprice"
	"o.o/backend/com/main/shipmentpricing/shipmentservice"
	"o.o/backend/com/main/shipmentpricing/shopshipmentpricelist"
	"o.o/backend/com/main/shipping/carrier"
	query4 "o.o/backend/com/main/shipping/query"
	query3 "o.o/backend/com/main/shippingcode/query"
	"o.o/backend/com/shopping/customering/query"
	query5 "o.o/backend/com/shopping/setting/query"
	"o.o/backend/com/shopping/setting/util"
	"o.o/backend/pkg/common/apifw/health"
	"o.o/backend/pkg/common/bus"
	"o.o/backend/pkg/common/redis"
	"o.o/backend/pkg/etop/sqlstore"
)

// Injectors from wire.go:

func Build(ctx context.Context, cfg config.Config) (Output, func(), error) {
	redisRedis := cfg.Redis
	store := redis.Connect(redisRedis)
	service := health.New(store)
	miscService := &api.MiscService{}
	databases := cfg.Databases
	mainDB, err := com.BuildDatabaseMain(databases)
	if err != nil {
		return Output{}, nil, err
	}
	webhookDB, err := com.BuildDatabaseWebhook(databases)
	if err != nil {
		return Output{}, nil, err
	}
	changesStore := storage.NewChangesStore(webhookDB)
	partnerStore := sqlstore.BuildPartnerStore(mainDB)
	partnerStoreInterface := sqlstore.BindPartnerStore(partnerStore)
	webhookSender := sender.New(mainDB, store, changesStore, partnerStoreInterface)
	webhookService := &api.WebhookService{
		Sender: webhookSender,
	}
	servers := api.NewServers(miscService, webhookService)
	mainServer, err := BuildMainServer(cfg, service, servers)
	if err != nil {
		return Output{}, nil, err
	}
	v := BuildServers(mainServer)
	handler, err := BuildIntHandler(ctx, cfg)
	if err != nil {
		return Output{}, nil, err
	}
	customerQuery := query.NewCustomerQuery(mainDB)
	queryBus := query.CustomerQueryMessageBus(customerQuery)
	fbUserQuery := fbuser.NewFbUserQuery(mainDB, queryBus)
	fbuseringQueryBus := fbuser.FbUserQueryMessageBus(fbUserQuery)
	faboRedis := redis2.NewFaboRedis(store)
	fbMessagingQuery := fbmessaging.NewFbMessagingQuery(mainDB, faboRedis)
	fbmessagingQueryBus := fbmessaging.FbMessagingQueryMessageBus(fbMessagingQuery)
	fbPageUtil := fbpage.NewFbPageUtil(store)
	fbPageQuery := fbpage.NewFbPageQuery(mainDB, fbPageUtil)
	fbpagingQueryBus := fbpage.FbPageQueryMessageBus(fbPageQuery)
	queryService := ordering.NewQueryService(mainDB)
	orderingQueryBus := ordering.QueryServiceMessageBus(queryService)
	busBus := bus.New()
	locationQuery := location.New(mainDB)
	locationQueryBus := location.QueryMessageBus(locationQuery)
	identityQueryService := identity.NewQueryService(mainDB)
	identityQueryBus := identity.QueryServiceMessageBus(identityQueryService)
	mapShipmentServices := shipment_all.SupportedShipmentServices()
	connectionQuery := query2.NewConnectionQuery(mainDB, mapShipmentServices)
	connectioningQueryBus := query2.ConnectionQueryMessageBus(connectionQuery)
	connectionAggregate := aggregate.NewConnectionAggregate(mainDB, busBus)
	commandBus := aggregate.ConnectionAggregateMessageBus(connectionAggregate)
	queryQueryService := query3.NewQueryService(mainDB)
	shippingcodeQueryBus := query3.QueryServiceMessageBus(queryQueryService)
	shipmentserviceQueryService := shipmentservice.NewQueryService(mainDB, store)
	shipmentserviceQueryBus := shipmentservice.QueryServiceMessageBus(shipmentserviceQueryService)
	pricelistQueryService := pricelist.NewQueryService(mainDB, store)
	pricelistQueryBus := pricelist.QueryServiceMessageBus(pricelistQueryService)
	shopshipmentpricelistQueryService := shopshipmentpricelist.NewQueryService(mainDB, store)
	shopshipmentpricelistQueryBus := shopshipmentpricelist.QueryServiceMessageBus(shopshipmentpricelistQueryService)
	shipmentpriceQueryService := shipmentprice.NewQueryService(mainDB, store, locationQueryBus, pricelistQueryBus, shopshipmentpricelistQueryBus)
	shipmentpriceQueryBus := shipmentprice.QueryServiceMessageBus(shipmentpriceQueryService)
	pricelistpromotionQueryService := pricelistpromotion.NewQueryService(mainDB, store, locationQueryBus, identityQueryBus, shopshipmentpricelistQueryBus, pricelistQueryBus)
	pricelistpromotionQueryBus := pricelistpromotion.QueryServiceMessageBus(pricelistpromotionQueryService)
	driver := shipment_all.SupportedCarrierDriver()
	connectionManager := manager.NewConnectionManager(store, connectioningQueryBus)
	addressStore := &sqlstore.AddressStore{
		DB: mainDB,
	}
	userStore := sqlstore.BuildUserStore(mainDB)
	userStoreInterface := sqlstore.BindUserStore(userStore)
	accountStore := &sqlstore.AccountStore{
		DB:           mainDB,
		EventBus:     busBus,
		AddressStore: addressStore,
		UserStore:    userStoreInterface,
	}
	accountStoreInterface := sqlstore.BindAccountStore(accountStore)
	shopStore := &sqlstore.ShopStore{
		DB: mainDB,
	}
	shopStoreInterface := sqlstore.BindShopStore(shopStore)
	orderStore := &sqlstore.OrderStore{
		DB:           mainDB,
		LocationBus:  locationQueryBus,
		AccountStore: accountStoreInterface,
		ShopStore:    shopStoreInterface,
	}
	orderStoreInterface := sqlstore.BindOrderStore(orderStore)
	shipmentManager, err := carrier.NewShipmentManager(busBus, locationQueryBus, identityQueryBus, connectioningQueryBus, commandBus, shippingcodeQueryBus, shipmentserviceQueryBus, shipmentpriceQueryBus, pricelistpromotionQueryBus, driver, connectionManager, orderStoreInterface)
	if err != nil {
		return Output{}, nil, err
	}
	queryService2 := query4.NewQueryService(mainDB, shipmentManager, connectioningQueryBus)
	shippingQueryBus := query4.QueryServiceMessageBus(queryService2)
	appConfig := cfg.FacebookApp
	fbClient := fbclient.New(appConfig)
	shopSettingUtil := util.NewShopSettingUtil(store)
	shopSettingQuery := query5.NewShopSettingQuery(mainDB, shopSettingUtil)
	settingQueryBus := query5.ShopSettingQueryMessageBus(shopSettingQuery)
	fbExternalMessagingAggregate := fbmessaging.NewFbExternalMessagingAggregate(mainDB, busBus, fbClient)
	fbmessagingCommandBus := fbmessaging.FbExternalMessagingAggregateMessageBus(fbExternalMessagingAggregate)
	webhookHandler := webhook.NewWebhookHandler(mainDB, faboRedis, fbClient, fbmessagingQueryBus, fbmessagingCommandBus, fbpagingQueryBus, fbuseringQueryBus)
	handlerHandler, err := BuildWebhookHandler(ctx, cfg, mainDB, fbuseringQueryBus, fbmessagingQueryBus, fbpagingQueryBus, orderingQueryBus, shippingQueryBus, fbClient, identityQueryBus, settingQueryBus, webhookHandler)
	if err != nil {
		return Output{}, nil, err
	}
	v2 := BuildWaiters(handler, handlerHandler)
	pgeventService, err := BuildPgEventService(ctx, cfg)
	if err != nil {
		return Output{}, nil, err
	}
	onesignalConfig := cfg.Onesignal
	notifier, err := BuildOneSignal(onesignalConfig)
	if err != nil {
		return Output{}, nil, err
	}
	notifierDB, err := com.BuildDatabaseNotifier(databases)
	if err != nil {
		return Output{}, nil, err
	}
	v3, err := BuildHandlers(ctx, cfg, mainDB, notifierDB)
	if err != nil {
		return Output{}, nil, err
	}
	fbExternalPageAggregate := fbpage.NewFbPageAggregate(mainDB, fbPageUtil, busBus)
	fbpagingCommandBus := fbpage.FbExternalPageAggregateMessageBus(fbExternalPageAggregate)
	fbUserAggregate := fbuser.NewFbUserAggregate(mainDB, fbpagingCommandBus, queryBus)
	fbuseringCommandBus := fbuser.FbUserAggregateMessageBus(fbUserAggregate)
	processManager := fbmessaging.NewProcessManager(busBus, fbmessagingQueryBus, fbmessagingCommandBus, fbpagingQueryBus, fbuseringQueryBus, fbuseringCommandBus, faboRedis)
	fbpageProcessManager := fbpage.NewProcessManager(busBus, fbPageUtil)
	output := Output{
		Servers:        v,
		Waiters:        v2,
		PgService:      pgeventService,
		WhSender:       webhookSender,
		Notifier:       notifier,
		Handlers:       v3,
		Health:         service,
		_fbMessagingPM: processManager,
		_fbPagePM:      fbpageProcessManager,
	}
	return output, func() {
	}, nil
}
