package business

import (
	_ "embed"
	"fmt"
	"os"
	"path"
	"text/template"

	log "github.com/sirupsen/logrus"
)

func (cba *CBA) RenderingReport(templateFile string, reportFile string) error {
	log.Debugf("rendering template ( %s )", templateFile)
	name := path.Base(templateFile)
	t, err := template.New(name).ParseFiles(templateFile)
	if err != nil {
		return fmt.Errorf("failed to parse template ( %s ) - %v", templateFile, err)
	}
	f, err := os.Create(reportFile)
	if err != nil {
		return fmt.Errorf("failed to create target report file ( %s ) - %v", reportFile, err)
	}
	defer f.Close()
	return t.Execute(f, cba.Report)
}
