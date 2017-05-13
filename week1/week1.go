package week1

import (
	"bytes"
	"math/big"
)

// norm will add leading 0 on x, y s.t. len(x) = len(y)
func norm(x, y []byte) (a, b []byte) {
	// compute max length of x, y
	xLen := len(x)
	yLen := len(y)
	length := xLen
	if length < yLen {
		length = yLen
		a := append(bytes.Repeat([]byte{byte(0)}, yLen-xLen), x...)
		b := y
		return a, b
	} else if length > yLen {
		b := append(bytes.Repeat([]byte{byte(0)}, xLen-yLen), y...)
		a := x
		return a, b
	}
	return x, y

}

var karatsubaThreshold = 40

// KaratsubaMultiply multiplies x*y and leaves result in z
func KaratsubaMultiply(x, y, z *big.Int) {
	xBytes := x.Bytes()
	yBytes := y.Bytes()

	// normalize xBytes and yBytes
	xBytes, yBytes = norm(xBytes, yBytes)
	// compute max length of x, y
	length := len(xBytes)

	//
	if length <= karatsubaThreshold {
		z.Mul(x, y)
		return
	}

	// split x, y s.t. x = B^{\lceil n/2 \rceil }a + b, y = B^{\lceil n/2 \rceil }c + d
	halfLength := length / 2

	a := new(big.Int).SetBytes(xBytes[:halfLength])
	b := new(big.Int).SetBytes(xBytes[halfLength:])

	c := new(big.Int).SetBytes(yBytes[:halfLength])
	d := new(big.Int).SetBytes(yBytes[halfLength:])

	ac := new(big.Int)
	bd := new(big.Int)
	abcd := new(big.Int)

	// compute ac = a*c, bd = b*d, abcd = (a+b)*(c+d) - ac - bd
	KaratsubaMultiply(a, c, ac)
	KaratsubaMultiply(b, d, bd)
	KaratsubaMultiply(a.Add(a, b), c.Add(c, d), abcd)
	abcd.Sub(abcd, ac).Sub(abcd, bd)

	shift := 8 * uint(length-halfLength)

	ac.Lsh(ac, shift*2)
	abcd.Lsh(abcd, shift)
	z.Add(ac, abcd).Add(z, bd)

}

func Mul(x, y, z *big.Int) {
	z.Mul(x, y)
}
