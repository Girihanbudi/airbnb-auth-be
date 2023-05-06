package env

import (
	"golang.org/x/exp/slices"
)

type Stage string

const (
	Prefix string = "STAGE"

	StageLocal       Stage = "local"
	StageDevelopment Stage = "development"
	StageStaging     Stage = "staging"
	StageProduction  Stage = "production"
)

var Stages = []Stage{
	StageLocal,
	StageDevelopment,
	StageStaging,
	StageProduction,
}

var (
	stageMap = map[string]Stage{
		string(StageLocal):       StageLocal,
		string(StageDevelopment): StageDevelopment,
		string(StageStaging):     StageStaging,
		string(StageProduction):  StageProduction,
	}
)

func ParseString(str string) Stage {
	stage, exist := stageMap[str]
	if !exist {
		stage = StageLocal
	}

	return stage
}

func (s Stage) ValidStage() bool {
	return slices.Contains(Stages, s)
}
