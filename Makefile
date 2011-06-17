include $(GOROOT)/src/Make.inc

TARG=raw

GOFILES=\
	container.go\
	mapping.go\
	sequence.go\
	byte_slice.go\
	slice.go\
	map.go\
	channel.go\
	reslice.go\
	runtime.go\
	intmap.go

include $(GOROOT)/src/Make.pkg