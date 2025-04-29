package models

type Matricula struct {
	ID      uint    `gorm:"primaryKey" json:"id"`
	AlunoID uint    `json:"aluno_id"`
	CursoID uint    `json:"curso_id"`
	Nota    float64 `json:"nota"`

	Aluno  Aluno  `gorm:"foreignKey:AlunoID" json:"-"`
	Curso  Curso  `gorm:"foreignKey:CursoID" json:"-"`
}