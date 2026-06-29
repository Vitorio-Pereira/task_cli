# task — CLI To-Do List em Go

## Objetivo Principal
Criar uma ferramenta de linha de comando em Go para gerenciar tarefas. O usuário adiciona, lista, marca como concluída e remove tarefas via terminal. As tarefas são persistidas em um arquivo JSON local.

## Arquitetura de Pastas
```
task_cli/
├── go.mod
├── go.sum
├── cmd/
│   └── main.go
└── internal/
    ├── cmd/
    │   ├── root.go
    │   ├── add.go
    │   ├── list.go
    │   ├── done.go
    │   └── remove.go
    └── store/
        └── store.go
```

## Requisitos Técnicos
- **CLI**: `github.com/spf13/cobra`
- **Persistência**: arquivo JSON em `~/.task/tasks.json`
- **Armazenamento**: leitura/escrita em arquivo, usando `encoding/json`
- **Concorrência**: não necessária (operações síncronas no CLI)

---

## Estrutura de Dados

### Task
```go
type Task struct {
    ID        int       `json:"id"`
    Title     string    `json:"title"`
    Done      bool      `json:"done"`
    CreatedAt time.Time `json:"created_at"`
}
```

---

## Comandos

| Comando               | O que faz                                 |
|-----------------------|-------------------------------------------|
| `task add "titulo"`   | Adiciona uma nova tarefa                  |
| `task list`           | Lista todas as tarefas                    |
| `task done <id>`      | Marca uma tarefa como concluída           |
| `task remove <id>`    | Remove uma tarefa pelo ID                 |

---

## Checklist de Implementação

- [ ] **1. Inicialização do Projeto**
  - Criar o módulo Go com `go mod init`
  - Instalar Cobra com `go get github.com/spf13/cobra`
  - Criar estrutura de pastas

- [ ] **2. Store (`internal/store/store.go`)**
  - Struct `Task` com campos `ID`, `Title`, `Done`, `CreatedAt`
  - Struct `Store` com campo `Tasks []Task`
  - `New() *Store` — construtor, define o caminho `~/.task/tasks.json`
  - `Load() error` — lê o arquivo JSON e popula o store
  - `Save() error` — serializa o store e escreve no arquivo JSON
  - `Add(title string) error` — cria nova task com ID auto-incrementado
  - `List() []Task` — retorna todas as tasks
  - `Done(id int) error` — marca task como concluída pelo ID
  - `Remove(id int) error` — remove task pelo ID

- [ ] **3. Root Command (`internal/cmd/root.go`)**
  - Inicializa o comando raiz com `cobra.Command`
  - Define nome, versão e descrição do CLI
  - Função `Execute()` exportada — chamada pelo `main.go`

- [ ] **4. Subcomandos (`internal/cmd/`)**
  - `add.go` — recebe o título como argumento, chama `store.Add()`
  - `list.go` — chama `store.List()` e imprime as tasks formatadas
  - `done.go` — recebe o ID como argumento, chama `store.Done()`
  - `remove.go` — recebe o ID como argumento, chama `store.Remove()`

- [ ] **5. Entrypoint (`cmd/main.go`)**
  - Inicializa o store
  - Chama `cmd.Execute()`

---

## Conceitos Envolvidos
- CLI com Cobra — comandos, subcomandos, argumentos
- Leitura e escrita de arquivos (`os`, `encoding/json`)
- Manipulação de slices em Go
- Tratamento de erros idiomático
- `os.UserHomeDir()` para caminho portável do arquivo
- `time.Time` para timestamp de criação