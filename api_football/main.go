package main

import (
	"fmt"
	"io"
	"net/http"
)

// type standing struct {
// }

func getEplStandings() {

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

	fmt.Println(string(body))

}

func main() {
	getEplStandings()
}
