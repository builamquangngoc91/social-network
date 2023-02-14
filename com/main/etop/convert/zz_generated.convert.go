// +build !generator

// Code generated by generator convert. DO NOT EDIT.

package convert

import (
	meta "o.o/api/meta"
	conversion "o.o/backend/pkg/common/conversion"
	etopmodel "o.o/backend/pkg/etop/model"
)

/*
Custom conversions:
    Error           // in use
    ErrorToModel    // in use

Ignored functions:
    Status3           // params are not pointer to named types
    Status3ToModel    // params are not pointer to named types
    Status4           // params are not pointer to named types
    Status4ToModel    // params are not pointer to named types
    Status5           // params are not pointer to named types
    Status5ToModel    // params are not pointer to named types
*/

func RegisterConversions(s *conversion.Scheme) {
	registerConversions(s)
}

func registerConversions(s *conversion.Scheme) {
	s.Register((*etopmodel.Error)(nil), (*meta.Error)(nil), func(arg, out interface{}) error {
		Convert_etopmodel_Error_meta_Error(arg.(*etopmodel.Error), out.(*meta.Error))
		return nil
	})
	s.Register(([]*etopmodel.Error)(nil), (*[]*meta.Error)(nil), func(arg, out interface{}) error {
		out0 := Convert_etopmodel_Errors_meta_Errors(arg.([]*etopmodel.Error))
		*out.(*[]*meta.Error) = out0
		return nil
	})
	s.Register((*meta.Error)(nil), (*etopmodel.Error)(nil), func(arg, out interface{}) error {
		Convert_meta_Error_etopmodel_Error(arg.(*meta.Error), out.(*etopmodel.Error))
		return nil
	})
	s.Register(([]*meta.Error)(nil), (*[]*etopmodel.Error)(nil), func(arg, out interface{}) error {
		out0 := Convert_meta_Errors_etopmodel_Errors(arg.([]*meta.Error))
		*out.(*[]*etopmodel.Error) = out0
		return nil
	})
}

//-- convert o.o/api/meta.Error --//

func Convert_etopmodel_Error_meta_Error(arg *etopmodel.Error, out *meta.Error) *meta.Error {
	return Error(arg)
}

func convert_etopmodel_Error_meta_Error(arg *etopmodel.Error, out *meta.Error) {
	out.Code = arg.Code // simple assign
	out.Msg = arg.Msg   // simple assign
	out.Meta = nil      // types do not match
}

func Convert_etopmodel_Errors_meta_Errors(args []*etopmodel.Error) (outs []*meta.Error) {
	if args == nil {
		return nil
	}
	tmps := make([]meta.Error, len(args))
	outs = make([]*meta.Error, len(args))
	for i := range tmps {
		outs[i] = Convert_etopmodel_Error_meta_Error(args[i], &tmps[i])
	}
	return outs
}

func Convert_meta_Error_etopmodel_Error(arg *meta.Error, out *etopmodel.Error) *etopmodel.Error {
	return ErrorToModel(arg)
}

func convert_meta_Error_etopmodel_Error(arg *meta.Error, out *etopmodel.Error) {
	out.Code = arg.Code // simple assign
	out.Msg = arg.Msg   // simple assign
	out.Meta = nil      // types do not match
}

func Convert_meta_Errors_etopmodel_Errors(args []*meta.Error) (outs []*etopmodel.Error) {
	if args == nil {
		return nil
	}
	tmps := make([]etopmodel.Error, len(args))
	outs = make([]*etopmodel.Error, len(args))
	for i := range tmps {
		outs[i] = Convert_meta_Error_etopmodel_Error(args[i], &tmps[i])
	}
	return outs
}
