package utils

import (
	"strconv"
	"fmt"
	"os"
	"strings"
	"net/http"
	"spider.bing89.com/hotbaidu/books/types"
	"github.com/PuerkitoBio/goquery"
	"regexp"
)

func GetHTML(url string)(string, string, error){
	resp, err := http.Get(url)
	if err != nil {
		return "0", "", err
	}
	html := []byte{}
	defer resp.Body.Close()
	_, err = resp.Body.Read(html)
	if err != nil {
		return "0", "", err
	}
	ids := strings.Split(url, "/")
	idStr := ids[len(ids) -1 ]
	
	return idStr, string(html), nil
}

func SaveHTML(fileName string, htmlStr string)error{
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil{
		return err
	}
	defer file.Close()
	_, err = file.WriteString(htmlStr)
	return err
}

func ParseBook(b *types.Book)error{
	doc, err := goquery.NewDocument(b.URL)
	if err != nil {
		fmt.Println(err)
		return  err
	}
	s := doc.Find(".post-single-content").First()
	b.Image, _ = s.Find("img").First().Attr("src")
	outline := s.Find("p").Text()
	s.Find("a").Each(func(n int, a *goquery.Selection){
		href, _ := a.Attr("href")
		// html, _ := s.Html()
		// b.Outline = strings.ReplaceAll(b.Outline, html, "")
		disk := a.Text()
		if strings.Index(disk, "百度") > -1 {
			b.DiskBaiduURL = href
		} else if strings.Index(disk, "城通") > -1 {
			b.CTFileURL = href
		} else if strings.Index(disk, "360") > -1{
			b.Disk360URL = href
		} else {
			b.DiskOtherURL = href
		}
	})
	outlines := strings.Split(outline, "\n")
	endLine := len(outlines)
	for n, line := range outlines {
		if strings.Index(line, "下载地址") > -1 {
			endLine = n
		}
		if n < endLine{
			b.Outline += line
		}
		if strings.Index(line, "提取密码") > -1 || strings.Index(line, "百度网盘提取码") > -1{
			pass := strings.Split(line, "：")
			if len(pass) > 1 {
				b.DiskBaiduPass = pass[1]
			}
		}
		if strings.Index(line, "提取密码") > -1 || strings.Index(line, "360网盘提取码") > -1{
			pass := strings.Split(line, "：")
			if len(pass) > 1 {
				b.DiskBaiduPass = pass[1]
			}
		}

	}
	html, _ := doc.Html()
	SaveHTML(fmt.Sprintf("/opt/data/spider/hotbaidu.com/books/html/book_%d.html", b.ID), html)
	return  nil
}

func ParsePage(url string)(*types.Page, error){
	p := types.Page{}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return &p, err
	}
	doc.Find("article").Each(func(n int, s *goquery.Selection){
		b := types.Book{}
		a := s.Find("a").First()
		href, _ := a.Attr("href")
		// fmt.Println(a.Text(), "\t", href)
		b.GetIDbyURL(href)
		b.Name = a.Text()
		b.URL = href
		p.Books = append(p.Books, b)
	})
	s := doc.Find(".total").First()
	span := s.Find("span").First()
	total := span.Text()
	reg, err := regexp.Compile("[0-9]+")
	if err != nil {
		p.ID=1
		p.Total = 1
	}
	res := reg.FindAllString(total, -1)
	if len(res) > 0{
		cur, err := strconv.Atoi(res[0])
		if err != nil {
			cur = 1
		}
		p.ID = cur
	}
	if len(res) >1 {
		t , err := strconv.Atoi(res[1])
		if err != nil {
			t = 1
		}
		p.Total = t
	}
	
	
	fmt.Println(p.ID, p.Total)
	html, _ := doc.Html()
	SaveHTML(fmt.Sprintf("/opt/data/spider/hotbaidu.com/books/html/page_%d.html", p.ID), html)
	return &p, nil
}