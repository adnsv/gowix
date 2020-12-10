package wix

import (
	"encoding/xml"
	"os"
	"os/exec"
	"path/filepath"
)

// Builder produces MSI files
type Builder struct {
	WorkDir       string
	WixBinaryPath string // path to wix binary dir
	WsxFile       string
	WixobjFile    string
	CandleArgs    []string
	LightArgs     []string
}

// NewBuilder constructs a new builder instance
func NewBuilder(workdir string, fnbase string) *Builder {
	wbp := os.Getenv("WIX")
	if wbp != "" {
		wbp = filepath.Join(wbp, "bin")
	}
	ret := &Builder{WorkDir: workdir, WixBinaryPath: wbp}
	ret.WsxFile = fnbase + ".wsx"
	ret.WixobjFile = fnbase + ".wixobj"
	return ret
}

// AddCandleArgs can be used to add flags and arguments for Wix Candle
func (b *Builder) AddCandleArgs(args ...string) {
	b.CandleArgs = append(b.CandleArgs, args...)
}

// AddLightArgs can be used to add flags and arguments for Wix Light
func (b *Builder) AddLightArgs(args ...string) {
	b.LightArgs = append(b.LightArgs, args...)
}

// Run builds the MSI package
func (b *Builder) Run(w *Wix) error {
	j, err := xml.MarshalIndent(w, "", "  ")
	if err != nil {
		return err
	}
	err = WriteFileIfChanged(filepath.Join(b.WorkDir, b.WsxFile), j)
	if err != nil {
		return err
	}

	candle := exec.Command(filepath.Join(b.WixBinaryPath, "candle.exe"), append(b.CandleArgs, b.WsxFile)...)
	candle.Dir = b.WorkDir
	candle.Stdout = os.Stdout
	candle.Stderr = os.Stderr
	err = candle.Run()
	if err != nil {
		return err
	}

	light := exec.Command(filepath.Join(b.WixBinaryPath, "light.exe"), append(b.LightArgs, b.WixobjFile)...)
	light.Dir = b.WorkDir
	light.Stdout = os.Stdout
	light.Stderr = os.Stderr
	err = light.Run()
	if err != nil {
		return err
	}

	return nil
}
