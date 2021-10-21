package template

import (
	"fmt"
	"github.com/jasperstritzke/cubid/pkg/config"
	"github.com/jasperstritzke/cubid/pkg/console/color"
	"github.com/jasperstritzke/cubid/pkg/console/logger"
	"github.com/jasperstritzke/cubid/pkg/model"
	"github.com/jasperstritzke/cubid/pkg/util/fileutil"
	"io/ioutil"
	"os"
	"strings"
)

const (
	templateFolder     = "templates/"
	templateDataFile   = "template.json"
	templateExecutable = "executable.jar"
)

var templates map[string][]model.Template

func LoadTemplates() {
	templates = make(map[string][]model.Template)

	_ = os.Mkdir(templateFolder, os.ModePerm)

	logger.Info("Loading templates...")
	loadTemplatesFromFolder(templateFolder)
}

func SetEnabled(group, templateName string, state bool) (bool, error) {
	return recursiveSetEnabled(templateFolder, group, templateName, state)
}

func recursiveSetEnabled(pth, group, templateName string, state bool) (bool, error) {
	files, err := ioutil.ReadDir(pth)

	if err != nil {
		return false, err
	}

	for _, file := range files {
		if file.IsDir() {
			success, rErr := recursiveSetEnabled(pth+file.Name()+"/", group, templateName, state)
			if rErr != nil {
				return false, rErr
			}

			if success {
				return success, nil
			}
		}

		if file.Name() == templateDataFile {
			if fileutil.ExistsFile(pth+file.Name()) && strings.Contains(pth, "/"+group+"/"+templateName+"/") {
				rErr := setEnabledLiteral(pth+file.Name(), state)
				return rErr == nil, rErr
			}
		}
	}

	return false, nil
}

func setEnabledLiteral(dataFilePath string, state bool) error {
	var scopeTemplate model.Template
	err := config.LoadConfig(dataFilePath, &scopeTemplate)

	if err != nil {
		logger.Error("Template not found. (if folder structure is deeply nested, it's required to toggle manually)")
		return err
	}

	scopeTemplate.Enabled = state

	err = config.WriteConfig(dataFilePath, &scopeTemplate)
	if err != nil {
		return err
	}

	if state {
		logger.Info("Template was enabled.")
	} else {
		logger.Warn("Template was disabled.")
	}

	return nil
}

func loadTemplatesFromFolder(pth string) {
	files, err := ioutil.ReadDir(pth)

	if err != nil {
		logger.Error("Unable to load templates: " + err.Error())
		os.Exit(1)
	}

	for _, file := range files {
		if file.IsDir() {
			if fileutil.ExistsFile(pth + file.Name() + "/" + templateDataFile) {
				loadTemplate(file.Name(), pth+file.Name()+"/")
				continue
			}

			loadTemplatesFromFolder(pth + file.Name() + "/")
		}
	}
}

func loadTemplate(folderName, pth string) {
	dataFile := pth + templateDataFile

	var template model.Template
	err := config.LoadConfig(dataFile, &template)

	var groupName = folderName
	pathSplit := strings.Split(pth, "/")

	if len(pathSplit) >= 2 {
		groupName = pathSplit[len(pathSplit)-3]
	}

	template.Group = groupName

	if !template.Enabled {
		fmt.Println(color.Yellow + "(i) Skipping template " + template.Group + "/" + template.Name + " because it's disabled.")
		return
	}

	if err != nil {
		logger.Error("Unable to load template at " + pth + ". Please re-initialize template or delete it.")
		return
	}

	if !fileutil.ExistsFile(pth + templateExecutable) {
		if len(template.Version.BuildURL) == 0 {
			logger.Error("Unable to load template " + pth + ": Executable with name " + templateExecutable + " not found and no build URL exists.")
			logger.Warn("Please move a executable jar file with the name " + templateExecutable + " into the template.")
			return
		}

		if len(template.Version.BuildURL) > 0 {
			logger.Warn("Downloading missing executable for template " + pth + "...")
			err = template.Version.DownloadTo(pth + templateExecutable)

			if err != nil {
				logger.Error("Unable to download version " + template.Version.Display + " for template " + pth + ".")
				logger.Error("Error: " + err.Error())
				return
			}
		}
	}

	logger.Info("Successfully loaded template " + groupName + "/" + template.Name + ".")
	templates[template.Group] = append(templates[template.Group], template)
}

func CreateTemplate(templateGroup, name string, version model.VersionValue) error {
	template := model.Template{
		Name:    name,
		Version: version,
		Enabled: true,
	}

	path := templateFolder + "/" + templateGroup + "/" + name + "/"
	_ = os.MkdirAll(path, os.ModePerm)

	file, err := os.Create(path + templateDataFile)

	if err != nil {
		return err
	}

	encoder := fileutil.NewPrettyEncoder(file)

	err = encoder.Encode(template)
	if err != nil {
		return err
	}

	err = file.Close()

	logger.Info("Template " + templateGroup + "/" + name + " created.")

	return err
}
