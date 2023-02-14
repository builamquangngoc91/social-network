// +build !generator

// Code generated by generator convert. DO NOT EDIT.

package convert

import (
	time "time"

	fbpaging "o.o/api/fabo/fbpaging"
	fbpagemodel "o.o/backend/com/fabo/main/fbpage/model"
	conversion "o.o/backend/pkg/common/conversion"
)

/*
Custom conversions: (none)

Ignored functions: (none)
*/

func RegisterConversions(s *conversion.Scheme) {
	registerConversions(s)
}

func registerConversions(s *conversion.Scheme) {
	s.Register((*fbpagemodel.ExternalCategory)(nil), (*fbpaging.ExternalCategory)(nil), func(arg, out interface{}) error {
		Convert_fbpagemodel_ExternalCategory_fbpaging_ExternalCategory(arg.(*fbpagemodel.ExternalCategory), out.(*fbpaging.ExternalCategory))
		return nil
	})
	s.Register(([]*fbpagemodel.ExternalCategory)(nil), (*[]*fbpaging.ExternalCategory)(nil), func(arg, out interface{}) error {
		out0 := Convert_fbpagemodel_ExternalCategories_fbpaging_ExternalCategories(arg.([]*fbpagemodel.ExternalCategory))
		*out.(*[]*fbpaging.ExternalCategory) = out0
		return nil
	})
	s.Register((*fbpaging.ExternalCategory)(nil), (*fbpagemodel.ExternalCategory)(nil), func(arg, out interface{}) error {
		Convert_fbpaging_ExternalCategory_fbpagemodel_ExternalCategory(arg.(*fbpaging.ExternalCategory), out.(*fbpagemodel.ExternalCategory))
		return nil
	})
	s.Register(([]*fbpaging.ExternalCategory)(nil), (*[]*fbpagemodel.ExternalCategory)(nil), func(arg, out interface{}) error {
		out0 := Convert_fbpaging_ExternalCategories_fbpagemodel_ExternalCategories(arg.([]*fbpaging.ExternalCategory))
		*out.(*[]*fbpagemodel.ExternalCategory) = out0
		return nil
	})
	s.Register((*fbpagemodel.FbExternalPage)(nil), (*fbpaging.FbExternalPage)(nil), func(arg, out interface{}) error {
		Convert_fbpagemodel_FbExternalPage_fbpaging_FbExternalPage(arg.(*fbpagemodel.FbExternalPage), out.(*fbpaging.FbExternalPage))
		return nil
	})
	s.Register(([]*fbpagemodel.FbExternalPage)(nil), (*[]*fbpaging.FbExternalPage)(nil), func(arg, out interface{}) error {
		out0 := Convert_fbpagemodel_FbExternalPages_fbpaging_FbExternalPages(arg.([]*fbpagemodel.FbExternalPage))
		*out.(*[]*fbpaging.FbExternalPage) = out0
		return nil
	})
	s.Register((*fbpaging.FbExternalPage)(nil), (*fbpagemodel.FbExternalPage)(nil), func(arg, out interface{}) error {
		Convert_fbpaging_FbExternalPage_fbpagemodel_FbExternalPage(arg.(*fbpaging.FbExternalPage), out.(*fbpagemodel.FbExternalPage))
		return nil
	})
	s.Register(([]*fbpaging.FbExternalPage)(nil), (*[]*fbpagemodel.FbExternalPage)(nil), func(arg, out interface{}) error {
		out0 := Convert_fbpaging_FbExternalPages_fbpagemodel_FbExternalPages(arg.([]*fbpaging.FbExternalPage))
		*out.(*[]*fbpagemodel.FbExternalPage) = out0
		return nil
	})
	s.Register((*fbpaging.CreateFbExternalPageArgs)(nil), (*fbpaging.FbExternalPage)(nil), func(arg, out interface{}) error {
		Apply_fbpaging_CreateFbExternalPageArgs_fbpaging_FbExternalPage(arg.(*fbpaging.CreateFbExternalPageArgs), out.(*fbpaging.FbExternalPage))
		return nil
	})
	s.Register((*fbpagemodel.FbExternalPageInternal)(nil), (*fbpaging.FbExternalPageInternal)(nil), func(arg, out interface{}) error {
		Convert_fbpagemodel_FbExternalPageInternal_fbpaging_FbExternalPageInternal(arg.(*fbpagemodel.FbExternalPageInternal), out.(*fbpaging.FbExternalPageInternal))
		return nil
	})
	s.Register(([]*fbpagemodel.FbExternalPageInternal)(nil), (*[]*fbpaging.FbExternalPageInternal)(nil), func(arg, out interface{}) error {
		out0 := Convert_fbpagemodel_FbExternalPageInternals_fbpaging_FbExternalPageInternals(arg.([]*fbpagemodel.FbExternalPageInternal))
		*out.(*[]*fbpaging.FbExternalPageInternal) = out0
		return nil
	})
	s.Register((*fbpaging.FbExternalPageInternal)(nil), (*fbpagemodel.FbExternalPageInternal)(nil), func(arg, out interface{}) error {
		Convert_fbpaging_FbExternalPageInternal_fbpagemodel_FbExternalPageInternal(arg.(*fbpaging.FbExternalPageInternal), out.(*fbpagemodel.FbExternalPageInternal))
		return nil
	})
	s.Register(([]*fbpaging.FbExternalPageInternal)(nil), (*[]*fbpagemodel.FbExternalPageInternal)(nil), func(arg, out interface{}) error {
		out0 := Convert_fbpaging_FbExternalPageInternals_fbpagemodel_FbExternalPageInternals(arg.([]*fbpaging.FbExternalPageInternal))
		*out.(*[]*fbpagemodel.FbExternalPageInternal) = out0
		return nil
	})
	s.Register((*fbpaging.CreateFbExternalPageInternalArgs)(nil), (*fbpaging.FbExternalPageInternal)(nil), func(arg, out interface{}) error {
		Apply_fbpaging_CreateFbExternalPageInternalArgs_fbpaging_FbExternalPageInternal(arg.(*fbpaging.CreateFbExternalPageInternalArgs), out.(*fbpaging.FbExternalPageInternal))
		return nil
	})
}

//-- convert o.o/api/fabo/fbpaging.ExternalCategory --//

func Convert_fbpagemodel_ExternalCategory_fbpaging_ExternalCategory(arg *fbpagemodel.ExternalCategory, out *fbpaging.ExternalCategory) *fbpaging.ExternalCategory {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &fbpaging.ExternalCategory{}
	}
	convert_fbpagemodel_ExternalCategory_fbpaging_ExternalCategory(arg, out)
	return out
}

func convert_fbpagemodel_ExternalCategory_fbpaging_ExternalCategory(arg *fbpagemodel.ExternalCategory, out *fbpaging.ExternalCategory) {
	out.ID = arg.ID     // simple assign
	out.Name = arg.Name // simple assign
}

func Convert_fbpagemodel_ExternalCategories_fbpaging_ExternalCategories(args []*fbpagemodel.ExternalCategory) (outs []*fbpaging.ExternalCategory) {
	if args == nil {
		return nil
	}
	tmps := make([]fbpaging.ExternalCategory, len(args))
	outs = make([]*fbpaging.ExternalCategory, len(args))
	for i := range tmps {
		outs[i] = Convert_fbpagemodel_ExternalCategory_fbpaging_ExternalCategory(args[i], &tmps[i])
	}
	return outs
}

func Convert_fbpaging_ExternalCategory_fbpagemodel_ExternalCategory(arg *fbpaging.ExternalCategory, out *fbpagemodel.ExternalCategory) *fbpagemodel.ExternalCategory {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &fbpagemodel.ExternalCategory{}
	}
	convert_fbpaging_ExternalCategory_fbpagemodel_ExternalCategory(arg, out)
	return out
}

func convert_fbpaging_ExternalCategory_fbpagemodel_ExternalCategory(arg *fbpaging.ExternalCategory, out *fbpagemodel.ExternalCategory) {
	out.ID = arg.ID     // simple assign
	out.Name = arg.Name // simple assign
}

func Convert_fbpaging_ExternalCategories_fbpagemodel_ExternalCategories(args []*fbpaging.ExternalCategory) (outs []*fbpagemodel.ExternalCategory) {
	if args == nil {
		return nil
	}
	tmps := make([]fbpagemodel.ExternalCategory, len(args))
	outs = make([]*fbpagemodel.ExternalCategory, len(args))
	for i := range tmps {
		outs[i] = Convert_fbpaging_ExternalCategory_fbpagemodel_ExternalCategory(args[i], &tmps[i])
	}
	return outs
}

//-- convert o.o/api/fabo/fbpaging.FbExternalPage --//

func Convert_fbpagemodel_FbExternalPage_fbpaging_FbExternalPage(arg *fbpagemodel.FbExternalPage, out *fbpaging.FbExternalPage) *fbpaging.FbExternalPage {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &fbpaging.FbExternalPage{}
	}
	convert_fbpagemodel_FbExternalPage_fbpaging_FbExternalPage(arg, out)
	return out
}

func convert_fbpagemodel_FbExternalPage_fbpaging_FbExternalPage(arg *fbpagemodel.FbExternalPage, out *fbpaging.FbExternalPage) {
	out.ID = arg.ID                             // simple assign
	out.ExternalID = arg.ExternalID             // simple assign
	out.ExternalUserID = arg.ExternalUserID     // simple assign
	out.ShopID = arg.ShopID                     // simple assign
	out.ExternalName = arg.ExternalName         // simple assign
	out.ExternalCategory = arg.ExternalCategory // simple assign
	out.ExternalCategoryList = Convert_fbpagemodel_ExternalCategories_fbpaging_ExternalCategories(arg.ExternalCategoryList)
	out.ExternalTasks = arg.ExternalTasks             // simple assign
	out.ExternalPermissions = arg.ExternalPermissions // simple assign
	out.ExternalImageURL = arg.ExternalImageURL       // simple assign
	out.Status = arg.Status                           // simple assign
	out.ConnectionStatus = arg.ConnectionStatus       // simple assign
	out.CreatedAt = arg.CreatedAt                     // simple assign
	out.UpdatedAt = arg.UpdatedAt                     // simple assign
}

func Convert_fbpagemodel_FbExternalPages_fbpaging_FbExternalPages(args []*fbpagemodel.FbExternalPage) (outs []*fbpaging.FbExternalPage) {
	if args == nil {
		return nil
	}
	tmps := make([]fbpaging.FbExternalPage, len(args))
	outs = make([]*fbpaging.FbExternalPage, len(args))
	for i := range tmps {
		outs[i] = Convert_fbpagemodel_FbExternalPage_fbpaging_FbExternalPage(args[i], &tmps[i])
	}
	return outs
}

func Convert_fbpaging_FbExternalPage_fbpagemodel_FbExternalPage(arg *fbpaging.FbExternalPage, out *fbpagemodel.FbExternalPage) *fbpagemodel.FbExternalPage {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &fbpagemodel.FbExternalPage{}
	}
	convert_fbpaging_FbExternalPage_fbpagemodel_FbExternalPage(arg, out)
	return out
}

func convert_fbpaging_FbExternalPage_fbpagemodel_FbExternalPage(arg *fbpaging.FbExternalPage, out *fbpagemodel.FbExternalPage) {
	out.ID = arg.ID                             // simple assign
	out.ShopID = arg.ShopID                     // simple assign
	out.ExternalID = arg.ExternalID             // simple assign
	out.ExternalUserID = arg.ExternalUserID     // simple assign
	out.ExternalName = arg.ExternalName         // simple assign
	out.ExternalTasks = arg.ExternalTasks       // simple assign
	out.ExternalCategory = arg.ExternalCategory // simple assign
	out.ExternalCategoryList = Convert_fbpaging_ExternalCategories_fbpagemodel_ExternalCategories(arg.ExternalCategoryList)
	out.ExternalPermissions = arg.ExternalPermissions // simple assign
	out.ExternalImageURL = arg.ExternalImageURL       // simple assign
	out.ConnectionStatus = arg.ConnectionStatus       // simple assign
	out.Status = arg.Status                           // simple assign
	out.CreatedAt = arg.CreatedAt                     // simple assign
	out.UpdatedAt = arg.UpdatedAt                     // simple assign
	out.DeletedAt = time.Time{}                       // zero value
}

func Convert_fbpaging_FbExternalPages_fbpagemodel_FbExternalPages(args []*fbpaging.FbExternalPage) (outs []*fbpagemodel.FbExternalPage) {
	if args == nil {
		return nil
	}
	tmps := make([]fbpagemodel.FbExternalPage, len(args))
	outs = make([]*fbpagemodel.FbExternalPage, len(args))
	for i := range tmps {
		outs[i] = Convert_fbpaging_FbExternalPage_fbpagemodel_FbExternalPage(args[i], &tmps[i])
	}
	return outs
}

func Apply_fbpaging_CreateFbExternalPageArgs_fbpaging_FbExternalPage(arg *fbpaging.CreateFbExternalPageArgs, out *fbpaging.FbExternalPage) *fbpaging.FbExternalPage {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &fbpaging.FbExternalPage{}
	}
	apply_fbpaging_CreateFbExternalPageArgs_fbpaging_FbExternalPage(arg, out)
	return out
}

func apply_fbpaging_CreateFbExternalPageArgs_fbpaging_FbExternalPage(arg *fbpaging.CreateFbExternalPageArgs, out *fbpaging.FbExternalPage) {
	out.ID = arg.ID                                     // simple assign
	out.ExternalID = arg.ExternalID                     // simple assign
	out.ExternalUserID = arg.ExternalUserID             // simple assign
	out.ShopID = arg.ShopID                             // simple assign
	out.ExternalName = arg.ExternalName                 // simple assign
	out.ExternalCategory = arg.ExternalCategory         // simple assign
	out.ExternalCategoryList = arg.ExternalCategoryList // simple assign
	out.ExternalTasks = arg.ExternalTasks               // simple assign
	out.ExternalPermissions = arg.ExternalPermissions   // simple assign
	out.ExternalImageURL = arg.ExternalImageURL         // simple assign
	out.Status = arg.Status                             // simple assign
	out.ConnectionStatus = arg.ConnectionStatus         // simple assign
	out.CreatedAt = time.Time{}                         // zero value
	out.UpdatedAt = time.Time{}                         // zero value
}

//-- convert o.o/api/fabo/fbpaging.FbExternalPageInternal --//

func Convert_fbpagemodel_FbExternalPageInternal_fbpaging_FbExternalPageInternal(arg *fbpagemodel.FbExternalPageInternal, out *fbpaging.FbExternalPageInternal) *fbpaging.FbExternalPageInternal {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &fbpaging.FbExternalPageInternal{}
	}
	convert_fbpagemodel_FbExternalPageInternal_fbpaging_FbExternalPageInternal(arg, out)
	return out
}

func convert_fbpagemodel_FbExternalPageInternal_fbpaging_FbExternalPageInternal(arg *fbpagemodel.FbExternalPageInternal, out *fbpaging.FbExternalPageInternal) {
	out.ID = arg.ID                 // simple assign
	out.ExternalID = arg.ExternalID // simple assign
	out.Token = arg.Token           // simple assign
	out.UpdatedAt = arg.UpdatedAt   // simple assign
}

func Convert_fbpagemodel_FbExternalPageInternals_fbpaging_FbExternalPageInternals(args []*fbpagemodel.FbExternalPageInternal) (outs []*fbpaging.FbExternalPageInternal) {
	if args == nil {
		return nil
	}
	tmps := make([]fbpaging.FbExternalPageInternal, len(args))
	outs = make([]*fbpaging.FbExternalPageInternal, len(args))
	for i := range tmps {
		outs[i] = Convert_fbpagemodel_FbExternalPageInternal_fbpaging_FbExternalPageInternal(args[i], &tmps[i])
	}
	return outs
}

func Convert_fbpaging_FbExternalPageInternal_fbpagemodel_FbExternalPageInternal(arg *fbpaging.FbExternalPageInternal, out *fbpagemodel.FbExternalPageInternal) *fbpagemodel.FbExternalPageInternal {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &fbpagemodel.FbExternalPageInternal{}
	}
	convert_fbpaging_FbExternalPageInternal_fbpagemodel_FbExternalPageInternal(arg, out)
	return out
}

func convert_fbpaging_FbExternalPageInternal_fbpagemodel_FbExternalPageInternal(arg *fbpaging.FbExternalPageInternal, out *fbpagemodel.FbExternalPageInternal) {
	out.ID = arg.ID                 // simple assign
	out.ExternalID = arg.ExternalID // simple assign
	out.Token = arg.Token           // simple assign
	out.UpdatedAt = arg.UpdatedAt   // simple assign
}

func Convert_fbpaging_FbExternalPageInternals_fbpagemodel_FbExternalPageInternals(args []*fbpaging.FbExternalPageInternal) (outs []*fbpagemodel.FbExternalPageInternal) {
	if args == nil {
		return nil
	}
	tmps := make([]fbpagemodel.FbExternalPageInternal, len(args))
	outs = make([]*fbpagemodel.FbExternalPageInternal, len(args))
	for i := range tmps {
		outs[i] = Convert_fbpaging_FbExternalPageInternal_fbpagemodel_FbExternalPageInternal(args[i], &tmps[i])
	}
	return outs
}

func Apply_fbpaging_CreateFbExternalPageInternalArgs_fbpaging_FbExternalPageInternal(arg *fbpaging.CreateFbExternalPageInternalArgs, out *fbpaging.FbExternalPageInternal) *fbpaging.FbExternalPageInternal {
	if arg == nil {
		return nil
	}
	if out == nil {
		out = &fbpaging.FbExternalPageInternal{}
	}
	apply_fbpaging_CreateFbExternalPageInternalArgs_fbpaging_FbExternalPageInternal(arg, out)
	return out
}

func apply_fbpaging_CreateFbExternalPageInternalArgs_fbpaging_FbExternalPageInternal(arg *fbpaging.CreateFbExternalPageInternalArgs, out *fbpaging.FbExternalPageInternal) {
	out.ID = arg.ID                 // simple assign
	out.ExternalID = arg.ExternalID // simple assign
	out.Token = arg.Token           // simple assign
	out.UpdatedAt = time.Time{}     // zero value
}