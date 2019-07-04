package postgreSQL

import (
	"fmt"
	"time"
)

type Conditions struct {
	Time        time.Time
	Location    string
	Temperature float32
}

func (c *Conditions) All() {
	var infos []*Conditions
	if err := DB.Distribution.Find(&infos).Error; err != nil {
		fmt.Println(infos)
	}
	fmt.Println(infos)
}

func (c *Conditions) TableName() string {
	return "conditions"
}
