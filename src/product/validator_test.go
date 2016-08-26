package product

import (
    "testing"
)

type getProductNameTest struct {
    in, out string
}
var getProductNameTests = []getProductNameTest{
    getProductNameTest{"α5100 ILCE-5100L パワーズームレンズキット", "α5100 ILCE-5100L"},
    getProductNameTest{"EOS 5D Mark III ボディ", "EOS 5D Mark III"},
}
func TestGetProductName(t *testing.T){
    for _, dt := range getProductNameTests {
		v := getProductName(dt.in)
		if v != dt.out {
			t.Errorf("getProductName(%d) = %d, want %d.", dt.in, v, dt.out)
		}
	}
}

type getPixelsValuesTest struct {
    in1 string
    out1, out2 int
}
var getPixelsValuesTests = []getPixelsValuesTest{
    getPixelsValuesTest{"1649万画素(総画素)1628万画素(有効画素)", 1649, 1628},
    getPixelsValuesTest{"2430万画素(有効画素)", -1, 2430},
    getPixelsValuesTest{"2430万画素(総画素)", 2430, -1},
}
func TestGetPixelsValues(t *testing.T){
    for _, dt := range getPixelsValuesTests {
		v1, v2 := getPixelsValues(dt.in1)
		if v1 != dt.out1 || v2 != dt.out2{
			t.Errorf("getPixelsValue(%d) = (%d, %d) , want %d, %d.", dt.in1, v1, v2, dt.out1, dt.out2)
		}
	}
}