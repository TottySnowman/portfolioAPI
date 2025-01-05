package vectorRepo

import (
	"context"
	"github.com/qdrant/go-client/qdrant"
	chatModel "portfolioAPI/chat/models"
)

type VectorRepo struct {
	client *qdrant.Client
}

const collectionName = "portfolio"
const vectorSize = uint64(384)

func NewVectorRepo() *VectorRepo {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "qdrant",
		Port: 6334,
	})

	if err != nil {
		panic(err.Error())
	}

	return &VectorRepo{
		client: client,
	}
}

func (repo *VectorRepo) DoesCollectionExist() bool {
	exists, err := repo.client.CollectionExists(context.Background(), collectionName)

	if err != nil {
		panic("Failed to connect to vector DB")
	}

	return exists
}

func (repo *VectorRepo) CreateCollection() error {
	err := repo.client.CreateCollection(context.Background(), &qdrant.CreateCollection{
		CollectionName: collectionName,
		VectorsConfig: qdrant.NewVectorsConfig(&qdrant.VectorParams{
			Size:     vectorSize,
			Distance: qdrant.Distance_Cosine,
		}),
	})

	return err
}

func (repo *VectorRepo) UpsertVector(vector chatModel.FeatureExtractionResponse, promptModel chatModel.PromptModel) error {

	_, err := repo.client.Upsert(context.Background(), &qdrant.UpsertPoints{
		CollectionName: collectionName,
		Points: []*qdrant.PointStruct{
			{
				Id:      qdrant.NewIDNum(uint64(promptModel.ProjectId)),
				Vectors: qdrant.NewVectors(vector...),
				Payload: qdrant.NewValueMap(map[string]any{"text": promptModel.Prompt}),
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}

func (repo *VectorRepo) FullResetDatabase() error {
	return repo.client.DeleteCollection(context.Background(), collectionName)
}

func (repo *VectorRepo) ResetProject() error {
	return nil
}

func (repo *VectorRepo) ResetPersonal() error {
	return nil
}
