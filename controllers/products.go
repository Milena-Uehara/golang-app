package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Milena-Uehara/golang-app/models"
	"go.uber.org/zap"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func Index(w http.ResponseWriter, r *http.Request) {
	todosProdutos := models.SearchProducts()
	temp.ExecuteTemplate(w, "Index", todosProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		priceConvertido, err := strconv.ParseFloat(price, 64)
		if err != nil {
			zap.S().Error("Error converting price to float:", err)
		}

		quantityConvertido, err := strconv.Atoi(quantity)
		if err != nil {
			zap.S().Error("Error converting quantity to int:", err)
		}

		models.CriarNovoProduto(name, description, priceConvertido, quantityConvertido)
		zap.L().Info("New product added successfully",
			zap.String("Method", "POST"),
			zap.String("Path", r.URL.Path),
		)
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
	zap.L().Info("Product deleted",
		zap.String("Method", "POST"),
		zap.String("Path", r.URL.Path),
	)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.EditProduto(idProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		quantity := r.FormValue("quantity")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			zap.S().Error("Error converting id to int:", err)
		}

		priceConvertidoParaFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			zap.S().Error("Error converting price to float64:", err)
		}

		quantityConvertidaParaInt, err := strconv.Atoi(quantity)
		if err != nil {
			zap.S().Error("Error converting quantity to int:", err)
		}

		models.AtualizaProduto(idConvertidaParaInt, name, description, priceConvertidoParaFloat, quantityConvertidaParaInt)
		zap.L().Info("Product information updated successfully",
			zap.String("Method", "POST"),
			zap.String("Path", r.URL.Path),
		)
	}
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
