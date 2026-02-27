package service

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

type Rander struct {
	ProjectName string
	ProjectPort int
	verbose     bool
}

type RanderOption func(*Rander)

func WithVerbose(v bool) RanderOption {
	return func(r *Rander) {
		r.verbose = v
	}
}

func NewRander(projectName string, projectPort int, opts ...RanderOption) *Rander {
	r := &Rander{
		ProjectName: projectName,
		ProjectPort: projectPort,
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

func (r *Rander) log(format string, args ...interface{}) {
	if r.verbose {
		fmt.Printf(format+"\n", args...)
	}
}

func (r *Rander) InitDir() error {
	dirs := []string{"cmd", "pkg", "etc", "script"}
	for _, dir := range dirs {
		curPath := filepath.Join(r.ProjectName, dir)
		if err := os.MkdirAll(curPath, os.ModePerm); err != nil {
			return ErrorWithCode(ErrCreateDir, fmt.Errorf("path: %s, error: %w", curPath, err))
		}
		r.log("create dir: %s", curPath)
	}
	return nil
}

func (r *Rander) InitPkg() error {
	command := fmt.Sprintf("cd %s/pkg && mkdir -p {config,controller,model,router,service,util/cm}", r.ProjectName)
	if err := exec.Command("/bin/bash", "-c", command).Run(); err != nil {
		return ErrorWithCode(ErrCreatePkg, err)
	}
	r.log("pkg init success")

	tplFS := GetTemplateFS()

	templates := []templateItem{
		{"cmd/server.go.tmpl", "cmd/server.go"},
		{"config/dev.yaml.tmpl", "etc/dev.yaml"},
		{"config/config.go.tmpl", "pkg/config/config.go"},
		{"config/log.go.tmpl", "pkg/config/log.go"},
		{"config/define.go.tmpl", "pkg/config/define.go"},
		{"service/service.go.tmpl", "pkg/service/service.go"},
		{"controller/typedef.go.tmpl", "pkg/controller/typedef.go"},
		{"controller/controller.go.tmpl", "pkg/controller/controller.go"},
		{"router/url.go.tmpl", "pkg/router/url.go"},
		{"router/middleware.go.tmpl", "pkg/router/middleware.go"},
		{"util/cm/cm.go.tmpl", "pkg/util/cm/cm.go"},
		{"model/base.go.tmpl", "pkg/model/base.go"},
	}

	for _, item := range templates {
		if err := r.renderTemplate(tplFS, item.source, item.target); err != nil {
			return err
		}
	}

	return nil
}

type templateItem struct {
	source string
	target string
}

func (r *Rander) renderTemplate(tplFS fs.FS, source, target string) error {
	tplContent, err := fs.ReadFile(tplFS, source)
	if err != nil {
		return ErrorWithCode(ErrReadTemplate, fmt.Errorf("template: %s, error: %w", source, err))
	}

	rendered, err := r.GenerateFile(string(tplContent))
	if err != nil {
		return err
	}

	targetFile := filepath.Join(r.ProjectName, target)
	if err := os.WriteFile(targetFile, rendered, 0644); err != nil {
		return ErrorWithCode(ErrWriteFile, fmt.Errorf("file: %s, error: %w", targetFile, err))
	}

	r.log("render: %s -> %s", source, target)
	return nil
}

func (r *Rander) GenerateFile(tpl string) ([]byte, error) {
	tmpl, err := template.New("yolo").Parse(tpl)
	if err != nil {
		return nil, ErrorWithCode(ErrParseTemplate, err)
	}

	var tplOutput bytes.Buffer
	if err := tmpl.Execute(&tplOutput, r); err != nil {
		return nil, ErrorWithCode(ErrRenderTemplate, err)
	}
	return tplOutput.Bytes(), nil
}

func (r *Rander) RunGoMod() error {
	command := fmt.Sprintf("cd %s && go mod init %s", r.ProjectName, r.ProjectName)
	if err := exec.Command("/bin/bash", "-c", command).Run(); err != nil {
		return ErrorWithCode(ErrGoMod, fmt.Errorf("go mod init failed: %w", err))
	}
	r.log("go mod init %s", r.ProjectName)

	command = fmt.Sprintf("cd %s && go mod tidy", r.ProjectName)
	if err := exec.Command("/bin/bash", "-c", command).Run(); err != nil {
		return ErrorWithCode(ErrGoMod, fmt.Errorf("go mod tidy failed: %w", err))
	}
	r.log("go mod tidy")
	return nil
}
