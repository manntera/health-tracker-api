package main

import (
    "log"
    "net/http"
    "os"
    "time"

    "manntera.com/health-tracker-api/cmd/database"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("open port: %s", port)

    e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"https://ogt-home-platform-front-ggqlr5vn4a-uc.a.run.app",
			"http://localhost:3000", // ローカル開発用のオリジンも含む
		},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
	}))

    e.POST("/addHealth", addHealthHandler)
    e.POST("/getHealth", getHealthHandler)

    e.Logger.Fatal(e.Start(":" + port))
}

type addHealthParam struct {
	User        string
	HealthScore int
	Comment     string
}

func addHealthHandler(c echo.Context) error {
	timestamp := time.Now().Unix()
	param := new(addHealthParam)
	bindErr := c.Bind(&param)
	if bindErr != nil {
		log.Fatalf("Failed send health data: %v", bindErr)
		return c.JSON(http.StatusInternalServerError, bindErr.Error())
	}

	err := database.AddHealthData(param.User, timestamp, param.HealthScore, param.Comment)
	if err != nil {
		log.Fatalf("Failed send health data: %v", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, param)
}

type getHealthParam struct {
	User      string
	StartTime int64
	EndTime   int64
}

func getHealthHandler(c echo.Context) error {
	param := new(getHealthParam)
	bindErr := c.Bind(&param)
	if bindErr != nil {
		log.Fatalf("Failed get health data: %v", bindErr)
		return c.JSON(http.StatusInternalServerError, bindErr.Error())
	}

	healths, err := database.GetHealthData(param.User, param.StartTime, param.EndTime)
	if err != nil {
		log.Fatalf("Failed get health data: %v", err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, healths)
}
