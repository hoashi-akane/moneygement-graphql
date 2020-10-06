package main

import (
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/hoashi-akane/moneygement-graphql/graph"
	"github.com/hoashi-akane/moneygement-graphql/graph/generated"
)


func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	savingDb, err := gorm.Open("mysql", dataSourceSavings)
	usrDb, err := gorm.Open("mysql", dataSourceUser)

	if err != nil{
		panic(err)
	}
	if savingDb == nil || usrDb == nil{
		panic(err)
	}
	defer func(){
		if savingDb != nil{
			if err := savingDb.Close(); err != nil{
				panic(err)
			}
		}
		if usrDb != nil{
			if err := usrDb.Close(); err != nil{
				panic(err)
			}
		}
	}()
	savingDb.LogMode(true)
	usrDb.LogMode(true)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{SAVDB: savingDb, USRDB: usrDb}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
