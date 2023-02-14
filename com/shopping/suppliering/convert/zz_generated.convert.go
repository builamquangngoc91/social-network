// +build !generator

// Code generated by generator convert. DO NOT EDIT.

package convert

import (
	time "time"

	suppliering "o.o/api/shopping/suppliering"
	status3 "o.o/api/top/types/etc/status3"
	supplieringmodel "o.o/backend/com/shopping/suppliering/model"
	conversion "o.o/backend/pkg/common/conversion"
)

/*
Custom conversions:
    createShopSupplier    // in use
    updateShopSupplier    // in use

Ignored functions:
    GenerateCode    // params are not pointer to named types
*/

func RegisterConversions(s *conversion.Scheme) {
	registerConversions(s)
}

func registerConversions(s *conversion.Scheme) {
	s.Register((*supplieringmodel.ShopSupplier)(nil), (*suppliering.ShopSupplier)(nil), func(arg, out interface{}) error {
		Convert_supplieringmodel_ShopSupplier_suppliering_ShopSupplier(arg.(*supplieringmodel.ShopSupplier), out.(*suppliering.ShopSupplier))
		return nil
	})
	s.Register(([]*supplieringmodel.ShopSupplier)(nil), (*[]*suppliering.ShopSupplier)(nil), func(arg, out interface{}) error {
		out0 := Convert_supplieringmodel_ShopSuppliers_suppliering_ShopSuppliers(arg.([]*supplieringmodel.ShopSupplier))
		*out.(*[]*suppliering.ShopSupplier) = out0
		return nil
	})
	s.Register((*suppliering.ShopSupplier)(nil), (*supplieringmodel.ShopSupplier)(nil), func(arg, out interface{}) error {
		Convert_suppliering_ShopSupplier_supplieringmodel_ShopSupplier(arg.(*suppliering.ShopSupplier), out.(*supplieringmodel.ShopSupplier))
		return nil
	})
	s.Register(([]*suppliering.ShopSupplier)(nil), (*[]*supplieringmodel.ShopSupplier)(nil), func(arg, out interface{}) error {
		out0 := Convert_suppliering_ShopSuppliers_supplieringmodel_ShopSuppliers(arg.([]*suppliering.ShopSupplier))
		*out.(*[]*supplieringmodel.ShopSupplier) = out0
		return nil
	})
	s.Register((*suppliering.CreateSupplierArgs)(nil), (*suppliering.ShopSupplier)(nil), func(arg, out interface{}) error {
		Apply_suppliering_CreateSupplierArgs_suppliering_ShopSupplier(arg.(*suppliering.CreateSupplierArgs), out.(*suppliering.ShopSupplier))
		return nil
	})
	s.Register((*suppliering.UpdateSupplierArgs)(nil), (*suppliering.ShopSupplier)(nil), func(arg, out interface{}) error {
		Apply_suppliering_UpdateSupplierArgs_suppliering_ShopSupplier(arg.(*suppliering.UpdateSupplierArgs), out.(*suppliering.ShopSupplier))
		return nil
	})
}

//-- convert o.o/api/shopping/suppliering.ShopSupplier --//

func Convert_supplieringmodel_ShopSupplier_suppliering_ShopSupplier(arg *supplieringmodel.ShopSupplier, out *suppliering.ShopSupplier) *suppliering.ShopSupplier {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &suppliering.ShopSupplier{}
	}
	convert_supplieringmodel_ShopSupplier_suppliering_ShopSupplier(arg, out)
	return out
}

func convert_supplieringmodel_ShopSupplier_suppliering_ShopSupplier(arg *supplieringmodel.ShopSupplier, out *suppliering.ShopSupplier) {
	out.ID = arg.ID                               // simple assign
	out.ShopID = arg.ShopID                       // simple assign
	out.FullName = arg.FullName                   // simple assign
	out.Phone = arg.Phone                         // simple assign
	out.Code = arg.Code                           // simple assign
	out.CodeNorm = arg.CodeNorm                   // simple assign
	out.Email = arg.Email                         // simple assign
	out.CompanyName = arg.CompanyName             // simple assign
	out.TaxNumber = arg.TaxNumber                 // simple assign
	out.HeadquaterAddress = arg.HeadquaterAddress // simple assign
	out.Note = arg.Note                           // simple assign
	out.Status = status3.Status(arg.Status)       // simple conversion
	out.CreatedAt = arg.CreatedAt                 // simple assign
	out.UpdatedAt = arg.UpdatedAt                 // simple assign
}

func Convert_supplieringmodel_ShopSuppliers_suppliering_ShopSuppliers(args []*supplieringmodel.ShopSupplier) (outs []*suppliering.ShopSupplier) {
	if args == nil {
		return nil
	}
	tmps := make([]suppliering.ShopSupplier, len(args))
	outs = make([]*suppliering.ShopSupplier, len(args))
	for i := range tmps {
		outs[i] = Convert_supplieringmodel_ShopSupplier_suppliering_ShopSupplier(args[i], &tmps[i])
	}
	return outs
}

func Convert_suppliering_ShopSupplier_supplieringmodel_ShopSupplier(arg *suppliering.ShopSupplier, out *supplieringmodel.ShopSupplier) *supplieringmodel.ShopSupplier {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &supplieringmodel.ShopSupplier{}
	}
	convert_suppliering_ShopSupplier_supplieringmodel_ShopSupplier(arg, out)
	return out
}

func convert_suppliering_ShopSupplier_supplieringmodel_ShopSupplier(arg *suppliering.ShopSupplier, out *supplieringmodel.ShopSupplier) {
	out.ID = arg.ID                               // simple assign
	out.ShopID = arg.ShopID                       // simple assign
	out.FullName = arg.FullName                   // simple assign
	out.Phone = arg.Phone                         // simple assign
	out.Email = arg.Email                         // simple assign
	out.Code = arg.Code                           // simple assign
	out.CodeNorm = arg.CodeNorm                   // simple assign
	out.CompanyName = arg.CompanyName             // simple assign
	out.CompanyNameNorm = ""                      // zero value
	out.TaxNumber = arg.TaxNumber                 // simple assign
	out.HeadquaterAddress = arg.HeadquaterAddress // simple assign
	out.Note = arg.Note                           // simple assign
	out.FullNameNorm = ""                         // zero value
	out.PhoneNorm = ""                            // zero value
	out.Status = int(arg.Status)                  // simple conversion
	out.CreatedAt = arg.CreatedAt                 // simple assign
	out.UpdatedAt = arg.UpdatedAt                 // simple assign
	out.DeletedAt = time.Time{}                   // zero value
	out.Rid = 0                                   // zero value
}

func Convert_suppliering_ShopSuppliers_supplieringmodel_ShopSuppliers(args []*suppliering.ShopSupplier) (outs []*supplieringmodel.ShopSupplier) {
	if args == nil {
		return nil
	}
	tmps := make([]supplieringmodel.ShopSupplier, len(args))
	outs = make([]*supplieringmodel.ShopSupplier, len(args))
	for i := range tmps {
		outs[i] = Convert_suppliering_ShopSupplier_supplieringmodel_ShopSupplier(args[i], &tmps[i])
	}
	return outs
}

func Apply_suppliering_CreateSupplierArgs_suppliering_ShopSupplier(arg *suppliering.CreateSupplierArgs, out *suppliering.ShopSupplier) *suppliering.ShopSupplier {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &suppliering.ShopSupplier{}
	}
	createShopSupplier(arg, out)
	return out
}

func apply_suppliering_CreateSupplierArgs_suppliering_ShopSupplier(arg *suppliering.CreateSupplierArgs, out *suppliering.ShopSupplier) {
	out.ID = 0                                    // zero value
	out.ShopID = arg.ShopID                       // simple assign
	out.FullName = arg.FullName                   // simple assign
	out.Phone = arg.Phone                         // simple assign
	out.Code = ""                                 // zero value
	out.CodeNorm = 0                              // zero value
	out.Email = arg.Email                         // simple assign
	out.CompanyName = arg.CompanyName             // simple assign
	out.TaxNumber = arg.TaxNumber                 // simple assign
	out.HeadquaterAddress = arg.HeadquaterAddress // simple assign
	out.Note = arg.Note                           // simple assign
	out.Status = 0                                // zero value
	out.CreatedAt = time.Time{}                   // zero value
	out.UpdatedAt = time.Time{}                   // zero value
}

func Apply_suppliering_UpdateSupplierArgs_suppliering_ShopSupplier(arg *suppliering.UpdateSupplierArgs, out *suppliering.ShopSupplier) *suppliering.ShopSupplier {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &suppliering.ShopSupplier{}
	}
	updateShopSupplier(arg, out)
	return out
}

func apply_suppliering_UpdateSupplierArgs_suppliering_ShopSupplier(arg *suppliering.UpdateSupplierArgs, out *suppliering.ShopSupplier) {
	out.ID = out.ID                                                            // identifier
	out.ShopID = out.ShopID                                                    // identifier
	out.FullName = arg.FullName.Apply(out.FullName)                            // apply change
	out.Phone = arg.Phone.Apply(out.Phone)                                     // apply change
	out.Code = out.Code                                                        // no change
	out.CodeNorm = out.CodeNorm                                                // no change
	out.Email = arg.Email.Apply(out.Email)                                     // apply change
	out.CompanyName = arg.CompanyName.Apply(out.CompanyName)                   // apply change
	out.TaxNumber = arg.TaxNumber.Apply(out.TaxNumber)                         // apply change
	out.HeadquaterAddress = arg.HeadquaterAddress.Apply(out.HeadquaterAddress) // apply change
	out.Note = arg.Note.Apply(out.Note)                                        // apply change
	out.Status = out.Status                                                    // no change
	out.CreatedAt = out.CreatedAt                                              // no change
	out.UpdatedAt = out.UpdatedAt                                              // no change
}