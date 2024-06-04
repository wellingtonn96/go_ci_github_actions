# Endpoint: Documentação da API

curl -X GET http://localhost:8000/nome/Guilherme

# Endpoint: TodosAlunos
curl -X GET http://localhost:8000/alunos

# Endpoint: CriarNovoAluno
curl -X POST http://localhost:8000/alunos -d '{"nome": "Fulano", "idade": 25, "cpf": "12345678900", "rg": "12345641512"}' -H "Content-Type: application/json"

# Endpoint: BuscarAlunoPorID
curl -X GET http://localhost:8000/alunos/1

# Endpoint: DeletarAluno
curl -X DELETE http://localhost:8000/alunos/1

# Endpoint: EditarAluno
curl -X PATCH http://localhost:8000/alunos/1 -d '{"nome": "Fulano Editado", "idade": 30, "cpf": "12345678901"}' -H "Content-Type: application/json"

# Endpoint: BuscaAlunoPorCPF
curl -X GET http://localhost:8000/alunos/cpf/12345678900

# Endpoint: ExibePaginaIndex (HTML)
curl -X GET http://localhost:8000/index

# Endpoint: RotaNaoEncontrada (404)
curl -X GET http://localhost:8000/notfound