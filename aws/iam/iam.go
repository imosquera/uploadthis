package iam

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"launchpad.net/goamz/aws"
	"github.com/imosquera/uploadthis/util"
	log "github.com/cihub/seelog"
)

type RolesFields struct {
	Code string
	LastUpdated string
	Type string
	AccessKeyId string
	SecretAccessKey string
	Token string
	Expiration string
}

func NewRolesFields(IamUrl string) (aws.Auth) {
	log.Info(IamUrl)
	resp, err := http.Get(IamUrl)
	util.LogPanic(err)
	defer resp.Body.Close()
	log.Info(resp)
	body, err := ioutil.ReadAll(resp.Body)
	util.LogPanic(err)
	log.Info(body)
	var roleFields RolesFields
	json.Unmarshal(body, &roleFields)
	//util.LogPanic(err)
	log.Info(roleFields)
	auth := aws.Auth{roleFields.AccessKeyId, roleFields.SecretAccessKey}
	log.Info(auth)
	return auth
}
