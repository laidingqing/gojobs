package mongo

import (
	"context"

	"github.com/laidingqing/gojobs/accountservice/conf"

	"github.com/Sirupsen/logrus"
	"github.com/laidingqing/gojobs/accountservice/model"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var databaseName = "gojobs"
var accountCollection = "accounts"

//New mongo storage implemention.
func New() *Storage {
	return &Storage{}
}

//Storage ...
type Storage struct {
	session *mgo.Session
}

//OpenSession open mongo client session.
func (mc *Storage) OpenSession() {
	var err error
	mc.session, err = mgo.Dial(conf.AccountConf.Core.MongoURI)
	if err != nil {
		logrus.Fatal(err)
	}
}

//CopySession copy open session.
func (mc *Storage) CopySession() *mgo.Session {
	return mc.session.Copy()
}

//Seed init some data.
func (mc *Storage) Seed() {

}

//Check is a naive healthcheck, just make sure a db connection has been initialized.
func (mc *Storage) Check() bool {
	return mc.session != nil
}

//CreateAccount create acct on storage
func (mc *Storage) CreateAccount(account model.Account) (string, error) {
	copySession := mc.CopySession()
	defer copySession.Close()
	query := copySession.DB(databaseName).C(accountCollection)
	account.ID = bson.NewObjectId().Hex()
	err := query.Insert(account)
	if err != nil {
		return "", err
	}
	return account.ID, nil
}

//QueryAccount lets us query an account object.
func (mc *Storage) QueryAccount(ctx context.Context, accountID string) (model.Account, error) {
	account := model.Account{}
	copySession := mc.CopySession()
	defer copySession.Close()
	query := copySession.DB(databaseName).C(accountCollection)
	err := query.Find(bson.M{"_id": accountID}).One(&account)
	if err != nil {
		return model.Account{}, err
	}
	return account, nil
}
