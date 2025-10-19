package app

// Terminal progress bar sequences without the BEL (bell) terminator.
// We use the ST (String Terminator) sequence so they don't trigger the
// system bell when refreshed frequently.
const (
	oscStringTerminator           = "\x1b\\"
	quietSetIndeterminateProgress = "\x1b]9;4;3" + oscStringTerminator
	quietResetProgress            = "\x1b]9;4;0" + oscStringTerminator
)
