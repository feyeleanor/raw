include $(GOROOT)/src/Make.inc

TARG=raw

GOFILES=\
	byte_slice.go\
	reslice.go\
	runtime.go

include $(GOROOT)/src/Make.pkg