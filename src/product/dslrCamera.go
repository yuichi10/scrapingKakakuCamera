package product

import (
	_ "fmt"
	"strings"
)

//一眼レフカメラ
type DslrCameraInfo struct {
	Name                 string  `json:"name"`                   //製品名
	ProductType          int     `json:"product_type"`           //プロダクトタイプ
	Company              string  `json:"company"`                //制作会社
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
	ShutterSpeedMin      int     `json:"shutter_speed_min"`      //シャッタースピードの最小
	ShutterSpeedMax      int     `json:"shutter_speed_max"`      //シャッタースピードの最大
	MonitorSize          float32 `json:"monitor_size"`           //液晶モニターのサイズ
	MonitorDot           string  `json:"monitor_dot"`            //液晶モニターのドット
	FinderType           string  `json:"finder_type"`            //ファインダー形式
	FinderMagnification  int     `json:"finder_magnification"`   //ファインダー倍率
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
	BootTime             float32 `json:"boot_time"`              //起動時間
	Video4k              bool    `json:"video_4k"`               //4K
	VideoPixelNumber     string  `json:"video_pixel_number"`     //動画記録画素数
	VideoFps             int     `json:"video_fps"`              //fps
	VideoFileType        string  `json:"video_file_type"`        //ビデオファイル形式
	VideoCompressionWay  string  `json:"video_compression_way"`  //映像圧縮方式
	VideoAudioMemoryWay  string  `json:"video_audio_memory_way"` //音声記録方式
	Nfc                  bool    `json:"nfc"`                    //NFC
	Wifi                 bool    `json:"wifi"`                   //WIFI
	WifiDirect           bool    `json:"wifi_direct"`            //wifi_direct
	Width                int     `json:"width"`                  //幅
	Height               int     `json:"height"`                 //高さ
	Depth                int     `json:"depth"`                  //奥行き
	Weight               int     `json:"weight"`                 //重さ
	Color                string  `json:"color"`                  //色
}

func SetDslrCameraInfo(data ...string) *DslrCameraInfo {
	cameraInfo := new(DslrCameraInfo)
	for i := range data {
		switch i {
		case 0:
		case 1:
			cameraInfo.Company = strings.TrimSpace(data[i])
		case 2:
			cameraInfo.Name = getProductName(strings.TrimSpace(data[i]))
		case 3:
			cameraInfo.ProductType = getProductType(strings.TrimSpace(data[i]))
		case 4:
			cameraInfo.LensMount = getLensMountType(strings.TrimSpace(data[i]))
		case 5:
			total, valid := getPixelsValues(strings.TrimSpace(data[i]))
			cameraInfo.TotalPixels = total
			cameraInfo.ValidPixels = valid
		case 6:
			imageSensor := getImageSensorType(strings.TrimSpace(data[i]))
			cameraInfo.ImageSensor = imageSensor
		case 7:
			cameraInfo.LowPassFilter = boolCheck(strings.TrimSpace(data[i]))
		case 8:
			bMin, bMax, eMin, eMax := getIsoValues(strings.TrimSpace(data[i]))
			cameraInfo.IsoMin = bMin
			cameraInfo.IsoMax = bMax
			cameraInfo.IsoExtentionMin = eMin
			cameraInfo.IsoExtentionMax = eMax
		case 9:
			cameraInfo.MemoryFormat = strings.TrimSpace(data[i])
		case 10:
			cameraInfo.ContinuousShooting = strings.TrimSpace(data[i])
		case 11:
		case 12:
		case 13:
		case 14:
		case 15:
		case 16:
		case 17:
		case 18:
		case 19:
		case 20:
		case 21:
		case 22:
		case 23:
		case 24:
		case 25:
		case 26:
		case 27:
		case 28:
		case 29:
		case 30:
		case 31:
		case 32:
		case 33:
		case 34:
		case 35:
		case 36:
		case 37:
		case 38:
		case 39:
		case 40:
		case 41:
		case 42:
		case 43:
		case 44:
		case 45:
		case 46:
		case 47:
		case 48:
		case 49:
		case 50:
		case 51:
		case 52:
		case 53:
		case 54:
		default:
		}
	}
	return cameraInfo
}
