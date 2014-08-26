package main

import "os"

func fileExists(path string) (bool, error) {
	_, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
