package request

type UserUpdate struct {
	Name     string `form:"name" validate:"max=100"`
	Email    string `form:"email" validate:"email,max=100"`
	TopicIDs []int  `form:"topic_ids[]" validate:"lt=7"`
}
