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

func stringToIntPositive(val string) int {
	result := -1
	num, err := strconv.Atoi(strings.TrimSpace(val))
	if err == nil {
		result = num
	}
	return result
}

func stringToFloat64Positive(val string) float64 {
	var result float64 = -1
	num, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
	if err == nil {
		result = num
	}
	return result
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

//液晶のサイズとドットを取得
func getMoniterSizeAndDot(val string) (float64, string) {
	inch := -1.0
	dot := ""
	if strings.Index(val, MonitorCheckKeyList["inch"]) >= 0 {
		vals := strings.Split(val, MonitorCheckKeyList["inch"])
		inchSize, err := strconv.ParseFloat(vals[0], 64)
		if err == nil {
			inch = inchSize
		}
		if len(vals) > 1 {
			dot = vals[len(vals)-1]
		}
	}
	return inch, dot
}

//ファインダー倍率を追加
func getFinderMagnification(val string) float64 {
	magni := -1.0
	if strings.Index(val, FinderCheckKeyList["magnification"]) >= 0 {
		vals := strings.Split(val, FinderCheckKeyList["magnification"])
		magnitu, err := strconv.ParseFloat(strings.TrimSpace(vals[0]), 64)
		if err == nil {
			magni = magnitu
		}
	}
	return magni
}

//撮影枚数を取得
func getNumberOfShots(val string) (int, int) {
	lcd := -1
	finder := -1
	val = strings.Replace(val, NumberOfShotCheckKeyList["unit"], NumberOfShotCheckKeyList["unit"]+" ", -1)
	vals := strings.Split(val, " ")
	for i := range vals {
		if strings.Index(vals[i], NumberOfShotCheckKeyList["lcd"]) >= 0 {
			//液晶の枚数
			lcdStr := strings.Replace(vals[i], NumberOfShotCheckKeyList["lcd"], "", -1)
			lcdStr = strings.Replace(lcdStr, NumberOfShotCheckKeyList["unit"], "", -1)
			lcdNum, err := strconv.Atoi(strings.TrimSpace(lcdStr))
			if err == nil {
				lcd = lcdNum
			}
		} else if strings.Index(vals[i], NumberOfShotCheckKeyList["finder"]) >= 0 {
			//finderの枚数
			finderStr := strings.Replace(vals[i], NumberOfShotCheckKeyList["finder"], "", -1)
			finderStr = strings.Replace(finderStr, NumberOfShotCheckKeyList["unit"], "", -1)
			finderNum, err := strconv.Atoi(strings.TrimSpace(finderStr))

			if err == nil {
				finder = finderNum
			}
		}
	}
	return lcd, finder
}

//起動時間の取得
func getBootTime(val string) float64 {
	time := -1.0
	if strings.Index(val, BootTimeCheckKeyList["unit"]) >= 0 {
		val = strings.Replace(val, BootTimeCheckKeyList["unit"], "", -1)
		booltime, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
		if err == nil {
			time = booltime
		}
	}
	return time
}

//video fpsの取得
func getVideoFps(val string) float64 {
	fps := -1.0
	if strings.Index(val, VideoFpsCheckKeyList["unit"]) >= 0 {
		val = strings.Replace(val, VideoFpsCheckKeyList["unit"], "", -1)
		fpsNum, err := strconv.ParseFloat(strings.TrimSpace(val), 64)
		if err == nil {
			fps = fpsNum
		}
	}
	return fps
}

//大きさの取得
func getSizes(val string) (float64, float64, float64) {
	var w, h, d float64 = -1, -1, -1
	val = strings.Replace(val, SizeCheckKeyList["unit"], "", -1)
	vals := strings.Split(val, SizeCheckKeyList["delimiter"])
	for i := range vals {
		switch i {
		case 0:
			w = stringToFloat64Positive(vals[i])
		case 1:
			h = stringToFloat64Positive(vals[i])
		case 2:
			d = stringToFloat64Positive(vals[i])
		}
	}
	return w, h, d
}

//重さを取得
func getWeight(val string) float64 {
	val = strings.Replace(val, WeightCheckKeyList["unit"], "", -1)
	weight := stringToFloat64Positive(val)
	return weight
}

//焦点距離を渡す(一つしか値がなかった時は両方に同じ値を)
func getFocalLengths(val string) (int, int) {
	var nums [] int
	val = strings.TrimSpace(strings.Replace(val, FocalLengthCheckKeyList["unit"], "", -1))
	vals := strings.Split(val, FocalLengthCheckKeyList["delimiter"])
	for i := range vals {
		if num := stringToIntPositive(vals[i]); num != -1 {
			nums = append(nums, num)
		}
	}
	if len(nums) == 0 {
		return -1, -1
	} else if len(nums) == 1 {
		return nums[0], nums[0]
	} else if len(nums) == 2 {
		num := nums[0]
		min := -1
		max := -1
		if num > nums[1] {
			min = nums[1]
			max = num
		} else {
			min = num
			max = nums[1]
		}
		return min, max
	}
	return -1, -1
}

//焦点距離を取得
func getFocusDistance(val string) int {
	return stringToIntPositive(strings.Replace(val, FocusDistanceCheckKeyList["unit"], "", -1))
}

//f値の取得
func getFs(val string) (float64, float64) {
	var nums []float64
	var min, max float64
	vals := strings.Split(val, FCheckKeyList["delimiter"])
	for i := range vals {
		if num := stringToFloat64Positive(strings.Replace(vals[i], FCheckKeyList["unit"], "", -1)); num != -1 {
			nums = append(nums, num)
		}
	}
	if len(nums) == 0 {
		return -1, -1
	} else if len(nums) == 1 {
		return nums[0], nums[0]
	} else if len(nums) == 2 {
		num := nums[0]
		if num > nums[1] {
			min = nums[1]
			max = num
		} else {
			min = num
			max = nums[1]
		}
		return min, max
	}
	return -1, -1
}