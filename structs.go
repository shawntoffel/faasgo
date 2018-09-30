package faasgo

type Function struct {
	Annotations  []string          `json:"annotations"`
	Constraints  []string          `json:"constraints"`
	EnvProcess   string            `json:"envProcess"`
	EnvVars      map[string]string `json:"envVars,omitempty"`
	Image        string            `json:"image"`
	Labels       []string          `json:"labels"`
	Limits       FunctionLimits    `json:"limits,omitempty"`
	Network      string            `json:"network,omitempty"`
	RegistryAuth string            `json:"registryAuth,omitempty"`
	Requests     FunctionRequests  `json:"requests,omitempty"`
	Secrets      []string          `json:"secrets"`
	Service      string            `json:"service"`
}

type FunctionLimits struct {
	CPU    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

type FunctionRequests struct {
	CPU    string `json:"cpu,omitempty"`
	Memory string `json:"memory,omitempty"`
}

type FunctionListEntry struct {
	Annotations       map[string]string `json:"annotations,omitempty"`
	AvailableReplicas int64             `json:"availableReplicas"`
	EnvProcess        string            `json:"envProcess"`
	Image             string            `json:"image"`
	InvocationCount   int64             `json:"invocationCount"`
	Labels            map[string]string `json:"labels"`
	Name              string            `json:"name"`
	Replicas          int64             `json:"replicas"`
}

type ScaleFunctionRequest struct {
	Service  string `json:"service"`
	Replicas int    `json:"replicas"`
}

type DeleteFunctionRequest struct {
	FunctionName string `json:"functionName"`
}

type Info struct {
	Provider InfoProvider `json:"provider"`
	Version  InfoVersion  `json:"version"`
}

type InfoProvider struct {
	Orchestration string              `json:"orchestration,omitempty"`
	Provider      string              `json:"provider,omitempty"`
	Version       InfoProviderVersion `json:"version,omitempty"`
}

type InfoProviderVersion struct {
	CommitMessage string `json:"commit_message,omitempty"`
	Release       string `json:"release,omitempty"`
	Sha           string `json:"sha,omitempty"`
}

type InfoVersion struct {
	CommitMessage string `json:"commit_message,omitempty"`
	Release       string `json:"release,omitempty"`
	Sha           string `json:"sha,omitempty"`
}
