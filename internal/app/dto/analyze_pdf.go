package dto

// AnalyzePDFInputDTO represents the external input passed
// into the AnalyzePDFUseCase. It is stable, explicit,
// and independent from HTTP or file system concerns.
type AnalyzePDFInputDTO struct {
	FilePath string
}

// AnalyzePDFOutputDTO is the structure returned by the
// use case, without exposing domain internals.
type AnalyzePDFOutputDTO struct {
	Content   string
	WordCount int
}
