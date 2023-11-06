package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"

	"manntera.com/health-tracker-api/pkg/repository/healthRepository"
	"manntera.com/health-tracker-api/pkg/repository/userRepository"
)

const baseURL = "https://health-tracker-api-ggqlr5vn4a-uc.a.run.app"

func TestAddUser(t *testing.T) {
	url := baseURL + "/user/add"
	testid := "manntera"
	reqBody, _ := json.Marshal(
		userRepository.User{
			Id:    testid,
			Email: "testEmail",
			Name:  "testName",
		})

	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(reqBody))

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	log.Printf("【StatusCode】: %v", resp.StatusCode)
	log.Printf("【Response】: %v", resp)
	var respBody userRepository.User

	json.NewDecoder(resp.Body).Decode(&respBody)

	if respBody.Id != testid {
		t.Errorf("Expected id to be testId, got %s", respBody.Id)
	}

	if respBody.Email != "testEmail" {
		t.Errorf("Expected email to be testEmail, got %s", respBody.Email)
	}

	if respBody.Name != "testName" {
		t.Errorf("Expected name to be testName, got %s", respBody.Name)
	}
	log.Printf("【ResponseBody】: %v", respBody)
}

func TestGetUser(t *testing.T) {
	url := baseURL + "/user/get"
	testid := "manntera"
	reqBody, _ := json.Marshal(
		struct {
			Id string `json:"id"`
		}{
			Id: testid,
		})

	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(reqBody))

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	log.Printf("【StatusCode】: %v", resp.StatusCode)
	log.Printf("【Response】: %v", resp)
	var respBody userRepository.User

	json.NewDecoder(resp.Body).Decode(&respBody)

	log.Printf("【ResponseBody】: %v", respBody)

	if respBody.Id != testid {
		t.Errorf("Expected id to be testId, got %s", respBody.Id)
	}

	if respBody.Email != "testEmail" {
		t.Errorf("Expected email to be testEmail, got %s", respBody.Email)
	}

	if respBody.Name != "testName" {
		t.Errorf("Expected name to be testName, got %s", respBody.Name)
	}
}

func TestDeleteUser(t *testing.T) {
	url := baseURL + "/user/delete"
	testid := "manntera"
	reqBody, _ := json.Marshal(
		struct {
			Id string `json:"id"`
		}{
			Id: testid,
		})

	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(reqBody))

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
	log.Printf("【Response】: %v", resp)

	var respBody UserIdRequest
	json.NewDecoder(resp.Body).Decode(&respBody)

	log.Printf("【StatusCode】: %v", resp.StatusCode)
	log.Printf("【ResponseBody】: %v", respBody)
}

func TestAddHealth(t *testing.T) {
	url := baseURL + "/health/add"

	userId := "manntera"
	score := 5
	comment := "元気かな？"
	reqBody, _ := json.Marshal(
		struct {
			UserId      string `json:"userId"`
			HealthScore int    `json:"healthScore"`
			Comment     string `json:"comment"`
		}{
			UserId:      userId,
			HealthScore: score,
			Comment:     comment,
		})

	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(reqBody))

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	log.Printf("【StatusCode】: %v", resp.StatusCode)
	log.Printf("【Response】: %v", resp)

	var respBody healthRepository.Health

	json.NewDecoder(resp.Body).Decode(&respBody)
	log.Printf("【ResponseBody】: %v", respBody)

	if respBody.HealthScore != score {
		t.Errorf("Expected healthScore to be 1, got %d", respBody.HealthScore)
	}

	if respBody.Comment != comment {
		t.Errorf("Expected comment to be testComment, got %s", respBody.Comment)
	}
}

func TestGetHealth(t *testing.T) {
	url := baseURL + "/health/get"

	userId := "manntera"
	startTime := int64(0)
	endTime := int64(1800000000)
	reqBody, _ := json.Marshal(
		struct {
			UserId    string
			StartTime int64
			EndTime   int64
		}{
			UserId:    userId,
			StartTime: startTime,
			EndTime:   endTime,
		})

	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(reqBody))

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	log.Printf("【StatusCode】: %v", resp.StatusCode)
	log.Printf("【Response】: %v", resp)

	var respBody []healthRepository.Health

	json.NewDecoder(resp.Body).Decode(&respBody)
	log.Printf("【ResponseBody】: %v", respBody)
}

func TestDeleteHealth(t *testing.T) {
	url := baseURL + "/health/delete"

	userId := "manntera"
	uuid := "V1NhhmlvrsH5N8Zctgr6"
	reqBody, _ := json.Marshal(
		struct {
			UserId string
			Uuid   string
		}{
			UserId: userId,
			Uuid:   uuid,
		})

	resp, _ := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	log.Printf("【Response】: %v", resp)

	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}

	log.Printf("【StatusCode】: %v", resp.StatusCode)

	var respBody healthRepository.Health
	json.NewDecoder(resp.Body).Decode(&respBody)
	log.Printf("【ResponseBody】: %v", respBody)
}
