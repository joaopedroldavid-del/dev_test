import requests #type: ignore

BASE_URL = "http://localhost:8080/cursos"

def test_criar_curso():
    payload = {
        "nome": "Curso Teste",
        "descricao": "Descrição Teste"
    }
    print(f"\n[POST] {BASE_URL} → {payload}")
    response = requests.post(BASE_URL, json=payload)
    print(f"Status: {response.status_code}")
    print("Response:", response.json())
    
    assert response.status_code == 200 or response.status_code == 201, f"Expected 200 or 201, got {response.status_code}"
    data = response.json()
    assert data["nome"] == payload["nome"], f"Expected name {payload['nome']}, got {data['nome']}"
    assert data["descricao"] == payload["descricao"], f"Expected description {payload['descricao']}, got {data['descricao']}"
    return data["id"] 

def test_listar_cursos():
    print(f"\n[GET] {BASE_URL}")
    response = requests.get(BASE_URL)
    print(f"Status: {response.status_code}")
    print("Response:", response.json())
    
    assert response.status_code == 200, f"Expected 200, got {response.status_code}"
    data = response.json()
    assert isinstance(data, list), f"Expected list, got {type(data)}"

def test_buscar_por_id(curso_id):
    url = f"{BASE_URL}/{curso_id}"
    print(f"\n[GET] {url}")
    response = requests.get(url)
    print(f"Status: {response.status_code}")
    print("Response:", response.json())
    
    assert response.status_code == 200, f"Expected 200, got {response.status_code}"
    data = response.json()
    assert data["id"] == curso_id, f"Expected id {curso_id}, got {data['id']}"

def test_atualizar_curso(curso_id):
    updated_payload = {
        "nome": "Curso Atualizado",
        "descricao": "Descrição Atualizada"
    }
    url = f"{BASE_URL}/{curso_id}"
    print(f"\n[PUT] {url} → {updated_payload}")
    response = requests.put(url, json=updated_payload)
    print(f"Status: {response.status_code}")
    print("Response:", response.json())
    
    assert response.status_code == 200, f"Expected 200, got {response.status_code}"
    data = response.json()
    assert data["nome"] == updated_payload["nome"], f"Expected name {updated_payload['nome']}, got {data['nome']}"
    assert data["descricao"] == updated_payload["descricao"], f"Expected description {updated_payload['descricao']}, got {data['descricao']}"

def test_calcular_media(curso_id):
    url = f"{BASE_URL}/{curso_id}/media"
    print(f"\n[GET] {url}")
    response = requests.get(url)
    print(f"Status: {response.status_code}")
    print("Response:", response.json())
    
    assert response.status_code == 200, f"Expected 200, got {response.status_code}"
    data = response.json()
    media = data.get("media") if isinstance(data, dict) else data
    assert isinstance(media, (float, int)), f"Expected float or int for media, got {type(media)}"

def test_deletar_curso(curso_id):
    url = f"{BASE_URL}/{curso_id}"
    print(f"\n[DELETE] {url}")
    response = requests.delete(url)
    print(f"Status: {response.status_code}")
    assert response.status_code == 204, f"Expected 204 No Content, got {response.status_code}"

    print(f"[GET] {url} (confirm delete)")
    response = requests.get(url)
    print(f"Status: {response.status_code}")
    print("Response:", response.text)
    assert response.status_code == 404, f"Expected 404 after delete, got {response.status_code}"

def run_all_tests():
    curso_id = test_criar_curso()
    test_listar_cursos()
    test_buscar_por_id(curso_id)
    test_atualizar_curso(curso_id)
    test_calcular_media(curso_id)
    test_deletar_curso(curso_id)
    print("\n✅ All tests passed successfully!")

if __name__ == "__main__":
    run_all_tests()