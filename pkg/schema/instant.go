package schema

type UpsertInstantRequest struct {
	ID          uint   `json:"id"`
	InstantType int    `json:"instantType"`
	Content     string `json:"content"`
	Visibility  int    `json:"visibility"` // 0: public, 1: followers, 2: friends, 3: private
	RefOriginID uint   `json:"refOriginId"`
}
