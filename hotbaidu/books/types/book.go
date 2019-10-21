package types

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Book struct {
	ID            int    `json:"ID" mysql:"book_id" bson:"bookid"`
	Name          string `json:"Name" mysql:"name" bson:"name"`
	URL 		  string `json:"URL" mysql:"url" bson:"URL"`
	Image		  string `json:"Image" mysql:"image" bson:"image"`
	Outline       string `json:"Outline" mysql:"outline" bson:"outline"`
	DiskBaiduURL  string `json:"DiskBaiduURL" mysql:"disk_baidu_url" bson:"diskbaiduurl"`
	DiskBaiduPass string `json:"DiskBaiduPass" mysql:"disk_baidu_pass" bson:"diskbaidupass"`
	CTFileURL     string `json:"CTFileURL" mysql:"ctfileurl" bson:"ctfileurl"`
	Disk360URL    string `json:"Disk360URL" mysql:"disk360url" bson:"disk360url"`
	Disk360Pass	  string `json:"Disk360Pass" mysql:"disk360pass" bson:"disk360pass"`
	DiskOtherURL  string `json:"DiskOtherURL" mysql:"disk_other_url" bson:"diskotherurl"`
	UpdateTime    int64  `json:"UpdateTime" mysql:"update_time" bson:"updatetime"`
	Downloaded    int    `json:"Downloaded" mysql:"downloaded" bson:"downloaded"`
}

func New(id int) *Book {
	return &Book{
		ID: id,
	}
}

func (b *Book) ToJSON() string {
	data, err := json.MarshalIndent(b, "", "\t")
	if err != nil {
		return err.Error()
	}
	return string(data)
}

func (b *Book) ToCSV() ([]string, []string) {
	col := []string{
		"ID", "Name", "URL", "Image", "Outline", "DiskBaiduURL", "DiskBaiduPass", "CTFileURL", "Disk360URL", "Disk360Pass", "DiskOtherURL", "UpdateTime", "Downloaded",
	}
	val := []string{
		fmt.Sprintf("%d", b.ID),
		b.Name,
		b.URL,
		b.Image,
		b.Outline,
		b.DiskBaiduURL,
		b.DiskBaiduPass,
		b.CTFileURL,
		b.Disk360URL,
		b.Disk360Pass,
		b.DiskOtherURL,
		time.Unix(b.UpdateTime, 0).Format("2000-01-02 15:04:05"),
		fmt.Sprintf("%v", b.Downloaded == 1),
	}
	return col, val
}

func (b *Book) GetIDbyURL(url string) error {
	ids := strings.Split(url, "/")
	idStr := ids[len(ids)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		b.ID = 0
		return err
	}
	b.ID = id
	return nil
}
