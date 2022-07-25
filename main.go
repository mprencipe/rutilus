package main

import (
	"context"
	"fmt"

	log "github.com/sirupsen/logrus"

	"rutilus/internal/check"
	root_account_not_used "rutilus/internal/check/iam/1.1_root_account_not_used"
	user_mfa "rutilus/internal/check/iam/1.2_user_mfa"
	unused_accounts_disabled "rutilus/internal/check/iam/1.3_unused_accounts_disabled"
	rotated_access_keys "rutilus/internal/check/iam/1.4_rotated_access_keys"
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
	&unused_accounts_disabled.UnusedAccountsDisabled{},
	&rotated_access_keys.RotatedAccessKeys{},
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
