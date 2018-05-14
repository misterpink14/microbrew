package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"text/template"
)

type ITemplates interface {
	GetTemplateFileNames() ([]string, error)
	CopyTemplates()
}

type ITemplateFile interface {
	CopyTemplate()
}

type Templates struct {
	TemplatesPath   string
	TemplatesConfig map[string]string
}

type TemplateFile struct {
	UserHome       string
	SourceFilePath string
	DestFilePath   string
}

func NewTemplates(templatesPath string, templatesConfig map[string]string) ITemplates {
	return &Templates{
		TemplatesPath:   templatesPath,
		TemplatesConfig: templatesConfig}
}

func NewTemplateFile(sourceFilePath, destFilePath string) ITemplateFile {
	localUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return &TemplateFile{
		UserHome:       localUser.HomeDir,
		SourceFilePath: sourceFilePath,
		DestFilePath:   destFilePath}
}

func (t *Templates) GetTemplateFileNames() ([]string, error) {
	fileInfos, err := ioutil.ReadDir(t.TemplatesPath)
	if err != nil {
		return nil, err
	}

	fileNames := make([]string, len(fileInfos))
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			fileNames = append(fileNames, fileInfo.Name())
		}
	}
	return fileNames, nil
}

func (t *Templates) CopyTemplates() {
	templateNames, err := t.GetTemplateFileNames()
	if err != nil {
		log.Fatal(err)
	}

	for _, templateName := range templateNames {
		if dest, ok := t.TemplatesConfig[templateName]; ok {
			NewTemplateFile(
				t.TemplatesPath+"/"+templateName,
				ReplaceWithHome(dest)).CopyTemplate()
		}
	}
}

func (t *TemplateFile) CopyTemplate() {
	log.Printf("Copying [%s] to [%s]", t.SourceFilePath, t.DestFilePath)
	temp, err := template.ParseFiles(t.SourceFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	destFile, err := os.OpenFile(t.DestFilePath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	temp.Execute(destFile, struct{ Home string }{t.UserHome})
}
