// +build !generator

// Code generated by generator sqlgen. DO NOT EDIT.

package sqlstore

import (
	time "time"

	sq "o.o/backend/pkg/common/sql/sq"
	dot "o.o/capi/dot"
)

type ShopSupplierFilters struct{ prefix string }

func NewShopSupplierFilters(prefix string) ShopSupplierFilters {
	return ShopSupplierFilters{prefix}
}

func (ft *ShopSupplierFilters) Filter(pred string, args ...interface{}) sq.WriterTo {
	return sq.Filter(&ft.prefix, pred, args...)
}

func (ft ShopSupplierFilters) Prefix() string {
	return ft.prefix
}

func (ft *ShopSupplierFilters) ByID(ID dot.ID) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "id",
		Value:  ID,
		IsNil:  ID == 0,
	}
}

func (ft *ShopSupplierFilters) ByIDPtr(ID *dot.ID) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "id",
		Value:  ID,
		IsNil:  ID == nil,
		IsZero: ID != nil && (*ID) == 0,
	}
}

func (ft *ShopSupplierFilters) ByShopID(ShopID dot.ID) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "shop_id",
		Value:  ShopID,
		IsNil:  ShopID == 0,
	}
}

func (ft *ShopSupplierFilters) ByShopIDPtr(ShopID *dot.ID) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "shop_id",
		Value:  ShopID,
		IsNil:  ShopID == nil,
		IsZero: ShopID != nil && (*ShopID) == 0,
	}
}

func (ft *ShopSupplierFilters) ByFullName(FullName string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "full_name",
		Value:  FullName,
		IsNil:  FullName == "",
	}
}

func (ft *ShopSupplierFilters) ByFullNamePtr(FullName *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "full_name",
		Value:  FullName,
		IsNil:  FullName == nil,
		IsZero: FullName != nil && (*FullName) == "",
	}
}

func (ft *ShopSupplierFilters) ByPhone(Phone string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "phone",
		Value:  Phone,
		IsNil:  Phone == "",
	}
}

func (ft *ShopSupplierFilters) ByPhonePtr(Phone *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "phone",
		Value:  Phone,
		IsNil:  Phone == nil,
		IsZero: Phone != nil && (*Phone) == "",
	}
}

func (ft *ShopSupplierFilters) ByEmail(Email string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "email",
		Value:  Email,
		IsNil:  Email == "",
	}
}

func (ft *ShopSupplierFilters) ByEmailPtr(Email *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "email",
		Value:  Email,
		IsNil:  Email == nil,
		IsZero: Email != nil && (*Email) == "",
	}
}

func (ft *ShopSupplierFilters) ByCode(Code string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "code",
		Value:  Code,
		IsNil:  Code == "",
	}
}

func (ft *ShopSupplierFilters) ByCodePtr(Code *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "code",
		Value:  Code,
		IsNil:  Code == nil,
		IsZero: Code != nil && (*Code) == "",
	}
}

func (ft *ShopSupplierFilters) ByCodeNorm(CodeNorm int) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "code_norm",
		Value:  CodeNorm,
		IsNil:  CodeNorm == 0,
	}
}

func (ft *ShopSupplierFilters) ByCodeNormPtr(CodeNorm *int) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "code_norm",
		Value:  CodeNorm,
		IsNil:  CodeNorm == nil,
		IsZero: CodeNorm != nil && (*CodeNorm) == 0,
	}
}

func (ft *ShopSupplierFilters) ByCompanyName(CompanyName string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "company_name",
		Value:  CompanyName,
		IsNil:  CompanyName == "",
	}
}

func (ft *ShopSupplierFilters) ByCompanyNamePtr(CompanyName *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "company_name",
		Value:  CompanyName,
		IsNil:  CompanyName == nil,
		IsZero: CompanyName != nil && (*CompanyName) == "",
	}
}

func (ft *ShopSupplierFilters) ByCompanyNameNorm(CompanyNameNorm string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "company_name_norm",
		Value:  CompanyNameNorm,
		IsNil:  CompanyNameNorm == "",
	}
}

func (ft *ShopSupplierFilters) ByCompanyNameNormPtr(CompanyNameNorm *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "company_name_norm",
		Value:  CompanyNameNorm,
		IsNil:  CompanyNameNorm == nil,
		IsZero: CompanyNameNorm != nil && (*CompanyNameNorm) == "",
	}
}

func (ft *ShopSupplierFilters) ByTaxNumber(TaxNumber string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "tax_number",
		Value:  TaxNumber,
		IsNil:  TaxNumber == "",
	}
}

func (ft *ShopSupplierFilters) ByTaxNumberPtr(TaxNumber *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "tax_number",
		Value:  TaxNumber,
		IsNil:  TaxNumber == nil,
		IsZero: TaxNumber != nil && (*TaxNumber) == "",
	}
}

func (ft *ShopSupplierFilters) ByHeadquaterAddress(HeadquaterAddress string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "headquater_address",
		Value:  HeadquaterAddress,
		IsNil:  HeadquaterAddress == "",
	}
}

func (ft *ShopSupplierFilters) ByHeadquaterAddressPtr(HeadquaterAddress *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "headquater_address",
		Value:  HeadquaterAddress,
		IsNil:  HeadquaterAddress == nil,
		IsZero: HeadquaterAddress != nil && (*HeadquaterAddress) == "",
	}
}

func (ft *ShopSupplierFilters) ByNote(Note string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "note",
		Value:  Note,
		IsNil:  Note == "",
	}
}

func (ft *ShopSupplierFilters) ByNotePtr(Note *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "note",
		Value:  Note,
		IsNil:  Note == nil,
		IsZero: Note != nil && (*Note) == "",
	}
}

func (ft *ShopSupplierFilters) ByFullNameNorm(FullNameNorm string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "full_name_norm",
		Value:  FullNameNorm,
		IsNil:  FullNameNorm == "",
	}
}

func (ft *ShopSupplierFilters) ByFullNameNormPtr(FullNameNorm *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "full_name_norm",
		Value:  FullNameNorm,
		IsNil:  FullNameNorm == nil,
		IsZero: FullNameNorm != nil && (*FullNameNorm) == "",
	}
}

func (ft *ShopSupplierFilters) ByPhoneNorm(PhoneNorm string) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "phone_norm",
		Value:  PhoneNorm,
		IsNil:  PhoneNorm == "",
	}
}

func (ft *ShopSupplierFilters) ByPhoneNormPtr(PhoneNorm *string) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "phone_norm",
		Value:  PhoneNorm,
		IsNil:  PhoneNorm == nil,
		IsZero: PhoneNorm != nil && (*PhoneNorm) == "",
	}
}

func (ft *ShopSupplierFilters) ByStatus(Status int) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "status",
		Value:  Status,
		IsNil:  Status == 0,
	}
}

func (ft *ShopSupplierFilters) ByStatusPtr(Status *int) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "status",
		Value:  Status,
		IsNil:  Status == nil,
		IsZero: Status != nil && (*Status) == 0,
	}
}

func (ft *ShopSupplierFilters) ByCreatedAt(CreatedAt time.Time) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "created_at",
		Value:  CreatedAt,
		IsNil:  CreatedAt.IsZero(),
	}
}

func (ft *ShopSupplierFilters) ByCreatedAtPtr(CreatedAt *time.Time) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "created_at",
		Value:  CreatedAt,
		IsNil:  CreatedAt == nil,
		IsZero: CreatedAt != nil && (*CreatedAt).IsZero(),
	}
}

func (ft *ShopSupplierFilters) ByUpdatedAt(UpdatedAt time.Time) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "updated_at",
		Value:  UpdatedAt,
		IsNil:  UpdatedAt.IsZero(),
	}
}

func (ft *ShopSupplierFilters) ByUpdatedAtPtr(UpdatedAt *time.Time) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "updated_at",
		Value:  UpdatedAt,
		IsNil:  UpdatedAt == nil,
		IsZero: UpdatedAt != nil && (*UpdatedAt).IsZero(),
	}
}

func (ft *ShopSupplierFilters) ByDeletedAt(DeletedAt time.Time) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "deleted_at",
		Value:  DeletedAt,
		IsNil:  DeletedAt.IsZero(),
	}
}

func (ft *ShopSupplierFilters) ByDeletedAtPtr(DeletedAt *time.Time) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "deleted_at",
		Value:  DeletedAt,
		IsNil:  DeletedAt == nil,
		IsZero: DeletedAt != nil && (*DeletedAt).IsZero(),
	}
}

func (ft *ShopSupplierFilters) ByRid(Rid dot.ID) *sq.ColumnFilter {
	return &sq.ColumnFilter{
		Prefix: &ft.prefix,
		Column: "rid",
		Value:  Rid,
		IsNil:  Rid == 0,
	}
}

func (ft *ShopSupplierFilters) ByRidPtr(Rid *dot.ID) *sq.ColumnFilterPtr {
	return &sq.ColumnFilterPtr{
		Prefix: &ft.prefix,
		Column: "rid",
		Value:  Rid,
		IsNil:  Rid == nil,
		IsZero: Rid != nil && (*Rid) == 0,
	}
}