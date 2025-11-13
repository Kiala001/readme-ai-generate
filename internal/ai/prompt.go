package ai

import (
	"fmt"
	"strings"

	"github.com/Kiala001/readme-ai-generate/internal/analyzer"
)

func BuildPrompt(projectName string, files []*analyzer.FileAnalysis) string {
    sb := strings.Builder{}
    sb.WriteString(fmt.Sprintf("Você é um assistente especializado em Go. Gere um README completo para o projeto '%s'.\n", projectName))
    sb.WriteString("Inclua seções: Introdução, Instalação, Uso, Estrutura do Projeto, Tecnologias, Licença.\n\n")
    sb.WriteString("Resumo do projeto:\n")

    for _, f := range files {
        sb.WriteString(fmt.Sprintf("Arquivo: %s\n", f.Path))
        if len(f.Comments) > 0 {
            sb.WriteString("Comentários:\n")
            for _, c := range f.Comments {
                sb.WriteString("- " + strings.TrimPrefix(c, "//") + "\n")
            }
        }
        if len(f.Functions) > 0 {
            sb.WriteString("Funções públicas:\n")
            for _, fn := range f.Functions {
                sb.WriteString("- " + fn + "\n")
            }
        }
        sb.WriteString("\n")
    }

    sb.WriteString("Gere o README em Markdown.\n")
    return sb.String()
}
