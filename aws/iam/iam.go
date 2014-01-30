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
	resp, err := http.Get(IamUrl)
	util.LogPanic(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	util.LogPanic(err)
	var roleFields RolesFields
	json.Unmarshal(body, &roleFields)
	//util.LogPanic(err)
	log.Debug(roleFields)
	auth := aws.Auth{roleFields.AccessKeyId, roleFields.SecretAccessKey}
	log.Info("IAM role credentials", auth)
	return auth
}
