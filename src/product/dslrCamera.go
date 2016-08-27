package product

import (
	_ "fmt"
	"strings"
)

//一眼レフカメラ
type DslrCameraInfo struct {
	Name                 string  `json:"name"`         //製品名
	ProductType          int     `json:"product_type"` //プロダクトタイプ
	Company              string  `json:"company"`      //制作会社
	LensType             string  `json:"lens_type"`
	LensMount            int     `json:"lens_mount"`             //レンズアマウント
	TotalPixels          int     `json:"total_pixels"`           //総画素数
	ValidPixels          int     `json:"valid_pixels"`           //有効画素数
	ImageSensor          string  `json:"image_sensor"`           //撮像素子
	LowPassFilter        bool    `json:"low_pass_filter"`        //ローパスフィルターレス
	IsoMin               int     `json:"iso_min"`                //標準isoの最小
	IsoMax               int     `json:"iso_max"`                //標準isoの最大
	IsoExtentionMin      int     `json:"iso_extention_min"`      //拡張isoの最小
	IsoExtentionMax      int     `json:"iso_extention_max"`      //拡張isoの最大
	MemoryFormat         string  `json:"memory_format"`          //記憶フォーマット
	ContinuousShooting   string  `json:"continuous_shooting"`    //連写撮影
	ShutterSpeedMax      int     `json:"shutter_speed_max"`      //シャッタースピードの最大
	ShutterSpeedMin      int     `json:"shutter_speed_min"`      //シャッタースピードの最小
	MonitorSize          float64 `json:"monitor_size"`           //液晶モニターのサイズ
	MonitorDot           string  `json:"monitor_dot"`            //液晶モニターのドット
	FinderType           string  `json:"finder_type"`            //ファインダー形式
	FinderMagnification  float64 `json:"finder_magnification"`   //ファインダー倍率
	FinderFinderCoverage string  `json:"finder_finder_coverage"` //ファインダー視野率
	BatteryType          string  `json:"battery_type"`           //専用電池タイプ
	BatteryModel         string  `json:"battery_model"`          //専用電池の型番
	NumberOfShotLcd      int     `json:"number_of_shot_lcd"`     //液晶モニタ使用時の撮影枚数
	NumberOfShotFinder   int     `json:"number_of_shot_finder"`  //ファインダー使用時の撮影枚数
	StorageMedia         string  `json:"storage_media"`          //記憶メディア
	ImageStabilization   bool    `json:"image_stabilization"`    //手ぶれ補正
	DustRemoval          bool    `json:"dust_removal"`           //ゴミ取り機構
	Liveview             bool    `json:"liveview"`               //ライブビュー
	MovableMonitorType   string  `json:"movable_monitor_type"`   //可動式モニタ
	Selfie               bool    `json:"selfie"`                 //自撮り機能
	PictBridge           bool    `json:"pictBridge"`             //PictBridge対応
	BuiltInStrobo        bool    `json:"built_in_strobo"`        //内蔵カメラ
	Bulb                 bool    `json:"bulb"`                   //バルブ
	Proof                bool    `json:"proof"`                  //防塵防滴
	RawJpegSaveTogether  bool    `json:"raw_jpeg_save_together"` //RAW+JPEG同時記録
	Raw                  string  `json:"raw"`                    //RAW
	SelfTimer            string  `json:"self_timer"`             //セルフタイマー
	Timelaps             bool    `json:"timelaps"`               //タイムラプス
	UsbCharge            bool    `json:"usb_charge"`             //usb充電
	Interface            string  `json:"interface"`              //インターフェース
	Gps                  bool    `json:"gps"`                    //GPS
	BootTime             float64 `json:"boot_time"`              //起動時間
	Video4k              bool    `json:"video_4k"`               //4K
	VideoPixelNumber     string  `json:"video_pixel_number"`     //動画記録画素数
	VideoFps             float64 `json:"video_fps"`              //fps
	VideoFileType        string  `json:"video_file_type"`        //ビデオファイル形式
	VideoCompressionWay  string  `json:"video_compression_way"`  //映像圧縮方式
	VideoAudioMemoryWay  string  `json:"video_audio_memory_way"` //音声記録方式
	Nfc                  bool    `json:"nfc"`                    //NFC
	Wifi                 bool    `json:"wifi"`                   //WIFI
	WifiDirect           bool    `json:"wifi_direct"`            //wifi_direct
	Width                float64 `json:"width"`                  //幅
	Height               float64 `json:"height"`                 //高さ
	Depth                float64 `json:"depth"`                  //奥行き
	Weight               float64 `json:"weight"`                 //重さ
	Color                string  `json:"color"`                  //色
}

func SetDslrCameraInfo(data ...string) *DslrCameraInfo {
	cameraInfo := new(DslrCameraInfo)
	for i := range data {
		switch i {
		case 0:
			cameraInfo.Company = strings.TrimSpace(data[i])
		case 1:
			cameraInfo.Name = getProductName(strings.TrimSpace(data[i]))
		case 2:
			cameraInfo.ProductType = getProductType(strings.TrimSpace(data[i]))
		case 3:
			cameraInfo.LensMount = getLensMountType(strings.TrimSpace(data[i]))
		case 4:
			total, valid := getPixelsValues(strings.TrimSpace(data[i]))
			cameraInfo.TotalPixels = total
			cameraInfo.ValidPixels = valid
		case 5:
			imageSensor := getImageSensorType(strings.TrimSpace(data[i]))
			cameraInfo.ImageSensor = imageSensor
		case 6:
			cameraInfo.LowPassFilter = boolCheck(strings.TrimSpace(data[i]))
		case 7:
			bMin, bMax, eMin, eMax := getIsoValues(strings.TrimSpace(data[i]))
			cameraInfo.IsoMin = bMin
			cameraInfo.IsoMax = bMax
			cameraInfo.IsoExtentionMin = eMin
			cameraInfo.IsoExtentionMax = eMax
		case 8:
			cameraInfo.MemoryFormat = strings.TrimSpace(data[i])
		case 9:
			cameraInfo.ContinuousShooting = strings.TrimSpace(data[i])
		case 10:
			max, min := getShutterSpeeds(strings.TrimSpace(data[i]))
			cameraInfo.ShutterSpeedMax = max
			cameraInfo.ShutterSpeedMin = min
		case 11:
			//cameraInfo
			inch, dot := getMoniterSizeAndDot(strings.TrimSpace(data[i]))
			cameraInfo.MonitorSize = inch
			cameraInfo.MonitorDot = dot
		case 12:
			cameraInfo.ContinuousShooting = strings.TrimSpace(data[i])
		case 13:
			cameraInfo.FinderMagnification = getFinderMagnification(strings.TrimSpace(data[i]))
		case 14:
			cameraInfo.FinderFinderCoverage = strings.TrimSpace(data[i])
		case 15:
			cameraInfo.BatteryType = strings.TrimSpace(data[i])
		case 16:
			cameraInfo.BatteryModel = strings.TrimSpace(data[i])
		case 17:
			lcd, finder := getNumberOfShots(strings.TrimSpace(data[i]))
			cameraInfo.NumberOfShotLcd = lcd
			cameraInfo.NumberOfShotFinder = finder
		case 18:
			cameraInfo.StorageMedia = strings.TrimSpace(data[i])
		case 19:
		case 20:
			cameraInfo.ImageStabilization = boolCheck(strings.TrimSpace(data[i]))
		case 21:
			cameraInfo.DustRemoval = boolCheck(strings.TrimSpace(data[i]))
		case 22:
			cameraInfo.Liveview = boolCheck(strings.TrimSpace(data[i]))
		case 23:
			cameraInfo.MovableMonitorType = strings.TrimSpace(data[i])
		case 24:
			cameraInfo.Selfie = boolCheck(strings.TrimSpace(data[i]))
		case 25:
			cameraInfo.PictBridge = boolCheck(strings.TrimSpace(data[i]))
		case 26:
			cameraInfo.BuiltInStrobo = boolCheck(strings.TrimSpace(data[i]))
		case 27:
			cameraInfo.Bulb = boolCheck(strings.TrimSpace(data[i]))
		case 28:
			cameraInfo.Proof = boolCheck(strings.TrimSpace(data[i]))
		case 29:
			cameraInfo.RawJpegSaveTogether = boolCheck(strings.TrimSpace(data[i]))
		case 30:
			cameraInfo.Raw = strings.TrimSpace(data[i])
		case 31:
			cameraInfo.SelfTimer = strings.TrimSpace(data[i])
		case 32:
			cameraInfo.Timelaps = boolCheck(strings.TrimSpace(data[i]))
		case 33:
			cameraInfo.UsbCharge = boolCheck(strings.TrimSpace(data[i]))
		case 34:
			cameraInfo.Interface = strings.TrimSpace(data[i])
		case 35:
			cameraInfo.Gps = boolCheck(strings.TrimSpace(data[i]))
		case 36:
			cameraInfo.BootTime = getBootTime(strings.TrimSpace(data[i]))
		case 37:
		case 38:
			cameraInfo.Video4k = boolCheck(strings.TrimSpace(data[i]))
		case 39:
			cameraInfo.VideoPixelNumber = strings.TrimSpace(data[i])
		case 40:
			cameraInfo.VideoFps = getVideoFps(strings.TrimSpace(data[i]))
		case 41:
			cameraInfo.VideoFileType = strings.TrimSpace(data[i])
		case 42:
			cameraInfo.VideoCompressionWay = strings.TrimSpace(data[i])
		case 43:
			cameraInfo.VideoAudioMemoryWay = strings.TrimSpace(data[i])
		case 44:
			cameraInfo.Nfc = boolCheck(strings.TrimSpace(data[i]))
		case 45:
			cameraInfo.Wifi = boolCheck(strings.TrimSpace(data[i]))
		case 46:
			cameraInfo.WifiDirect = boolCheck(strings.TrimSpace(data[i]))
		case 47:
		case 48:
			w, h, d := getSizes(strings.TrimSpace(data[i]))
			cameraInfo.Width = w
			cameraInfo.Height = h
			cameraInfo.Depth = d
		case 49:
			cameraInfo.Weight = getWeight(strings.TrimSpace(data[i]))
		case 50:
		case 51:
		case 52:
			cameraInfo.Color = strings.TrimSpace(data[i])
		case 53:
		default:
		}
	}
	return cameraInfo
}
