package product

import (
	"fmt"
	"strconv"
	"strings"
)

func TestVlid() {
	a, b, c, d := getIsoValues("標準：ISO100～25600拡張：ISO50～102400")
	fmt.Printf("%v : %v : %v : %v", a, b, c, d)
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
	return strings.TrimSpace(formatName)
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

//ブールのチェック
func boolCheck(val string) bool {
	if val == "○" {
		return true
	}
	return false
}

//isoチェック
func getIsoValues(val string) (int, int, int, int) {
	minBasicIso := -1
	maxBasicIso := -1
	minExtendIso := -1
	maxExtenIso := -1
	//標準と拡張を分ける
	val = strings.Replace(val, IsoCheckKeyList["basic"], " "+IsoCheckKeyList["basic"], -1)
	val = strings.TrimSpace(strings.Replace(val, IsoCheckKeyList["extend"], " "+IsoCheckKeyList["extend"], -1))
	vals := strings.Split(val, " ")
	for i := range vals {
		if strings.Index(strings.TrimSpace(vals[i]), IsoCheckKeyList["basic"]) >= 0 {
			minBasicIso, maxBasicIso = findIsoValue(strings.TrimSpace(vals[i]))
		}
		if strings.Index(strings.TrimSpace(vals[i]), IsoCheckKeyList["extend"]) >= 0 {
			minExtendIso, maxExtenIso = findIsoValue(strings.TrimSpace(vals[i]))
		}
	}
	return minBasicIso, maxBasicIso, minExtendIso, maxExtenIso
}

//標準、拡張を指定してその最小　最大を返す
func findIsoValue(val string) (int, int) {
	//標準があった時の処理
	var min = -1
	var max = -1
	if strings.Index(val, IsoCheckKeyList["iso"]) >= 0 {
		val = strings.Replace(val, IsoCheckKeyList["iso"], IsoCheckKeyList["iso"]+" ", -1)
		vals := strings.Split(val, " ")
		min, max = splitIso(vals[len(vals)-1])
	}
	return min, max
}

//isoを複数パターンで分割
func splitIso(val string) (int, int) {
	min, max, isSuccess := splitIsoPatternA(val)
	if !isSuccess {
		min, max, isSuccess = splitIsoPatternB(val)
		if !isSuccess {
			min, max = -1, -1
		}
	}
	return min, max
}

//isoを〜で区切るパターン
func splitIsoPatternA(val string) (int, int, bool) {
	min := -1
	max := -1
	isSuccess := true
	vals := strings.Split(val, "～")
	iso, err := strconv.Atoi(vals[0])
	if err == nil {
		min = iso
	} else {
		isSuccess = false
	}
	iso, err = strconv.Atoi(vals[len(vals)-1])
	if err == nil {
		max = iso
	} else {
		isSuccess = false
	}
	return min, max, isSuccess
}

//isoを、で区切るパターン
func splitIsoPatternB(val string) (int, int, bool) {
	min := -1
	max := -1
	isSuccess := true
	vals := strings.Split(val, "、")
	iso, err := strconv.Atoi(vals[0])
	if err == nil {
		min = iso
	} else {
		isSuccess = false
	}
	iso, err = strconv.Atoi(vals[len(vals)-1])
	if err == nil {
		max = iso
	} else {
		isSuccess = false
	}
	return min, max, isSuccess
}

//シャッタースピードの取得
func getShutterSpeeds(val string) (int, int) {
    min := -1
    max := -1
    vals := strings.Split(val, "～")
    //速い方
    valMax := strings.TrimSpace(vals[0])
    valMaxs := strings.Split(valMax, "/")
    sMax := valMaxs[len(valMaxs)-1]
    if time, err := strconv.Atoi(sMax); err == nil {
        max = time
    }
    //遅い方
    valMin := strings.TrimSpace(vals[len(vals)-1])
    valMins := strings.Split(valMin, " ")
    sMin := valMins[0]
    if time, err := strconv.Atoi(sMin); err == nil {
        min = time
    }
    return max, min
}
