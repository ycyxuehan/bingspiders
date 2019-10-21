package main

import (
	"spider.bing89.com/hotbaidu/books/utils"
	"spider.bing89.com/hotbaidu/books/types"
	// "spider.bing89.com/hotbaidu/books"
	"os"
	"strings"
	"fmt"
)

const BOOK_BASE_URL="https://hotbaidu.com/category/uncategorized"


func main(){
	p, err := utils.ParsePage(BOOK_BASE_URL)
	if err!=nil{
		panic(err)
	}
	bookFile, err := os.OpenFile("/opt/data/spider/hotbaidu.com/books/books.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	b := types.Book{}
	col, _ := b.ToCSV()
	bookFile.WriteString(fmt.Sprintf("%s\n", strings.Join(col, ",")))
	defer bookFile.Close()
	for i:= 3004; i<= p.Total; i++ {
		fmt.Println("Get Page ", i, p.Total, " remaining...")
		URL := BOOK_BASE_URL + fmt.Sprintf("/page/%d", i)
		page, err := utils.ParsePage(URL)
		if err == nil {
			for _, book := range page.Books {
				fmt.Println("Get Book", book.Name, " URL: ", book.URL)
				err = utils.ParseBook(&book)
				if err == nil {
					_, data := book.ToCSV()
					bookFile.WriteString(fmt.Sprintf("%s\n", strings.Join(data, ",")))
				}
			}
		}
	}
	fmt.Println("done")
}
