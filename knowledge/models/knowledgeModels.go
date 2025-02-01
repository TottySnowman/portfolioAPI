package knowledgeModels
type DeleteModel struct{
  PointId string `json:"pointId"`
}

type KnowledgeDisplayModel struct{
  PointId string `json:"pointId"`
  Text string `json:"Prompt"`
}

type SyncModel struct {
	ResetProject  bool `json:"resetProject"`
	ResetPersonal bool `json:"resetPersonal"`
}
