package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/zbyju/todogo/internal/application/fdiscovery"
	"github.com/zbyju/todogo/internal/domain/contents"
	"github.com/zbyju/todogo/internal/domain/filesystem"
	"github.com/zbyju/todogo/internal/domain/parsing/regex"
)

var serveCmd = &cobra.Command{
	Use:   "discover",
	Short: "Discover command does stuff :)",
	Long:  `Let's see what this command actually does in practice.`,
	Run: func(cmd *cobra.Command, args []string) {
		wd, err := os.Getwd()
		if err != nil {
			log.Println("Cannot resolve the working directory: ", err)
			return
		}

		fd := fdiscovery.FileDiscoveryImpl{}
		folder, err := fd.Discover(wd)
		if err != nil {
			fmt.Println("Discovery error: ", err)
		}

		err = parseFolder(folder)
		if err != nil {
			fmt.Println("Parsing error: ", err)
		}
	},
}

func parseFolder(folder filesystem.Folder) error {
	for _, file := range folder.Files {
		err := parseFile(file)
		if err != nil {
			return err
		}
	}

	for _, folder := range folder.Folders {
		err := parseFolder(folder)
		if err != nil {
			return err
		}
	}

	return nil
}

func parseFile(file filesystem.File) error {
	if !file.IsKnown() {
		return nil
	}

	fullpath := filepath.Join(file.Path, file.Name)
	f, err := os.Open(fullpath)
	if err != nil {
		return errors.New("Cannot open file: " + fullpath)
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println("Cannot close file: ", err)
		}
	}()

	scanner := bufio.NewScanner(f)

	state := regex.NewState(file.Extension, false)
	annotations := []contents.Annotation{}
	for scanner.Scan() {
		line := scanner.Text()
		newAnnotations, newState := regex.ParseLine(line, state)
		state = newState
		annotations = append(annotations, newAnnotations...)
	}
	fmt.Println(fullpath, annotations)

	if err := scanner.Err(); err != nil {
		return errors.New("Cannot read file: " + fullpath)
	}

	return nil
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
