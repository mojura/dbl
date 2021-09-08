package mojura

import (
	"time"

	"github.com/hatchify/errors"
	"github.com/mojura-backends/bolt"
	"github.com/mojura/backend"
	"github.com/mojura/kiroku"
)

const (
	// DefaultMaxBatchCalls is the default maximum number of calls a batch will take
	DefaultMaxBatchCalls = 1024
	// DefaultMaxBatchDuration is the default maximum duration a batch will take to collect calls
	DefaultMaxBatchDuration = time.Millisecond * 10
	// DefaultRetryBatchFail is the default value for if a batch call will retry when a batch sibling fails
	DefaultRetryBatchFail = true
	// DefaultIndexLength is the default index length
	DefaultIndexLength = 8
)

const (
	// ErrEmptyEncoder is returned when an encoder is unset
	ErrEmptyEncoder = errors.Error("invalid encoder, cannot be empty")
)

var defaultOpts = Opts{
	MaxBatchCalls:    DefaultMaxBatchCalls,
	MaxBatchDuration: DefaultMaxBatchDuration,
	RetryBatchFail:   DefaultRetryBatchFail,

	Initializer: bolt.New(),
	Encoder:     &JSONEncoder{},
}

// Opts represent mojura options
type Opts struct {
	Initializer backend.Initializer
	Encoder     Encoder

	Importer func(kiroku.Processor)
	Exporter kiroku.Exporter

	IndexLength int

	MaxBatchCalls    int
	MaxBatchDuration time.Duration
	RetryBatchFail   bool

	kiroku.Options
}

// Validate will validate a set of Options
func (o *Opts) Validate() (err error) {
	o.init()
	return
}

func (o *Opts) init() {
	if o.Encoder == nil {
		o.Encoder = defaultOpts.Encoder
	}

	if o.Initializer == nil {
		o.Initializer = defaultOpts.Initializer
	}

	if o.MaxBatchCalls == 0 {
		o.MaxBatchCalls = DefaultMaxBatchCalls
	}

	if o.MaxBatchDuration == 0 {
		o.MaxBatchDuration = DefaultMaxBatchDuration
	}

	if o.IndexLength == 0 {
		o.IndexLength = DefaultIndexLength
	}
}
