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
	for i, team := range teams {
		ranking = append(ranking, model.RankingResponse{
			Rank:   i + 1,
			TeamID: team.ID,
			Name:   team.Name,
			Score:  team.Score,
		})
	}

	return ranking, nil
}
