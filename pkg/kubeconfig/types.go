package kubeconfig

// KubeConfig represents the structure of a Kubernetes config file
type KubeConfig struct {
	APIVersion     string        `yaml:"apiVersion"`
	Kind           string        `yaml:"kind"`
	Clusters       []ClusterItem `yaml:"clusters"`
	Contexts       []ContextItem `yaml:"contexts"`
	Users          []UserItem    `yaml:"users"`
	Preferences    interface{}   `yaml:"preferences"`
	CurrentContext string        `yaml:"current-context,omitempty"` // Add back with omitempty
}

// ClusterItem represents a cluster entry in KubeConfig
type ClusterItem struct {
	Name    string        `yaml:"name"`
	Cluster ClusterConfig `yaml:"cluster"`
}

// ClusterConfig holds the cluster configuration details
type ClusterConfig struct {
	Server                   string `yaml:"server"`
	CertificateAuthorityData string `yaml:"certificate-authority-data,omitempty"`
}

// ContextItem represents a context entry in KubeConfig
type ContextItem struct {
	Name    string        `yaml:"name"`
	Context ContextConfig `yaml:"context"`
}

// ContextConfig holds the context configuration details
type ContextConfig struct {
	Cluster string `yaml:"cluster"`
	User    string `yaml:"user"`
}

// UserItem represents a user entry in KubeConfig
type UserItem struct {
	Name string     `yaml:"name"`
	User UserConfig `yaml:"user"`
}

// UserConfig holds the user authentication details
type UserConfig struct {
	ClientCertificateData string `yaml:"client-certificate-data,omitempty"`
	ClientKeyData         string `yaml:"client-key-data,omitempty"`
	Token                 string `yaml:"token,omitempty"`
}
