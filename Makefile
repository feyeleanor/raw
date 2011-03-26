include $(GOROOT)/src/Make.inc

TARG=raw

GOFILES=\
	byte_slice.go\
	slice.go\
	map.go\
	channel.go\
	reslice.go\
	runtime.go

include $(GOROOT)/src/Make.pkg