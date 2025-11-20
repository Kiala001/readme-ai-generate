
# ✨ Readme AI Generate

An intelligent Go-based tool designed to automate the creation of comprehensive `README.md` files for your projects. By analyzing your Go codebase and leveraging the power of Google Gemini AI, `Readme AI Generate` crafts detailed and informative documentation, saving you valuable time and ensuring consistent project descriptions.

## 🌟 Introdução

Escrever um bom `README.md` é essencial para qualquer projeto, mas muitas vezes é uma tarefa demorada e repetitiva. `Readme AI Generate` surge para resolver esse problema, fornecendo uma solução automatizada que:

1.  **Analisa a estrutura e o código do seu projeto Go:** Entendendo os componentes e as funcionalidades do seu código.
2.  **Cria um prompt inteligente:** Baseado na análise do projeto, gerando uma requisição otimizada para a IA.
3.  **Utiliza a API Google Gemini:** Para gerar um `README.md` completo, claro e bem estruturado.

Com `Readme AI Generate`, você pode focar no desenvolvimento, sabendo que a documentação será gerada de forma eficiente e profissional.

### Como Funciona

O fluxo de trabalho é simples:

1.  Você aponta o `Readme AI Generate` para a raiz do seu projeto Go.
2.  O programa percorre seus diretórios e arquivos, extraindo informações cruciais (nomes de arquivos, funções, comentários, estrutura geral).
3.  Essas informações são usadas para construir um prompt detalhado e específico para o seu projeto.
4.  O prompt é enviado à API do Google Gemini.
5.  A IA processa o prompt e retorna um `README.md` formatado em Markdown, que é então exibido no console ou salvo em um arquivo.

## 🚀 Instalação

Para começar a usar `Readme AI Generate`, siga os passos abaixo:

### Pré-requisitos

*   **Go (versão 1.18 ou superior):** Certifique-se de ter o Go instalado em sua máquina. Você pode baixá-lo em [go.dev](https://go.dev/dl/).
*   **Google Gemini API Key:** Você precisará de uma chave de API válida para acessar o modelo Gemini. Obtenha uma em [Google AI Studio](https://makersuite.google.com/app/apikey).

### Passos de Instalação

1.  **Clone o repositório:**
    ```bash
    git clone https://github.com/seu-usuario/readme-ai-generate.git # Altere para o seu repositório
    cd readme-ai-generate
    ```

2.  **Instale o executável:**
    ```bash
    go install .
    ```
    Isso instalará o executável `readme-ai-generate` no seu `$GOPATH/bin` (ou `$GOBIN`), tornando-o acessível globalmente.

OU

1. **Execute o comando:**
```bash
    go get https://github.com/Kiala001/readme-ai-generate
```

## 💡 Uso

Para gerar um `README.md` para seu projeto, você precisa configurar sua chave de API do Gemini e então executar o comando.

1.  **Configure sua chave de API do Gemini:**
    Defina a chave de API como uma variável de ambiente:
    ```bash
    export GEMINI_API_KEY="SUA_CHAVE_DE_API_GEMINI"
    ```
    (Recomenda-se adicionar isso ao seu `.bashrc`, `.zshrc` ou equivalente para persistência).

2.  **Execute o gerador de README:**
    Navegue até o diretório raiz do seu projeto (o mesmo que você deseja documentar) e execute:
    ```bash
    readme-ai-generate .
    ```
    Ou, se o projeto estiver em outro lugar:
    ```bash
    readme-ai-generate /caminho/para/seu/projeto
    ```

    O `README.md` gerado será impresso na saída padrão. Você pode redirecioná-lo para um arquivo:
    ```bash
    readme-ai-generate . > README.md
    ```

## 🏗️ Estrutura do Projeto

O projeto `Readme AI Generate` é organizado de forma modular, seguindo as convenções de projetos Go para facilitar a manutenção e escalabilidade.

```
.
├── internal/
│   ├── ai/
│   │   ├── client.go         # Lógica para interagir com a API do Google Gemini.
│   │   └── prompt.go         # Funções para construir prompts inteligentes para a IA.
│   └── analyzer/
│       ├── extractor.go      # Extrai informações detalhadas de arquivos Go individuais.
│       └── walker.go         # Percorre a estrutura de diretórios do projeto.
└── main.go                 # Ponto de entrada principal da aplicação, orquestra o fluxo.
```

### Detalhes dos Módulos:

*   **`main.go`**:
    *   Este é o coração da aplicação CLI. Ele inicializa o processo, chama os componentes do `analyzer` para coletar dados do projeto e, em seguida, utiliza os módulos do `ai` para gerar o README.

*   **`internal/ai/`**:
    *   **`client.go`**:
        *   `GenerateReadmeWithGemini(prompt string) (string, error)`: Envia um prompt à API do Google Gemini e retorna o `README.md` gerado.
        *   `ListAvailableModels() ([]string, error)`: Lista os modelos de IA disponíveis através da API.
    *   **`prompt.go`**:
        *   `BuildPrompt(projectName string, projectInfo string) string`: Constrói um prompt detalhado para a IA, usando o nome do projeto e as informações extraídas do código.

*   **`internal/analyzer/`**:
    *   **`extractor.go`**:
        *   `ExtractFileInfo(filePath string) (FileInfo, error)`: Analisa um arquivo Go específico, extraindo informações relevantes como nome do arquivo, caminho, funções públicas, etc. (Note que apenas funções públicas/maiúsculas iniciais são consideradas).
    *   **`walker.go`**:
        *   `WalkerProject(root string) ([]FileInfo, error)`: Percorre recursivamente o diretório raiz de um projeto, coletando informações de todos os arquivos Go relevantes através do `extractor`.

## 🛠️ Tecnologias

*   **Go**: A linguagem de programação principal para o desenvolvimento da ferramenta.
*   **Google Gemini API**: Utilizada para a capacidade de geração de texto inteligente.
*   **Markdown**: O formato de saída padrão para o `README.md` gerado.

## 📄 Licença

Este projeto está sob a licença MIT. Consulte o arquivo [LICENSE](LICENSE) para mais detalhes.