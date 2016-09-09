package main

import (
	_ "database/sql"
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	iconv "github.com/djimenez/iconv-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"os"
	"product"
	"strings"
)

var db *gorm.DB

const (
	CameraUnknown = -1
	CameraBody    = 1
	CameraLens    = 2
	VideoCamera   = 3
	Flash = 4
)
const (
	KakakuURL       = "http://kakaku.com"
	KakakuURLSpec   = "spec"
	KakakuURLCamera = "camera"
	KakakuURLBody   = "digital-slr-camera"
	KakakuURLLens   = "camera-lens"
	KakakuURLVideo  = "video-camera"
	KakakuURLFlash = "camera-flash"
)

var ProductType = map[string]int{
	"Unknown":      CameraUnknown,
	KakakuURLBody:  CameraBody,
	KakakuURLLens:  CameraLens,
	KakakuURLVideo: VideoCamera,
	KakakuURLFlash: Flash,
}

//どのタイプをするかを保存
var itemType = CameraUnknown

//次のページが有るかどうかの検索
func isNextPage(url string) (string, bool) {
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", false
	}
	if _, is := doc.Find(".pagenation .pageNextOn").Attr("src"); is {
		urls := make([]string, 0)
		doc.Find(".pagenation .pageicon a").Each(func(_ int, s *goquery.Selection) {
			url, _ := s.Attr("href")
			urls = append(urls, url)
		})
		return KakakuURL + urls[len(urls)-1], true
	}
	return "", false
}

//商品詳細のリンクを取得
func getLinks(url string) []string {
	urls := make([]string, 0)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Printf("getLinks Error: %v \n", err)
		return nil
	}
	doc.Find(".ckitemLink .ckitanker").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("href")
		urls = append(urls, string(url+KakakuURLSpec))
	})
	return urls
}

//詳細ページからデータを取得
func GetPage(url string) {
	pInfo := make([]string, 0)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		fmt.Printf("url error : %d\n", err)
		return
	}
	if maker := doc.Find(".makerLabel .cateBoxPare").Text(); maker != "" {
		pInfo = append(pInfo, getUtf8String(doc.Find(".makerLabel .cateBoxPare").Text()))
	} else {
		pInfo = append(pInfo, getUtf8String(doc.Find(".makerLabel").Text()))
	}
	
	
	//製品名の取得
	pInfo = append(pInfo, getUtf8String(doc.Find("h2").Text()))
	//詳細取得
	doc.Find(".tblBorderGray td").Each(func(_ int, s *goquery.Selection) {
		//fmt.Printf("%v\n", getUtf8String(s.Text()))
		pInfo = append(pInfo, getUtf8String(s.Text()))
	})
	for i := range pInfo {
		fmt.Printf("%v: %v\n", i, pInfo[i])
	}
	saveProductInfo(pInfo...)
}

//プロダクト情報を保存する
func saveProductInfo(info ...string) {
	switch itemType {
	case CameraBody:
		camera := product.SetDslrCameraInfo(info...)
		db.Create(&camera)
	case CameraLens:
		lens := product.SetLensInfo(info...)
		db.Create(&lens)
	case VideoCamera:
		video := product.SetVideoCameraInfo(info...)
		db.Create(&video)
	}
}

func getUtf8String(str string) string {
	s, err := sjisToUtf8(str)
	if err != nil {
		output, err := iconv.ConvertString(s, "shift-jis", "utf-8")
		if err != nil {
			return str
		}
		return output
	} else {
		return s
	}
}

func sjisToUtf8(str string) (string, error) {
	str = strings.TrimSpace(str)
	//fmt.Printf("%v\n", str)
	ret, err := ioutil.ReadAll(transform.NewReader(strings.NewReader(str), japanese.ShiftJIS.NewDecoder()))
	if err != nil {
		return str, err
	}
	return string(ret), err
}

func getProductInfos(url string) {
	urls := getLinks(url)
	for i := range urls {
		GetPage(urls[i])
	}
	if next, is := isNextPage(url); is {
		getProductInfos(next)
	}
}

func Env_load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func openDB() {
	var err error
	sql := fmt.Sprintf("%v@%v/%v?charset=utf8&parseTime=True", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err = gorm.Open(os.Getenv("DB"), sql)
	if err != nil {
		fmt.Printf("エラー%v\n", err)
		return
	}
}

func checkUrl(url string) bool {
	if strings.Index(url, KakakuURL+"/"+KakakuURLCamera) == 0 {
		return true
	}
	return false
}

func whichItem(url string) int {
	urls := strings.Split(url, "/")
	for i := range urls {
		switch urls[i] {
		case KakakuURLLens:
			return ProductType[KakakuURLLens]
		case KakakuURLBody:
			return ProductType[KakakuURLBody]
		case KakakuURLVideo:
			return ProductType[KakakuURLVideo]
		case KakakuURLFlash:
			return ProductType[KakakuURLFlash]
		}
	}
	return CameraUnknown
}

func main() {
	Env_load()
	flag.Parse()
	url := flag.Arg(0)
	openDB()
	defer db.Close()
	db.DB()
	db.AutoMigrate(&product.DslrCameraInfo{})
	db.AutoMigrate(&product.LensInfo{})
	db.AutoMigrate(&product.VideoCameraInfo{})
	itemType = Flash
	GetPage("http://kakaku.com/item/K0000792966/spec/")
	//GetPage("http://kakaku.com/item/K0000891806/spec/")
	return
	if !checkUrl(url) {
		fmt.Println("url が不正です")
		return
	}
	itemType = whichItem(url)
	getProductInfos(url)
}
