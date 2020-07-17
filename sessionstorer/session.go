package sessionstorer

import "time"

type Session struct {
	Key        string
	CreateTime time.Time
}
