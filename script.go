package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/digitalocean/godo"
)

func main() {
	fmt.Println("Creating kube cluster...")
	client := godo.NewFromToken(os.Getenv("DIGITALOCEAN_ACCESS_TOKEN"))
	ctx := context.TODO()
	cluster := createCLuster(client, ctx)
	deployFoldingImage(cluster.ID)
	fmt.Println("Successfully created a folding cluster!")
}

// createCluster
func createCLuster(client *godo.Client, ctx context.Context) *godo.KubernetesCluster {
	createRequest := &godo.KubernetesClusterCreateRequest{
		Name:        "folding-cluster",
		RegionSlug:  "sfo2",
		VersionSlug: "1.19.3-do.2",
		Tags:        []string{"folding"},
		NodePools: []*godo.KubernetesNodePoolCreateRequest{
			{
				Name:   "folding-pool",
				Size:   "s-8vcpu-16gb",
				Count:  5,
				Tags:   []string{"folding"},
				Labels: map[string]string{"app": "folding"},
			},
		},
	}
	cluster, _, err := client.Kubernetes.Create(ctx, createRequest)
	handleError(err)
	return cluster
}

// deployFoldingImage
func deployFoldingImage(id string) {
	fmt.Println("Waiting for the cluster to be provisioned. This should take 4 minutes...")
	time.Sleep(4 * time.Minute)
	cmd := exec.Command("doctl", "kubernetes", "cluster", "kubeconfig", "save", id)
	err := cmd.Run()
	handleError(err)
	cmd = exec.Command("kubectl", "apply", "-f", "deployment.yml")
	err = cmd.Run()
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
