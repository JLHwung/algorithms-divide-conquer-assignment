package week1

import (
	"fmt"
	"math/big"
	"testing"
)

func ExampleKaratsubaMultiplyBig() {
	x, _ := new(big.Int).SetString("11717829880366207009516117596335367088558084999998952205599979459063929499736583746670572176471460312928594829675428279466566527115212748467589894601965568", 10)
	y, _ := new(big.Int).SetString("3239475104050450443565264378728065788649097520952449527834792452971981976143292558073856937958553180532878928001494706097394108577585732452307673444020333", 10)
	z := new(big.Int)
	KaratsubaMultiply(x, y, z)
	fmt.Printf("%d\n", z)
	// Output: 37959618170944795725324538013290906137018244916907906527287401498999829081375415975866536071466546752682477317662379832350293021462358879602173480320695337457396899676082676610811111847847061394808891260195708317334274309524938664839670497441006075850183805997651645775459987212670318230587996802917957894144
}

func BenchmarkKaratsubaMultiply(b *testing.B) {
	x, _ := new(big.Int).SetString("11717829880366207009516117596335367088558084999998952205599979459063929499736583746670572176471460312928594829675428279466566527115212748467589894601965568", 10)
	y, _ := new(big.Int).SetString("3239475104050450443565264378728065788649097520952449527834792452971981976143292558073856937958553180532878928001494706097394108577585732452307673444020333", 10)
	z := new(big.Int)
	for n := 0; n < b.N; n++ {
		KaratsubaMultiply(x, y, z)
	}
}

func BenchmarkMul(b *testing.B) {
	x, _ := new(big.Int).SetString("11717829880366207009516117596335367088558084999998952205599979459063929499736583746670572176471460312928594829675428279466566527115212748467589894601965568", 10)
	y, _ := new(big.Int).SetString("3239475104050450443565264378728065788649097520952449527834792452971981976143292558073856937958553180532878928001494706097394108577585732452307673444020333", 10)
	z := new(big.Int)
	for n := 0; n < b.N; n++ {
		Mul(x, y, z)
	}
}