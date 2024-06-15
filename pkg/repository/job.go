package repository

import (
	"encoding/json"
	"strings"

	requestmodels "github.com/Vajid-Hussain/machine-test/pkg/models/requestModels"
	responsemodels "github.com/Vajid-Hussain/machine-test/pkg/models/responseModels"
	interfaceRepository "github.com/Vajid-Hussain/machine-test/pkg/repository/interface"
	"gorm.io/gorm"
)

type jobRepository struct {
	DB *gorm.DB
}

func NewJobRepository(db *gorm.DB) interfaceRepository.IJobRepository {
	return &jobRepository{DB: db}
}

func (d *jobRepository) InsertResumereq(req *requestmodels.UserResumeData) (*responsemodels.UserResumeData, error) {
	// Convert skills data into array format
	skillsArray := "{" + strings.Join(req.Skills, ",") + "}"

	educationJSON, err := json.Marshal(req.Education)
	if err != nil {
		return nil, err
	}

	// Define the SQL query
	query := `
        INSERT INTO user_resume_data (name, email, phone, skills, education, experience, user_id, resume)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?)
        ON CONFLICT (user_id) DO UPDATE
        SET name = EXCLUDED.name, email = EXCLUDED.email, phone = EXCLUDED.phone,
            skills = EXCLUDED.skills, education = EXCLUDED.education, experience = EXCLUDED.experience, resume = EXCLUDED.resume
        RETURNING id, name, email, phone, skills, education, experience, user_id, resume
    `

	// Execute the raw SQL query
	var insertedData responsemodels.UserResumeData
	result := d.DB.Raw(query, req.Name, req.Email, req.Phone, skillsArray, educationJSON, req.Experience, req.UserID, req.Resume).Scan(&insertedData)
	if result.Error != nil {
		return nil, result.Error
	}

	return &insertedData, nil
}

func (d *jobRepository) ApplyJob(req *requestmodels.JobApplication) (*responsemodels.JobApplication, error) {
	var res responsemodels.JobApplication
	query := `
			INSERT INTO job_applies (user_id, job_id, apply_time) SELECT $1, $2, now() 
			WHERE NOT EXISTS (SELECT 1 FROM job_applies WHERE user_id = $1 AND job_id = $2)  AND EXISTS (SELECT 1 FROM jobs WHERE id = $2) RETURNING *
			`
	result := d.DB.Raw(query, req.UserID, req.JobApplicationID).Scan(&res)
	if result.Error != nil {
		return nil, responsemodels.ErrInternalServer
	}
	if result.RowsAffected == 0 {
		return nil, responsemodels.ErrJobAlreadyApplied
	}

	return &res, nil
}

func (d *jobRepository) GetAppliedJob(job *requestmodels.GetAppliedJob, pagination *requestmodels.Pagination) (*[]responsemodels.JobApplication, error) {
	var res []responsemodels.JobApplication
	query := `
			SELECT * FROM job_applies ja INNER JOIN users u ON u.id = ja.user_id INNER JOIN jobs j ON j.id = ja.job_id 
			WHERE u.id= $1 AND u.status= 'active' AND j.status='active' AND j.title ILIKE '%' || $2 || '%' 
			OFFSET $3 FETCH FIRST $4 ROWS ONLY
			`
	result := d.DB.Raw(query, job.UserID, job.Job, pagination.Offset, pagination.Limit).Scan(&res)
	if result.Error != nil {
		return nil, responsemodels.ErrInternalServer
	}

	return &res, nil
}

// ------------------------------------- Admin Job ----------------------------------------------

func (d *jobRepository) CreateJob(job *requestmodels.CreateJob) (*responsemodels.Job, error) {
	var res responsemodels.Job
	query := `INSERT INTO jobs (title, description, post_on, company_name, posted_by) VALUES($1, $2, now(), $3, $4) RETURNING *`
	result := d.DB.Raw(query, job.Title, job.Description, job.CompanyName, job.PostedBy).Scan(&res)
	if result.Error != nil {
		return nil, responsemodels.ErrInternalServer
	}

	return &res, nil
}

func (d *jobRepository) DeleteJob(req *requestmodels.DeleteJob) error {

	query := `DELETE FROM jobs WHERE id= $1 AND status= 'active'`
	result := d.DB.Exec(query, req.JobID)
	if result.Error != nil {
		return responsemodels.ErrInternalServer
	}
	if result.RowsAffected == 0 {
		return responsemodels.ErrNoActiveJob
	}

	return nil
}

func (d *jobRepository) GetJob(prefix *requestmodels.JobSearch, pagination *requestmodels.Pagination) (*[]responsemodels.Job, error) {
	var res []responsemodels.Job

	query := `SELECT * FROM jobs WHERE status='active' AND title ILIKE '%' || $1 || '%' OFFSET $2 FETCH FIRST $3 ROWS ONLY`
	result := d.DB.Raw(query, prefix.Job, pagination.Offset, pagination.Limit).Scan(&res)
	if result.Error != nil {
		return nil, responsemodels.ErrInternalServer
	}

	return &res, nil
}

func (d *jobRepository) GetJobDetails(job *requestmodels.JobID, pagination *requestmodels.Pagination) (*[]responsemodels.JobApplicationAdmin, error) {
	var res []responsemodels.JobApplicationAdmin
	query := `
			SELECT * FROM job_applies ja INNER JOIN users u ON u.id = ja.user_id INNER JOIN jobs j ON j.id = ja.job_id 
			WHERE j.id= $1 AND j.status='active'
			OFFSET $2 FETCH FIRST $3 ROWS ONLY
			`
	result := d.DB.Raw(query, job.JobID, pagination.Offset, pagination.Limit).Scan(&res)
	if result.Error != nil {
		return nil, responsemodels.ErrInternalServer
	}

	return &res, nil
}
