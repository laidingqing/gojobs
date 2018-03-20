package storage

import (
	"github.com/laidingqing/gojobs/resumeservice/model"
)

//Storage interface.
type Storage interface {
	OpenSession()
	Seed()
	Check() bool
	FindResume(accountID string) (model.Resume, error)
	CreateResume(r model.Resume) (string, error)
	UpdateBaseResume(name string, bio string, intentions []string, id string) error
	UpdateWorkResume(work model.Work, id string) error
	UpdateProjectResume(work model.Project, id string) error
	UpdateEducationResume(work model.Education, id string) error
}
