package demo2

type AnnouncementMode struct {
	Title       Language `json:"title" db:"title"`
	Content     Language `json:"content" db:"content"`
	StartingAt  int64    `json:"starting_at" db:"starting_at"`
	EndedAt     int64    `json:"ended_at" db:"ended_at"`
	Index       int      `json:"index" db:"index"`
	PublishName string   `json:"publish_name" db:"publish_name"`
	Id          int64    `json:"id"db:"id"`
	TouchAt
}
