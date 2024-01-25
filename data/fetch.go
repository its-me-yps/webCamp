package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	model "wingiesOrNot/models"
)

func Fetch(loginKey string, Batch string) model.Students {

	// Access Token
	accessToken := GetAccessToken(loginKey)

	reqBody, _ := json.Marshal(map[string]interface{}{
		"dataSource": "Cluster0",
		"database":   "student_data",
		"collection": "student_data",
		// filter for Batch
		"filter": map[string]interface{}{
			"i": map[string]interface{}{
				"$regex": "^" + Batch + "[0-9]{4}$",
			},
		},
		"limit": 1500,
	})

	req, err := http.NewRequest("POST", "https://ap-south-1.aws.data.mongodb-api.com/app/data-yubip/endpoint/data/v1/action/find",
		bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println("Error:", err)
		return nil
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println("Error:", err)
	}
	defer res.Body.Close()

	// Reading Data
	resBody, _ := io.ReadAll(res.Body)

	var resJson struct {
		Data model.Students `json:"documents"`
	}
	json.Unmarshal(resBody, &resJson)

	return resJson.Data
}

func GetAccessToken(loginKey string) string {
	reqBody, _ := json.Marshal(map[string]string{
		"key": loginKey,
	})

	res, err := http.Post("https://ap-south-1.aws.realm.mongodb.com/api/client/v2.0/app/data-yubip/auth/providers/api-key/login",
		"application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Println("Error:", err)
		return ""
	}
	defer res.Body.Close()

	// Reading response Body
	resBody, _ := io.ReadAll(res.Body)

	// There are other keys in response JSON but no need to parse them as they are not needed.
	var resJson struct {
		AccessToken string `json:"access_token"`
	}
	json.Unmarshal(resBody, &resJson)

	return resJson.AccessToken
}

func PrintData(data model.Students) {
	for i, student := range data {
		fmt.Println(i + 1)
		fmt.Println("-------------------------------")
		fmt.Printf("ID: %s\n", student.Id)
		fmt.Printf("Name: %s\n", student.Name)
		fmt.Printf("Gender: %s\n", student.Gender)
		fmt.Printf("Email: %s\n", student.Email)
		fmt.Printf("Hall: %s\n", student.Hall)
		fmt.Printf("Room: %s\n", student.Room)
		fmt.Println("-------------------------------")
	}
}
