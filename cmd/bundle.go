package main

import (
	"fmt"
	"os"

	esbuild "github.com/evanw/esbuild/pkg/api"
)

func Bundle() esbuild.BuildResult {
	return esbuild.Build(esbuild.BuildOptions{
		EntryPoints:       []string{"pkg/web/app/src/main.mjs"},
		Bundle:            true,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
		Outfile:           "pkg/web/app/dist/main.min.js",
		Write:             true,
	})
}

func main() {
	bundleWeb := Bundle()
	if len(bundleWeb.Errors) > 0 {
		fmt.Println(bundleWeb.Errors)
		os.Exit(1)
	}
	if len(bundleWeb.OutputFiles) > 0 {
		for _, files := range bundleWeb.OutputFiles {
			fmt.Println("Created: ", files.Path)
		}
	}
}
