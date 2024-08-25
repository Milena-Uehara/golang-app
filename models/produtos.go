package models

import (
	"github.com/Milena-Uehara/db"
	"go.uber.org/zap"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func SearchProducts() []Produto {
	db := db.ConnectDb()

	selectDeProdutos, err := db.Query("select * from produtos order by id asc")
	if err != nil {
		zap.L().Error(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			zap.L().Error(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectDb()

	insertData, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		zap.L().Error(err.Error())
	}

	insertData.Exec(nome, descricao, preco, quantidade)
	zap.L().Info("Dados inseridos com sucesso")

	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConnectDb()

	deletarProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		zap.L().Error(err.Error())
	}

	deletarProduto.Exec(id)
	zap.L().Info("Dados deletados com sucesso")

	defer db.Close()
}

func EditProduto(id string) Produto {
	db := db.ConnectDb()

	produtoDoBanco, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		zap.L().Error(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			zap.L().Error(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db := db.ConnectDb()

	AtualizaProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		zap.L().Error(err.Error())
	}
	AtualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	zap.L().Info("Dados atualizados com sucesso")
	defer db.Close()
}
