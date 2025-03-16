package main

import (
	"fmt"
	"os"
)

func ResolveKeyStorePath() (string, error) {
	// Define the directory path
	dirPath := CONFIG["KMS_STORE"]

	// Expand the tilde (~) to the full home directory path
	expandedDirPath, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		return "", err
	}
	dirPath = expandedDirPath + "/keystore"

	// Check if the directory exists
	_, err = os.Stat(dirPath)
	if os.IsNotExist(err) {
		// Create the directory if it doesn't exist
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return "", err
		}
		fmt.Println("Directory created:", dirPath)
	} else {
		fmt.Println("Directory already exists:", dirPath)
	}
	return dirPath, nil
}
