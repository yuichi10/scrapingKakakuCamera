package product
//プロダクトのタイプ一覧
var ProductTypeList = map[string]int{
	"other":  0,
	"一眼レフ":   1,
	"ミラーレス":  2,
	"単焦点":    3,
	"ズームレンズ": 4,
}
//レンズアマウントのタイプ一覧
var LensMountTypeList = map[string]int{
	"other": 0,
	"キヤノンEFマウント":     1,
	"ニコンFマウント":       2,
	"ペンタックスKマウント":    3,
	"マイクロフォーサーズマウント": 4,
	"α Aマウント":        5,
	"α Eマウント":        6,
	"シグマSAマウント":      7,
	"ペンタックス645マウント":  8,
	"Qマウント":          9,
	"ニコン1マウント":       10,
	"Xマウント":          11,
	"キヤノンEF-Mマウント":   12,
}
//画素数の検索データキー
var PixelsKeyList = map[string]string{
	"unit": "万画素",
	"total": "(総画素)",
	"valid": "(有効画素)",
}
//イメージセンサーの種類
//var ImageSensorTypeList
//製品名から消すデータの一覧
var ProductNameDeleteKeyList = []string{
	"ボディ",
	"キット",
}





