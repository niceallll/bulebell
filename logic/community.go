package logic

import (
	"bulebell/dao/mysql"
	"bulebell/models"
)

func GetCommunityList() ([]*models.Community, error) {
	//查数据库。查找所有的community 并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	//查数据库。查找所有的communityid 并返回
	return mysql.GetCommunityByID(id)
}
