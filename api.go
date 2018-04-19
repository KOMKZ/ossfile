package ossfile

import (
	"net/http"
	"encoding/json"
	"log"
	"strings"
	"fmt"
	"crypto/md5"
	"io"
	"path/filepath"
)


const (
	NO_ERROR = "0"
	DEFAULT_ERROR = "1"
	ERROR_RES = `{"code":"1","message":"服务器错误","data":null}`
)






type resData struct {
	Code string `json:"code"`
	Message error `json:"message"`
	Data interface{} `json:"data"`
}



type Api struct {
	
}


func (api Api) error(res http.ResponseWriter, message error, code string)  {
	d := resData{code, message, nil}
	r, err := json.Marshal(d)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r = []byte(ERROR_RES)
		log.Println(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(r)
	res.Write([]byte{'\n'})
}

func (api Api) innerError(res http.ResponseWriter){
	api.error(res, fmt.Errorf("server innernal error"), DEFAULT_ERROR)
}

func (api Api) succ(res http.ResponseWriter, data interface{}){
	d := resData{NO_ERROR, nil, data}
	r, err := json.Marshal(d)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		r = []byte(ERROR_RES)
		log.Println(err)
	}
	res.Header().Set("Content-Type", "application/json")
	res.Write(r)
	res.Write([]byte{'\n'})
}

type FileApi struct {
	Api
	srv *Server
}

type SiteApi struct {
	Api
	srv *Server
}

func (api SiteApi) Index(res http.ResponseWriter, req *http.Request)  {
	api.succ(res, "Welcome To OssFile")
}

// 上传一个文件
func (api FileApi) UploadFile(res http.ResponseWriter, req *http.Request)  {
	ct, ok := req.Header["Content-Type"]
	if !ok {
		api.error(res, fmt.Errorf("can't not detect content type of request header"), DEFAULT_ERROR)
		return
	}
	conType := ""
	for _, ctv := range ct {
		if strings.ContainsAny(ctv, "multipart/form-data"){
			conType = "multipart/form-data"
			break
		}
	}
	if conType == "" {
		api.error(res, fmt.Errorf("can't not detect content type of request header"), DEFAULT_ERROR)
		return
	}
	req.ParseForm()
	if conType == "multipart/form-data" {
		req.ParseMultipartForm(32 << 30)
		tmp, fh, err := req.FormFile("file")
		if err != nil{
			log.Println(err)
			api.error(res, err, DEFAULT_ERROR)
			return
		}
		defer tmp.Close()

		h := md5.New()
		if _, err = io.Copy(h, tmp); err != nil {
			log.Println(err)
			api.error(res, err, DEFAULT_ERROR)
			return
		}
		file := File{
			File_hash: fmt.Sprintf("%x", h.Sum(nil)),
			File_ext: filepath.Ext(fh.Filename),
			File_trace: 0,
		}
		api.succ(res, file)
		return
		/*
		测试保存
		tmpPath := filepath.Join(api.srv.RuntimeDir, strhelper.RandStringRunes(20) + filepath.Ext(fh.Filename))
		tmpFile, err := os.OpenFile(tmpPath, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			api.error(res, err, DEFAULT_ERROR)
			return
		}
		defer tmpFile.Close()
		io.Copy(tmpFile, file)
		log.Printf("uploaded file save in: %s\n", tmpPath)
		*/
	}
}

// 获取文件列表
func (api FileApi) GetFileList(res http.ResponseWriter, req *http.Request)  {

}

// 获取一个文件信息
func (api FileApi) GetFileInfo(res http.ResponseWriter, req *http.Request) {
	f := File{
			File_id:"oss:hash.jpg",
		}
	api.succ(res, f)
}

// 获取一个文件流
func (api FileApi) AccessFile(res http.ResponseWriter, req *http.Request) {

}

// 删除一个文件
func (api FileApi) DeleteFile(res http.ResponseWriter, req *http.Request) {

}

// 更新文件信息
func (api FileApi) UpdateFileInfo(res http.ResponseWriter, req *http.Request)  {

}
