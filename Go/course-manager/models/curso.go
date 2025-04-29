package models

type Curso struct {
	ID        uint     `gorm:"primaryKey" json:"id"`
	Nome      string   `json:"nome"`
	Descricao string   `json:"descricao"`
	Matriculas []Matricula `json:"-"`
}