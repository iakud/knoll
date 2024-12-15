/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/iakud/krocher/kds/kdsc/codegen"
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
	kdsPaths []string
	tplPath string
	out string
}

func execute(cmd *cobra.Command, args []string) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("kds error: %v\n%s", err, debug.Stack())
		}
	}()
	// 获取全部kds文件
	kdsFiles, err := globKdsFiles(args)
	if err != nil {
		panic(err)
	}
	// 生成kds
	if err := codegen.Parse(kdsFiles, kdsCommand.tplPath, kdsCommand.out); err != nil {
		panic(err)
	}
}

func globKdsFiles(patterns []string) ([]string, error) {
	var kdsFiles []string
	var unique = make(map[string]struct{})
	for _, pattern := range patterns {
		for _, kdsPath := range kdsCommand.kdsPaths {
			matcheFiles, err := filepath.Glob(filepath.Join(kdsPath, pattern))
			if err != nil {
				return nil, err
			}
			for _, matchFile := range matcheFiles {
				absFile, err := filepath.Abs(matchFile)
				if err != nil {
					return nil, err
				}
				if _, ok := unique[absFile]; ok {
					continue
				}
				if stat, err := os.Stat(matchFile); err != nil {
					return nil, err
				} else if stat.IsDir() {
					continue
				}
				kdsFiles = append(kdsFiles, matchFile)
				unique[absFile] = struct{}{}
			}
		}
	}
	return kdsFiles, nil
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
	rootCmd.Flags().StringArrayVarP(&kdsCommand.kdsPaths, "kds_path", "I", []string{""}, "Specify the directory in which to search for imports. May be specified multiple times; directories will be searched in order. If not given, the current working directory is used.")
	rootCmd.Flags().StringVar(&kdsCommand.tplPath, "tpl_path", "", "Template file.")
	rootCmd.Flags().StringVar(&kdsCommand.out, "out", "", "Generate file.")
}