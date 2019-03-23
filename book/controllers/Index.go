package controllers

import "fmt"

type indexController struct {

}

func NewIndex() *indexController {
	return new(indexController)
}

func (i *indexController) Index()  {
	fmt.Println("index controller method:index")
}
