package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/qdrant/go-client/qdrant"
)

type EmbedRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type EmbedResponse struct {
	Embedding []float32 `json:"embedding"`
}

func GetEmbedding(text string) ([]float32, error) {
	url := "http://localhost:11434/api/embeddings"
	reqBody, _ := json.Marshal(EmbedRequest{
		Model:  "nomic-embed-text",
		Prompt: text,
	})

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var embedResp EmbedResponse
	if err := json.NewDecoder(resp.Body).Decode(&embedResp); err != nil {
		return nil, err
	}

	return embedResp.Embedding, nil
}

func ChunkCode(content string, chunkSize int) []string {
	// Simple chunking logic (could be improved with AST)
	var chunks []string
	lines := strings.Split(content, "\n")
	
	currentChunk := ""
	for _, line := range lines {
		if len(currentChunk)+len(line) > chunkSize {
			chunks = append(chunks, currentChunk)
			currentChunk = ""
		}
		currentChunk += line + "\n"
	}
	if currentChunk != "" {
		chunks = append(chunks, currentChunk)
	}
	return chunks
}

func IndexProject(projectID string, rootPath string) error {
	// 1. Create collection if not exists
	ctx := context.Background()
	_ = qClient.CreateCollection(ctx, projectID, 768) // nomic-embed-text size is 768

	// 2. Scan files
	files, _ := ScanCodebase(rootPath)
	
	var points []*qdrant.PointStruct

	for _, relPath := range files {
		relPath = strings.TrimSpace(relPath)
		if strings.HasSuffix(relPath, "/") {
			continue
		}

		fullPath := fmt.Join(rootPath, relPath)
		content, err := os.ReadFile(fullPath)
		if err != nil {
			continue
		}

		chunks := ChunkCode(string(content), 1000)
		for i, chunk := range chunks {
			embedding, err := GetEmbedding(chunk)
			if err != nil {
				fmt.Printf("Error embedding chunk %d of %s: %v\n", i, relPath, err)
				continue
			}

			id := uuid.New().String()
			points = append(points, &qdrant.PointStruct{
				Id: qdrant.NewIDUUID(id),
				Vectors: qdrant.NewVectors(embedding...),
				Payload: map[string]*qdrant.Value{
					"file_path": qdrant.NewValueString(relPath),
					"content":   qdrant.NewValueString(chunk),
					"chunk_idx": qdrant.NewValueInt(int64(i)),
				},
			})
		}
	}

	// 3. Batch upsert
	if len(points) > 0 {
		return qClient.UpsertPoints(ctx, projectID, points)
	}
	return nil
}
