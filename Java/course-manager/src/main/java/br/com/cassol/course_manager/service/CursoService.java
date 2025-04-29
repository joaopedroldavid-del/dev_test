package br.com.cassol.course_manager.service;

import br.com.cassol.course_manager.entity.Curso;
import br.com.cassol.course_manager.entity.Matricula;
import br.com.cassol.course_manager.exception.ResourceNotFoundException;
import br.com.cassol.course_manager.repository.CursoRepository;
import br.com.cassol.course_manager.repository.MatriculaRepository;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class CursoService {

    private final CursoRepository cursoRepository;
    private final MatriculaRepository matriculaRepository;

    public CursoService(CursoRepository cursoRepository, MatriculaRepository matriculaRepository) {
        this.cursoRepository = cursoRepository;
        this.matriculaRepository = matriculaRepository;
    }

    public Curso salvar(Curso curso) {
        return cursoRepository.save(curso);
    }

    public List<Curso> listarTodos() {
        return cursoRepository.findAll();
    }

    public Curso buscarPorId(Long id) {
        return cursoRepository.findById(id)
                .orElseThrow(() -> new ResourceNotFoundException("Curso não encontrado com ID: " + id));
    }

    public Curso atualizar(Long id, Curso cursoAtualizado) {
        Curso curso = buscarPorId(id);
        curso.setNome(cursoAtualizado.getNome());
        curso.setDescricao(cursoAtualizado.getDescricao());
        return cursoRepository.save(curso);
    }

    public void deletar(Long id) {
        Curso curso = buscarPorId(id);
        cursoRepository.delete(curso);
    }

    public Double calcularNotaMedia(Long cursoId) {
        Curso curso = buscarPorId(cursoId);
        List<Matricula> matriculas = matriculaRepository.findByCurso(curso);
        return matriculas.stream()
                         .filter(m -> m.getNota() != null)
                         .mapToDouble(Matricula::getNota)
                         .average()
                         .orElse(0.0);
    }
}
