package br.com.cassol.course_manager.repository;

import br.com.cassol.course_manager.entity.Aluno;
import org.springframework.data.jpa.repository.JpaRepository;

public interface AlunoRepository extends JpaRepository<Aluno, Long> {
}
