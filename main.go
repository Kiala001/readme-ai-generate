package main

import (
	"fmt"

	"github.com/Kiala001/readme-ai-generate/internal/ai"
	"github.com/Kiala001/readme-ai-generate/internal/analyzer"
)

/*func main() {
	files, _ := analyzer.WalkerProject(".")
	for _, file := range files {
		info, _ := analyzer.ExtractFileInfo(file)
		fmt.Println("Arquivo:", info.Path)
		fmt.Println("Funções Exportadas:", info.Functions)
		fmt.Println("Comentários:", info.Comments)
		fmt.Println("---------------------------")
	}
}*/

func main() {
    // Analyzer
    files, _ := analyzer.WalkerProject(".")
    var analysis []*analyzer.FileAnalysis
    for _, f := range files {
        info, _ := analyzer.ExtractFileInfo(f)
        analysis = append(analysis, info)
    }

    // Build prompt
    prompt := ai.BuildPrompt("MeuProjetoGo", analysis)

    // Chamada Gemini
    readme, err := ai.GenerateReadmeWithGemini(prompt)
    if err != nil {
        panic(err)
    }

    fmt.Println("README gerado:\n", readme)
}
