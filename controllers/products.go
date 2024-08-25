package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Milena-Uehara/models"
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
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			zap.S().Error("Erro na conversão do preço:", err)
		}

		quantidadeConvertido, err := strconv.Atoi(quantidade)
		if err != nil {
			zap.S().Error("Erro na conversão da quantidade:", err)
		}

		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertido)
		zap.L().Info("Produto cadastrado com sucesso",
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
	zap.L().Info("Produto deletado com sucesso",
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
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			zap.S().Error("Erro na convesão do ID para int:", err)
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			zap.S().Error("Erro na convesão do preço para float64:", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			zap.S().Error("Erro na convesão da quantidade para int:", err)
		}

		models.AtualizaProduto(idConvertidaParaInt, nome, descricao, precoConvertidoParaFloat, quantidadeConvertidaParaInt)
		zap.L().Info("Produto atualizado com sucesso",
			zap.String("Method", "POST"),
			zap.String("Path", r.URL.Path),
		)
	}
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}
