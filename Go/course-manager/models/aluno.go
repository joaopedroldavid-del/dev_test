package models

type Aluno struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Nome  string `json:"nome"`
	Matriculas []Matricula `json:"-"`
}