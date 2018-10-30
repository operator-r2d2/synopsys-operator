package protoform-installer

// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This command encodes all the bootstrap components into the cluster.
// It should be run once : When the user first installs the blackduck operator.

import (
	"fmt"

	bdutil "github.com/blackducksoftware/perceptor-protoform/cmd/protoform-installer/blackduckctl/pkg/util"

	"github.com/blackducksoftware/perceptor-protoform/pkg/alert"
	"github.com/blackducksoftware/perceptor-protoform/pkg/hub"
	"github.com/sirupsen/logrus"

	horizondep "github.com/blackducksoftware/horizon/pkg/deployer"
	"github.com/blackducksoftware/perceptor-protoform/pkg/opssight"
	"github.com/blackducksoftware/perceptor-protoform/pkg/protoform"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
)

func StartBlackduckOperator(configPath string) {
	deployer, err := protoform.NewController(configPath)
	if err != nil {
		panic(err)
	}

	stopCh := make(chan struct{})

	alertController := alert.NewCRDInstaller(deployer.Config, deployer.KubeConfig, deployer.KubeClientSet, bdutil.GetAlertDefaultValue(), stopCh)
	deployer.AddController(alertController)

	hubController := hub.NewCRDInstaller(deployer.Config, deployer.KubeConfig, deployer.KubeClientSet, bdutil.GetHubDefaultValue(), stopCh)
	deployer.AddController(hubController)

	opssSightController, err := opssight.NewCRDInstaller(&opssight.Config{
		Config:        deployer.Config,
		KubeConfig:    deployer.KubeConfig,
		KubeClientSet: deployer.KubeClientSet,
		Defaults:      bdutil.GetOpsSightDefaultValue(),
		Threadiness:   deployer.Config.Threadiness,
		StopCh:        stopCh,
	})
	if err != nil {
		panic(err)
	}
	deployer.AddController(opssSightController)

	logrus.Info("Starting deployer.  All controllers have been added to horizon.")
	if err = deployer.Deploy(); err != nil {
		logrus.Errorf("ran into errors during deployment, but continuing anyway: %s", err.Error())
	}

	<-stopCh
}
