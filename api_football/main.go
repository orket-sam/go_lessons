package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	model "github.com/orket-sam/go_lessons/api_football/models"
)

// type standing struct {
// }

func getEplStandings() (data string) {

	url := "https://v3.football.api-sports.io/standings?league=39&season=2023"
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}
	req.Header.Add("x-rapidapi-key", "782dc71c25fc3d403321e94559d232a8")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	data = string(body)

	return data
}

func main() {
	data := getEplStandings()
	var apidata model.Apidata
	err := json.Unmarshal([]byte(data), &apidata)

	if err != nil {
		fmt.Println(err)
		return
	}

	standings := apidata.Response[0].League.Standings[0]
	fmt.Println("Epl 2023 standings ")

	for i, standing := range standings {
		space := len(strconv.Itoa(i + 1))
		space2 := strings.Repeat(" ", 4-space)
		line := strings.Repeat(" ", 25-len(standing.Team.Name))
		if standing.Team.Name == "Manchester United" {
			line = strings.Repeat("*", 25-len(standing.Team.Name))
		}
		fmt.Println(i+1, space2, standing.Team.Name, line, standing.Points)
	}

}
