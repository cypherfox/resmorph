/*
Copyright © 2025 Lutz Behnke <lutz.behnke@gmx.de>
This file is part of {{ .appName }}
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/cypherfox/resmorph/api/tracker"
	//"github.com/cypherfox/snacks-manager/internal/backend"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
)

// trackerCmd represents the tracker command
var trackerCmd = &cobra.Command{
	Use:   "tracker",
	Short: "start the tracker service",
	Long: `Start the tracker service to receive notifications about the existence and status of desired state resources.
	    While primarily focussed on Kubernetes resources, in principle any resource definition in YAML or JSON format may be
	    managed.`,
	Run: func(cmd *cobra.Command, args []string) {
		serverAddr := "0.0.0.0:8080"

		runListener(serverAddr)
	},
}

func runListener(addr string) {

	backend := backend.NewTrackerBackEnd()

	strict := tracker.NewApiHandler(backend)

	r := mux.NewRouter()

	// get an `http.Handler` that we can use
	h := oapi.HandlerFromMuxWithBaseURL(strict, r, "")

	fmt.Printf("starting server and listening on port %s", addr)

	s := &http.Server{
		Handler: h,
		Addr:    addr,
	}

	err := s.ListenAndServe()
	fmt.Printf("\n ended server \n")
	if err != nil {
		if err != nil {
			fmt.Printf("with error: %s", err.Error())
		}
	}
}

func init() {
	rootCmd.AddCommand(trackerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// trackerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// trackerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
