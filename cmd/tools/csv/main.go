package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"manntera.com/health-tracker-api/pkg/repository/healthRepository"
	"manntera.com/health-tracker-api/pkg/repository/userRepository"
	"manntera.com/health-tracker-api/pkg/usecase/healthUsecase"
)

type Data struct {
	score     int
	comment   string
	timeStamp int64
}

func main() {
	y := 2023
	m := 4
	for ; m <= 12; m++ {
		file, err := os.Open("csv-source/" + strconv.Itoa(y) + "_" + strconv.Itoa(m) + ".csv")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		r := csv.NewReader(file)
		rows, err := r.ReadAll()
		if err != nil {
			panic(err)
		}
		var data []Data

		for i := 2; i < len(rows); i++ {
			for j := 1; j < len(rows[i]); j = j + 2 {
				var t Data
				score, err := strconv.Atoi(rows[i][j])
				if err != nil {
					t.timeStamp = -1
					data = append(data, t)
					continue
				}
				t.score = score
				t.comment = rows[i][j+1]
				day := ((j - 1) / 2) + 1
				hour := i - 2
				t.timeStamp = time.Date(y, time.Month(m), day, hour, 0, 0, 0, time.Local).Unix()
				data = append(data, t)
			}
		}

		fmt.Printf("%+v\n", data)
		for _, d := range data {

			if d.timeStamp == -1 {
				continue
			}
			cxt := context.Background()

			userRepo, userRepoErr := userRepository.NewUserRepository(cxt)
			if userRepoErr != nil {
				panic(userRepoErr)
			}
			healthRepo, healthRepoErr := healthRepository.NewHealthRepository(cxt)
			if healthRepoErr != nil {
				panic(healthRepoErr)
			}
			healthUC := healthUsecase.NewHealthUsecase(healthRepo, userRepo)

			result, getErr := healthUC.AddData(cxt, "rVzkNtsT4Zfk4O1mrhUgkxJgdVi2", d.score, d.comment, d.timeStamp)
			if getErr != nil {
				panic(getErr)
			}
			println(result)
		}
	}
}
