package mysql

import (
	"bulebell/models"
	"database/sql"
	"go.uber.org/zap"
)

func GetCommunityList() (communitylist []*models.Community, err error) {
	sqlStr := "select community_id,community_name from community"
	if err := db.Select(&communitylist, sqlStr); err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
		}
	}
	return
}

//查询分类详情社区
func GetCommunityByID(id int64) (community *models.CommunityDetail, err error) {
	community = new(models.CommunityDetail)
	sqlStr := `select 
       community_id,community_name,introduction,create_time 
       from community where community_id = ?
       `
	if err := db.Get(community, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrorInvalidID
		}
	}
	return community, err
}
