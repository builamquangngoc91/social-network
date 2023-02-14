// +build !generator

// Code generated by generator sqlgen. DO NOT EDIT.

package sqlstore

import (
	time "time"

	sq "o.o/backend/pkg/common/sql/sq"
	dot "o.o/capi/dot"
)

type EtopAcountFilters struct{ prefix string }

func NewEtopAcountFilters(prefix string) EtopAcountFilters {
	return EtopAcountFilters{prefix}
}

func (ft *EtopAcountFilters) Filter(pred string, args ...interface{}) sq.WriterTo {
	return sq.Filter(&ft.prefix, pred, args...)
}

func (ft EtopAcountFilters) Prefix() string {
	return ft.prefix
}

func (ft *EtopAcountFilters) ByID(ID string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "id",
		Value:  ID,
		IsNil:  ID == "",
	}
}

func (ft *EtopAcountFilters) ByIDPtr(ID *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "id",
		Value:  ID,
		IsNil:  ID == nil,
		IsZero: ID != nil && (*ID) == "",
	}
}

func (ft *EtopAcountFilters) ByFullName(FullName string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "full_name",
		Value:  FullName,
		IsNil:  FullName == "",
	}
}

func (ft *EtopAcountFilters) ByFullNamePtr(FullName *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "full_name",
		Value:  FullName,
		IsNil:  FullName == nil,
		IsZero: FullName != nil && (*FullName) == "",
	}
}

func (ft *EtopAcountFilters) ByPhone(Phone string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "phone",
		Value:  Phone,
		IsNil:  Phone == "",
	}
}

func (ft *EtopAcountFilters) ByPhonePtr(Phone *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "phone",
		Value:  Phone,
		IsNil:  Phone == nil,
		IsZero: Phone != nil && (*Phone) == "",
	}
}

func (ft *EtopAcountFilters) ByEmail(Email string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "email",
		Value:  Email,
		IsNil:  Email == "",
	}
}

func (ft *EtopAcountFilters) ByEmailPtr(Email *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "email",
		Value:  Email,
		IsNil:  Email == nil,
		IsZero: Email != nil && (*Email) == "",
	}
}

func (ft *EtopAcountFilters) ByAccountID(AccountID string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "account_id",
		Value:  AccountID,
		IsNil:  AccountID == "",
	}
}

func (ft *EtopAcountFilters) ByAccountIDPtr(AccountID *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "account_id",
		Value:  AccountID,
		IsNil:  AccountID == nil,
		IsZero: AccountID != nil && (*AccountID) == "",
	}
}

func (ft *EtopAcountFilters) ByAcountName(AcountName string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "acount_name",
		Value:  AcountName,
		IsNil:  AcountName == "",
	}
}

func (ft *EtopAcountFilters) ByAcountNamePtr(AcountName *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "acount_name",
		Value:  AcountName,
		IsNil:  AcountName == nil,
		IsZero: AcountName != nil && (*AcountName) == "",
	}
}

func (ft *EtopAcountFilters) ByAccountType(AccountType string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "account_type",
		Value:  AccountType,
		IsNil:  AccountType == "",
	}
}

func (ft *EtopAcountFilters) ByAccountTypePtr(AccountType *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "account_type",
		Value:  AccountType,
		IsNil:  AccountType == nil,
		IsZero: AccountType != nil && (*AccountType) == "",
	}
}

func (ft *EtopAcountFilters) ByIsOperator(IsOperator bool) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "is_operator",
		Value:  IsOperator,
		IsNil:  bool(!IsOperator),
	}
}

func (ft *EtopAcountFilters) ByIsOperatorPtr(IsOperator *bool) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "is_operator",
		Value:  IsOperator,
		IsNil:  IsOperator == nil,
		IsZero: IsOperator != nil && bool(!(*IsOperator)),
	}
}

func (ft *EtopAcountFilters) ByVtigerAccount(VtigerAccount string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "vtiger_account",
		Value:  VtigerAccount,
		IsNil:  VtigerAccount == "",
	}
}

func (ft *EtopAcountFilters) ByVtigerAccountPtr(VtigerAccount *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "vtiger_account",
		Value:  VtigerAccount,
		IsNil:  VtigerAccount == nil,
		IsZero: VtigerAccount != nil && (*VtigerAccount) == "",
	}
}

type VtigerAccountFilters struct{ prefix string }

func NewVtigerAccountFilters(prefix string) VtigerAccountFilters {
	return VtigerAccountFilters{prefix}
}

func (ft *VtigerAccountFilters) Filter(pred string, args ...interface{}) sq.WriterTo {
	return sq.Filter(&ft.prefix, pred, args...)
}

func (ft VtigerAccountFilters) Prefix() string {
	return ft.prefix
}

func (ft *VtigerAccountFilters) ByID(ID string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "id",
		Value:  ID,
		IsNil:  ID == "",
	}
}

func (ft *VtigerAccountFilters) ByIDPtr(ID *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "id",
		Value:  ID,
		IsNil:  ID == nil,
		IsZero: ID != nil && (*ID) == "",
	}
}

func (ft *VtigerAccountFilters) ByUserName(UserName string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "user_name",
		Value:  UserName,
		IsNil:  UserName == "",
	}
}

func (ft *VtigerAccountFilters) ByUserNamePtr(UserName *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "user_name",
		Value:  UserName,
		IsNil:  UserName == nil,
		IsZero: UserName != nil && (*UserName) == "",
	}
}

func (ft *VtigerAccountFilters) ByFirstName(FirstName string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "first_name",
		Value:  FirstName,
		IsNil:  FirstName == "",
	}
}

func (ft *VtigerAccountFilters) ByFirstNamePtr(FirstName *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "first_name",
		Value:  FirstName,
		IsNil:  FirstName == nil,
		IsZero: FirstName != nil && (*FirstName) == "",
	}
}

func (ft *VtigerAccountFilters) ByRoleID(RoleID dot.ID) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "role_id",
		Value:  RoleID,
		IsNil:  RoleID == 0,
	}
}

func (ft *VtigerAccountFilters) ByRoleIDPtr(RoleID *dot.ID) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "role_id",
		Value:  RoleID,
		IsNil:  RoleID == nil,
		IsZero: RoleID != nil && (*RoleID) == 0,
	}
}

func (ft *VtigerAccountFilters) ByEmail1(Email1 string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "email1",
		Value:  Email1,
		IsNil:  Email1 == "",
	}
}

func (ft *VtigerAccountFilters) ByEmail1Ptr(Email1 *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "email1",
		Value:  Email1,
		IsNil:  Email1 == nil,
		IsZero: Email1 != nil && (*Email1) == "",
	}
}

func (ft *VtigerAccountFilters) BySecondaryemail(Secondaryemail string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "secondaryemail",
		Value:  Secondaryemail,
		IsNil:  Secondaryemail == "",
	}
}

func (ft *VtigerAccountFilters) BySecondaryemailPtr(Secondaryemail *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "secondaryemail",
		Value:  Secondaryemail,
		IsNil:  Secondaryemail == nil,
		IsZero: Secondaryemail != nil && (*Secondaryemail) == "",
	}
}

func (ft *VtigerAccountFilters) ByStatus(Status string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "status",
		Value:  Status,
		IsNil:  Status == "",
	}
}

func (ft *VtigerAccountFilters) ByStatusPtr(Status *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "status",
		Value:  Status,
		IsNil:  Status == nil,
		IsZero: Status != nil && (*Status) == "",
	}
}

type VtigerContactFilters struct{ prefix string }

func NewVtigerContactFilters(prefix string) VtigerContactFilters {
	return VtigerContactFilters{prefix}
}

func (ft *VtigerContactFilters) Filter(pred string, args ...interface{}) sq.WriterTo {
	return sq.Filter(&ft.prefix, pred, args...)
}

func (ft VtigerContactFilters) Prefix() string {
	return ft.prefix
}

func (ft *VtigerContactFilters) ByID(ID string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "id",
		Value:  ID,
		IsNil:  ID == "",
	}
}

func (ft *VtigerContactFilters) ByIDPtr(ID *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "id",
		Value:  ID,
		IsNil:  ID == nil,
		IsZero: ID != nil && (*ID) == "",
	}
}

func (ft *VtigerContactFilters) ByFirstname(Firstname string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "firstname",
		Value:  Firstname,
		IsNil:  Firstname == "",
	}
}

func (ft *VtigerContactFilters) ByFirstnamePtr(Firstname *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "firstname",
		Value:  Firstname,
		IsNil:  Firstname == nil,
		IsZero: Firstname != nil && (*Firstname) == "",
	}
}

func (ft *VtigerContactFilters) ByContactNo(ContactNo string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "contact_no",
		Value:  ContactNo,
		IsNil:  ContactNo == "",
	}
}

func (ft *VtigerContactFilters) ByContactNoPtr(ContactNo *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "contact_no",
		Value:  ContactNo,
		IsNil:  ContactNo == nil,
		IsZero: ContactNo != nil && (*ContactNo) == "",
	}
}

func (ft *VtigerContactFilters) ByPhone(Phone string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "phone",
		Value:  Phone,
		IsNil:  Phone == "",
	}
}

func (ft *VtigerContactFilters) ByPhonePtr(Phone *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "phone",
		Value:  Phone,
		IsNil:  Phone == nil,
		IsZero: Phone != nil && (*Phone) == "",
	}
}

func (ft *VtigerContactFilters) ByLastname(Lastname string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "lastname",
		Value:  Lastname,
		IsNil:  Lastname == "",
	}
}

func (ft *VtigerContactFilters) ByLastnamePtr(Lastname *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "lastname",
		Value:  Lastname,
		IsNil:  Lastname == nil,
		IsZero: Lastname != nil && (*Lastname) == "",
	}
}

func (ft *VtigerContactFilters) ByMobile(Mobile string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "mobile",
		Value:  Mobile,
		IsNil:  Mobile == "",
	}
}

func (ft *VtigerContactFilters) ByMobilePtr(Mobile *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "mobile",
		Value:  Mobile,
		IsNil:  Mobile == nil,
		IsZero: Mobile != nil && (*Mobile) == "",
	}
}

func (ft *VtigerContactFilters) ByEmail(Email string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "email",
		Value:  Email,
		IsNil:  Email == "",
	}
}

func (ft *VtigerContactFilters) ByEmailPtr(Email *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "email",
		Value:  Email,
		IsNil:  Email == nil,
		IsZero: Email != nil && (*Email) == "",
	}
}

func (ft *VtigerContactFilters) ByLeadsource(Leadsource string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "leadsource",
		Value:  Leadsource,
		IsNil:  Leadsource == "",
	}
}

func (ft *VtigerContactFilters) ByLeadsourcePtr(Leadsource *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "leadsource",
		Value:  Leadsource,
		IsNil:  Leadsource == nil,
		IsZero: Leadsource != nil && (*Leadsource) == "",
	}
}

func (ft *VtigerContactFilters) BySecondaryemail(Secondaryemail string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "secondaryemail",
		Value:  Secondaryemail,
		IsNil:  Secondaryemail == "",
	}
}

func (ft *VtigerContactFilters) BySecondaryemailPtr(Secondaryemail *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "secondaryemail",
		Value:  Secondaryemail,
		IsNil:  Secondaryemail == nil,
		IsZero: Secondaryemail != nil && (*Secondaryemail) == "",
	}
}

func (ft *VtigerContactFilters) ByAssignedUserID(AssignedUserID string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "assigned_user_id",
		Value:  AssignedUserID,
		IsNil:  AssignedUserID == "",
	}
}

func (ft *VtigerContactFilters) ByAssignedUserIDPtr(AssignedUserID *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "assigned_user_id",
		Value:  AssignedUserID,
		IsNil:  AssignedUserID == nil,
		IsZero: AssignedUserID != nil && (*AssignedUserID) == "",
	}
}

func (ft *VtigerContactFilters) ByCreatedAt(CreatedAt time.Time) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "created_at",
		Value:  CreatedAt,
		IsNil:  CreatedAt.IsZero(),
	}
}

func (ft *VtigerContactFilters) ByCreatedAtPtr(CreatedAt *time.Time) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "created_at",
		Value:  CreatedAt,
		IsNil:  CreatedAt == nil,
		IsZero: CreatedAt != nil && (*CreatedAt).IsZero(),
	}
}

func (ft *VtigerContactFilters) ByEtopUserID(EtopUserID dot.ID) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "etop_user_id",
		Value:  EtopUserID,
		IsNil:  EtopUserID == 0,
	}
}

func (ft *VtigerContactFilters) ByEtopUserIDPtr(EtopUserID *dot.ID) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "etop_user_id",
		Value:  EtopUserID,
		IsNil:  EtopUserID == nil,
		IsZero: EtopUserID != nil && (*EtopUserID) == 0,
	}
}

func (ft *VtigerContactFilters) ByUpdatedAt(UpdatedAt time.Time) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "updated_at",
		Value:  UpdatedAt,
		IsNil:  UpdatedAt.IsZero(),
	}
}

func (ft *VtigerContactFilters) ByUpdatedAtPtr(UpdatedAt *time.Time) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "updated_at",
		Value:  UpdatedAt,
		IsNil:  UpdatedAt == nil,
		IsZero: UpdatedAt != nil && (*UpdatedAt).IsZero(),
	}
}

func (ft *VtigerContactFilters) ByDescription(Description string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "description",
		Value:  Description,
		IsNil:  Description == "",
	}
}

func (ft *VtigerContactFilters) ByDescriptionPtr(Description *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "description",
		Value:  Description,
		IsNil:  Description == nil,
		IsZero: Description != nil && (*Description) == "",
	}
}

func (ft *VtigerContactFilters) BySource(Source string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "source",
		Value:  Source,
		IsNil:  Source == "",
	}
}

func (ft *VtigerContactFilters) BySourcePtr(Source *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "source",
		Value:  Source,
		IsNil:  Source == nil,
		IsZero: Source != nil && (*Source) == "",
	}
}

func (ft *VtigerContactFilters) ByUsedShippingProvider(UsedShippingProvider string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "used_shipping_provider",
		Value:  UsedShippingProvider,
		IsNil:  UsedShippingProvider == "",
	}
}

func (ft *VtigerContactFilters) ByUsedShippingProviderPtr(UsedShippingProvider *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "used_shipping_provider",
		Value:  UsedShippingProvider,
		IsNil:  UsedShippingProvider == nil,
		IsZero: UsedShippingProvider != nil && (*UsedShippingProvider) == "",
	}
}

func (ft *VtigerContactFilters) ByOrdersPerDay(OrdersPerDay string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "orders_per_day",
		Value:  OrdersPerDay,
		IsNil:  OrdersPerDay == "",
	}
}

func (ft *VtigerContactFilters) ByOrdersPerDayPtr(OrdersPerDay *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "orders_per_day",
		Value:  OrdersPerDay,
		IsNil:  OrdersPerDay == nil,
		IsZero: OrdersPerDay != nil && (*OrdersPerDay) == "",
	}
}

func (ft *VtigerContactFilters) ByCompany(Company string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "company",
		Value:  Company,
		IsNil:  Company == "",
	}
}

func (ft *VtigerContactFilters) ByCompanyPtr(Company *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "company",
		Value:  Company,
		IsNil:  Company == nil,
		IsZero: Company != nil && (*Company) == "",
	}
}

func (ft *VtigerContactFilters) ByCity(City string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "city",
		Value:  City,
		IsNil:  City == "",
	}
}

func (ft *VtigerContactFilters) ByCityPtr(City *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "city",
		Value:  City,
		IsNil:  City == nil,
		IsZero: City != nil && (*City) == "",
	}
}

func (ft *VtigerContactFilters) ByState(State string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "state",
		Value:  State,
		IsNil:  State == "",
	}
}

func (ft *VtigerContactFilters) ByStatePtr(State *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "state",
		Value:  State,
		IsNil:  State == nil,
		IsZero: State != nil && (*State) == "",
	}
}

func (ft *VtigerContactFilters) ByWebsite(Website string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "website",
		Value:  Website,
		IsNil:  Website == "",
	}
}

func (ft *VtigerContactFilters) ByWebsitePtr(Website *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "website",
		Value:  Website,
		IsNil:  Website == nil,
		IsZero: Website != nil && (*Website) == "",
	}
}

func (ft *VtigerContactFilters) ByLane(Lane string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "lane",
		Value:  Lane,
		IsNil:  Lane == "",
	}
}

func (ft *VtigerContactFilters) ByLanePtr(Lane *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "lane",
		Value:  Lane,
		IsNil:  Lane == nil,
		IsZero: Lane != nil && (*Lane) == "",
	}
}

func (ft *VtigerContactFilters) ByCountry(Country string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "country",
		Value:  Country,
		IsNil:  Country == "",
	}
}

func (ft *VtigerContactFilters) ByCountryPtr(Country *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "country",
		Value:  Country,
		IsNil:  Country == nil,
		IsZero: Country != nil && (*Country) == "",
	}
}

func (ft *VtigerContactFilters) BySearchNorm(SearchNorm string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "search_norm",
		Value:  SearchNorm,
		IsNil:  SearchNorm == "",
	}
}

func (ft *VtigerContactFilters) BySearchNormPtr(SearchNorm *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "search_norm",
		Value:  SearchNorm,
		IsNil:  SearchNorm == nil,
		IsZero: SearchNorm != nil && (*SearchNorm) == "",
	}
}

func (ft *VtigerContactFilters) ByVtigerCreatedAt(VtigerCreatedAt time.Time) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "vtiger_created_at",
		Value:  VtigerCreatedAt,
		IsNil:  VtigerCreatedAt.IsZero(),
	}
}

func (ft *VtigerContactFilters) ByVtigerCreatedAtPtr(VtigerCreatedAt *time.Time) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "vtiger_created_at",
		Value:  VtigerCreatedAt,
		IsNil:  VtigerCreatedAt == nil,
		IsZero: VtigerCreatedAt != nil && (*VtigerCreatedAt).IsZero(),
	}
}

func (ft *VtigerContactFilters) ByVtigerUpdatedAt(VtigerUpdatedAt time.Time) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "vtiger_updated_at",
		Value:  VtigerUpdatedAt,
		IsNil:  VtigerUpdatedAt.IsZero(),
	}
}

func (ft *VtigerContactFilters) ByVtigerUpdatedAtPtr(VtigerUpdatedAt *time.Time) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "vtiger_updated_at",
		Value:  VtigerUpdatedAt,
		IsNil:  VtigerUpdatedAt == nil,
		IsZero: VtigerUpdatedAt != nil && (*VtigerUpdatedAt).IsZero(),
	}
}