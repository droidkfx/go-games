package ttf

//goland:noinspection GoUnusedConst
const (
	sizeOfFixed32      = sizeOfUint16 * 2
	sizeShortFrac16    = sizeOfInt16
	sizeOfFWord        = sizeOfInt16
	sizeOfUFWord       = sizeOfUint16
	sizeOfF2Dot14      = sizeOfInt16
	sizeOfLongDateTime = sizeOfInt64
	sizeOfUint64       = 8
	sizeOfUint32       = 4
	sizeOfUint16       = 2
	sizeOfUint8        = 1
	sizeOfInt64        = 8
	sizeOfInt32        = 4
	sizeOfInt16        = 2
	sizeOfInt8         = 1
)

// Fixed32 16.16-bit signed fixed-point number
type Fixed32 struct {
	Top uint16
	Bot uint16
}

// ShortFrac16 is an int16_t with a bias of 14. This means it can represent numbers between
// 1.999 (0x7fff) and -2.0 (0x8000). 1.0 is stored as 16384 (0x4000) and -1.0 is stored as -16384 (0xc000).
type ShortFrac16 int16

// FWord 16-bit signed integer that describes a quantity in FUnits, the smallest measurable distance in em space.
type FWord int16

// UFWord 16-bit unsigned integer that describes a quantity in FUnits, the smallest measurable distance in em space.
type UFWord uint16

// F2Dot14 16-bit signed fixed number with the low 14 bits representing fraction.
type F2Dot14 int16

// LongDateTime The long internal format of a date in seconds since 12:00 midnight, January 1, 1904. It is represented
// as a signed 64-bit integer.
type LongDateTime int64
