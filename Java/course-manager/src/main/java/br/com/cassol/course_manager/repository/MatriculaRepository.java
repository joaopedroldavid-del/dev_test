package br.com.cassol.course_manager.repository;

import br.com.cassol.course_manager.entity.Matricula;
import br.com.cassol.course_manager.entity.Curso;
import org.springframework.data.jpa.repository.JpaRepository;
import java.util.List;

public interface MatriculaRepository extends JpaRepository<Matricula, Long> {
    List<Matricula> findByCurso(Curso curso);
}
