package commands

import (
	"context"
	"fmt"

	"github.com/la3mmchen/elastic-cluster-diff/internal/types"
	"github.com/r3labs/diff"
	pretty "github.com/kr/pretty"
	"github.com/urfave/cli"

	"github.com/olivere/elastic"
)

func compareCluster() cli.Command {

	cmd := cli.Command{
		Name:  "compare",
		Usage: "Compare indices on base of the number of documents.",
	}

	cmd.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "config",
			Destination: &PrintConfig,
			Usage:       "Print retrieved config",
		},
		cli.StringSliceFlag{
			Name:  "cluster",
			Usage: "Clusters to compare.",
			//Value: &cli.StringSlice{"localhost:9200", "localhost:8200"}, //TODO: default values? Drop them
		},
	}

	cmd.Action = func(c *cli.Context) error {
		var conf types.Config

		// populate config
		m := make(map[string]*types.Cluster)
		for _, item := range c.StringSlice("cluster") {
			// build cluster object
			clstr := types.Cluster{
				Name:   "undef",
				Remote: item,
			}
			m[item] = &clstr
		}
		conf.ElasticCluster = m

		// load basic settings
		clsts := conf.ElasticCluster
		for key, value := range clsts {
			_ = key
			_ = value

			// Create Client
			client, err := elastic.NewClient(elastic.SetURL("http://" + value.Remote))
			if err != nil {
				fmt.Printf("ERR %v not reachable \n", value.Remote)
				continue //FIXME: implement some proper error handling if cluster not reachable
			}

			// Get Cluster Version
			esversion, err := client.ElasticsearchVersion("http://" + value.Remote)
			if err != nil {
				// Handle error
				panic(err)
			}

			// Get cluster state
			clusterHealth, err := client.ClusterHealth().Do(context.Background())
			if err != nil {
				panic(err)
			}
			// tmpCluster := types.ClusterHealthResponse(clusterHealth) // FIXME: this would be cooler with embedidng the elastic.ClusterHealthResponse in types.ClusterHealthResponse and convert the struct later. But it does not work so we go the hard way
			tmpCluster := types.ClusterHealthResponse{
				ClusterName:                    clusterHealth.ClusterName,
				Status:                         clusterHealth.Status,
				TimedOut:                       clusterHealth.TimedOut,
				NumberOfNodes:                  clusterHealth.NumberOfNodes,
				NumberOfDataNodes:              clusterHealth.NumberOfDataNodes,
				ActivePrimaryShards:            clusterHealth.ActivePrimaryShards,
				ActiveShards:                   clusterHealth.ActiveShards,
				RelocatingShards:               clusterHealth.RelocatingShards,
				InitializingShards:             clusterHealth.InitializingShards,
				UnassignedShards:               clusterHealth.UnassignedShards,
				DelayedUnassignedShards:        clusterHealth.DelayedUnassignedShards,
				NumberOfPendingTasks:           clusterHealth.NumberOfPendingTasks,
				NumberOfInFlightFetch:          clusterHealth.NumberOfInFlightFetch,
				TaskMaxWaitTimeInQueueInMillis: clusterHealth.TaskMaxWaitTimeInQueueInMillis,
				ActiveShardsPercentAsNumber:    clusterHealth.ActiveShardsPercentAsNumber,
			}

			// Get all indices
			indices, err := client.IndexNames()
			if err != nil {
				panic(err)
			}

			// Get information about the indices
			docCount, err := elastic.NewCatCountService(client).Index(indices...).Pretty(true).Do(context.Background())
			var clusterIndices []types.Index
			// Load details of each index
			for _, idx := range indices {
				var index types.Index
				docs, _ := elastic.NewCatCountService(client).Columns("count").Index(idx).Pretty(true).Do(context.Background())
				index.Name = idx
				index.DocCount = docs[0].Count
				clusterIndices = append(clusterIndices, index)
			}

			// Persist values
			conf.ElasticCluster[key].Name = clusterHealth.ClusterName
			conf.ElasticCluster[key].Version = esversion
			conf.ElasticCluster[key].TotalDocuments = docCount[0].Count
			conf.ElasticCluster[key].ClusterHealth = tmpCluster
			conf.ElasticCluster[key].Indices = clusterIndices

			client.Stop()
		}

		if PrintConfig {
			fmt.Printf("\n")
			fmt.Printf("%# v", pretty.Formatter(conf))
			fmt.Printf("\n")
		}

		// Diff the information
		// TODO: I guess this is not the smartest solution to run a diff each entrys of a slice.
		var keys []string
		for k := range conf.ElasticCluster {
			keys = append(keys, k)
		}
		for i := 0; i < len(keys); i++ {
			if i < len(keys)-1 {
				tmp, _ := diff.Diff(conf.ElasticCluster[keys[i]], conf.ElasticCluster[keys[i+1]])
				fmt.Printf("\n \nDiff +%v with -%v \n", keys[i], keys[i+1])
				fmt.Printf("%# v", pretty.Formatter(tmp))
			}
		}
		return nil
	}
	return cmd
}
