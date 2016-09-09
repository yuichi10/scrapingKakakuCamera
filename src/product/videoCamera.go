package product

import (
    "strings"
)

type VideoCameraInfo struct {
	Name                  string  `json:"name"`                     //製品名
	ProductType           int     `json:"product_type"`             //プロダクトタイプ
	Company               string  `json:"company"`                  //制作会社
	Type                  string  `json:"type"`                     //タイプ
	ShootStyle            string  `json:"shoot_style"`              //撮影スタイル
	ImageSensor           string  `json:"image_sensor"`             //撮像素子
	TotalPixels           int     `json:"total_pixels"`             //総画素数
	ValidPixels           int     `json:"valid_pixels"`             //有効画素数
	ImageValidPixels      int     `json:"image_valid_pixels"`       //静止画有効画素数
	FullHd                bool    `json:"hd"`                       //ハイビジョン
	OpticalZoom           float64 `json:"optical_zoom"`             //光学ズーム
	DigitalZoom           float64 `json:"digital_zoom"`             //デジタルズーム
	MemoryFormat          string  `json:"memory_format"`            //記憶フォーマット
	MonitorSize           float64 `json:"monitor_size"`             //液晶モニターのサイズ
	ImageStabilizationWay string  `json:"image_stabilization"`      //手ぶれ補正
	FocalLengthMin        float64 `json:"focal_length_max"`         //最小の焦点距離
	FocalLengthMax        float64 `json:"focal_length_max"`         //最大の焦点距離
	FMin                  float64 `json:"f_min"`                    //開放F値最小
	FMax                  float64 `json:"f_max"`                    //開放F値最大
	TouchPanel            bool    `json:"touch_panel"`              //タッチパネル
	Shoot3D               bool    `json:"shoot_3d"`                 //3D
	Finder                bool    `json:"finder"`                   //ファインダー
	NightShoot            bool    `json:"night_shoot"`              //夜間撮影機能
	MinDepthOfField       int     `json:"min_depth_of_field"`       //最低被写体照度
	ImageResolution       string  `json:"image_resolution_type"`    //静止画解像度
	MemoryStaticImageType string  `json:"memory_static_image_type"` //メモリー静止画記録形式
	MemoryVideoResolution string  `json:"memory_video_resolution"`  //メモリー動画解像度
	Video4k               bool    `json:"video_4k"`                 //4K
	Interface             string  `json:"interface"`                //インターフェース
	WaterProof            int     `json:"water_proof"`              //防水性能
	BelongBattery         string  `json:"belong_battery"`           //付属バッテリー
	ShootTime             int     `json:"shoot_time"`               //撮影時間
	Weight                float64 `json:"weight"`                   //重さ
	Width                 float64 `json:"width"`                    //幅
	Height                float64 `json:"height"`                   //高さ
	Depth                 float64 `json:"depth"`                    //奥行き
	Color                 string  `json:"color"`                    //色
}

func SetVideoCameraInfo(data ...string) *VideoCameraInfo {
    videoInfo := new(VideoCameraInfo)
    for i := range data {
        switch i {
        case 0:
            videoInfo.Company = strings.TrimSpace(data[i])
        case 1:
            videoInfo.Name = getProductName(strings.TrimSpace(data[i]))
        case 2:
            videoInfo.ProductType = getProductType(strings.TrimSpace(data[i]))
        case 3:
            videoInfo.ShootStyle = strings.TrimSpace(data[i])
        case 4:
            videoInfo.ImageSensor = strings.TrimSpace(data[i])
        case 5:
            total, _ := getPixelsValues(strings.TrimSpace(data[i]))
            videoInfo.TotalPixels = total
        case 6:
            valid, _ := getPixelsValues(strings.TrimSpace(data[i]))
            videoInfo.ValidPixels = valid
        case 7:
            imageValid, _ := getPixelsValues(strings.TrimSpace(data[i]))
            videoInfo.ImageValidPixels = imageValid
        case 8:
            videoInfo.FullHd = boolCheck(strings.TrimSpace(data[i]))
        case 9:
            videoInfo.OpticalZoom = getMagnification(strings.TrimSpace(data[i]))
        case 10:
            videoInfo.DigitalZoom = getMagnification(strings.TrimSpace(data[i]))
        case 11:
            videoInfo.MemoryFormat = strings.TrimSpace(data[i])
        case 12:
            size, _ := getMoniterSizeAndDot(strings.TrimSpace(data[i]))
            videoInfo.MonitorSize = size
        case 13:
            videoInfo.ImageStabilizationWay = strings.TrimSpace(data[i])
        case 14:
            min, max := getFocalLengthsFloat(strings.TrimSpace(data[i]))
            videoInfo.FocalLengthMin = min
            videoInfo.FocalLengthMax = max
        case 15:
            min, max := getFs(strings.TrimSpace(data[i]))
            videoInfo.FMin = min
            videoInfo.FMax = max
        case 16:
            videoInfo.TouchPanel = boolCheck(strings.TrimSpace(data[i]))
        case 17:
            videoInfo.Shoot3D = boolCheck(strings.TrimSpace(data[i]))
        case 18:
            videoInfo.Finder = boolCheck(strings.TrimSpace(data[i]))
        case 19:
            videoInfo.NightShoot = boolCheck(strings.TrimSpace(data[i]))
        case 20:
            videoInfo.MinDepthOfField = getMinFieldOfDepth(strings.TrimSpace(data[i]))
        case 21:
            videoInfo.ImageResolution = strings.TrimSpace(data[i])
        case 22:
            videoInfo.MemoryStaticImageType = strings.TrimSpace(data[i])
        case 23:
            videoInfo.MemoryVideoResolution = strings.TrimSpace(data[i])
        case 24:
            videoInfo.Video4k = boolCheck(strings.TrimSpace(data[i]))
        case 25:
            videoInfo.Interface = strings.TrimSpace(data[i])
        case 26:
            videoInfo.WaterProof = getWaterProof(strings.TrimSpace(data[i]))
        case 27:
            videoInfo.BelongBattery = strings.TrimSpace(data[i])
        case 28:
            videoInfo.ShootTime = getShootTime(strings.TrimSpace(data[i]))
        case 29:
            w, h, d := getSizes(strings.TrimSpace(data[i]))
            videoInfo.Width = w
            videoInfo.Height = h
            videoInfo.Depth = d
        case 30:
            weight := getWeight(strings.TrimSpace(data[i]))
            videoInfo.Weight = weight
        case 31:
            videoInfo.Color = strings.TrimSpace(data[i])
        case 32:
        case 33:
        case 34:
        case 35:
        }
    }
    return videoInfo
}