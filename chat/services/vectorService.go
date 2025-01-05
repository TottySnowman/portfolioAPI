package chatService

import (
	"context"
	chatModel "portfolioAPI/chat/models"
"math/rand/v2"
	"github.com/qdrant/go-client/qdrant"
)

type VectorService struct {
	vectorClient *qdrant.Client
}

const collectionName = "portfolio"
const vectorSize = uint64(384)

func NewVectorService() *VectorService {
	client, _ := qdrant.NewClient(&qdrant.Config{
		Host: "qdrant",
		Port: 6334,
	})

	return &VectorService{
		vectorClient: client,
	}
}

func (service *VectorService) CreateCollectionIfNeeded() {
	exists, err := service.vectorClient.CollectionExists(context.Background(), collectionName)

	if err != nil {
		panic("Failed to connect to vector DB")
	}

	if !exists {
		service.createCollection()
	}
}

func (service *VectorService) createCollection() {
	service.vectorClient.CreateCollection(context.Background(), &qdrant.CreateCollection{
		CollectionName: collectionName,
		VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
			Size:     vectorSize,
			Distance: qdrant.Distance_Cosine,
		}),
	})
}

func (service *VectorService) InsertVector(vector chatModel.FeatureExtractionResponse, query string) {
	service.CreateCollectionIfNeeded()

	_, err := service.vectorClient.Upsert(context.Background(), &qdrant.UpsertPoints{
		CollectionName: collectionName,
		Points: []*qdrant.PointStruct{
			{
				Id:      qdrant.NewIDNum(uint64(rand.IntN(10000))),
				Vectors: qdrant.NewVectors(vector...),
				Payload: qdrant.NewValueMap(map[string]any{"text": query}),
			},
		},
	})

	if err != nil {
		return
	}
}
