package ossfile



type File struct {
	ID        uint `json:"-" gorm:"primary_key"`
	File_id string `json:"_id" gorm:"index"`
	File_ext string `json:"file_ext" gorm:"size:15"`
	File_comment string `json:"file_comment"`
	File_trace int `json:"file_trace" gorm:"default:0"`
	Is_tmp int `json:"is_tmp" gorm:"default:0"`
	Save_prefix string `json:"save_prefix" gorm:"size:64"`
	Save_bucket string `json:"save_bucket" gorm:"size:64"`
	Created_at int `json:"created_at"`
	File_size int `json:"file_size"`
	Mine_type string `json:"mine_type"`
	Tps_type string `json:"tps_type" gorm:"size(15)"`
	Is_private int `json:"is_private"`
	Is_unique int `json:"is_unique"`
	File_hash string `json:"file_hash" gorm:"index"`
}