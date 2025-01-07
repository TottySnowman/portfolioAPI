package chatModel

import projectModel "portfolioAPI/project/models"

type PromptModel struct {
	Prompt    string
	ProjectId int
}

type ModifyProjectModel struct {
  ProjectPayload projectModel.ProjectDisplay
  ProjectTags []string
  Vector FeatureExtractionResponse
}

type ModifyPersonalModel struct {
  Text string
  Vector FeatureExtractionResponse
  PointId *uint64
}

type SyncModel struct {
	ResetProject  bool
	ResetPersonal bool
}
