package nks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testAwsCluster = Cluster{
	Name:               "Test AWS Cluster Go SDK " + getTicks(),
	Provider:           "aws",
	MasterCount:        1,
	MasterSize:         "t2.medium",
	WorkerCount:        2,
	WorkerSize:         "t2.medium",
	Region:             "us-east-1",
	Zone:               "us-east-1a",
	ProviderNetworkID:  "__new__",
	ProviderNetworkCdr: "172.23.0.0/16",
	ProviderSubnetID:   "__new__",
	ProviderSubnetCidr: "172.23.1.0/24",
	KubernetesVersion:  "v1.13.1",
	RbacEnabled:        true,
	DashboardEnabled:   true,
	EtcdType:           "classic",
	Platform:           "coreos",
	Channel:            "stable",
	Solutions:          []Solution{Solution{Solution: "helm_tiller"}},
}

var testEKSCluster = Cluster{
	Name:               "Test EKS Cluster Go SDK " + getTicks(),
	Provider:           "eks",
	MasterCount:        1,
	MasterSize:         "t2.medium",
	WorkerCount:        2,
	WorkerSize:         "t2.medium",
	Region:             "us-east-1",
	Zone:               "us-east-1a",
	ProviderNetworkID:  "__new__",
	ProviderNetworkCdr: "172.23.0.0/16",
	ProviderSubnetID:   "__new__",
	ProviderSubnetCidr: "172.23.1.0/24",
	KubernetesVersion:  "v1.13.1",
	RbacEnabled:        true,
	DashboardEnabled:   true,
	EtcdType:           "classic",
	Platform:           "coreos",
	Channel:            "stable",
	Solutions:          []Solution{Solution{Solution: "helm_tiller"}},
}

var testAzureCluster = Cluster{
	Name:               "Test Azure Cluster Go SDK " + getTicks(),
	Provider:           "azure",
	MasterCount:        1,
	MasterSize:         "standardA2",
	WorkerCount:        2,
	WorkerSize:         "standardA2",
	Region:             "eastus",
	ProviderResourceGp: "__new__",
	ProviderNetworkID:  "__new__",
	ProviderNetworkCdr: "10.0.0.0/16",
	ProviderSubnetID:   "__new__",
	ProviderSubnetCidr: "10.0.0.0/24",
	KubernetesVersion:  "v1.13.1",
	RbacEnabled:        true,
	DashboardEnabled:   true,
	EtcdType:           "classic",
	Platform:           "coreos",
	Channel:            "stable",
	Solutions:          []Solution{Solution{Solution: "helm_tiller"}},
}

var testAKSCluster = Cluster{
	Name:               "Test AKS Cluster Go SDK " + getTicks(),
	Provider:           "aks",
	MasterCount:        1,
	MasterSize:         "standardA2",
	WorkerCount:        2,
	WorkerSize:         "standardA2",
	Region:             "eastus",
	ProviderResourceGp: "__new__",
	ProviderNetworkID:  "__new__",
	ProviderNetworkCdr: "10.0.0.0/16",
	ProviderSubnetID:   "__new__",
	ProviderSubnetCidr: "10.0.0.0/24",
	KubernetesVersion:  "v1.13.1",
	RbacEnabled:        true,
	DashboardEnabled:   true,
	EtcdType:           "classic",
	Platform:           "coreos",
	Channel:            "stable",
	Solutions:          []Solution{Solution{Solution: "helm_tiller"}},
}

var testGKECluster = Cluster{
	Name:               "Test GKE Cluster Go SDK " + getTicks(),
	Provider:           "gke",
	MasterCount:        1,
	MasterSize:         "n1-standard-1",
	WorkerCount:        2,
	WorkerSize:         "n1-standard-1",
	Region:             "us-east1-c",
	ProviderNetworkID:  "__new__",
	ProviderNetworkCdr: "172.23.0.0/16",
	ProviderSubnetID:   "__new__",
	ProviderSubnetCidr: "172.23.1.0/24",
	KubernetesVersion:  "latest",
	RbacEnabled:        true,
	DashboardEnabled:   true,
	EtcdType:           "classic",
	Platform:           "gci",
	Channel:            "stable",
	Solutions:          []Solution{Solution{Solution: "helm_tiller"}},
}

var testGCECluster = Cluster{
	Name:               "Test GCE Cluster Go SDK " + getTicks(),
	Provider:           "gce",
	MasterCount:        1,
	MasterSize:         "n1-standard-1",
	WorkerCount:        2,
	WorkerSize:         "n1-standard-1",
	Region:             "us-east1-c",
	ProviderNetworkID:  "__new__",
	ProviderNetworkCdr: "172.23.0.0/16",
	ProviderSubnetID:   "__new__",
	ProviderSubnetCidr: "172.23.1.0/24",
	KubernetesVersion:  "v1.13.1",
	RbacEnabled:        true,
	DashboardEnabled:   true,
	EtcdType:           "classic",
	Platform:           "coreos",
	Channel:            "stable",
	Solutions:          []Solution{Solution{Solution: "helm_tiller"}},
}

var clusterIds = make([]int, 0)

func TestClusterCreateAWS(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	sshKeysetID, err := GetIDFromEnv("NKS_SSH_KEYSET")
	if err != nil {
		t.Error(err)
	}

	awsKeysetID, err := GetIDFromEnv("NKS_AWS_KEYSET")
	if err != nil {
		t.Error(err)
	}

	testAwsCluster.ProviderKey = awsKeysetID
	testAwsCluster.SSHKeySet = sshKeysetID

	cluster, err := c.CreateCluster(orgID, testAwsCluster)
	if err != nil {
		t.Error(err)
	}

	clusterIds = append(clusterIds, cluster.ID)

	timeout := 1200
	c.WaitClusterRunning(orgID, cluster.ID, true, timeout)
	if err != nil {
		t.Error(err)
	}
}

func TestClusterCreateEKS(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	sshKeysetID, err := GetIDFromEnv("NKS_SSH_KEYSET")
	if err != nil {
		t.Error(err)
	}

	awsKeysetID, err := GetIDFromEnv("NKS_AWS_KEYSET")
	if err != nil {
		t.Error(err)
	}

	testEKSCluster.ProviderKey = awsKeysetID
	testEKSCluster.SSHKeySet = sshKeysetID

	cluster, err := c.CreateCluster(orgID, testEKSCluster)
	if err != nil {
		t.Error(err)
	}

	clusterIds = append(clusterIds, cluster.ID)

	timeout := 1200
	c.WaitClusterRunning(orgID, cluster.ID, true, timeout)
	if err != nil {
		t.Error(err)
	}
}

func TestClusterCreateAzure(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	sshKeysetID, err := GetIDFromEnv("NKS_SSH_KEYSET")
	if err != nil {
		t.Error(err)
	}

	azureKeysetID, err := GetIDFromEnv("NKS_AZR_KEYSET")
	if err != nil {
		t.Error(err)
	}

	testAzureCluster.ProviderKey = azureKeysetID
	testAzureCluster.SSHKeySet = sshKeysetID

	cluster, err := c.CreateCluster(orgID, testAzureCluster)
	if err != nil {
		t.Error(err)
	}

	clusterIds = append(clusterIds, cluster.ID)

	timeout := 1200
	c.WaitClusterRunning(orgID, cluster.ID, true, timeout)
	if err != nil {
		t.Error(err)
	}
}

func TestClusterCreateAKS(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	sshKeysetID, err := GetIDFromEnv("NKS_SSH_KEYSET")
	if err != nil {
		t.Error(err)
	}

	azureKeysetID, err := GetIDFromEnv("NKS_AZR_KEYSET")
	if err != nil {
		t.Error(err)
	}

	testAKSCluster.ProviderKey = azureKeysetID
	testAKSCluster.SSHKeySet = sshKeysetID

	cluster, err := c.CreateCluster(orgID, testAKSCluster)
	if err != nil {
		t.Error(err)
	}

	clusterIds = append(clusterIds, cluster.ID)

	timeout := 1200
	c.WaitClusterRunning(orgID, cluster.ID, true, timeout)
	if err != nil {
		t.Error(err)
	}
}

func TestClusterCreateGCE(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	sshKeysetID, err := GetIDFromEnv("NKS_SSH_KEYSET")
	if err != nil {
		t.Error(err)
	}

	gceKeysetID, err := GetIDFromEnv("NKS_GCE_KEYSET")
	if err != nil {
		t.Error(err)
	}

	testGCECluster.ProviderKey = gceKeysetID
	testGCECluster.SSHKeySet = sshKeysetID

	cluster, err := c.CreateCluster(orgID, testGCECluster)
	if err != nil {
		t.Error(err)
	}

	clusterIds = append(clusterIds, cluster.ID)

	timeout := 1200
	c.WaitClusterRunning(orgID, cluster.ID, true, timeout)
	if err != nil {
		t.Error(err)
	}
}

func TestClusterCreateGKE(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	sshKeysetID, err := GetIDFromEnv("NKS_SSH_KEYSET")
	if err != nil {
		t.Error(err)
	}

	gkeKeysetID, err := GetIDFromEnv("NKS_GKE_KEYSET")
	if err != nil {
		t.Error(err)
	}

	testGKECluster.ProviderKey = gkeKeysetID
	testGKECluster.SSHKeySet = sshKeysetID

	cluster, err := c.CreateCluster(orgID, testGKECluster)
	if err != nil {
		t.Error(err)
	}

	clusterIds = append(clusterIds, cluster.ID)

	timeout := 1200
	c.WaitClusterRunning(orgID, cluster.ID, true, timeout)
	if err != nil {
		t.Error(err)
	}
}
func TestClusterList(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	clusters, err := c.GetClusters(orgID)
	if err != nil {
		t.Error(err)
	}

	assert.True(t, len(clusters) > 0, "There should be at lease one cluster")
}

func TestClusterGet(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	cluster, err := c.GetCluster(orgID, clusterIds[0])
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, cluster, "Cluster does not exists")
}

func TestClusterDelete(t *testing.T) {
	c, err := NewClientFromEnv()
	if err != nil {
		t.Error(err)
	}
	orgID, err := GetIDFromEnv("NKS_ORG_ID")
	if err != nil {
		t.Error(err)
	}

	for _, clusterID := range clusterIds {
		err = c.DeleteCluster(orgID, clusterID)
		if err != nil {
			t.Error(err)
		}
	}
}
