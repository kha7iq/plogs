package main

import (
	"log"
	"os"
	"sync"

	"github.com/kha7iq/pl/internal/client"
	pl "github.com/kha7iq/pl/internal/plogs"
	"github.com/urfave/cli/v2"
)

var (
	markWord      string // Word to mark in logs
	namespace     string // Kubernetes namespace
	containerName string // Container name within pod
	label         string // Labels to match
	podName       string // Pod name
	version       string // App version
	followLogs    bool   // Follow logs option
	tailLines     int64  // Tail lines option
)

func main() {
	app := &cli.App{
		Name:      "plogs",
		Usage:     "Retrieve and manage Kubernetes pod logs with advanced filtering and highlighting.",
		UsageText: "kubectl plogs --pod mypod --namespace default --container nginx --tail 10 --follow ",
		Version:   version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "mark",
				Aliases:     []string{"m"},
				Value:       "",
				Usage:       "Marks the specified word or sentence within the logs, highlighting occurrences.",
				Destination: &markWord,
			},
			&cli.StringFlag{
				Name:        "namespace",
				Aliases:     []string{"n"},
				Value:       "",
				Usage:       "Sets the Kubernetes namespace to retrieve logs from.",
				Destination: &namespace,
				EnvVars:     []string{"PLOGS_NAMESPACE"},
			},
			&cli.StringFlag{
				Name:        "container",
				Aliases:     []string{"c"},
				Value:       "",
				Usage:       "Specifies the container name within the pod to retrieve logs from incase of multiple containers.",
				Destination: &containerName,
			},
			&cli.StringFlag{
				Name:        "label",
				Aliases:     []string{"l"},
				Value:       "",
				Usage:       "Filters pods using labels to match for log retrieval.",
				Destination: &label,
			},
			&cli.StringFlag{
				Name:        "pod",
				Aliases:     []string{"p"},
				Value:       "",
				Usage:       "Specifies the pod name to retrieve logs from.",
				Destination: &podName,
			},
			&cli.BoolFlag{
				Name:        "follow",
				Aliases:     []string{"f"},
				Value:       false,
				Usage:       "Follows log output after retrieving logs.",
				Destination: &followLogs,
				EnvVars:     []string{"PLOGS_FOLLOW"},
			},
			&cli.Int64Flag{
				Name:        "tail",
				Aliases:     []string{"t"},
				Usage:       "Specifies the number of lines to retrieve from the end of the logs.",
				Destination: &tailLines,
				EnvVars:     []string{"PLOGS_TAIL_LINES"},
			},
		},

		Action: func(c *cli.Context) error {
			// Initialize Kubernetes client
			cs, err := client.New("")
			if err != nil {
				return err
			}

			// Check and set followLogs flag
			if c.IsSet("follow") {
				followLogs = true
			}

			// If namespace not provided, retrieve the current namespace
			if len(namespace) == 0 {
				namespace, err = cs.GetCurrentNamespace()
			}
			if err != nil {
				return err
			}

			// Manage a waitgroup for concurrent goroutines
			var wg sync.WaitGroup

			// Retrieve and display pod logs based on provided parameters
			if err := pl.GetPodLogs(cs, namespace, label, markWord, podName, containerName, followLogs, tailLines, &wg); err != nil {
				return err
			}

			// Wait for all goroutines to finish before exiting
			wg.Wait()
			return nil
		},
	}

	// Execute the CLI application
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
