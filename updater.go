package main

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const repoApi = "https://api.github.com/repos/ur-wesley/csvParser/releases/latest"

var currentVersion string

type GithubRepo struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name        string `json:"name"`
		DownloadUrl string `json:"browser_download_url"`
	} `json:"assets"`
}

func UpdateAvailable() (bool, error) {
	repo, err := fetchRepo()
	if err != nil {
		return false, err
	}

	return repo.TagName != currentVersion, nil
}

func DownloadUpdate() error {
	repo, err := fetchRepo()
	if err != nil {
		return err
	}

	asset := repo.Assets[0]
	resp, err := http.Get(asset.DownloadUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	tmpFile, err := os.CreateTemp("", asset.Name)
	if err != nil {
		return err
	}
	defer os.Remove(tmpFile.Name())

	_, err = io.Copy(tmpFile, resp.Body)
	if err != nil {
		return err
	}

	err = tmpFile.Close()
	if err != nil {
		return err
	}

	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	err = unzip(tmpFile.Name(), currentDir)
	if err != nil {
		return err
	}

	log.Println("Update downloaded and applied successfully")
	return nil
}

func fetchRepo() (GithubRepo, error) {
	resp, err := http.Get(repoApi)
	if err != nil {
		return GithubRepo{}, err
	}
	defer resp.Body.Close()

	var repo GithubRepo
	err = json.NewDecoder(resp.Body).Decode(&repo)
	if err != nil {
		return GithubRepo{}, err
	}

	return repo, nil
}

func Update() {}

func unzip(src, dest string) error {
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, f := range r.File {
		fpath := filepath.Join(dest, f.Name)
		fmt.Println(dest, f.Name)
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", fpath)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		if err := os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return err
		}

		rc, err := f.Open()
		if err != nil {
			return err
		}

		_, err = io.Copy(outFile, rc)
		outFile.Close()
		rc.Close()

		if err != nil {
			return err
		}
	}
	return nil
}
