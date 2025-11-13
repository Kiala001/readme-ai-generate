package analyzer

import (
	"bufio"
	"os"
	"strings"
)

type FileAnalyzer interface {
	AnalyzeFile(filePath string) (FileAnalysis, error)
}

type FileAnalysis struct {
	Path string
	Functions []string
	Comments []string
}

func ExtractFileInfo(filePath string) (*FileAnalysis, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	analysis := &FileAnalysis{	Path: filePath	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "//") {
			analysis.Comments = append(analysis.Comments, line)
		}
		if strings.HasPrefix(line, "func ") {
			// Apenas funções públicas (maiuscula inicial)
			funcName := strings.Fields(line)[1]
			if len(funcName) > 0 && funcName[0] >= 'A' && funcName[0] <= 'Z' {
				analysis.Functions = append(analysis.Functions, funcName)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return analysis, nil
}