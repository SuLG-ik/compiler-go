package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type FileResult struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

func (a *App) OpenFile() (*FileResult, error) {
	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Открыть файл",
		Filters: []runtime.FileFilter{
			{DisplayName: "Текстовые файлы (*.txt;*.go;*.c;*.cpp;*.h;*.pas;*.cs;*.py)", Pattern: "*.txt;*.go;*.c;*.cpp;*.h;*.pas;*.cs;*.py"},
			{DisplayName: "Все файлы (*.*)", Pattern: "*.*"},
		},
	})
	if err != nil {
		return nil, err
	}
	if path == "" {
		return nil, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл: %w", err)
	}

	return &FileResult{
		Path:    path,
		Content: string(data),
	}, nil
}

func (a *App) ReadFileByPath(path string) (*FileResult, error) {
	if path == "" {
		return nil, nil
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("не удалось прочитать файл: %w", err)
	}
	return &FileResult{
		Path:    path,
		Content: string(data),
	}, nil
}

func (a *App) SaveFile(path string, content string) error {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("не удалось сохранить файл: %w", err)
	}
	return nil
}

func (a *App) SaveFileAs(currentPath string, content string) (string, error) {
	defaultName := "untitled.txt"
	if currentPath != "" {
		defaultName = filepath.Base(currentPath)
	}

	path, err := runtime.SaveFileDialog(a.ctx, runtime.SaveDialogOptions{
		Title:           "Сохранить как",
		DefaultFilename: defaultName,
		Filters: []runtime.FileFilter{
			{DisplayName: "Текстовые файлы (*.txt;*.go;*.c;*.cpp;*.h;*.pas)", Pattern: "*.txt;*.go;*.c;*.cpp;*.h;*.pas"},
			{DisplayName: "Все файлы (*.*)", Pattern: "*.*"},
		},
	})
	if err != nil {
		return "", err
	}
	if path == "" {
		return "", nil
	}

	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("не удалось сохранить файл: %w", err)
	}
	return path, nil
}
