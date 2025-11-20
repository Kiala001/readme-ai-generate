package main

import (
	"fmt"
	"os"

	"github.com/Kiala001/readme-ai-generate/internal/ai"
	"github.com/Kiala001/readme-ai-generate/internal/analyzer"
)

func main() {
	// Analyzer the files in the current project directory
	files, _ := analyzer.WalkerProject(".")
	var analysis []*analyzer.FileAnalysis
	for _, f := range files {
		info, _ := analyzer.ExtractFileInfo(f)
		analysis = append(analysis, info)
	}

	prompt := ai.BuildPrompt("Readme AI Generate", analysis)

	readme, err := ai.GenerateReadmeWithGemini(prompt)
	if err != nil {
		panic(err)
	}

	readmeFile := "README.md"
	file, err := os.Create(readmeFile)
	if err != nil {
		fmt.Printf("Erro ao criar o arquivo %s: %v\n", readmeFile, err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(readme)
	if err != nil {
		fmt.Printf("Erro ao escrever no arquivo %s: %v\n", readmeFile, err)
		return
	}

	fmt.Printf("README gerado com sucesso: %s\n", readmeFile)
}
