package main

import (
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
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
	baseDb, err := gorm.Open("mysql", dataSourceLedger)

	if err != nil{
		panic(err)
	}
	if savingDb == nil || usrDb == nil || baseDb == nil{
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
		if baseDb != nil{
			if err := baseDb.Close(); err != nil{
				panic(err)
			}
		}
	}()
	savingDb.LogMode(true)
	usrDb.LogMode(true)
	baseDb.LogMode(true)

	//opt := option.WithCredentialsFile("./moneygement-b0bff-firebase-adminsdk-zddus-3b4e07930b.json")
	//app, err := firebase.NewApp(context.Background(), nil, opt)
	//if err != nil {
	//	fmt.Errorf("error initializing app: %v", err)
	//}
	//
	////ctx := context.Background()
	//client, err := app.Messaging(ctx)
	//registrationToken := []string{
	//
	//}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{SAVDB: savingDb, USRDB: usrDb, BASEDB: baseDb}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
