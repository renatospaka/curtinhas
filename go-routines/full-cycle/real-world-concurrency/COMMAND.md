## 10.000 cliques com 100 acessos/usuários simultâneos
- ab -n 10000 -c 100 http://localhost:3000/
- go run -race main.go => verifica em tempo de execução se há race conditions