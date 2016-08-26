package product

import (
    "strings"
    "strconv"
    _ "fmt"
)

func TestVlid(){
    getPixelsValues("1649万画素(総画素)1628万画素(有効画素)")
}

//プロダクトのタイプを取得
func getProductType(val string) int {
	for key, value := range ProductTypeList {
		if key == val {
			return value
		}
	}
	return ProductTypeList["other"]
}
//レンズマウントタイプを取得
func getLensMountType(val string) int {
	for key, value := range LensMountTypeList {
		if key == val {
			return value
		}
	}
	return LensMountTypeList["other"]
}
//商品名を修正したものを返す
func getProductName(name string) string {
    var formatName string
    names := strings.Split(name, " ")
    for i := range names {
        for j := range ProductNameDeleteKeyList {
            if strings.Index(names[i], ProductNameDeleteKeyList[j]) >= 0 {
                names[i] = ""
            }
        }
    }
    for i := range names {
        if names[i] != "" {
            if i == 0 {
                formatName += names[i]
            } else {
            formatName += " " + names[i]
            }
        }
    }
    return formatName
}
//画素数の取り出し
func getPixelsValues(pixel string) (int, int) {
    total := -1
    valid := -1
    var err error
    pix := strings.Replace(pixel, PixelsKeyList["unit"], " ", -1)
    pix = strings.Replace(pix, PixelsKeyList["total"], PixelsKeyList["total"]+" ", -1)
    pix = strings.Replace(pix, PixelsKeyList["valid"], PixelsKeyList["valid"]+" ", -1)
    pixs := strings.Split(pix, " ")
    for i := range pixs {
        if i+1 < len(pixs) {
            if pixs[i+1] == PixelsKeyList["total"] {
                total, err = strconv.Atoi(pixs[i])
                if err != nil {
                    total = -1
                }
            } else if pixs[i+1] == PixelsKeyList["valid"] {
                valid, err = strconv.Atoi(pixs[i])
                if err != nil {
                    valid = -1
                }
            } 
        }
    }
    return total, valid
}
//イメージセンサーの取得
func getImageSensorType(sensor string) string {
    for i := range ImageSensorTypeList {
        if strings.Index(sensor, ImageSensorTypeList[i]) >= 0 {
            return ImageSensorTypeList[i]
        }
    }
    return ImageSensorTypeList[0]
}