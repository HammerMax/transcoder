package transcoder

import (
	"context"
	"io"
)

type FileOption struct {
	Filename string
	Option Options
}

// Transcoder ...
type Transcoder interface {
	Start(opts Options) (<-chan Progress, error)
	Input(f *FileOption) Transcoder
	InputPipe(w *io.WriteCloser, r *io.ReadCloser) Transcoder
	Output(f *FileOption) Transcoder
	OutputPipe(w *io.WriteCloser, r *io.ReadCloser) Transcoder
	WithOptions(opts Options) Transcoder
	WithAdditionalOptions(opts Options) Transcoder
	WithContext(ctx *context.Context) Transcoder
	GetMetadata() (Metadata, error)
}
