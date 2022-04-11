package main

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"rutilus/internal/check"
	"rutilus/internal/check/iam/root_account_not_used"
	"rutilus/internal/check/iam/user_mfa"
	rutilusConfig "rutilus/internal/config"
	"rutilus/internal/util"

	"github.com/aws/aws-sdk-go-v2/config"
)

const BANNER = `
               |   _)  |              
   __|  |   |  __|  |  |  |   |   __| 
  |     |   |  |    |  |  |   | \__ \ 
 _|    \__,_| \__| _| _| \__,_| ____/ 
                                      `

var checks = []check.Check{
	&root_account_not_used.RootAccountNotUsed{},
	&user_mfa.UsersWithPasswordsHaveMFA{},
}

func main() {
	fmt.Println(BANNER)

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	rutilusConfig.Config = cfg

	util.GenerateCredentialReport()

	for _, c := range checks {
		log.Println(c.Describe())
		checkResult, err := c.Check()
		if err != nil {
			log.Warn("Error in test")
		}
		if checkResult == check.Success {
			log.Println("Success")
		} else if checkResult == check.Failure {
			log.Println("Failure")
		} else if checkResult == check.Warning {
			log.Println("Warning")
		}
	}
}
