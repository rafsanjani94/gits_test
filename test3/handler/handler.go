package handler

import (
	"bytes"
	bookRepo "gits/test3/repository/book"
	userRepo "gits/test3/repository/user"
	"gits/test3/request"
	bookService "gits/test3/service/book"
	loginService "gits/test3/service/login"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/gin-gonic/gin"
)

var BookRepo = bookRepo.NewBookRepository()
var UserRepo = userRepo.NewUserRepository()
var BookService = bookService.NewBookService(BookRepo, UserRepo)

func Login(c *gin.Context) {
	var req request.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	login, token, err := loginService.Login(req)
	if err != nil {
		if err.Error() == "401" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  false,
				"message": "invalid credential",
			})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"token":  token,
		"data":   login,
	})
}

func CreateBook(c *gin.Context) {
	var req request.BookReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userInfo, _ := c.Get("userInfo")
	create, err := BookService.CreateBook(req, userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": true,
		"data":   create,
	})
}

func ListBook(c *gin.Context) {
	userInfo, _ := c.Get("userInfo")
	list, err := BookService.ListBook(userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   list,
	})
}

func UpdateBook(c *gin.Context) {
	var req request.BookReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userInfo, _ := c.Get("userInfo")
	create, err := BookService.UpdateBook(c.Param("id"), req, userInfo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   create,
	})
}

func DeleteBook(c *gin.Context) {
	err := BookService.DeleteBook(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  false,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": true,
	})
}

func PdfBook(c *gin.Context) {
	wkhtmltopdf.SetPath("C:/Program Files/wkhtmltopdf/bin/wkhtmltopdf")
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	data := map[string]string{
		"foo": "bar",
	}
	tmpl, _ := template.ParseFiles("views/input.html")
	var body bytes.Buffer
	if err = tmpl.Execute(&body, data); err != nil {
		log.Fatal(err)
	}

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(body.String())))
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Dpi.Set(300)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	// err = pdfg.WriteFile("views/output.pdf")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	log.Println("Done")

	// c.Writer.Header().Set("Content-type", "application/pdf")
	c.Writer.Write(pdfg.Bytes())
}
