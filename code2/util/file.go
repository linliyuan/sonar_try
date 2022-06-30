package util

import "io/ioutil"

func Write(path string, data []byte) error {
	return ioutil.WriteFile(path, data, 0644)
}

func Read(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}
