package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Dependency struct {
	GroupID    string `xml:"groupId"`
	ArtifactID string `xml:"artifactId"`
	Version    string `xml:"version"`
}

type Project struct {
	Dependencies []Dependency `xml:"dependencies>dependency"`
}

func main() {
	// Replace "path/to/your/pom.xml" with the actual path to your pom.xml file
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run read_maven_classpath.go <path/to/project>")
		os.Exit(1)
	}

	pomPath := os.Args[1]

	file, err := os.Open(pomPath)
	if err != nil {
		fmt.Println("Error opening pom.xml:", err)
		return
	}
	defer file.Close()

	var project Project
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&project)
	if err != nil {
		fmt.Println("Error decoding pom.xml:", err)
		return
	}

	// Print the dependencies
	fmt.Println("Dependencies:")
	for _, dep := range project.Dependencies {
		fmt.Printf("%s:%s:%s\n", dep.GroupID, dep.ArtifactID, dep.Version)
	}
}
