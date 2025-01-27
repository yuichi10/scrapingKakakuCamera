package product

//プロダクトのタイプ一覧
var ProductTypeList = map[string]int{
	"other": 0,
	"一眼レフ":  1,
	"ミラーレス": 2,
	"レンズ":   3,
	"ハンディカメラ": 4,
	"アクションカメラ": 5,
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
	"unit":  "万画素",
	"total": "(総画素)",
	"valid": "(有効画素)",
}

//イメージセンサーの種類
var ImageSensorTypeList = []string{
	"その他",
	"フルサイズ",
	"APS-C",
	"ローパスフィルターレス",
	"フォーサーズ",
}

//製品名から消すデータの一覧
var ProductNameDeleteKeyList = []string{
	"ボディ",
	"キット",
}

//isoを確認するためのkey
var IsoCheckKeyList = map[string]string{
	"basic":  "標準：",
	"extend": "拡張：",
	"iso":    "ISO",
}

//液晶の情報を取得するKey
var MonitorCheckKeyList = map[string]string{
	"inch": "インチ",
}

//ファインダーの情報を取得するKey
var FinderCheckKeyList = map[string]string{
	"magnification": "倍",
}

//撮影枚数を取得するためのkey
var NumberOfShotCheckKeyList = map[string]string{
	"lcd":    "液晶モニタ使用時：",
	"finder": "ファインダー使用時：",
	"unit":   "枚",
}

//起動時間を取得するためのKey
var BootTimeCheckKeyList = map[string]string{
	"unit": "秒",
}

//起動時間を取得するためのKey
var VideoFpsCheckKeyList = map[string]string{
	"unit": "fps",
}

//サイズを取得するためのKey
var SizeCheckKeyList = map[string]string{
	"unit":      "mm",
	"delimiter": "x",
}

//重さを取得するためのKey
var WeightCheckKeyList = map[string]string{
	"unit": "g",
}

//焦点距離を取得するためのkey
var FocalLengthCheckKeyList = map[string]string{
	"unit":      "mm",
	"delimiter": "～",
}

//最大撮影距離を取得するためのkey
var FocusDistanceCheckKeyList = map[string]string{
	"unit": "m",
}

//F値を取得するためのkey
var FCheckKeyList = map[string]string{
	"unit":      "F",
	"delimiter": "-",
}

//倍率を求める
var MagnificationKeyList = map[string]string{
	"magnification": "倍",
} 

//被写体照度の取得
var DeapthOfFielsKeyList = map[string]string{
	"magnification": "ルクス",
}

//防水機能
var WaterProofKeyList = map[string]string{
	"magnification": "m",
}

//撮影時間
var ShootTimeProofKeyList = map[string]string{
	"magnification": "分",
}
