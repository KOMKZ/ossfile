package osssrv

import (
	"net/http"
	"encoding/json"
	"log"
)

const NO_ERROR = "0"
const ERROR_RES = `{"code":"1","message":"服务器错误","data":null}`



type FileInfo struct {
	File_id string `json:"_id"`
}


type resData struct {
	Code string `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}



type Api struct {
	
}

func (api Api) error(res http.ResponseWriter, message string, code string)  {
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

func (api Api) succ(res http.ResponseWriter, data interface{}){
	d := resData{NO_ERROR, "", data}
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
}

// 上传一个文件
func (api FileApi) UploadFile(res http.ResponseWriter, req *http.Request)  {

}

// 获取文件列表
func (api FileApi) GetFileList(res http.ResponseWriter, req *http.Request)  {

}

// 获取一个文件信息
func (api FileApi) GetFileInfo(res http.ResponseWriter, req *http.Request) {
	f := FileInfo{"oss:hash.jpg"}
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
