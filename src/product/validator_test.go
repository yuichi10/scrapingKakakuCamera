package product

import (
    "testing"
)

//商品名
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

//画素数
type getPixelsValuesTest struct {
    in1 string
    out1, out2 int
}
var getPixelsValuesTests = []getPixelsValuesTest{
    getPixelsValuesTest{"1649万画素(総画素)1628万画素(有効画素)", 1649, 1628},
    getPixelsValuesTest{"2430万画素(有効画素)", -1, 2430},
    getPixelsValuesTest{"2430万画素(総画素)", 2430, -1},
    getPixelsValuesTest{"", -1, -1},
}
func TestGetPixelsValues(t *testing.T){
    for _, dt := range getPixelsValuesTests {
		v1, v2 := getPixelsValues(dt.in1)
		if v1 != dt.out1 || v2 != dt.out2{
			t.Errorf("getPixelsValue(%d) = (%d, %d) , want %d, %d.", dt.in1, v1, v2, dt.out1, dt.out2)
		}
	}
}

//イメージセンサー
type getImageSensorTypeTest struct {
    in, out string
}
var getImageSensorTypeTests = []getImageSensorTypeTest{
    getImageSensorTypeTest{"APS-C23.7mm×15.7mmCMOS", "APS-C"},
    getImageSensorTypeTest{"フルサイズ35.8mm×23.9mmCMOS", "フルサイズ"},
    getImageSensorTypeTest{"", "その他"},
}
func TestGetImageSensorType(t *testing.T){
    for _, dt := range getImageSensorTypeTests {
		v := getImageSensorType(dt.in)
		if v != dt.out {
			t.Errorf("getProductName(%d) = %d, want %d.", dt.in, v, dt.out)
		}
	}
}