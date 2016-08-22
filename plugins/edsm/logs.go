package edsm

// CommanderPosition Object representing the commander position on EDSM.
type CommanderPosition struct {
	MSGNum        int64  `json:"msgnum"`
	Message       string `json:"msg,omitempty"`
	System        string `json:"system,omitempty"`
	FirstDiscover bool   `json:"firstDiscover"`
	Date          string `json:"date,omitempty"`
}
