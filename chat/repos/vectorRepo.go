package vectorRepo

import (
	"context"
	chatModel "portfolioAPI/chat/models"
	projectModel "portfolioAPI/project/models"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/qdrant/go-client/qdrant"
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

func (repo *VectorRepo) UpsertProject(modifyProjectModel chatModel.ModifyProjectModel) error {
	projectPointId := repo.getExistingProjectPoint(modifyProjectModel.ProjectPayload.ProjectID)

	if projectPointId == nil {
		projectPointId = qdrant.NewID(uuid.NewString())
	}

	convertedProject := convertProjectDisplayToPayload(modifyProjectModel.ProjectPayload, modifyProjectModel.ProjectTags)

	err := repo.upsertVector(projectPointId, modifyProjectModel.Vector, convertedProject)
	if err != nil {
		return err
	}

	return nil
}

func (repo *VectorRepo) getExistingProjectPoint(projectId int) *qdrant.PointId {
	filter := &qdrant.Filter{
		Must: []*qdrant.Condition{
			qdrant.NewMatch("projectId", strconv.Itoa(projectId)),
		},
	}

	point, err := repo.client.Scroll(context.Background(), &qdrant.ScrollPoints{
		CollectionName: collectionName,
		Filter:         filter,
	})

	if err != nil {
		panic(err)
	}
	if point != nil && len(point) > 0 {

		return point[0].Id
	}
	return nil
}

func convertProjectDisplayToPayload(projectDisplay projectModel.ProjectDisplay, tags []string) map[string]interface{} {
	return map[string]interface{}{
		"project_id":  projectDisplay.ProjectID,
		"status":      projectDisplay.Status.Status,
		"name":        projectDisplay.Name,
		"about":       projectDisplay.About,
		"github_link": projectDisplay.Github_Link,
		"demo_link":   projectDisplay.Demo_Link,
		"logo_path":   projectDisplay.Logo_Path,
		"tags":        strings.Join(tags, ","),
		"dev_date":    projectDisplay.DevDate.Format(time.RFC3339),
		"hidden":      projectDisplay.Hidden,
	}
}

func (repo *VectorRepo) upsertVector(pointId *qdrant.PointId, vector chatModel.FeatureExtractionResponse, payload map[string]interface{}) error {
	_, err := repo.client.Upsert(context.Background(), &qdrant.UpsertPoints{
		CollectionName: collectionName,
		Points: []*qdrant.PointStruct{
			{
				Id:      pointId,
				Vectors: qdrant.NewVectors(vector...),
				Payload: qdrant.NewValueMap(payload),
			},
		},
	})

	return err
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
