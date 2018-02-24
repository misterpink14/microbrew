package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"testing"
)

const (
	TempFolderName1 = "tempfolder1"
	TempFolderName2 = "tempfolder2"
	TempFileName1   = "tempfile1"
	TempFileName2   = "tempfile2"

	EmptyFolderName = "emptyfolder"
)

func CreateFile(filePath, fileContents string) {
	if file, err := os.Create(filePath); err != nil {
		log.Fatalln(err)
	} else {
		if fileContents != "" {
			file.Write([]byte(fileContents))
		}
		defer file.Close()
	}
}

func SetupTempDir(dirPath string, fileNames ...string) {
	CleanupTempDir(dirPath)
	if err := os.Mkdir(dirPath, 0755); err != nil {
		log.Fatalln(err)
	}
	for _, fileName := range fileNames {
		CreateFile(fmt.Sprintf("%s/%s", dirPath, fileName), "")
	}
}

func CleanupTempDir(dirPath string) error {
	if err := os.RemoveAll(dirPath); err != nil {
		return err
	}
	return nil
}

func TestTemplate_CopyTemplate_Home(t *testing.T) {
	dirPath1 := os.TempDir() + TempFolderName1
	dirPath2 := os.TempDir() + TempFolderName2
	sourceFilePath := fmt.Sprintf("%s/%s", dirPath1, TempFileName1)
	destFilePath := fmt.Sprintf("%s/%s", dirPath2, TempFileName1)
	localUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	SetupTempDir(dirPath1)
	SetupTempDir(dirPath2)

	CreateFile(sourceFilePath, "{{.Home}}")
	NewTemplateFile(sourceFilePath, destFilePath).CopyTemplate()

	fileContents, err := ioutil.ReadFile(destFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	if _, err := os.Stat(destFilePath); err != nil {
		t.Errorf("The destination file path does not exist")
	}
	expected := localUser.HomeDir
	if string(fileContents) != expected {
		t.Errorf("%s\nis not the same as\n%s", string(fileContents), expected)
	}
	CleanupTempDir(dirPath1)
	CleanupTempDir(dirPath2)
}

func TestTemplates_GetTemplateFileNames_EmptyDir(t *testing.T) {
	emptyPath := os.TempDir() + EmptyFolderName
	SetupTempDir(emptyPath)

	// Validate that an empty array is returned for empty folder
	templates := NewTemplates(emptyPath, map[string]string{})
	emptyTemplateFileNames, err := templates.GetTemplateFileNames()
	if err != nil {
		t.Error(err)
	}
	if len(emptyTemplateFileNames) != 0 || err != nil {
		log.Println(err)
		t.Errorf("Expecting 0 file names, got %d", len(emptyTemplateFileNames))
	}
	CleanupTempDir(emptyPath)
}

func TestTemplates_GetTemplateFileNames_NonEmptyDir(t *testing.T) {
	dirPath := os.TempDir() + TempFolderName1
	SetupTempDir(dirPath, TempFileName1, TempFileName2)

	// Create a Templates object and validate that the temp files exist
	templates := NewTemplates(dirPath, map[string]string{})
	templateFileNames, err := templates.GetTemplateFileNames()
	if err != nil {
		t.Error(err)
	}

	if len(templateFileNames) != 4 || err != nil {
		log.Println(err)
		t.Errorf("Expecting 2 file names, got %d", len(templateFileNames)-2)
	}
	if !Includes(templateFileNames, TempFileName1) {
		log.Println(err)
		t.Errorf("File is not included: %s", TempFileName1)
	}
	if !Includes(templateFileNames, TempFileName2) {
		log.Println(err)
		t.Errorf("File is not included: %s", TempFileName2)
	}
	CleanupTempDir(dirPath)
}

func TestTemplates_CopyTemplates(t *testing.T) {
	dirPath := os.TempDir() + TempFolderName1
	SetupTempDir(dirPath)
	templatesConfig := map[string]string{
		"init.vim":     dirPath + "/init.vim",
		"vimrc":        dirPath + "/vimrc",
		"bash_profile": dirPath + "/bash_profile",
	}

	templates := NewTemplates(TemplatesPath, templatesConfig)
	templates.CopyTemplates()

	if files, err := ioutil.ReadDir(dirPath); err != nil {
		t.Error(err)
	} else {
		if len(files) != 3 {
			t.Errorf("Expecting %d files, got %d", len(templatesConfig), len(files))
		}
		for _, fileInfo := range files {
			if _, ok := templatesConfig[fileInfo.Name()]; !ok {
				t.Errorf("%s is not in the templates config: %s", fileInfo.Name(), templatesConfig)
			}
		}
	}
	CleanupTempDir(dirPath)
}
