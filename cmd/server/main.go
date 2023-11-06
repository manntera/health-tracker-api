package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"manntera.com/health-tracker-api/pkg/repository/healthRepository"
	"manntera.com/health-tracker-api/pkg/repository/userRepository"
	"manntera.com/health-tracker-api/pkg/usecase/healthUsecase"
	"manntera.com/health-tracker-api/pkg/usecase/userUsecase"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"https://ogt-home-platform-front-ggqlr5vn4a-uc.a.run.app",
			"http://localhost:3000",
		},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

	e.POST("/user/add", addUser)
	e.POST("/user/get", getUser)
	e.POST("/user/delete", deleteUser)
	e.POST("/health/add", addHealth)
	e.POST("/health/get", getHealth)
	e.POST("/health/delete", deleteHealth)

	e.Logger.Fatal(e.Start(":" + port))
}

func addUser(c echo.Context) error {
	userRepo, userRepoErr := userRepository.NewUserRepository(c.Request().Context())

	if userRepoErr != nil {
		return echo.NewHTTPError(500, userRepoErr)
	}

	userUC := userUsecase.NewUserUsecase(userRepo)
	var userData userRepository.User
	if err := c.Bind(&userData); err != nil {
		return echo.NewHTTPError(500, err)
	}

	getErr := userUC.AddUser(c.Request().Context(), &userData)
	if getErr != nil {
		return echo.NewHTTPError(500, getErr)
	}
	log.Print(userData)
	return c.JSON(200, userData)
}

type UserIdRequest struct {
	Id string `json:"id"`
}

func getUser(c echo.Context) error {
	userRepo, userRepoErr := userRepository.NewUserRepository(c.Request().Context())

	if userRepoErr != nil {
		return echo.NewHTTPError(500, userRepoErr)
	}

	userUC := userUsecase.NewUserUsecase(userRepo)

	var getUserRequest UserIdRequest
	if err := c.Bind(&getUserRequest); err != nil {
		return echo.NewHTTPError(500, err)
	}

	result, getErr := userUC.GetUser(c.Request().Context(), getUserRequest.Id)

	if getErr != nil {
		return echo.NewHTTPError(500, getErr)
	}
	log.Print(result)

	return c.JSON(200, result)

}

func deleteUser(c echo.Context) error {
	userRepo, userRepoErr := userRepository.NewUserRepository(c.Request().Context())

	if userRepoErr != nil {
		return echo.NewHTTPError(500, userRepoErr)
	}

	userUC := userUsecase.NewUserUsecase(userRepo)

	var getUserRequest UserIdRequest
	if err := c.Bind(&getUserRequest); err != nil {
		return echo.NewHTTPError(500, err)
	}

	deleteErr := userUC.DeleteUser(c.Request().Context(), getUserRequest.Id)

	if deleteErr != nil {
		return echo.NewHTTPError(500, deleteErr)
	}

	return c.JSON(200, getUserRequest)
}

type healthAddReq struct {
	UserId      string `json:"userId"`
	HealthScore int    `json:"healthScore"`
	Comment     string `json:"comment"`
}

func addHealth(c echo.Context) error {
	userRepo, userRepoErr := userRepository.NewUserRepository(c.Request().Context())

	if userRepoErr != nil {
		return echo.NewHTTPError(500, userRepoErr)
	}

	healthRepo, healthRepoErr := healthRepository.NewHealthRepository(c.Request().Context())

	if healthRepoErr != nil {
		return echo.NewHTTPError(500, healthRepoErr)
	}

	healthUC := healthUsecase.NewHealthUsecase(healthRepo, userRepo)

	var healthData healthAddReq
	if err := c.Bind(&healthData); err != nil {
		return echo.NewHTTPError(500, err)
	}

	result, getErr := healthUC.AddData(c.Request().Context(), healthData.UserId, healthData.HealthScore, healthData.Comment)

	if getErr != nil {
		return echo.NewHTTPError(500, getErr)
	}

	return c.JSON(200, result)
}

type healthGetReq struct {
	UserId    string
	StartTime int64
	EndTime   int64
}

func getHealth(c echo.Context) error {
	userRepo, userRepoErr := userRepository.NewUserRepository(c.Request().Context())

	if userRepoErr != nil {
		return echo.NewHTTPError(500, userRepoErr)
	}

	healthRepo, healthRepoErr := healthRepository.NewHealthRepository(c.Request().Context())

	if healthRepoErr != nil {
		return echo.NewHTTPError(500, healthRepoErr)
	}

	healthUC := healthUsecase.NewHealthUsecase(healthRepo, userRepo)

	var healthData healthGetReq
	if err := c.Bind(&healthData); err != nil {
		return echo.NewHTTPError(500, err)
	}

	result, getErr := healthUC.GetData(c.Request().Context(), healthData.UserId, healthData.StartTime, healthData.EndTime)

	if getErr != nil {
		return echo.NewHTTPError(500, getErr)
	}

	return c.JSON(200, result)
}

type healthDeleteReq struct {
	UserId string
	Uuid   string
}

func deleteHealth(c echo.Context) error {
	userRepo, userRepoErr := userRepository.NewUserRepository(c.Request().Context())

	if userRepoErr != nil {
		return echo.NewHTTPError(500, userRepoErr)
	}

	healthRepo, healthRepoErr := healthRepository.NewHealthRepository(c.Request().Context())

	if healthRepoErr != nil {
		return echo.NewHTTPError(500, healthRepoErr)
	}

	healthUC := healthUsecase.NewHealthUsecase(healthRepo, userRepo)

	var healthData healthDeleteReq
	if err := c.Bind(&healthData); err != nil {
		return echo.NewHTTPError(500, err)
	}

	deletedData, deleteErr := healthUC.DeleteData(c.Request().Context(), healthData.UserId, healthData.Uuid)

	if deleteErr != nil {
		return echo.NewHTTPError(500, deleteErr)
	}
	return c.JSON(200, deletedData)
}
