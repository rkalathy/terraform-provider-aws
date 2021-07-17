package waiter

// https://github.com/hashicorp/terraform-provider-aws/pull/20172#issuecomment-881017208

import (
	"time"

	"github.com/aws/aws-sdk-go/service/appconfig"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

const DeploymentCreatedTimeout = 5 * time.Minute

func DeploymentCreated(conn *appconfig.AppConfig, appID, envID string, deployNum int64) error {
	stateConf := &resource.StateChangeConf{
		Pending: []string{appconfig.DeploymentStateBaking, appconfig.DeploymentStateRollingBack, appconfig.DeploymentStateValidating, appconfig.DeploymentStateDeploying},
		Target:  []string{appconfig.DeploymentStateComplete},
		Refresh: DeploymentStatus(conn, appID, envID, deployNum),
		Timeout: DeploymentCreatedTimeout,
	}

	_, err := stateConf.WaitForState()

	return err
}
