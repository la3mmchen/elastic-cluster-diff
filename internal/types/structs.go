package types

// Config <tbd>
type Config struct {
	ElasticCluster map[string]*Cluster
}

// Cluster <tbd>
type Cluster struct {
	Name           string                `diff:"-"`
	Remote         string                `diff:"-"`
	Version        string                `diff:"Version"`
	TotalDocuments int                   `diff:"TotalDocuments"`
	ClusterHealth  ClusterHealthResponse `diff:"ClusterHealth"`
	Indices        []Index               `diff:"Indices"`
}

// ClusterHealthResponse <tbd>
type ClusterHealthResponse struct {
	ClusterName                    string  `diff:"-"`
	Status                         string  `diff:"Status"`
	TimedOut                       bool    `diff:"TimedOut"`
	NumberOfNodes                  int     `diff:"NumberOfNodes"`
	NumberOfDataNodes              int     `diff:"NumberOfDataNodes"`
	ActivePrimaryShards            int     `diff:"ActivePrimaryShards"`
	ActiveShards                   int     `diff:"ActiveShards"`
	RelocatingShards               int     `diff:"RelocatingShards"`
	InitializingShards             int     `diff:"InitializingShards"`
	UnassignedShards               int     `diff:"UnassignedShards"`
	DelayedUnassignedShards        int     `diff:"DelayedUnassignedShards"`
	NumberOfPendingTasks           int     `diff:"NumberOfPendingTasks"`
	NumberOfInFlightFetch          int     `diff:"NumberOfInFlightFetch"`
	TaskMaxWaitTimeInQueueInMillis int     `diff:"TaskMaxWaitTimeInQueueInMillis"`
	ActiveShardsPercentAsNumber    float64 `diff:"ActiveShardsPercentAsNumber"`
}

// Index <tdb>
type Index struct {
	Name     string `diff:"Name"`
	DocCount int    `diff:"DocCount"`
}
