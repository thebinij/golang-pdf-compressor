// shrink-utils.go

package pdfcompresser

import (
	"os/exec"
	"path/filepath"
)

func executeShrinkPDF(scriptDir, dpi string, grayscale bool, inputFilePath string) error {
	cmdArgs := []string{filepath.Join(scriptDir, "shrinkpdf.sh")}

	if grayscale {
		cmdArgs = append(cmdArgs, "-g")
	}

	if dpi != "" {
		cmdArgs = append(cmdArgs, "-r", dpi)
	}

	cmdArgs = append(cmdArgs, "-o", compressedPDFName, inputFilePath)

	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)
	cmd.Dir = scriptDir

	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
