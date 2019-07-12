package entity

type AbstractEntity interface {
	GetId() uint64
	SetId(id uint64)
	GetSId() string
	SetSId(sId string)
}

type Entity struct {
	id  uint64
	sId string
}

func (e *Entity) GetId() uint64 {
	return e.id
}
func (e *Entity) SetId(id uint64) {
	e.id = id
}
func (e *Entity) GetSId() string {
	return e.sId
}
func (e *Entity) SetSId(sId string) {
	e.sId = sId
}
