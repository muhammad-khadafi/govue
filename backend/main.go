package main

import (
	"backend/config"
	"backend/handler"
	middlewares "backend/middleware"
	"backend/repository"
	"backend/service"
	"backend/util"
	"flag"
	"fmt"
	"github.com/Graylog2/go-gelf/gelf"
	"github.com/gin-gonic/gin"
	_ "github.com/godror/godror"
	"io"
	"log"
	"os"
)

/*
 * Created by muhammad.khadafi on 27/05/2022
 */

func main() {
	//s := gocron.NewScheduler(time.Local)
	//s.Every(1).Day().At("14:23").Do(func() {
	//	fmt.Println("tes scheduler")
	//})
	//s.Every(1).Day().At("14:23").Do(func() {
	//	fmt.Println("tes scheduler")
	//})
	//s.StartAsync()

	// Graylog
	var graylogAddr string

	flag.StringVar(&graylogAddr, "graylog", "", "graylog server addr")
	flag.Parse()
	log.Println(graylogAddr)
	if graylogAddr != "" {
		fmt.Println("logging graylog")
		fmt.Println(graylogAddr)
		gelfWriter, err := gelf.NewWriter(graylogAddr)
		if err != nil {
			log.Fatalf("gelf.NewWriter: %s", err)
		}
		// log to both stderr and graylog2
		log.SetOutput(io.MultiWriter(os.Stderr, gelfWriter))
		log.Printf("logging to stderr & graylog2@'%s'", graylogAddr)
	} else {
		log.Printf("logging not '%s'", graylogAddr)
	}

	DB, errDB := config.ConnectDatabase()
	if errDB != nil {
		fmt.Println("DB Connection Failed..")
		fmt.Println(errDB)
		return
	}
	fmt.Println("DB Connection Establish..")

	scheduler := util.NewSchedulerRepository(DB)
	scheduler.RunningScheduler()

	userRepository := repository.NewUserRepository(DB)
	anggotaRepository := repository.NewAnggotaRepository(DB)
	institusiRepository := repository.NewInstitusiRepository(DB)

	userService := service.NewUserService(userRepository)
	anggotaService := service.NewAnggotaService(anggotaRepository)
	institusiService := service.NewInstitusiService(institusiRepository)

	userHandler := handler.NewUserHandler(userService)
	loginHandler := handler.NewLoginHandler(userService)
	anggotaHandler := handler.NewAnggotaHandler(anggotaService)
	institusiHandler := handler.NewInstitusiHandler(institusiService)

	router := gin.Default()
	apiPublic := router.Group("/api")
	apiPublic.POST("/login", loginHandler.Login)

	apiProtected := router.Group("/api")
	apiProtected.Use(middlewares.JwtAuthMiddleware(DB))
	routeUser := apiProtected.Group("/user")
	{
		routeUser.GET("/get/:id", userHandler.FindUserByID)
		routeUser.GET("/current", loginHandler.CurrentUser)
		routeUser.POST("/insert", userHandler.InsertUser)
	}

	routeAnggota := apiProtected.Group("/anggota")
	{
		routeAnggota.GET("/get/:id", anggotaHandler.FindAnggotaById)
		routeAnggota.POST("/insert", anggotaHandler.InsertAnggota)
		routeAnggota.POST("/add", anggotaHandler.AddAnggota) //nyoba generic json parser
		routeAnggota.PUT("/update", anggotaHandler.UpdateAnggota)
		routeAnggota.DELETE("/delete/:id", anggotaHandler.DeleteAnggotaById)
		routeAnggota.GET("/list", anggotaHandler.ListAnggota)
	}

	routeInstitusi := apiProtected.Group("/institusi")
	{
		routeInstitusi.POST("/add", institusiHandler.AddInstitusi) //nyoba generic json parse
		routeInstitusi.GET("/list", institusiHandler.ListInstitusi)
	}

	apiProtected.POST("/logout", loginHandler.Logout)

	router.Run("127.0.0.1:8080")

}
