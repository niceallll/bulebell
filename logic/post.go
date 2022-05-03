package logic

import (
	"bulebell/dao/mysql"
	"bulebell/models"
	"bulebell/pkg/snowflake"
	"go.uber.org/zap"
)

func CreatePost(p *models.Post) (err error) {
	//生产postid
	p.ID = snowflake.GenID()
	//保存到数据库
	return mysql.CreatePost(p)
	//返回
}
func GetPostById(pid int64) (data *models.ApiPostDetail, err error) {
	//取出post_id
	post, err := mysql.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) fialed",
			zap.Int64("pid", pid),
			zap.Error(err))
		return
	}
	user, err := mysql.GetUserId(post.AuthorID)
	if err != nil {
		zap.L().Error("logic.GetUserId(post.AuthorID) fialed",
			zap.Int64("AuthorIDd", post.AuthorID),
			zap.Error(err))
		return
	}
	//根据社区id查询社区详情
	community, err := mysql.GetCommunityByID(post.CommunityID)
	if err != nil {
		zap.L().Error("logic.GetCommunityByID(post.CommunityID) fialed",
			zap.Int64("post.CommunityID", post.CommunityID),
			zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		user.Username,
		post,
		community,
	}
	data.AuthorName = user.Username
	data.CommunityDetail = community
	return
}

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		user, err := mysql.GetUserId(post.AuthorID)
		if err != nil {
			zap.L().Error("logic.GetUserId(post.AuthorID) fialed",
				zap.Int64("AuthorIDd", post.AuthorID),
				zap.Error(err))
			continue
		}
		//根据社区id查询社区详情
		community, err := mysql.GetCommunityByID(post.CommunityID)
		if err != nil {
			zap.L().Error("logic.GetCommunityByID(post.CommunityID) fialed",
				zap.Int64("post.CommunityID", post.CommunityID),
				zap.Error(err))
			continue
		}
		postdetail := &models.ApiPostDetail{
			user.Username,
			post,
			community,
		}
		data = append(data, postdetail)

	}
	return
}
