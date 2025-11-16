package dto

import "errors"

var ErrInvalidPath = errors.New("file path cannot be empty")

// Validate checks whether the external input is minimally correct.
func (in AnalyzePDFInputDTO) Validate() error {
	if in.FilePath == "" {
		return ErrInvalidPath
	}
	return nil
}
