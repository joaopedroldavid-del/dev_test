package br.com.cassol.course_manager.entity;

import jakarta.persistence.*;
import lombok.Data;

@Entity
@Data
public class Matricula {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @ManyToOne(optional = false)
    private Aluno aluno;

    @ManyToOne(optional = false)
    private Curso curso;

    private Double nota;
}
