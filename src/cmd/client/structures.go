package main

type projectConfiguration struct {
	// Name - The name of the project.
	Name string `yaml:"name"`

	// Description - The description of the project.
	Description string `yaml:"description"`
}

type mainConfiguration struct {
	// Version - The version of the configuration file.
	Version float64 `yaml:"version"`

	// Project - Basic information about the project.
	Project *projectConfiguration `yaml:"project"`
}
