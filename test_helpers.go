package raw

type Fatal interface {
	Fatal(...interface{})
	Fatalf(string, ...interface{})
}

func FailOnBadBufferSize(f Fatal, good_length, good_capacity bool) {
	switch {
	case !good_length:
		f.Fatalf("byte buffer lengths differ")
	case !good_capacity:
		f.Fatalf("byte buffer capacities differ")
	}
}