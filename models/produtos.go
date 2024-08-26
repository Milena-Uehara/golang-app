package models

import (
	"github.com/Milena-Uehara/golang-app/db"
	"go.uber.org/zap"
)

type Produto struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

func init() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction()))
}

func SearchProducts() []Produto {
	db := db.ConnectDb()

	selectDeProdutos, err := db.Query("select * from products order by id asc")
	if err != nil {
		zap.L().Error(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeProdutos.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = selectDeProdutos.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			zap.L().Error(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CriarNovoProduto(name, description string, price float64, quantity int) {
	db := db.ConnectDb()

	insertData, err := db.Prepare("insert into products(name, description, price, quantity) values($1, $2, $3, $4)")
	if err != nil {
		zap.L().Error(err.Error())
	}

	insertData.Exec(name, description, price, quantity)
	zap.L().Info("New data inserted into the database.")

	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConnectDb()

	deletarProduto, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		zap.L().Error(err.Error())
	}

	deletarProduto.Exec(id)
	zap.L().Info("Data deleted from the database.")

	defer db.Close()
}

func EditProduto(id string) Produto {
	db := db.ConnectDb()

	produtoDoBanco, err := db.Query("select * from products where id=$1", id)
	if err != nil {
		zap.L().Error(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = produtoDoBanco.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			zap.L().Error(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Name = name
		produtoParaAtualizar.Description = description
		produtoParaAtualizar.Price = price
		produtoParaAtualizar.Quantity = quantity
	}
	defer db.Close()
	return produtoParaAtualizar
}

func AtualizaProduto(id int, name, description string, price float64, quantity int) {
	db := db.ConnectDb()

	AtualizaProduto, err := db.Prepare("update products set name=$1, description=$2, price=$3, quantity=$4 where id=$5")
	if err != nil {
		zap.L().Error(err.Error())
	}
	AtualizaProduto.Exec(name, description, price, quantity, id)
	zap.L().Info("Data updated in the database.")
	defer db.Close()
}
