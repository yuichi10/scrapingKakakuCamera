package product

import (
    _ "strings"
)

type FlashInfo struct {
    Name                  string  `json:"name"`                     //製品名
	ProductType           int     `json:"product_type"`             //プロダクトタイプ
	Company               string  `json:"company"`                  //制作会社
    GuideNumber int
    IlluminatingAngleMin int
    IlluminatingAngleMax int
    DimmerData string
    BatteryType string
    BatteryNumber string
    Manual bool
    AutoZoom bool
    Wireless bool 
    Bounce bool
    ExternalPhotometry bool
    ConsecutiveEmission bool
    Modeling bool
    Multi bool 
    RedEyeRelief bool
    FPEmission bool
    TrailingCurtainSynchronization bool
    ExternalBattery bool
    Width                float64 `json:"width"`                  //幅
	Height               float64 `json:"height"`                 //高さ
	Depth                float64 `json:"depth"`                  //奥行き
	Weight               float64 `json:"weight"`                 //重さ
}

func SetFlashInfo(data ...string) *FlashInfo {
    flashInfo := new(FlashInfo)
    for i := range data {
        switch i {
        case 0:
            //flashInfo.Name = 
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
        }
    }
    return flashInfo
}