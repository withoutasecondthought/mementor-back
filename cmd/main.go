package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	mementor_back "mementor-back"
	"mementor-back/config"
	"mementor-back/pkg/handler"
	"mementor-back/pkg/repository"
	"mementor-back/pkg/service"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title       Mementor back
// @version     1.0
// @description Best backend ever.

// @contact.name  @withoutasecondthought
// @contact.email mrmarkeld@gmail.com

// @host     api.ilyaprojects.com/
// @BasePath /mementor

//@securityDefinitions.apikey ApiKeyAUth
//@in header
//@name Authorization

func main() {
	err := config.InitConfig()
	if err != nil {
		logrus.Fatalf("cannot init config %s", err)
	}

	db, err := InitDB()
	if err != nil {
		logrus.Fatalf("Error init database %s", err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(mementor_back.Server)
	logrus.Println("Server started")

	go func() {
		logrus.Fatal(srv.Run(viper.GetString("port"), handlers.InitRoutes()))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Server shutting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}

func InitDB() (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(viper.GetString("db.uri")))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return client.Database(viper.GetString("db.database")), nil
}
