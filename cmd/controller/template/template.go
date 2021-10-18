package template

import (
	"encoding/json"
	"github.com/jasperstritzke/cubid/pkg/config"
	"github.com/jasperstritzke/cubid/pkg/console/logger"
	"github.com/jasperstritzke/cubid/pkg/model"
	"github.com/jasperstritzke/cubid/pkg/util/fileutil"
	"io/ioutil"
	"os"
)

const (
	templateFolder     = "templates/"
	templateDataFile   = "template.json"
	templateExecutable = "server.jar"
)

var templates []model.Template

func LoadTemplates() {
	templates = nil

	_ = os.Mkdir(templateFolder, os.ModePerm)

	loadTemplatesFromFolder(templateFolder)

	if len(templates) == 0 {
		logger.Warn("No templates found. Creating default templates...")

		err := CreateTemplate("Proxy", "default", true, model.Version.Waterfall)
		if err != nil {
			logger.Error("Error creating default proxy group: " + err.Error())
			return
		}

		err = CreateTemplate("Lobby", "default", true, model.Version.Paper17)
		if err != nil {
			logger.Error("Error creating default lobby group: " + err.Error())
			return
		}
	}
}

func loadTemplatesFromFolder(pth string) {
	files, err := ioutil.ReadDir(templateFolder)

	if err != nil {
		logger.Error("Unable to load templates: " + err.Error())
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			if fileutil.ExistsFile(pth + templateDataFile) {
				loadTemplate(file.Name(), pth+file.Name())
				continue
			}

			loadTemplatesFromFolder(pth + file.Name() + "/")
		}
	}
}

func loadTemplate(folderName, pth string) {
	logger.Info("Loading template " + pth + "...")

	dataFile := pth + templateDataFile

	var template model.Template
	err := config.LoadConfig(dataFile, &template)

	if err != nil {
		logger.Error("Unable to load template at " + pth + ". Please re-initialize template or delete it.")
		return
	}

	if !fileutil.ExistsFile(pth + templateExecutable) {
		if len(template.Version.Display) == 0 {
			logger.Error("Unable to load template " + pth + ": Executable with name " + templateExecutable + " not found.")
			logger.Warn("Please move a executable jar file with the name " + templateExecutable + " into the template.")
			return
		}

		logger.Warn("Downloading missing executable for template " + pth + "...")
		err = template.Version.DownloadTo(pth + templateExecutable)

		if err != nil {
			logger.Error("Unable to download version " + template.Version.Display + " for template " + pth + ".")
			return
		}
	}

	template.Group = folderName

	logger.Info("Successfully loaded template " + pth + ".")
	templates = append(templates, template)
}

func CreateTemplate(templateGroup, name string, proxy bool, version model.VersionValue) error {
	template := model.Template{
		Name:    name,
		Proxy:   proxy,
		Version: version,
	}

	path := templateFolder + "/" + templateGroup + "/" + name + "/"
	_ = os.MkdirAll(path, os.ModePerm)

	file, err := os.Create(path + templateDataFile)

	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)

	err = encoder.Encode(template)
	if err != nil {
		return err
	}

	err = file.Close()

	logger.Info("Template " + templateGroup + "/" + name + " created.")

	return err
}
