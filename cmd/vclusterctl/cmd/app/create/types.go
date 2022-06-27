package create

// CreateOptions holds the create cmd options
type CreateOptions struct {
	KubeConfigContextName string
	ChartVersion          string
	ChartName             string
	ChartRepo             string
	LocalChartDir         string
	K3SImage              string
	Distro                string
	CIDR                  string
	ExtraValues           []string

	KubernetesVersion string

	CreateNamespace    bool
	DisableIngressSync bool
	CreateClusterRole  bool
	UpdateCurrent      bool
	Expose             bool
	ExposeLocal        bool

	Connect       bool
	Upgrade       bool
	Isolate       bool
	ReleaseValues string
}
