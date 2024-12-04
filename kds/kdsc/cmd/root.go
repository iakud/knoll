/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"

	"github.com/iakud/keeper/kds/kdsc/codegen"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kds",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
	Run: execute,
}

var kdsCommand KdsCommand

type KdsCommand struct {
	kdsPath []string
	tplPath string
	out string
}

func execute(cmd *cobra.Command, args []string) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("kds error: %v\n%s", err, debug.Stack())
		}
	}()
	const suffix = ".kds"
	var files []string
	var unique = make(map[string]struct{})
	log.Println("args:", args)
	for _, pattern := range args {
		for _, kdsPath := range kdsCommand.kdsPath {
			matches, err := filepath.Glob(filepath.Join(kdsPath, pattern))
			if err != nil {
				panic(err)
			}
			for _, match := range matches {
				abs, err := filepath.Abs(match)
				if err != nil {
					panic(err)
				}
				cleanpath := filepath.Clean(abs)
				if _, ok := unique[cleanpath]; ok {
					continue
				}
				if stat, err := os.Stat(match); err != nil {
					panic(err)
				} else if stat.IsDir() {
					continue
				} else if suffix != "" && !strings.HasSuffix(match, suffix) {
					continue
				}
				files = append(files, match)
				unique[cleanpath] = struct{}{}
			}
		}
	}

	// templates
	filenames, err := filepath.Glob(filepath.Join(kdsCommand.tplPath, "*.tpl"))
	if err != nil {
		panic(err)
	}
	codegen.Parse(files, filenames, kdsCommand.out)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kds.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringArrayVarP(&kdsCommand.kdsPath, "kds_path", "I", []string{""}, "Specify the directory in which to search for imports. May be specified multiple times; directories will be searched in order. If not given, the current working directory is used.")
	rootCmd.Flags().StringVar(&kdsCommand.tplPath, "tpl_path", "", "Template file.")
	rootCmd.Flags().StringVar(&kdsCommand.out, "out", "", "Generate file.")
}