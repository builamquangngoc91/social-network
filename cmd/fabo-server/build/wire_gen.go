// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package build

import (
	"context"
	"o.o/api/main/accountshipnow"
	"o.o/api/main/transaction"
	"o.o/api/services/affiliate"
	"o.o/backend/cmd/fabo-server/config"
	"o.o/backend/cogs/config/_server"
	"o.o/backend/cogs/database/_min"
	fabo4 "o.o/backend/cogs/server/fabo"
	"o.o/backend/cogs/server/shop"
	"o.o/backend/cogs/shipment/_fabo"
	v2_2 "o.o/backend/cogs/shipment/ghn/v2"
	"o.o/backend/cogs/sms/_min"
	"o.o/backend/cogs/storage/_all"
	"o.o/backend/cogs/uploader"
	"o.o/backend/com/etc/logging/shippingwebhook"
	"o.o/backend/com/etc/logging/smslog/aggregate"
	"o.o/backend/com/eventhandler/fabo/publisher"
	"o.o/backend/com/eventhandler/handler"
	"o.o/backend/com/eventhandler/notifier"
	sqlstore2 "o.o/backend/com/eventhandler/notifier/sqlstore"
	"o.o/backend/com/fabo/main/fbcustomerconversationsearch"
	"o.o/backend/com/fabo/main/fbmessagetemplate"
	pm6 "o.o/backend/com/fabo/main/fbmessagetemplate/pm"
	"o.o/backend/com/fabo/main/fbmessaging"
	"o.o/backend/com/fabo/main/fbpage"
	"o.o/backend/com/fabo/main/fbuser"
	pm5 "o.o/backend/com/fabo/main/fbuser/pm"
	"o.o/backend/com/fabo/pkg/fbclient"
	redis2 "o.o/backend/com/fabo/pkg/redis"
	"o.o/backend/com/fabo/pkg/webhook"
	"o.o/backend/com/main/address"
	aggregate3 "o.o/backend/com/main/authorization/aggregate"
	aggregate4 "o.o/backend/com/main/catalog/aggregate"
	"o.o/backend/com/main/catalog/pm"
	query3 "o.o/backend/com/main/catalog/query"
	aggregate6 "o.o/backend/com/main/connectioning/aggregate"
	"o.o/backend/com/main/connectioning/manager"
	query7 "o.o/backend/com/main/connectioning/query"
	"o.o/backend/com/main/identity"
	pm2 "o.o/backend/com/main/identity/pm"
	"o.o/backend/com/main/inventory/aggregatex"
	query6 "o.o/backend/com/main/inventory/query"
	aggregate2 "o.o/backend/com/main/invitation/aggregate"
	"o.o/backend/com/main/invitation/query"
	"o.o/backend/com/main/location"
	"o.o/backend/com/main/ordering"
	pm3 "o.o/backend/com/main/ordering/pm"
	query4 "o.o/backend/com/main/receipting/query"
	"o.o/backend/com/main/shipmentpricing/pricelist"
	"o.o/backend/com/main/shipmentpricing/pricelistpromotion"
	"o.o/backend/com/main/shipmentpricing/shipmentprice"
	"o.o/backend/com/main/shipmentpricing/shipmentservice"
	"o.o/backend/com/main/shipmentpricing/shopshipmentpricelist"
	"o.o/backend/com/main/shipnow"
	aggregate7 "o.o/backend/com/main/shipping/aggregate"
	"o.o/backend/com/main/shipping/carrier"
	pm4 "o.o/backend/com/main/shipping/pm"
	query9 "o.o/backend/com/main/shipping/query"
	query8 "o.o/backend/com/main/shippingcode/query"
	query5 "o.o/backend/com/main/stocktaking/query"
	aggregate5 "o.o/backend/com/shopping/customering/aggregate"
	query2 "o.o/backend/com/shopping/customering/query"
	aggregate8 "o.o/backend/com/shopping/setting/aggregate"
	query10 "o.o/backend/com/shopping/setting/query"
	"o.o/backend/com/shopping/setting/util"
	query11 "o.o/backend/com/summary/fabo/query"
	"o.o/backend/pkg/common/apifw/captcha"
	"o.o/backend/pkg/common/apifw/health"
	auth2 "o.o/backend/pkg/common/authorization/auth"
	"o.o/backend/pkg/common/bus"
	"o.o/backend/pkg/common/mq"
	"o.o/backend/pkg/common/redis"
	"o.o/backend/pkg/etop/api/root"
	"o.o/backend/pkg/etop/api/root/fabo"
	"o.o/backend/pkg/etop/api/sadmin"
	"o.o/backend/pkg/etop/api/sadmin/_fabo"
	"o.o/backend/pkg/etop/api/shop"
	fabo2 "o.o/backend/pkg/etop/api/shop/_min/fabo"
	"o.o/backend/pkg/etop/api/shop/account"
	"o.o/backend/pkg/etop/api/shop/authorize"
	"o.o/backend/pkg/etop/api/shop/category"
	"o.o/backend/pkg/etop/api/shop/collection"
	"o.o/backend/pkg/etop/api/shop/connection"
	"o.o/backend/pkg/etop/api/shop/customer"
	"o.o/backend/pkg/etop/api/shop/customergroup"
	"o.o/backend/pkg/etop/api/shop/fulfillment"
	"o.o/backend/pkg/etop/api/shop/history"
	"o.o/backend/pkg/etop/api/shop/notification"
	"o.o/backend/pkg/etop/api/shop/order"
	"o.o/backend/pkg/etop/api/shop/product"
	"o.o/backend/pkg/etop/api/shop/setting"
	"o.o/backend/pkg/etop/api/shop/shipment"
	"o.o/backend/pkg/etop/authorize/auth"
	"o.o/backend/pkg/etop/authorize/middleware"
	"o.o/backend/pkg/etop/authorize/tokens"
	"o.o/backend/pkg/etop/eventstream"
	imcsv3 "o.o/backend/pkg/etop/logic/fulfillments/imcsv"
	"o.o/backend/pkg/etop/logic/orders"
	"o.o/backend/pkg/etop/logic/orders/imcsv"
	imcsv2 "o.o/backend/pkg/etop/logic/products/imcsv"
	"o.o/backend/pkg/etop/sqlstore"
	fabo3 "o.o/backend/pkg/fabo"
	"o.o/backend/pkg/fabo/faboinfo"
	"o.o/backend/pkg/integration/email"
	"o.o/backend/pkg/integration/shipping/ghn/webhook/v2"
	"o.o/backend/pkg/integration/sms"
)

// Injectors from wire.go:

func Build(ctx context.Context, cfg config.Config, consumer mq.KafkaConsumer) (Output, func(), error) {
	redisRedis := cfg.Redis
	store := redis.Connect(redisRedis)
	service := health.New(store)
	miscService := &root.MiscService{}
	policy := ProvidePolicy()
	authorizer := auth.New(policy)
	sharedConfig := cfg.SharedConfig
	sAdminToken := config_server.WireSAdminToken(sharedConfig)
	database_minConfig := cfg.Databases
	databases, err := database_min.BuildDatabases(database_minConfig)
	if err != nil {
		return Output{}, nil, err
	}
	mainDB := databases.Main
	queryService := identity.NewQueryService(mainDB)
	queryBus := identity.QueryServiceMessageBus(queryService)
	tokenStore := tokens.NewTokenStore(store)
	busBus := bus.New()
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
	partnerStore := sqlstore.BuildPartnerStore(mainDB)
	partnerStoreInterface := sqlstore.BindPartnerStore(partnerStore)
	accountUserStore := &sqlstore.AccountUserStore{
		DB: mainDB,
	}
	accountUserStoreInterface := sqlstore.BindAccountUserStore(accountUserStore)
	shopStore := &sqlstore.ShopStore{
		DB: mainDB,
	}
	shopStoreInterface := sqlstore.BindShopStore(shopStore)
	sessionStarter := &middleware.SessionStarter{
		SAdminToken:      sAdminToken,
		IdentityQuery:    queryBus,
		TokenStore:       tokenStore,
		AccountStore:     accountStoreInterface,
		UserStore:        userStoreInterface,
		PartnerStore:     partnerStoreInterface,
		AccountUserStore: accountUserStoreInterface,
		ShopStore:        shopStoreInterface,
	}
	session := config_server.NewSession(authorizer, sessionStarter, userStoreInterface, accountUserStoreInterface, sharedConfig, store)
	identityAggregate := identity.NewAggregate(mainDB, busBus)
	commandBus := identity.AggregateMessageBus(identityAggregate)
	invitationQuery := query.NewInvitationQuery(mainDB)
	invitationQueryBus := query.InvitationQueryMessageBus(invitationQuery)
	notifierQueryService := notifier.NewQueryService(mainDB)
	notifyQueryBus := notifier.QueryServiceNotifyBus(notifierQueryService)
	notifierAggregate := notifier.NewNotifyAggregate(mainDB)
	notifyCommandBus := notifier.NewNotifyAggregateMessageBus(notifierAggregate)
	generator := auth2.NewGenerator(store)
	smsConfig := cfg.SMS
	v := sms_min.SupportedSMSDrivers(smsConfig)
	logDB := databases.Log
	smsLogAggregate := aggregate.NewSmsLogAggregate(busBus, logDB)
	smslogCommandBus := aggregate.SmsLogAggregateMessageBus(smsLogAggregate)
	client := sms.New(smsConfig, v, smslogCommandBus)
	smtpConfig := cfg.SMTP
	emailClient := email.New(smtpConfig)
	userStoreFactory := sqlstore.NewUserStore(mainDB)
	login := &sqlstore.Login{
		UserStore: userStoreInterface,
	}
	loginInterface := sqlstore.BindLogin(login)
	webphonePublicKey := cfg.WebphonePublicKey
	userService := &root.UserService{
		Session:           session,
		IdentityAggr:      commandBus,
		IdentityQuery:     queryBus,
		InvitationQuery:   invitationQueryBus,
		NotifyQuery:       notifyQueryBus,
		NotifyAggr:        notifyCommandBus,
		EventBus:          busBus,
		AuthStore:         generator,
		TokenStore:        tokenStore,
		RedisStore:        store,
		SMSClient:         client,
		EmailClient:       emailClient,
		UserStore:         userStoreFactory,
		UserStoreIface:    userStoreInterface,
		ShopStore:         shopStoreInterface,
		AccountUserStore:  accountUserStoreInterface,
		LoginIface:        loginInterface,
		WebphonePublicKey: webphonePublicKey,
	}
	partnerStoreFactory := sqlstore.NewPartnerStore(mainDB)
	accountService := &root.AccountService{
		Session:           session,
		PartnerStore:      partnerStoreFactory,
		AccountStore:      accountStoreInterface,
		AccountUserStore:  accountUserStoreInterface,
		PartnerStoreIface: partnerStoreInterface,
	}
	locationQuery := location.New(mainDB)
	locationQueryBus := location.QueryMessageBus(locationQuery)
	locationService := &root.LocationService{
		Session:       session,
		LocationQuery: locationQueryBus,
	}
	bankService := &root.BankService{
		Session: session,
	}
	addressAggregate := address.NewAggregateAddress(busBus, mainDB, locationQueryBus)
	addressCommandBus := address.AddressAggregateMessageBus(addressAggregate)
	addressQueryService := address.NewQueryAddress(mainDB, busBus)
	addressQueryBus := address.QueryServiceMessageBus(addressQueryService)
	addressService := &root.AddressService{
		Session:     session,
		AddressAggr: addressCommandBus,
		AddressQS:   addressQueryBus,
	}
	invitationConfig := cfg.Invitation
	customerQuery := query2.NewCustomerQuery(mainDB)
	customeringQueryBus := query2.CustomerQueryMessageBus(customerQuery)
	invitationAggregate := aggregate2.NewInvitationAggregate(mainDB, invitationConfig, customeringQueryBus, queryBus, busBus, client, emailClient, accountUserStoreInterface, shopStoreInterface, userStoreInterface, store)
	invitationCommandBus := aggregate2.InvitationAggregateMessageBus(invitationAggregate)
	authorizationAggregate := &aggregate3.AuthorizationAggregate{
		Auth:             authorizer,
		AccountUserStore: accountUserStoreInterface,
		ShopStore:        shopStoreInterface,
	}
	authorizationCommandBus := aggregate3.AuthorizationAggregateMessageBus(authorizationAggregate)
	accountRelationshipService := &root.AccountRelationshipService{
		Session:           session,
		InvitationAggr:    invitationCommandBus,
		InvitationQuery:   invitationQueryBus,
		AuthorizationAggr: authorizationCommandBus,
		UserStore:         userStoreFactory,
		AccountUserStore:  accountUserStoreInterface,
	}
	userRelationshipService := &root.UserRelationshipService{
		Session:                session,
		InvitationAggr:         invitationCommandBus,
		InvitationQuery:        invitationQueryBus,
		AuthorizationAggregate: authorizationCommandBus,
		ShopStore:              shopStoreInterface,
		UserStore:              userStoreInterface,
	}
	ecomService := &root.EcomService{
		Session:        session,
		SessionStarter: sessionStarter,
	}
	emailConfig := cfg.Email
	servers, cleanup := fabo.NewServers(miscService, userService, accountService, locationService, bankService, addressService, accountRelationshipService, userRelationshipService, ecomService, store, emailConfig, smsConfig)
	shopMiscService := &shop.MiscService{
		Session: session,
	}
	accountshipnowQueryBus := _wireQueryBusValue
	accountshipnowCommandBus := _wireCommandBusValue
	accountAccountService := &account.AccountService{
		Session:             session,
		IdentityAggr:        commandBus,
		IdentityQuery:       queryBus,
		AddressQuery:        addressQueryBus,
		AddressAggr:         addressCommandBus,
		UserStore:           userStoreFactory,
		AccountStore:        accountStoreInterface,
		UserStoreIface:      userStoreInterface,
		AccountshipnowQuery: accountshipnowQueryBus,
		AccountshipnowAggr:  accountshipnowCommandBus,
	}
	queryQueryService := query3.New(mainDB)
	catalogQueryBus := query3.QueryServiceMessageBus(queryQueryService)
	aggregateAggregate := aggregate4.New(busBus, mainDB)
	catalogCommandBus := aggregate4.AggregateMessageBus(aggregateAggregate)
	collectionService := &collection.CollectionService{
		Session:      session,
		CatalogQuery: catalogQueryBus,
		CatalogAggr:  catalogCommandBus,
	}
	customerAggregate := aggregate5.NewCustomerAggregate(busBus, mainDB)
	customeringCommandBus := aggregate5.CustomerAggregateMessageBus(customerAggregate)
	aggregateAddressAggregate := aggregate5.NewAddressAggregate(mainDB, locationQueryBus)
	addressingCommandBus := aggregate5.AddressAggregateMessageBus(aggregateAddressAggregate)
	addressQuery := query2.NewAddressQuery(mainDB)
	addressingQueryBus := query2.AddressQueryMessageBus(addressQuery)
	orderingQueryService := ordering.NewQueryService(mainDB)
	orderingQueryBus := ordering.QueryServiceMessageBus(orderingQueryService)
	receiptQuery := query4.NewReceiptQuery(mainDB, orderingQueryBus)
	receiptingQueryBus := query4.ReceiptQueryMessageBus(receiptQuery)
	customerService := &customer.CustomerService{
		Session:       session,
		LocationQuery: locationQueryBus,
		CustomerQuery: customeringQueryBus,
		CustomerAggr:  customeringCommandBus,
		AddressAggr:   addressingCommandBus,
		AddressQuery:  addressingQueryBus,
		OrderQuery:    orderingQueryBus,
		ReceiptQuery:  receiptingQueryBus,
	}
	customerGroupService := &customergroup.CustomerGroupService{
		Session:       session,
		CustomerAggr:  customeringCommandBus,
		CustomerQuery: customeringQueryBus,
	}
	stocktakeQuery := query5.NewQueryStocktake(mainDB)
	stocktakingQueryBus := query5.StocktakeQueryMessageBus(stocktakeQuery)
	inventoryQueryService := query6.NewQueryInventory(stocktakingQueryBus, busBus, mainDB)
	inventoryQueryBus := query6.InventoryQueryServiceMessageBus(inventoryQueryService)
	shopVariantStore := &sqlstore.ShopVariantStore{
		DB: mainDB,
	}
	shopVariantStoreInterface := sqlstore.BindShopVariantStore(shopVariantStore)
	productService := &product.ProductService{
		Session:          session,
		CatalogQuery:     catalogQueryBus,
		CatalogAggr:      catalogCommandBus,
		InventoryQuery:   inventoryQueryBus,
		ShopVariantStore: shopVariantStoreInterface,
	}
	categoryService := &category.CategoryService{
		Session:      session,
		CatalogQuery: catalogQueryBus,
		CatalogAggr:  catalogCommandBus,
	}
	orderingAggregate := ordering.NewAggregate(busBus, mainDB)
	orderingCommandBus := ordering.AggregateMessageBus(orderingAggregate)
	mapShipmentServices := shipment_all.SupportedShipmentServices()
	connectionQuery := query7.NewConnectionQuery(mainDB, mapShipmentServices)
	connectioningQueryBus := query7.ConnectionQueryMessageBus(connectionQuery)
	connectionAggregate := aggregate6.NewConnectionAggregate(mainDB, busBus)
	connectioningCommandBus := aggregate6.ConnectionAggregateMessageBus(connectionAggregate)
	queryService2 := query8.NewQueryService(mainDB)
	shippingcodeQueryBus := query8.QueryServiceMessageBus(queryService2)
	shipmentserviceQueryService := shipmentservice.NewQueryService(mainDB, store)
	shipmentserviceQueryBus := shipmentservice.QueryServiceMessageBus(shipmentserviceQueryService)
	pricelistQueryService := pricelist.NewQueryService(mainDB, store)
	pricelistQueryBus := pricelist.QueryServiceMessageBus(pricelistQueryService)
	shopshipmentpricelistQueryService := shopshipmentpricelist.NewQueryService(mainDB, store)
	shopshipmentpricelistQueryBus := shopshipmentpricelist.QueryServiceMessageBus(shopshipmentpricelistQueryService)
	shipmentpriceQueryService := shipmentprice.NewQueryService(mainDB, store, locationQueryBus, pricelistQueryBus, shopshipmentpricelistQueryBus)
	shipmentpriceQueryBus := shipmentprice.QueryServiceMessageBus(shipmentpriceQueryService)
	pricelistpromotionQueryService := pricelistpromotion.NewQueryService(mainDB, store, locationQueryBus, queryBus, shopshipmentpricelistQueryBus, pricelistQueryBus)
	pricelistpromotionQueryBus := pricelistpromotion.QueryServiceMessageBus(pricelistpromotionQueryService)
	driver := shipment_all.SupportedCarrierDriver()
	connectionManager := manager.NewConnectionManager(store, connectioningQueryBus)
	orderStore := &sqlstore.OrderStore{
		DB:           mainDB,
		LocationBus:  locationQueryBus,
		AccountStore: accountStoreInterface,
		ShopStore:    shopStoreInterface,
	}
	orderStoreInterface := sqlstore.BindOrderStore(orderStore)
	shipmentManager, err := carrier.NewShipmentManager(busBus, locationQueryBus, queryBus, connectioningQueryBus, connectioningCommandBus, shippingcodeQueryBus, shipmentserviceQueryBus, shipmentpriceQueryBus, pricelistpromotionQueryBus, driver, connectionManager, orderStoreInterface)
	if err != nil {
		cleanup()
		return Output{}, nil, err
	}
	addressStoreInterface := sqlstore.BindAddressStore(addressStore)
	flagFaboOrderAutoConfirmPaymentStatus := cfg.FlagFaboOrderAutoConfirmPaymentStatus
	orderLogic := &orderS.OrderLogic{
		CatalogQuery:                          catalogQueryBus,
		OrderAggr:                             orderingCommandBus,
		CustomerAggr:                          customeringCommandBus,
		CustomerQuery:                         customeringQueryBus,
		TraderAddressAggr:                     addressingCommandBus,
		TraderAddressQuery:                    addressingQueryBus,
		LocationQuery:                         locationQueryBus,
		EventBus:                              busBus,
		ShipmentManager:                       shipmentManager,
		AddressStore:                          addressStoreInterface,
		OrderStore:                            orderStoreInterface,
		FlagFaboOrderUpdatePaymentSatusConfig: flagFaboOrderAutoConfirmPaymentStatus,
	}
	faboRedis := redis2.NewFaboRedis(store)
	fbMessagingQuery := fbmessaging.NewFbMessagingQuery(mainDB, faboRedis)
	fbmessagingQueryBus := fbmessaging.FbMessagingQueryMessageBus(fbMessagingQuery)
	orderService := &order.OrderService{
		Session:          session,
		OrderAggr:        orderingCommandBus,
		CustomerQuery:    customeringQueryBus,
		OrderQuery:       orderingQueryBus,
		ReceiptQuery:     receiptingQueryBus,
		OrderLogic:       orderLogic,
		FbMessagingQuery: fbmessagingQueryBus,
		OrderStore:       orderStoreInterface,
	}
	queryService3 := query9.NewQueryService(mainDB, shipmentManager, connectioningQueryBus)
	shippingQueryBus := query9.QueryServiceMessageBus(queryService3)
	fulfillmentService := &fulfillment.FulfillmentService{
		Session:         session,
		ShipmentManager: shipmentManager,
		ShippingQuery:   shippingQueryBus,
		OrderStore:      orderStoreInterface,
	}
	historyStore := &sqlstore.HistoryStore{
		DB: mainDB,
	}
	historyStoreInterface := sqlstore.BindHistoryStore(historyStore)
	historyService := &history.HistoryService{
		Session:      session,
		HistoryStore: historyStoreInterface,
	}
	notifierDB := databases.Notifier
	notificationStore := sqlstore2.NewNotificationStore(notifierDB, accountUserStoreInterface)
	deviceStore := sqlstore2.NewDeviceStore(notifierDB)
	notificationService := &notification.NotificationService{
		Session:           session,
		NotificationStore: notificationStore,
		DeviceStore:       deviceStore,
	}
	authorizeService := &authorize.AuthorizeService{
		Session:      session,
		PartnerStore: partnerStoreInterface,
	}
	aggregate9 := aggregate7.NewAggregate(mainDB, busBus, locationQueryBus, orderingQueryBus, shipmentManager, connectioningQueryBus, queryBus, addressQueryBus)
	shippingCommandBus := aggregate7.AggregateMessageBus(aggregate9)
	shipmentService := &shipment.ShipmentService{
		Session:           session,
		ShipmentManager:   shipmentManager,
		ShippingAggregate: shippingCommandBus,
		OrderStore:        orderStoreInterface,
	}
	shopSettingUtil := util.NewShopSettingUtil(store)
	shopSettingQuery := query10.NewShopSettingQuery(mainDB, shopSettingUtil)
	settingQueryBus := query10.ShopSettingQueryMessageBus(shopSettingQuery)
	shopSettingAggregate := aggregate8.NewShopSettingAggregate(mainDB, addressCommandBus, shopSettingUtil)
	settingCommandBus := aggregate8.ShopSettingAggregateMessageBus(shopSettingAggregate)
	settingService := &setting.SettingService{
		Session:      session,
		SettingQuery: settingQueryBus,
		SettingAggr:  settingCommandBus,
		AddressQ:     addressQueryBus,
	}
	connectionService := &connection.ConnectionService{
		Session:            session,
		ShipmentManager:    shipmentManager,
		ConnectionQuery:    connectioningQueryBus,
		ConnectionAggr:     connectioningCommandBus,
		IdentityQuery:      queryBus,
		AccountshipnowAggr: accountshipnowCommandBus,
	}
	shopServers := fabo2.NewServers(store, shopMiscService, accountAccountService, collectionService, customerService, customerGroupService, productService, categoryService, orderService, fulfillmentService, historyService, notificationService, authorizeService, shipmentService, settingService, connectionService)
	fbPageUtil := fbpage.NewFbPageUtil(store)
	fbPageQuery := fbpage.NewFbPageQuery(mainDB, fbPageUtil)
	fbpagingQueryBus := fbpage.FbPageQueryMessageBus(fbPageQuery)
	fbUserQuery := fbuser.NewFbUserQuery(mainDB, customeringQueryBus)
	fbuseringQueryBus := fbuser.FbUserQueryMessageBus(fbUserQuery)
	faboPagesKit := &faboinfo.FaboPagesKit{
		FBPageQuery: fbpagingQueryBus,
		FBUserQuery: fbuseringQueryBus,
	}
	fbExternalPageAggregate := fbpage.NewFbPageAggregate(mainDB, fbPageUtil, busBus)
	fbpagingCommandBus := fbpage.FbExternalPageAggregateMessageBus(fbExternalPageAggregate)
	fbUserAggregate := fbuser.NewFbUserAggregate(mainDB, fbpagingCommandBus, customeringQueryBus)
	fbuseringCommandBus := fbuser.FbUserAggregateMessageBus(fbUserAggregate)
	appConfig := cfg.FacebookApp
	fbClient := fbclient.New(appConfig)
	pageService := &fabo3.PageService{
		Session:             session,
		FaboInfo:            faboPagesKit,
		FBMessagingQuery:    fbmessagingQueryBus,
		FBExternalUserQuery: fbuseringQueryBus,
		FBExternalUserAggr:  fbuseringCommandBus,
		FBExternalPageQuery: fbpagingQueryBus,
		FBExternalPageAggr:  fbpagingCommandBus,
		FBClient:            fbClient,
	}
	demoService := &fabo3.DemoService{
		Session:  session,
		FBClient: fbClient,
	}
	fbExternalMessagingAggregate := fbmessaging.NewFbExternalMessagingAggregate(mainDB, busBus, fbClient)
	fbmessagingCommandBus := fbmessaging.FbExternalMessagingAggregateMessageBus(fbExternalMessagingAggregate)
	fbSearchService := fbcustomerconversationsearch.NewFbSearchServiceQuery(mainDB)
	fbcustomerconversationsearchQueryBus := fbcustomerconversationsearch.FbSearchQueryMessageBus(fbSearchService)
	fbMessageTemplateQuery := fbmessagetemplate.NewFbMessagingQuery(mainDB)
	fbmessagetemplateQueryBus := fbmessagetemplate.FbMessagingQueryMessageBus(fbMessageTemplateQuery)
	fbMessageTemplateAggregate := fbmessagetemplate.NewFbMessageTemplateAggregate(mainDB)
	fbmessagetemplateCommandBus := fbmessagetemplate.FbMessageTemplateAggregateMessageBus(fbMessageTemplateAggregate)
	customerConversationService := &fabo3.CustomerConversationService{
		Session:                session,
		FaboPagesKit:           faboPagesKit,
		FBClient:               fbClient,
		FBMessagingQuery:       fbmessagingQueryBus,
		FBMessagingAggr:        fbmessagingCommandBus,
		FBPagingQuery:          fbpagingQueryBus,
		FBUserQuery:            fbuseringQueryBus,
		FbSearchQuery:          fbcustomerconversationsearchQueryBus,
		FbMessageTemplateQuery: fbmessagetemplateQueryBus,
		FbMessageTemplateAggr:  fbmessagetemplateCommandBus,
	}
	faboCustomerService := &fabo3.CustomerService{
		Session:        session,
		CustomerQuery:  customeringQueryBus,
		FBUseringQuery: fbuseringQueryBus,
		FBUseringAggr:  fbuseringCommandBus,
	}
	shopService := &fabo3.ShopService{
		Session:             session,
		FBExternalUserQuery: fbuseringQueryBus,
		FBExternalUserAggr:  fbuseringCommandBus,
	}
	extraShipmentService := &fabo3.ExtraShipmentService{
		Session:      session,
		ShippingQS:   shippingQueryBus,
		ConnectionQS: connectioningQueryBus,
	}
	dashboardQuery := query11.NewDashboardQuery(mainDB, store)
	summaryQueryBus := query11.DashboardQueryMessageBus(dashboardQuery)
	summaryService := &fabo3.SummaryService{
		Session:      session,
		SummaryQuery: summaryQueryBus,
	}
	faboServers := fabo3.NewServers(pageService, demoService, customerConversationService, faboCustomerService, shopService, extraShipmentService, summaryService, store)
	webhookCallbackService := sadmin.NewWebhookCallbackService(store)
	webhookService := &sadmin.WebhookService{
		Session:                session,
		WebhookCallbackService: webhookCallbackService,
	}
	sadminServers := _fabo.NewServers(webhookService)
	captchaConfig := cfg.Captcha
	captchaCaptcha := captcha.New(captchaConfig)
	intHandlers, err := BuildIntHandlers(servers, shopServers, faboServers, sadminServers, captchaCaptcha)
	if err != nil {
		cleanup()
		return Output{}, nil, err
	}
	dirConfigs := cfg.UploadDirs
	driverConfig := cfg.StorageDriver
	bucket, err := storage_all.Build(ctx, driverConfig)
	if err != nil {
		cleanup()
		return Output{}, nil, err
	}
	uploader, err := _uploader.NewUploader(ctx, dirConfigs, bucket)
	if err != nil {
		cleanup()
		return Output{}, nil, err
	}
	exportAttemptStore := sqlstore.BuildExportAttemptStore(mainDB)
	exportAttemptStoreInterface := sqlstore.BindExportAttemptStore(exportAttemptStore)
	imcsvImport, cleanup2 := imcsv.New(authorizer, locationQueryBus, store, uploader, mainDB, orderStoreInterface, exportAttemptStoreInterface)
	categoryStore := &sqlstore.CategoryStore{
		DB: mainDB,
	}
	categoryStoreInterface := sqlstore.BindCategoryStore(categoryStore)
	import2, cleanup3 := imcsv2.New(store, uploader, mainDB, exportAttemptStoreInterface, categoryStoreInterface, shopStoreInterface)
	import3, cleanup4 := imcsv3.New(store, uploader, exportAttemptStoreInterface)
	importHandler := server_shop.BuildImportHandler(imcsvImport, import2, import3, session)
	eventStream := eventstream.New(ctx)
	eventStreamHandler := server_shop.BuildEventStreamHandler(eventStream, session)
	downloadHandler := server_shop.BuildDownloadHandler()
	faboImageHandler := fabo4.BuildFaboImageHandler()
	mainServer := BuildMainServer(service, intHandlers, sharedConfig, importHandler, eventStreamHandler, downloadHandler, faboImageHandler)
	shipment_allConfig := cfg.Shipment
	webhookConfig := shipment_allConfig.GHNWebhook
	shippingwebhookAggregate := shippingwebhook.NewAggregate(logDB)
	v2Webhook := v2.New(mainDB, shipmentManager, queryBus, shippingCommandBus, shippingwebhookAggregate, orderStoreInterface)
	ghnWebhookServer := v2_2.NewGHNWebhookServer(webhookConfig, shipmentManager, queryBus, shippingCommandBus, v2Webhook)
	configWebhookConfig := cfg.Webhook
	kafkaProducer, err := BuildPgProducer(ctx, cfg)
	if err != nil {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
		return Output{}, nil, err
	}
	kafka := cfg.Kafka
	webhookWebhook := webhook.New(logDB, store, configWebhookConfig, faboRedis, kafkaProducer, kafka)
	fbWebhookServer := BuildWebhookServer(configWebhookConfig, webhookWebhook)
	v3 := BuildServers(mainServer, ghnWebhookServer, fbWebhookServer)
	handlerHandler := handler.New(consumer, kafka)
	publisherPublisher := publisher.New(eventStream)
	processManager := pm.New(busBus, catalogQueryBus, catalogCommandBus)
	pmProcessManager := pm2.New(busBus, queryBus, commandBus, invitationQueryBus, addressQueryBus, addressCommandBus, accountUserStoreInterface)
	affiliateCommandBus := _wireAffiliateCommandBusValue
	inventoryAggregate := aggregatex.NewAggregateInventory(busBus, mainDB, stocktakingQueryBus, catalogQueryBus, orderStoreInterface)
	inventoryCommandBus := aggregatex.InventoryAggregateMessageBus(inventoryAggregate)
	shipnowQueryService := shipnow.NewQueryService(mainDB)
	shipnowQueryBus := shipnow.QueryServiceMessageBus(shipnowQueryService)
	processManager2 := pm3.New(busBus, orderingCommandBus, affiliateCommandBus, receiptingQueryBus, inventoryCommandBus, orderingQueryBus, customeringQueryBus, shipnowQueryBus)
	transactionQueryBus := _wireTransactionQueryBusValue
	processManager3 := pm4.New(busBus, shippingQueryBus, shippingCommandBus, store, connectioningQueryBus, shopStoreInterface, transactionQueryBus)
	processManager4 := pm5.New(busBus, fbuseringCommandBus)
	fbmessagingProcessManager := fbmessaging.NewProcessManager(busBus, fbmessagingQueryBus, fbmessagingCommandBus, fbpagingQueryBus, fbuseringQueryBus, fbuseringCommandBus, faboRedis)
	processManager5 := pm6.NewProcessManager(busBus, fbmessagetemplateCommandBus)
	fbpageProcessManager := fbpage.NewProcessManager(busBus, fbPageUtil)
	output := Output{
		Servers:              v3,
		EventStream:          eventStream,
		Health:               service,
		Handler:              handlerHandler,
		Publisher:            publisherPublisher,
		FbClient:             fbClient,
		_catalogPM:           processManager,
		_identityPM:          pmProcessManager,
		_orderPM:             processManager2,
		_shippingPM:          processManager3,
		_fbuserPM:            processManager4,
		_fbMessagingPM:       fbmessagingProcessManager,
		_fbMessageTemplatePM: processManager5,
		_fbPagePM:            fbpageProcessManager,
	}
	return output, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

var (
	_wireQueryBusValue            = accountshipnow.QueryBus{}
	_wireCommandBusValue          = accountshipnow.CommandBus{}
	_wireAffiliateCommandBusValue = affiliate.CommandBus{}
	_wireTransactionQueryBusValue = transaction.QueryBus{}
)