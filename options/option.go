package options

import "github.com/RangelReale/instruct"

type AnyOption[IT any, DC instruct.DecodeContext] interface {
	isAnyOption()
}

type AnyTypeOption[IT any, DC instruct.DecodeContext] interface {
	isAnyTypeOption()
}

type DefaultOption[IT any, DC instruct.DecodeContext, DOPTT any] interface {
	AnyOption[IT, DC]
	ApplyDefaultOption(*DOPTT)
}

type TypeDefaultOption[IT any, DC instruct.DecodeContext, TOPTT any] interface {
	AnyTypeOption[IT, DC]
	ApplyTypeDefaultOption(*TOPTT)
}

type DecodeOption[IT any, DC instruct.DecodeContext, COPTT any] interface {
	AnyOption[IT, DC]
	ApplyDecodeOption(*COPTT)
}

type TypeDecodeOption[IT any, DC instruct.DecodeContext, COPTT any] interface {
	AnyTypeOption[IT, DC]
	ApplyTypeDecodeOption(*COPTT)
}

type DefaultAndTypeDefaultOption[IT any, DC instruct.DecodeContext, DOPTT any, TOPTT any] interface {
	DefaultOption[IT, DC, DOPTT]
	TypeDefaultOption[IT, DC, TOPTT]
}

type DefaultAndDecodeOption[IT any, DC instruct.DecodeContext, DOPTT any, COPTT any] interface {
	DefaultOption[IT, DC, DOPTT]
	DecodeOption[IT, DC, COPTT]
}

type TypeDefaultAndTypeDecodeOption[IT any, DC instruct.DecodeContext, TOPTT any, COPTT any] interface {
	TypeDefaultOption[IT, DC, TOPTT]
	TypeDecodeOption[IT, DC, COPTT]
}

type DefaultAndTypeDefaultDecodeOption[IT any, DC instruct.DecodeContext, DCOPTT any, TCOPTT any] interface {
	DecodeOption[IT, DC, DCOPTT]
	TypeDecodeOption[IT, DC, TCOPTT]
}

type TypeDefaultAndDecodeOption[IT any, DC instruct.DecodeContext, TOPTT any, DCOPTT any] interface {
	TypeDefaultOption[IT, DC, TOPTT]
	DecodeOption[IT, DC, DCOPTT]
}

type FullOption[IT any, DC instruct.DecodeContext, DOPTT any, TOPTT any, DCOPTT any, TCOPTT any] interface {
	DefaultOption[IT, DC, DOPTT]
	TypeDefaultOption[IT, DC, TOPTT]
	DecodeOption[IT, DC, DCOPTT]
	TypeDecodeOption[IT, DC, TCOPTT]
}
