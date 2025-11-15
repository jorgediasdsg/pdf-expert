# 📄 PDF Analyzer — Fase 1 (versão simples e acoplada)

Este projeto começa com uma implementação **mínima e propositalmente acoplada**, cujo objetivo é:

1. **Enviar um PDF**
2. **Abrir o arquivo**
3. **Ler o texto**
4. **Contar o número de palavras**
5. **Retornar o resultado**

Nada de rotas elegantes.  
Nada de camadas.  
Nada de arquitetura limpa.  
Nada de boas práticas.

A intenção nesta fase é **validar a tecnologia de leitura/extração de PDF** e estabelecer uma base mínima de funcionamento, antes de evoluir para uma arquitetura mais robusta.

---

## 🚦 Fases do Projeto

### **Fase 1 — Versão acoplada e simples (atual)**
- Um único arquivo Go (`main.go`)
- Rota única `/analyze` (POST)
- Recebe o PDF via upload
- Extrai texto usando biblioteca PDF
- Conta palavras
- Retorna JSON
- Sem separação de camadas
- Sem testes
- Sem estruturação
- Sem interfaces
- Sem worker pool
- Sem repositório
- Sem padrões

É propositalmente uma versão **feia porém funcional**.

---

### **Fase 2 — Separação mínima de responsabilidades**
- Criar pacotes simples (pdf, http, service)
- Melhorar organização
- Rotas mais claras
- Código reaproveitável
- Reduzir acoplamento
- Começar a introduzir testes

---

### **Fase 3 — Arquitetura limpa**
- Domínio separado  
- Interfaces (ports)  
- Adapters para HTTP, PDF e Storage  
- Camada de aplicação (serviços)  
- Erros estruturados  
- Configuração por env  
- Logging melhorado  

---

### **Fase 4 — Sistema assíncrono com Worker Pool**
- Fila interna baseada em canais  
- Workers concorrentes  
- Retry / backoff  
- Timeout via context  
- Storage real (SQLite ou Redis)  
- Métricas e healthcheck  
- ADRs completas  
- Docker / docker-compose  
- Observabilidade  

---

## 🧭 Objetivo final

Construir um sistema que demonstre claramente:

- evolução técnica,  
- domínio de refatoração,  
- compreensão de boas práticas,  
- capacidade de justificar decisões,  
- crescimento real de arquitetura.

---

## 🛠 Tecnologias

Nesta fase:

- **Go**
- Biblioteca simples de PDF (ex: `ledongthuc/pdf`)
- `net/http`
- Apenas o básico para rodar

---

## ▶️ Como executar

```bash
go run main.go

curl -X POST -F "file=@meuarquivo.pdf" http://localhost:8080/analyze
```

## Resposta esperada (simplificada):

```json
{
  "word_count": 1234
}
```

# 📌 Futuro (resumo de tudo que virá)

- Worker Pool
- Fila via channels
- Camadas limpas
- Interfaces (ports)
- Repos adaptáveis
- Observabilidade real
- Retry com backoff exponencial
- Timeout por job
- Storage (SQLite/Redis)
- Logs estruturados (slog)
- Dockerfile + compose
- ADRs documentando cada escolha