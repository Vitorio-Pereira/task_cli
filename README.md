# task_cli

> **Projeto task_cli | Feito por Vitório Alberti**
---
## Sobre
O task_cli é um To Do simples e fácil de ser utilizado no terminal.

---
## Stack

| Camada  | Tecnologia      |
|---------|-----------------|
| Backend | Golang          |
| Libs    | Cobra + stdlib  |

---
## Quick Start

```bash
# 1. Clone
git clone https://github.com/Vitorio-Pereira/task_cli.git
cd task_cli

# Setup local
go build -o task ./cmd/main.go

# Mover para PATH
mv task /usr/local/bin/task

# Funcional para qualquer terminal
task add "comprar café"

```

### How to use

Temos 4 funções:
- task add "titulo"   → adiciona nova task
- task list           → lista todas as tasks  
- task done <id>      → marca como concluída
- task remove <id>    → remove pelo ID