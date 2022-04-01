package ttf

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"
)

func parse[T any](data io.Reader) T {
	var val T
	readStructure(&val, data)
	return val
}

func readStructure(f any, data io.Reader) {
	t := reflect.ValueOf(f).Elem()
	for i := 0; i < t.NumField(); i++ {
		k := t.Field(i)
		switch k.Kind() {
		case reflect.Uint8:
			k.Set(reflect.ValueOf(ui8(data)))
		case reflect.Uint16:
			k.Set(reflect.ValueOf(ui16(data)))
		case reflect.Uint32:
			k.Set(reflect.ValueOf(ui32(data)))
		case reflect.Uint64:
			k.Set(reflect.ValueOf(ui64(data)))
		default:
			t2 := t.Type()
			panic(fmt.Sprintf("unknown mapping type `%v` to set `%v.%v`", k.Kind(), t2.Name(), t2.Field(i).Name))
		}
	}
}

func ui64(data io.Reader) uint64 {
	return binary.BigEndian.Uint64(readNBytes(sizeOfUint64, data))
}

func ui32(data io.Reader) uint32 {
	return binary.BigEndian.Uint32(readNBytes(sizeOfUint32, data))
}

func ui16(data io.Reader) uint16 {
	return binary.BigEndian.Uint16(readNBytes(sizeOfUint16, data))
}

func ui8(data io.Reader) uint8 {
	return readNBytes(sizeOfUint8, data)[0]
}

func readNBytes(n int, r io.Reader) []byte {
	buff := make([]byte, n)
	if na, rErr := r.Read(buff); rErr != nil {
		panic(rErr)
	} else if na != n {
		panic(errors.New(fmt.Sprintf("unable to read %d bytes from the header, actually read %d", n, na)))
	}
	return buff
}

func uint32ToString(n uint32) string {
	buff := make([]byte, 4)
	binary.BigEndian.PutUint32(buff, n)
	return string(buff)
}
