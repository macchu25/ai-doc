package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/qdrant/go-client/qdrant"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type QdrantClient struct {
	client *qdrant.PointsClient
	conn   *grpc.ClientConn
}

func NewQdrantClient(addr string) (*QdrantClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := qdrant.NewPointsClient(conn)
	return &QdrantClient{
		client: &client,
		conn:   conn,
	}, nil
}

func (q *QdrantClient) CreateCollection(ctx context.Context, name string, size uint64) error {
	colClient := qdrant.NewCollectionsClient(q.conn)
	_, err := colClient.Create(ctx, &qdrant.CreateCollection{
		CollectionName: name,
		VectorsConfig: &qdrant.VectorsConfig{
			Config: &qdrant.VectorsConfig_Params{
				Params: &qdrant.VectorParams{
					Size:     size,
					Distance: qdrant.Distance_Cosine,
				},
			},
		},
	})
	return err
}

func (q *QdrantClient) UpsertPoints(ctx context.Context, collection string, points []*qdrant.PointStruct) error {
	_, err := (*q.client).Upsert(ctx, &qdrant.UpsertPoints{
		CollectionName: collection,
		Points:         points,
	})
	return err
}

func (q *QdrantClient) Close() {
	if q.conn != nil {
		q.conn.Close()
	}
}

// Global instance helper
var qClient *QdrantClient

func InitQdrant() {
	var err error
	qClient, err = NewQdrantClient("localhost:6334") // gRPC port
	if err != nil {
		log.Printf("Failed to connect to Qdrant: %v", err)
	} else {
		fmt.Println("Successfully connected to Qdrant gRPC")
	}
}
