// lolgo project main.go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

var (
	TEST_SUMMONERS             = false
	TEST_MATCHLIST             = false
	TEST_MATCH                 = false
	TEST_GAMES                 = false
	TEST_STATIC_DATA_CHAMPIONS = false
	TEST_ITEMS                 = false
	TEST_MASTERIES             = false
	TEST_RUNES                 = false
	TEST_SUMMONER_SPELLS       = false
	TEST_CHAMPIONS             = false
	TEST_RANKED_STATS          = true
	TEST_SUMMARY_STATS         = true
)

func main() {
	fileData, err := ioutil.ReadFile("key.txt")
	if err != nil {
		log.Fatal(err)
	}
	apiKey := string(fileData)
	riotClient := RiotClient{apiKey, "na", 10}

	if TEST_SUMMONERS || TEST_MATCHLIST || TEST_MATCH || TEST_GAMES || TEST_RANKED_STATS || TEST_SUMMARY_STATS {
		summoners := riotClient.GetSummoners("drunk7irishman", "rastarockit", "ohsnap62")
		fmt.Printf("Summoners by name: %+v\n", summoners)
		sId := summoners[0].Id

		if TEST_MATCHLIST || TEST_MATCH {
			matchList := riotClient.GetRankedMatchList(sId)
			fmt.Printf("Match list for %d: %+v\n", sId, matchList)

			if TEST_MATCH {
				mId := matchList.Matches[0].MatchId
				match := riotClient.GetMatch(mId)
				fmt.Printf("Match details for match %d: %+v\n", mId, match)
			}
		}

		if TEST_GAMES {
			games := riotClient.GetRecentGames(sId)
			fmt.Printf("Recent games for %d: %+v\n", sId, games)
		}

		if TEST_RANKED_STATS {
			rankedStats := riotClient.GetRecentGames(sId)
			fmt.Printf("Ranked stats for %d: %+v\n", sId, rankedStats)
		}

		if TEST_SUMMARY_STATS {
			summaryStats := riotClient.GetRecentGames(sId)
			fmt.Printf("Summary stats for %d: %+v\n", sId, summaryStats)
		}
	}

	if TEST_STATIC_DATA_CHAMPIONS {
		staticChampions := riotClient.GetAllChampionsStaticData(false)
		fmt.Printf("All staticchampions: %+v\n", staticChampions)

		staticThresh := riotClient.GetChampionStaticData(412, true)
		fmt.Printf("Static Thresh: %+v\n", staticThresh)
	}

	if TEST_ITEMS {
		items := riotClient.GetAllItems(false)
		fmt.Printf("All items: %+v\n", items)

		zekes := riotClient.GetItem(3050, true)
		fmt.Printf("Zeke's: %+v\n", zekes)
	}

	if TEST_MASTERIES {
		masteries := riotClient.GetAllMasteries(false)
		fmt.Printf("All masteries: %+v\n", masteries)

		thunderlords := riotClient.GetMastery(6362, true)
		fmt.Printf("Thunderlord's: %+v\n", thunderlords)
	}

	if TEST_RUNES {
		runes := riotClient.GetAllRunes(false)
		fmt.Printf("All runes: %+v\n", runes)
	}

	if TEST_SUMMONER_SPELLS {
		summonerSpells := riotClient.GetAllSummonerSpells(false)
		fmt.Printf("All summoner spells: %+v\n", summonerSpells)
	}

	if TEST_CHAMPIONS {
		champions := riotClient.GetAllChampions(false)
		fmt.Printf("All champions: %+v\n", champions)

		thresh := riotClient.GetChampion(412)
		fmt.Printf("Thresh: %+v\n", thresh)
	}
}