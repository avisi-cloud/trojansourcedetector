package trojansourcedetector

// ErrorCode is a code indicating an error.
type ErrorCode string

// ErrUnicode indicates a disallowed unicode character.
const ErrUnicode ErrorCode = "non-ASCII unicode character"

// ErrBIDI indicates a disallowed BIDI character.
const ErrBIDI ErrorCode = "bidirectional control character"

// ErrIOFile indicates a read failure on a file.
const ErrIOFile ErrorCode = "failed to read file"

// ErrIOSeek indicates a seek error in the file.
const ErrIOSeek ErrorCode = "failed to seek in file"

// ErrIODirectory indicates a failure to scan the specified directory.
const ErrIODirectory ErrorCode = "failed to scan directory"
