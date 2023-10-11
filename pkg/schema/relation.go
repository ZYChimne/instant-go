package schema

type UpdateFollowingRequest struct {
	TargetID uint   `json:"targetID"`
}

type FollowingResponse struct {
	FollowingID uint `json:"followingID"`
	TargetID    uint `json:"targetID"`
	IsFriend    bool `json:"isFriend"`
	TargetType int  `json:"targetType"`
}

type JointFollowing struct {
	ID          uint   `json:"followingID"`
	UserID      uint   `json:"userID"`
	TargetID    uint   `json:"targetID"`
	IsFriend    bool   `json:"isFriend"`
	TargetType int    `json:"accountType"`
	Username    string `json:"username"`
	Nickname	string `json:"nickname"`
	Avatar      string `json:"avatar"`
}