include $(GOROOT)/src/Make.inc

TARG=raw

GOFILES=\
	type.go\
	runtime.go\
	byte_slice.go\
	reslice.go

include $(GOROOT)/src/Make.pkg