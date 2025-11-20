package main

import (
	"fmt"
	"os"

	"github.com/Kiala001/readme-ai-generate/internal/ai"
	"github.com/Kiala001/readme-ai-generate/internal/analyzer"
	"github.com/spf13/cobra"
)

var projectName string

func main() {

	var rootCmd = &cobra.Command{
		Use:   "readme-generate",
		Short: "Gere READMEs automaticamente para seus projetos Go",
	}

	var generateCmd = &cobra.Command{
		Use:   "generate-readme",
		Short: "Gera um README para o projeto especificado",
		Run: func(cmd *cobra.Command, args []string) {

			// Analyzer
			files, _ := analyzer.WalkerProject(".")
			var analysis []*analyzer.FileAnalysis
			for _, f := range files {
				info, _ := analyzer.ExtractFileInfo(f)
				analysis = append(analysis, info)
			}

			prompt := ai.BuildPrompt(projectName, analysis)

			readme, err := ai.GenerateReadmeWithGemini(prompt)
			if err != nil {
				fmt.Printf("Erro ao gerar README: %v\n", err)
				os.Exit(1)
			}

			readmeFile := "README.md"
			file, err := os.Create(readmeFile)
			if err != nil {
				fmt.Printf("Erro ao criar o arquivo %s: %v\n", readmeFile, err)
				os.Exit(1)
			}
			defer file.Close()

			_, err = file.WriteString(readme)
			if err != nil {
				fmt.Printf("Erro ao escrever no arquivo %s: %v\n", readmeFile, err)
				os.Exit(1)
			}

			fmt.Printf("README gerado com sucesso: %s\n", readmeFile)
		},
	}

	generateCmd.Flags().StringVar(&projectName, "project-name", "", "Nome do projeto para o README")
	generateCmd.MarkFlagRequired("project-name")

	rootCmd.AddCommand(generateCmd)
	rootCmd.Execute()
}
