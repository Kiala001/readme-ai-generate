

# Readme AI Generate

![Go](https://img.shields.io/badge/Go-1.20%2B-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Gemini API](https://img.shields.io/badge/Gemini_API-Enabled-FF6838?style=for-the-badge&logo=google&logoColor=white)
![License](https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge)

## 📖 Introdução

`Readme AI Generate` é uma ferramenta CLI poderosa e inteligente, desenvolvida em Go, projetada para automatizar e aprimorar o processo de criação de arquivos `README.md` para seus projetos Go. Ao analisar a estrutura do seu projeto, arquivos Go e funções, esta ferramenta constrói um contexto rico que é então utilizado por um modelo de inteligência artificial (Gemini) para gerar um `README` completo, detalhado e preciso.

**Chega de começar do zero!** Deixe o `Readme AI Generate` fazer o trabalho pesado, permitindo que você se concentre no desenvolvimento do seu código.

### ✨ Funcionalidades

*   **Análise de Projeto:** Percorre a estrutura de diretórios do seu projeto Go, identificando arquivos `.go`.
*   **Extração de Informações:** Analisa arquivos Go para extrair funções públicas, nomes de pacotes e comentários relevantes.
*   **Construção de Prompt Inteligente:** Gera um prompt detalhado para a IA, incorporando as informações extraídas do projeto.
*   **Geração de README com IA (Gemini):** Utiliza a API do Google Gemini para criar um `README.md` abrangente com base no prompt.
*   **Listagem de Modelos AI:** Capacidade de listar os modelos Gemini disponíveis para geração.

## 🚀 Instalação

Para instalar `Readme AI Generate`, certifique-se de ter o Go instalado (versão 1.20 ou superior).

1.  **Clone o repositório:**
    ```bash
    git clone https://github.com/your-username/readme-ai-generate.git # Substitua pelo caminho real do seu repositório
    cd readme-ai-generate
    ```

2.  **Instale a ferramenta:**
    ```bash
    go install .
    ```
    Ou, se quiser instalar a partir de qualquer diretório:
    ```bash
    go install github.com/your-username/readme-ai-generate@latest # Substitua pelo caminho real do seu repositório
    ```

3.  **Configurar Chave da API Gemini:**
    `Readme AI Generate` requer uma chave de API do Google Gemini. Obtenha sua chave no [Google AI Studio](https://aistudio.google.com/app/apikey).
    Defina a chave como uma variável de ambiente:
    ```bash
    export GEMINI_API_KEY="SUA_CHAVE_API_AQUI"
    ```
    Recomenda-se adicionar esta linha ao seu arquivo `.bashrc`, `.zshrc` ou equivalente para que a chave esteja sempre disponível.

4.  **Verificar a instalação:**
    ```bash
    readme-ai-generate --help
    ```
    Se a instalação foi bem-sucedida, você verá as opções de ajuda da ferramenta.

## 💡 Uso

Para gerar um `README.md` para o seu projeto, navegue até o diretório raiz do projeto e execute a ferramenta.

1.  **Navegue até o diretório do seu projeto Go:**
    ```bash
    cd /path/to/your/go/project
    ```

2.  **Certifique-se de que a variável de ambiente `GEMINI_API_KEY` esteja configurada.**

3.  **Gere o README:**
    ```bash
    readme-ai-generate
    ```
    A ferramenta irá analisar seu projeto e criar um arquivo `README.md` no diretório atual.

4.  **Listar modelos Gemini disponíveis (opcional):**
    ```bash
    readme-ai-generate models
    ```
    Este comando listará os modelos de IA do Gemini que podem ser usados para geração de conteúdo.

**Observação:** O README gerado pela IA é um excelente ponto de partida. Sempre revise e refine-o para garantir que ele atenda perfeitamente às suas necessidades e reflita com precisão seu projeto.

## 🏗️ Estrutura do Projeto

A arquitetura do `Readme AI Generate` é organizada em pacotes internos para modularidade e clareza.

```
.
├── go.mod                  # Módulos Go do projeto
├── go.sum                  # Checksums das dependências
├── main.go                 # Ponto de entrada principal da aplicação. Orquestra a análise do projeto e a geração do README.
└── internal/               # Pacotes internos do projeto, não destinados ao consumo externo
    ├── ai/                 # Lida com a interação com modelos de IA e a construção de prompts.
    │   ├── client.go       # Cliente para interagir com a API do Google Gemini.
    │   │   └── Funcs: GenerateReadmeWithGemini, ListAvailableModels
    │   └── prompt.go       # Constrói o prompt detalhado para a IA.
    │       └── Funcs: BuildPrompt
    └── analyzer/           # Componentes responsáveis pela análise e extração de informações do projeto Go.
        ├── extractor.go    # Extrai informações detalhadas de arquivos Go individuais.
        │   └── Funcs: ExtractFileInfo (apenas funções públicas são consideradas)
        └── walker.go       # Percorre a estrutura de diretórios do projeto para encontrar arquivos Go.
            └── Funcs: WalkerProject
```

### Detalhes dos Componentes:

*   **`main.go`**:
    É o coração da aplicação. Ele inicia o processo, chamando o `walker` para encontrar arquivos, o `extractor` para analisá-los, o `prompt` para construir a requisição e, finalmente, o `ai client` para gerar o `README`. Comentário: "Analyzer the files in the current project directory" refere-se à orquestração dessas etapas.

*   **`internal/ai/client.go`**:
    Este arquivo é responsável por toda a comunicação com a API do Google Gemini.
    -   `GenerateReadmeWithGemini(prompt string)`: Envia o prompt de texto para o modelo Gemini e retorna o `README` gerado.
    -   `ListAvailableModels()`: Consulta a API para obter uma lista dos modelos Gemini disponíveis.

*   **`internal/ai/prompt.go`**:
    Encapsula a lógica para formatar e construir o prompt que será enviado à IA.
    -   `BuildPrompt(projectName string, projectInfo *analyzer.ProjectInfo)`: Recebe o nome do projeto e informações detalhadas (obtidas do `analyzer`) e constrói um prompt textual compreensível para a IA.

*   **`internal/analyzer/extractor.go`**:
    Um parser focado em arquivos Go individuais.
    -   `ExtractFileInfo(filePath string)`: Lê um arquivo `.go`, analisa seu Abstract Syntax Tree (AST) e extrai informações cruciais como nome do pacote, dependências, funções públicas (nome, parâmetros, retornos, comentários), structs, interfaces, etc.

*   **`internal/analyzer/walker.go`**:
    Responsável por percorrer recursivamente o sistema de arquivos.
    -   `WalkerProject(root string)`: Inicia a varredura a partir de um diretório raiz, identifica todos os arquivos `.go` e os passa para o `extractor` para análise. Retorna uma estrutura consolidada com todas as informações do projeto.

## 🛠️ Tecnologias

*   **Go**: Linguagem de programação principal.
*   **Google Gemini API**: Para capacidades de geração de texto de IA.
*   **Go Standard Library**: Utilização extensiva de pacotes como `os`, `path/filepath`, `go/ast`, `go/parser`, `go/token` para análise de código.

## 📄 Licença

Este projeto está licenciado sob a Licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---