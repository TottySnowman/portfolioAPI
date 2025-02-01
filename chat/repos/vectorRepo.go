package vectorRepo

import (
	"context"
	"encoding/json"
	"fmt"
	chatModel "portfolioAPI/chat/models"
	projectService "portfolioAPI/project/services"
	"strconv"

	"github.com/google/uuid"
	"github.com/qdrant/go-client/qdrant"
)

type VectorRepo struct {
	client         *qdrant.Client
	projectService *projectService.ProjectService
}

const collectionName = "portfolio"
const vectorSize = uint64(384)

func NewVectorRepo(projectService *projectService.ProjectService) *VectorRepo {
	client, err := qdrant.NewClient(&qdrant.Config{
		Host: "qdrant",
		Port: 6334,
	})

	if err != nil {
		panic(err.Error())
	}

	return &VectorRepo{
		client:         client,
		projectService: projectService,
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

func (repo *VectorRepo) UpsertText(vector chatModel.FeatureExtractionResponse, text string, textId string) (*qdrant.ScoredPoint, error) {
	textPoint := qdrant.NewID(uuid.NewString())

	if textId != "" {
		existingPoint := repo.GetExistingPoint(textId).Id
		if existingPoint != nil {
			textPoint = qdrant.NewIDUUID(textId)
		}
	}

	convertedText := map[string]interface{}{
		"text": text,
	}

	if err := repo.upsertVector(textPoint, vector, convertedText); err != nil {
		return nil, err
	}

	point := repo.GetExistingPoint(textPoint.GetUuid())

	return point, nil
}

func (repo *VectorRepo) GetExistingPoint(pointId string) *qdrant.ScoredPoint {
	point, err := repo.client.Query(context.Background(), &qdrant.QueryPoints{
		CollectionName: collectionName,
		Query:          qdrant.NewQueryID(qdrant.NewID(pointId)),
		WithPayload:    qdrant.NewWithPayload(true),
	})

	if err != nil {
		panic(err)
	}
	if point != nil && len(point) > 0 {
		return point[0]
	}
	return nil
}

func (repo *VectorRepo) UpsertProject(modifyProjectModel chatModel.ModifyProjectModel) error {
	projectPointId := repo.getExistingProjectPoint(modifyProjectModel.ProjectPayload.ProjectID)

	if projectPointId == nil {
		projectPointId = qdrant.NewID(uuid.NewString())
	}

	convertedProject := map[string]interface{}{
		"project_id": strconv.Itoa(modifyProjectModel.ProjectPayload.ProjectID),
	}

	err := repo.upsertVector(projectPointId, modifyProjectModel.Vector, convertedProject)
	if err != nil {
		return err
	}

	return nil
}

func (repo *VectorRepo) DeleteProjectPoint(projectId int) error {
	projectPointId := repo.getExistingProjectPoint(projectId)
	if projectPointId == nil {
		fmt.Println("No point found for the given project ID.")
		return nil
	}

	repo.client.Delete(context.Background(), &qdrant.DeletePoints{
		CollectionName: collectionName,
		Points: qdrant.NewPointsSelector(
			qdrant.NewID(projectPointId.GetUuid())),
	})

	fmt.Println("Point successfully deleted.")
	return nil
}

func (repo *VectorRepo) getExistingProjectPoint(projectId int) *qdrant.PointId {
	filter := &qdrant.Filter{
		Must: []*qdrant.Condition{
			qdrant.NewMatch("project_id", strconv.Itoa(projectId)),
		},
	}

	point, err := repo.client.Scroll(context.Background(), &qdrant.ScrollPoints{
		CollectionName: collectionName,
		Filter:         filter,
		WithPayload:    qdrant.NewWithPayload(true),
	})

	if err != nil {
		panic(err)
	}
	if point != nil && len(point) > 0 {
		return point[0].Id
	}
	return nil
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

func (repo *VectorRepo) SearchSimilarity(vector chatModel.FeatureExtractionResponse) ([]string, error) {
	response := make([]string, 0)
	similarVectors, err := repo.getSimilarVectors(vector)
	if err != nil {
		return response, err
	}

	response = repo.convertFoundVectorsToPromptableResponse(similarVectors)

	return response, nil
}
func (repo *VectorRepo) getSimilarVectors(vector chatModel.FeatureExtractionResponse) ([]*qdrant.ScoredPoint, error) {
	var limit uint64 = 3

	return repo.client.Query(context.Background(), &qdrant.QueryPoints{
		CollectionName: collectionName,
		Query:          qdrant.NewQuery(vector...),
		WithPayload:    qdrant.NewWithPayload(true),
		Limit:          &limit,
	})
}

func (repo *VectorRepo) convertFoundVectorsToPromptableResponse(foundVectors []*qdrant.ScoredPoint) []string {
	response := make([]string, 0)
	for _, result := range foundVectors {
		payload := result.GetPayload()

		if projectId, ok := payload["project_id"]; ok {
			convertedProjectId, _ := strconv.Atoi(projectId.GetStringValue())
			project := repo.getStringyfiedProject(int64(convertedProjectId))

			response = append(response, project)
		} else {
			qdrantText, _ := payload["text"]
			response = append(response, qdrantText.GetStringValue())
		}
	}

	return response
}

func (repo *VectorRepo) getStringyfiedProject(projectId int64) string {
	projectDetails, err := repo.projectService.GetProjectById(int(projectId), false)
	if err != nil {
		panic(err.Error())
	}

	jsonProject, err := json.Marshal(projectDetails)
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)

	}

	stringyfiedProject := string(jsonProject)

	return stringyfiedProject
}

func (repo *VectorRepo) DeleteSinglePoint(pointId string) error {
	_, err := repo.client.Delete(context.Background(), &qdrant.DeletePoints{
		CollectionName: collectionName,
		Points: qdrant.NewPointsSelector(
			qdrant.NewIDUUID(pointId),
		),
	})

	return err
}

func (repo *VectorRepo) GetFullKnowledgeBase() []*qdrant.RetrievedPoint {
	filter := &qdrant.Filter{
		Must: []*qdrant.Condition{
			qdrant.NewIsEmpty("project_id"),
		},
	}

	points, err := repo.client.Scroll(context.Background(), &qdrant.ScrollPoints{
		CollectionName: collectionName,
		Filter:         filter,
		WithPayload:    qdrant.NewWithPayload(true),
	})

	if err != nil {
		panic(err)
	}
	return points
}
func (repo *VectorRepo) FullResetDatabase() error {
	return repo.client.DeleteCollection(context.Background(), collectionName)
}

func (repo *VectorRepo) ResetProject() error {
	filter := &qdrant.Filter{
		MustNot: []*qdrant.Condition{
			qdrant.NewIsEmpty("project_id"),
		},
	}
	_, err := repo.client.Delete(context.Background(), &qdrant.DeletePoints{
		CollectionName: collectionName,
		Points: qdrant.NewPointsSelectorFilter(
      filter,
		),
	})
	return err
}

func (repo *VectorRepo) ResetPersonal() error {
	filter := &qdrant.Filter{
		Must: []*qdrant.Condition{
			qdrant.NewIsEmpty("project_id"),
		},
	}
	_, err := repo.client.Delete(context.Background(), &qdrant.DeletePoints{
		CollectionName: collectionName,
		Points: qdrant.NewPointsSelectorFilter(
      filter,
		),
	})
	return err
}
