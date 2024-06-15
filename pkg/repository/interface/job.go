package interfaceRepository

import (
	requestmodels "github.com/Vajid-Hussain/machine-test/pkg/models/requestModels"
	responsemodels "github.com/Vajid-Hussain/machine-test/pkg/models/responseModels"
)

type IJobRepository interface {
	InsertResumereq(*requestmodels.UserResumeData) (*responsemodels.UserResumeData, error)
	CreateJob(*requestmodels.CreateJob) (*responsemodels.Job, error)
	DeleteJob(*requestmodels.DeleteJob) error
	GetJob(*requestmodels.JobSearch, *requestmodels.Pagination) (*[]responsemodels.Job, error)
}
