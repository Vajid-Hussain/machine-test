package interfaceUseCase

import (
	requestmodels "github.com/Vajid-Hussain/machine-test/pkg/models/requestModels"
	responsemodels "github.com/Vajid-Hussain/machine-test/pkg/models/responseModels"
)

type IJobUseCase interface {
	DecodeResume(reqResume requestmodels.Resume) (*responsemodels.UserResumeData, error)
	CreateJob(*requestmodels.CreateJob) (*responsemodels.Job, error)
	DeleteJob(*requestmodels.DeleteJob) error
	GetJob(*requestmodels.JobSearch, *requestmodels.Pagination) (*[]responsemodels.Job, error)
}
