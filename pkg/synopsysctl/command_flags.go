// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

package synopsysctl

import (
	alertv1 "github.com/blackducksoftware/synopsys-operator/pkg/api/alert/v1"
	blackduckv1 "github.com/blackducksoftware/synopsys-operator/pkg/api/blackduck/v1"
	opssightv1 "github.com/blackducksoftware/synopsys-operator/pkg/api/opssight/v1"
)

// synopsysctl Defaults
var namespace = ""

// Start Command Defaults
var startNamespace = "synopsys-operator"
var startSynopsysOperatorImage = "docker.io/blackducksoftware/synopsys-operator:2019.2.0-RC"
var startPrometheusImage = "docker.io/prom/prometheus:v2.1.0"
var startBlackduckRegistrationKey = ""
var startDockerConfigPath = ""
var startSecretName = "blackduck-secret"
var startSecretType = "Opaque"
var startSecretAdminPassword = "YmxhY2tkdWNr"
var startSecretPostgresPassword = "YmxhY2tkdWNr"
var startSecretUserPassword = "YmxhY2tkdWNr"
var startSecretBlackduckPassword = "YmxhY2tkdWNr"

// Stop Command Defaults
var stopNamespace = "synopsys-operator"

// Create Blackduck Command Defaults
var defaultBlackduckSpec = &blackduckv1.BlackduckSpec{}
var defaultBlackduckName = "blackduck"

var createBlackduckSize = ""
var createBlackduckDbPrototype = ""
var createBlackduckExternalPostgresPostgresHost = ""
var createBlackduckExternalPostgresPostgresPort = 0
var createBlackduckExternalPostgresPostgresAdmin = ""
var createBlackduckExternalPostgresPostgresUser = ""
var createBlackduckExternalPostgresPostgresSsl = false
var createBlackduckExternalPostgresPostgresAdminPassword = ""
var createBlackduckExternalPostgresPostgresUserPassword = ""
var createBlackduckPvcStorageClass = ""
var createBlackduckLivenessProbes = false
var createBlackduckScanType = ""
var createBlackduckPersistentStorage = false
var createBlackduckPVCJSONSlice = []string{}
var createBlackduckCertificateName = ""
var createBlackduckCertificate = ""
var createBlackduckCertificateKey = ""
var createBlackduckProxyCertificate = ""
var createBlackduckType = ""
var createBlackduckDesiredState = ""
var createBlackduckEnvirons = []string{}
var createBlackduckImageRegistries = []string{}
var createBlackduckImageUIDMapJSONSlice = []string{}
var createBlackduckLicenseKey = ""

// Create OpsSight Command Defaults
var defaultOpsSightSpec = &opssightv1.OpsSightSpec{}
var defaultOpsSightName = "opssight"

var createOpssightPerceptorName = ""
var createOpssightPerceptorImage = ""
var createOpssightPerceptorPort = 0
var createOpssightPerceptorCheckForStalledScansPauseHours = 0
var createOpssightPerceptorStalledScanClientTimeoutHours = 0
var createOpssightPerceptorModelMetricsPauseSeconds = 0
var createOpssightPerceptorUnknownImagePauseMilliseconds = 0
var createOpssightPerceptorClientTimeoutMilliseconds = 0
var createOpssightScannerPodName = ""
var createOpssightScannerPodScannerName = ""
var createOpssightScannerPodScannerImage = ""
var createOpssightScannerPodScannerPort = 0
var createOpssightScannerPodScannerClientTimeoutSeconds = 0
var createOpssightScannerPodImageFacadeName = ""
var createOpssightScannerPodImageFacadeImage = ""
var createOpssightScannerPodImageFacadePort = 0
var createOpssightScannerPodImageFacadeInternalRegistriesJSONSlice = []string{}
var createOpssightScannerPodImageFacadeImagePullerType = ""
var createOpssightScannerPodImageFacadeServiceAccount = ""
var createOpssightScannerPodReplicaCount = 0
var createOpssightScannerPodImageDirectory = ""
var createOpssightPerceiverEnableImagePerceiver = false
var createOpssightPerceiverEnablePodPerceiver = false
var createOpssightPerceiverImagePerceiverName = ""
var createOpssightPerceiverImagePerceiverImage = ""
var createOpssightPerceiverPodPerceiverName = ""
var createOpssightPerceiverPodPerceiverImage = ""
var createOpssightPerceiverPodPerceiverNamespaceFilter = ""
var createOpssightPerceiverAnnotationIntervalSeconds = 0
var createOpssightPerceiverDumpIntervalMinutes = 0
var createOpssightPerceiverServiceAccount = ""
var createOpssightPerceiverPort = 0
var createOpssightPrometheusName = ""
var createOpssightPrometheusImage = ""
var createOpssightPrometheusPort = 0
var createOpssightEnableSkyfire = false
var createOpssightSkyfireName = ""
var createOpssightSkyfireImage = ""
var createOpssightSkyfirePort = 0
var createOpssightSkyfirePrometheusPort = 0
var createOpssightSkyfireServiceAccount = ""
var createOpssightSkyfireHubClientTimeoutSeconds = 0
var createOpssightSkyfireHubDumpPauseSeconds = 0
var createOpssightSkyfireKubeDumpIntervalSeconds = 0
var createOpssightSkyfirePerceptorDumpIntervalSeconds = 0
var createOpssightBlackduckHosts = []string{}
var createOpssightBlackduckUser = ""
var createOpssightBlackduckPort = 0
var createOpssightBlackduckConcurrentScanLimit = 0
var createOpssightBlackduckTotalScanLimit = 0
var createOpssightBlackduckPasswordEnvVar = ""
var createOpssightBlackduckInitialCount = 0
var createOpssightBlackduckMaxCount = 0
var createOpssightBlackduckDeleteHubThresholdPercentage = 0
var createOpssightEnableMetrics = false
var createOpssightDefaultCPU = ""
var createOpssightDefaultMem = ""
var createOpssightLogLevel = ""
var createOpssightConfigMapName = ""
var createOpssightSecretName = ""

// Create Alert Command Defaults
var defaultAlertSpec = &alertv1.AlertSpec{}
var defaultAlertName = "alert"

var createAlertRegistry = ""
var createAlertImagePath = ""
var createAlertAlertImageName = ""
var createAlertAlertImageVersion = ""
var createAlertCfsslImageName = ""
var createAlertCfsslImageVersion = ""
var createAlertBlackduckHost = ""
var createAlertBlackduckUser = ""
var createAlertBlackduckPort = 0
var createAlertPort = 0
var createAlertStandAlone = false
var createAlertAlertMemory = ""
var createAlertCfsslMemory = ""
var createAlertState = ""