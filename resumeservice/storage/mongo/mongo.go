package mongo

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/laidingqing/gojobs/resumeservice/model"
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var databaseName = "gojobs"
var resumeCollection = "resumes"

//New mongo storage implemention.
func New() *Storage {
	return &Storage{}
}

//Storage ...
type Storage struct {
	session *mgo.Session
}

//OpenSession open mongo client session.
func (s *Storage) OpenSession() {
	var err error
	s.session, err = mgo.Dial(fmt.Sprintf("%s:%s", viper.GetString("mongo_host"), viper.GetString("mongo_port")))
	if err != nil {
		logrus.Fatal(err)
	}
}

//CopySession copy open session.
func (s *Storage) CopySession() *mgo.Session {
	return s.session.Copy()
}

//Seed init some data.
func (s *Storage) Seed() {

}

//Check is a naive healthcheck, just make sure a db connection has been initialized.
func (s *Storage) Check() bool {
	return s.session != nil
}

//CreateResume create resume on storage
func (s *Storage) CreateResume(resume model.Resume) (string, error) {
	copySession := s.CopySession()
	defer copySession.Close()
	query := copySession.DB(databaseName).C(resumeCollection)
	resume.ID = bson.NewObjectId().Hex()
	err := query.Insert(resume)
	if err != nil {
		return "", err
	}
	return resume.ID, nil
}

//FindResume find a resume
func (s *Storage) FindResume(id string) (model.Resume, error) {
	copySession := s.CopySession()
	defer copySession.Close()
	query := copySession.DB(databaseName).C(resumeCollection)
	var resume model.Resume
	err := query.FindId(id).One(&resume)
	if err != nil {
		return model.Resume{}, err
	}
	return resume, nil
}

//updateResume ...
func (s *Storage) updateResume(resume model.Resume) error {
	copySession := s.CopySession()
	defer copySession.Close()
	query := copySession.DB(databaseName).C(resumeCollection)
	err := query.UpdateId(resume.ID, resume)
	return err
}

//UpdateBaseResume update base resume.
func (s *Storage) UpdateBaseResume(name string, bio string, intentions []string, id string) error {
	resume, err := s.FindResume(id)
	if err != nil {
		return err
	}
	resume.Name = name
	resume.Bio = bio
	resume.Intentions = intentions
	err = s.updateResume(resume)
	return err
}

//UpdateWorkResume ...
func (s *Storage) UpdateWorkResume(work model.Work, id string) error {
	resume, err := s.FindResume(id)
	if err != nil {
		return err
	}

	for i := range resume.Works {
		if resume.Works[i].ID == work.ID {
			resume.Works[i] = work
		}
	}
	err = s.updateResume(resume)
	return err
}

//UpdateProjectResume ...
func (s *Storage) UpdateProjectResume(project model.Project, id string) error {
	resume, err := s.FindResume(id)
	if err != nil {
		return err
	}

	for i := range resume.Projects {
		if resume.Projects[i].ID == project.ID {
			resume.Projects[i] = project
		}
	}
	err = s.updateResume(resume)
	return err
}

//UpdateEducationResume ...
func (s *Storage) UpdateEducationResume(edu model.Education, id string) error {
	resume, err := s.FindResume(id)
	if err != nil {
		return err
	}

	for i := range resume.Educations {
		if resume.Educations[i].ID == edu.ID {
			resume.Educations[i] = edu
		}
	}
	err = s.updateResume(resume)
	return err
}
