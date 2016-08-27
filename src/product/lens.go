package product

import (
	_ "fmt"
	"strings"
)

//一眼レフカメラ
type LensInfo struct {
	Name                     string  `json:"name"`                        //製品名
	ProductType              int     `json:"product_type"`                //プロダクトタイプ
	Company                  string  `json:"company"`                     //制作会社
	LensMount                int     `json:"lens_mount"`                  //レンズアマウント
	Focus                    string  `json:"focus"`                       //focus
	LensType                 string  `json:"lens_type"`                   //レンズタイプ
	SpecificLensType         string  `json:"specific_lens_type"`          //詳細レンズタイプ
	FullSize                 bool    `json:"full_size"`                   //フルサイズ
	LensConstitution         string  `json:"lens_constitution"`           //レンズ構成
	ApertureBladesNumber     string  `json:"aperture_blades_number"`      //絞り羽根枚数
	FocalLengthMin           int     `json:"focal_length_max"`            //最小の焦点距離
	FocalLengthMax           int     `json:"focal_length_max"`            //最大の焦点距離
	FocusDistance            int     `json:"focus_distance"`              //最短撮影距離
	MagnificationRatio       string  `json:"magnification_ratio"`         //最大撮影倍率
	FMin                     float64 `json:"f_min"`                       //開放F値最小
	FMax                     float64 `json:"f_max"`                       //開放F値最大
	AngleOfView              string  `json:"angle_of_view"`               //画角
	ImageStabilization       bool    `json:"image_stabilization"`         //手振れ防止
	Shoot3D                  bool    `json:"shoot_3d"`                    //3D
	FishEye                  bool    `json:"fish_eye"`                    //フィッシュアイ
	Macro                    bool    `json:"macro"`                       //マクロ
	SwingAndTilt             bool    `json:"swing_and_tilt"`              //アオリ撮影
	Telephoto                bool    `json:"telephoto"`                   //望遠
	HighMagnification        bool    `json:"high_magnification"`          //高倍率
	WideAngle                bool    `json:"wide_angle"`                  //広角
	Mirror                   bool    `json:"mirror"`                      //ミラー
	Pancake                  bool    `json:"pancake"`                     //パンケーキ
	LargeAperture            bool    `json:"large_aperture"`              //大口径
	MaximumDiameterAndLength string  `json:"maximum_diameter_and_length"` //最大径x長さ
	Filter                   string  `json:"filter"`                      //フィルター
	Weight                   int     `json:"weight"`                      //重さ
}

func SetLensInfo(data ...string) *LensInfo {
	lensInfo := new(LensInfo)
	for i := range data {
        switch i {
        case 0:
            lensInfo.Company = strings.TrimSpace(data[i])
        case 1:
            lensInfo.Name = getProductName(strings.TrimSpace(data[i]))
        case 2:
            lensInfo.LensMount = getLensMountType(strings.TrimSpace(data[i]))
        case 3:
            lensInfo.Focus = strings.TrimSpace(data[i])
        case 4:
            lensInfo.ProductType = ProductTypeList["レンズ"]
            lensInfo.LensType = strings.TrimSpace(data[i])
        case 5:
            lensInfo.SpecificLensType = strings.TrimSpace(data[i])
        case 6:
            lensInfo.FullSize = boolCheck(strings.TrimSpace(data[i]))
        case 7:
            lensInfo.LensConstitution = strings.TrimSpace(data[i])
        case 8:
            lensInfo.ApertureBladesNumber = strings.TrimSpace(data[i])
        case 9:
            min, max := getFocalLengths(strings.TrimSpace(data[i]))
            lensInfo.FocalLengthMin = min
            lensInfo.FocalLengthMax = max
        case 10:
            lensInfo.FocusDistance = getFocusDistance(strings.TrimSpace(data[i]))
        case 11:
            lensInfo.MagnificationRatio = strings.TrimSpace(data[i])
        case 12:
            min, max := getFs(strings.TrimSpace(data[i]))
            lensInfo.FMin = min
            lensInfo.FMax = max
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
        }
    }
    return lensInfo
}