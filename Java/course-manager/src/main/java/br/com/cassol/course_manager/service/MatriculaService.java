package br.com.cassol.course_manager.service;

import br.com.cassol.course_manager.entity.Matricula;
import br.com.cassol.course_manager.exception.ResourceNotFoundException;
import br.com.cassol.course_manager.repository.MatriculaRepository;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class MatriculaService {

    private final MatriculaRepository matriculaRepository;

    public MatriculaService(MatriculaRepository matriculaRepository) {
        this.matriculaRepository = matriculaRepository;
    }

    public Matricula salvar(Matricula matricula) {
        return matriculaRepository.save(matricula);
    }

    public List<Matricula> listarTodos() {
        return matriculaRepository.findAll();
    }

    public Matricula buscarPorId(Long id) {
        return matriculaRepository.findById(id)
                .orElseThrow(() -> new ResourceNotFoundException("Matricula não encontrado com ID: " + id));
    }

    public Matricula atualizar(Long id, Matricula matriculaAtualizada) {
        Matricula matricula = buscarPorId(id);
        matricula.setNota(matriculaAtualizada.getNota());
        return matriculaRepository.save(matricula);
    }
    

    public void deletar(Long id) {
        Matricula matricula = buscarPorId(id);
        matriculaRepository.delete(matricula);
    }
}
