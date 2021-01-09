package serializer

import "shManager/model"

// User 用户序列化器
type User struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Email     string `json:"email"`
	CreatedAt int64  `json:"createdAt"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		Nickname:  user.Nickname,
		Avatar:    user.Avatar,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Unix(),
	}
}
