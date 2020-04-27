package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"

	"github.com/gorilla/mux"
)

func Intial() {

	r := mux.NewRouter()

	r.HandleFunc("/download", Download)
	r.HandleFunc("/upload", Upload)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	accessKey := "sdZhovYDxtvtdPmHfFfoP-2cIFIeMdFgVLLeFs_i"
	secretKey := "M2qCMCQsI-iBBsaRPZ0xvLOOUpxPZ8EJhjx38tPI"

	localFile := "/Users/xujinzhao/Desktop/WechatIMG3.jpeg"
	bucket := "microservice-exam-images"
	key := "WechatIMG3.jpeg"

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ret.Key, ret.Hash)
}

func Download(w http.ResponseWriter, r *http.Request) {
	domain := "q6tsx2p5b.bkt.clouddn.com"
	key := "WechatIMG3.jpeg"
	publicAccessURL := storage.MakePublicURL(domain, key)
	fmt.Println(publicAccessURL)
	http.Redirect(w, r, publicAccessURL, http.StatusFound)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/download", Download)
	r.HandleFunc("/upload", Upload)
	http.ListenAndServe(":8000", r)
}
