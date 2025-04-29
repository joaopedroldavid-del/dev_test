import requests

BASE_URL = "http://localhost:8080/cursos"

def test_criar_curso():
    payload = {
        "nome": "Curso Teste",
        "descricao": "Descrição Teste"
    }
    response = requests.post(BASE_URL, json=payload)
    assert response.status_code == 200, f"Expected 200, got {response.status_code}"
    data = response.json()
    assert data["nome"] == payload["nome"], f"Expected name {payload['nome']}, got {data['nome']}"
    assert data["descricao"] == payload["descricao"], f"Expected description {payload['descricao']}, got {data['descricao']}"
    return data["id"]  # return id to use in other tests

def test_listar_cursos():
    response = requests.get(BASE_URL)
    assert response.status_code == 200, f"Expected 200, got {response.status_code}"
    data = response.json()
    assert isinstance(data, list), f"Expected list, got {type(data)}"

def test_buscar_por_id(curso_id):
    response = requests.get(f"{BASE_URL}/{curso_id}")
    assert response.status_code == 200, f"Expected 200, got {response.status_code}"
    data = response.json()
    assert data["id"] == curso_id, f"Expected id {curso_id}, got {data['id']}"

def test_atualizar_curso(curso_id):
    updated_payload = {
        "nome": "Curso Atualizado",
        "descricao": "Descrição Atualizada"
    }
    response = requests.put(f"{BASE_URL}/{curso_id}", json=updated_payload)
    assert response.status_code == 200, f"Expected 200, got {response.status_code}"
    data = response.json()
    assert data["nome"] == updated_payload["nome"], f"Expected name {updated_payload['nome']}, got {data['nome']}"
    assert data["descricao"] == updated_payload["descricao"], f"Expected description {updated_payload['descricao']}, got {data['descricao']}"

def test_calcular_media(curso_id):
    response = requests.get(f"{BASE_URL}/{curso_id}/media")
    assert response.status_code == 200, f"Expected 200, got {response.status_code}"
    data = response.json()
    assert isinstance(data, float) or isinstance(data, int), f"Expected float or int for media, got {type(data)}"

def test_deletar_curso(curso_id):
    response = requests.delete(f"{BASE_URL}/{curso_id}")
    assert response.status_code == 204, f"Expected 204 No Content, got {response.status_code}"
    # Verify if the course really was deleted
    response = requests.get(f"{BASE_URL}/{curso_id}")
    assert response.status_code == 404, f"Expected 404 after delete, got {response.status_code}"

def run_all_tests():
    curso_id = test_criar_curso()
    test_listar_cursos()
    test_buscar_por_id(curso_id)
    test_atualizar_curso(curso_id)
    test_calcular_media(curso_id)
    test_deletar_curso(curso_id)
    print("✅ All tests passed successfully!")

if __name__ == "__main__":
    run_all_tests()
