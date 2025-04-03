package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
)

type Game struct {
	AppId           int    `json:"appid"`
	PlaytimeForever int    `json:"playtime_forever"`
	Name            string `json:"name"`
	ImgIconUrl      string `json:"img_icon_url"`
}

type GameResponse struct {
	Response struct {
		GameCount int    `json:"game_count"`
		Games     []Game `json:"games"`
	} `json:"response"`
}

func GetGamesHandler(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Access-Control-Allow-Origin", "*")

	steamApiKey := os.Getenv("STEAM_API_KEY")
	steamProfileId := os.Getenv("STEAM_PROFILE_ID")

	client := http.Client{}

	ownedGamesRes := GameResponse{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%v&steamid=%v&format=json&include_appinfo=true", steamApiKey, steamProfileId), nil)
	if err != nil {
		log.Fatal(err.Error())
	}

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}

	decoder := json.NewDecoder(res.Body)
	defer res.Body.Close()

	err = decoder.Decode(&ownedGamesRes)
	if err != nil {
		log.Fatal(err.Error())
	}

	sort.Slice(ownedGamesRes.Response.Games, func(i, j int) bool {
		return ownedGamesRes.Response.Games[i].PlaytimeForever > ownedGamesRes.Response.Games[j].PlaytimeForever
	})

	userGamesRes, err := json.Marshal(ownedGamesRes.Response)
	if err != nil {
		log.Fatal(err.Error())
	}

	rw.Write(userGamesRes)
}
