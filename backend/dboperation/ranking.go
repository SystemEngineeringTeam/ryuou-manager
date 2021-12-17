package dboperation

import "github.com/SystemEngineeringTeam/ryuou-manager/model"

func SelectRankingData() ([]model.RankingResponse, error) {
	db := gormConnect()
	defer db.Close()

	var teams []model.Team
	if err := db.Model(&teams).Order("score desc").Scan(&teams).Error; err != nil {
		return nil, err
	}

	var ranking []model.RankingResponse
	prevScore := 0
	rank := 0
	for i, team := range teams {
		if prevScore != team.Score {
			rank = i + 1
		}

		ranking = append(ranking, model.RankingResponse{
			Rank:   rank,
			TeamID: team.ID,
			Name:   team.Name,
			Score:  team.Score,
		})
		prevScore = team.Score
	}

	return ranking, nil
}
