package product

type VideoCameraInfo struct {
	Name                  string  `json:"name"`                     //製品名
	ProductType           int     `json:"product_type"`             //プロダクトタイプ
	Company               string  `json:"company"`                  //制作会社
	Type                  string  `json:"type"`                     //タイプ
	ShootStyle            string  `json:"shoot_style"`              //撮影スタイル
	ImageSensor           string  `json:"image_sensor"`             //撮像素子
	TotalPixels           int     `json:"total_pixels"`             //総画素数
	ValidPixels           int     `json:"valid_pixels"`             //有効画素数
	ImageValidPixels      string  `json:"image_valid_pixels"`       //静止画有効画素数
	Hd                    bool    `json:"hd"`                       //ハイビジョン
	OpticalZoom           int     `json:"optical_zoom"`             //光学ズーム
	DigitalZoom           int     `json:"digital_zoom"`             //デジタルズーム
	MemoryFormat          string  `json:"memory_format"`            //記憶フォーマット
	MonitorSize           float64 `json:"monitor_size"`             //液晶モニターのサイズ
	ImageStabilizationWay string  `json:"image_stabilization"`      //手ぶれ補正
	FocalLengthMin        int     `json:"focal_length_max"`         //最小の焦点距離
	FocalLengthMax        int     `json:"focal_length_max"`         //最大の焦点距離
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
        case 1:
        case 2:
        case 3:
        case 4:
        case 5:
        case 6:
        case 7:
        case 8:
        case 9:
        case 10:
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
        }
    }
    return videoInfo
}