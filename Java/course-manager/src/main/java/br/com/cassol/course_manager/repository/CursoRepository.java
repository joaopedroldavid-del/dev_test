package br.com.cassol.course_manager.repository;

import br.com.cassol.course_manager.entity.Curso;
import org.springframework.data.jpa.repository.JpaRepository;

public interface CursoRepository extends JpaRepository<Curso, Long> {
}
