package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Vajid-Hussain/machine-test/pkg/intenals/config"
	requestmodels "github.com/Vajid-Hussain/machine-test/pkg/models/requestModels"
	responsemodels "github.com/Vajid-Hussain/machine-test/pkg/models/responseModels"
	interfaceRepository "github.com/Vajid-Hussain/machine-test/pkg/repository/interface"
	interfaceUseCase "github.com/Vajid-Hussain/machine-test/pkg/usecase/interface"
	"github.com/Vajid-Hussain/machine-test/pkg/utils"
)

type jobUseCase struct {
	repo interfaceRepository.IJobRepository
	s3   config.S3Bucket
}

func NewJobUseCase(repository interfaceRepository.IJobRepository, s3 config.S3Bucket) interfaceUseCase.IJobUseCase {
	return &jobUseCase{repo: repository, s3: s3}
}

func (d jobUseCase) DecodeResume(reqResume requestmodels.Resume) (*responsemodels.UserResumeData, error) {
	var (
		resumeData requestmodels.UserResumeData
		typeDocs   = "application/docs"
		typepdf    = "application/pdf"
	)

	// Check the content type
	contentType := reqResume.Resume.Header.Get("Content-Type")
	if contentType != typeDocs && contentType != typepdf {
		return nil, responsemodels.ErrWrongContecntTypeResume
	}

	fmt.Println("reqResume.Resume.Header ", reqResume.Resume.Header)
	file, err := reqResume.Resume.Open()
	if err != nil {
		return nil, err
	}

	// Request third party api for decode resume details
	clind := http.Client{}

	req, err := http.NewRequest("POST", "https://api.apilayer.com/resume_parser/upload", file)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/octet-stream")
	req.Header.Set("apikey", "gNiXyflsFu3WNYCz1ZCxdWDb7oQg1Nl1")

	result, err := clind.Do(req)
	if err != nil {
		return nil, err
	}

	fmt.Println("result ", result.Status, req.Header)

	body, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(body, &resumeData)
	resumeData.UserID = reqResume.UserID

	// Seek back to the beginning of the file
	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	// fetch the file and create byte veriable for s3
	byteFile, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// resume upload to s3
	s3Session := utils.CreateSession(d.s3)
	url, err := utils.UploadImageToS3(byteFile, s3Session)
	if err != nil {
		return nil, err
	}

	resumeData.Resume = url
	return d.repo.InsertResumereq(&resumeData)
}

func (d *jobUseCase) CreateJob(job *requestmodels.CreateJob) (*responsemodels.Job, error) {
	return d.repo.CreateJob(job)
}

func (d *jobUseCase) GetJob(prefix *requestmodels.JobSearch, pagination *requestmodels.Pagination) (res *[]responsemodels.Job, err error) {
	pagination.Offset, err = utils.Pagination(pagination.Limit, pagination.Offset)
	return d.repo.GetJob(prefix, pagination)
}

func (d *jobUseCase) DeleteJob(req *requestmodels.DeleteJob) error {
	return d.repo.DeleteJob(req)
}
