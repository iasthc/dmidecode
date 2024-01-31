package tpm_test

import (
	"testing"

	"github.com/iasthc/dmidecode/parser/tpm"
	"github.com/iasthc/dmidecode/smbios"
	"github.com/stretchr/testify/assert"
)

var (
	s = &smbios.Structure{
		Header: smbios.Header{
			Type:   0x2b,
			Length: 0x1F,
			Handle: 0x3C,
		},
		Formatted: []byte{0x00, 0x58, 0x46, 0x49, 0x02, 0x00, 0x3E, 0x00, 0x05, 0x00,
			0x00, 0x36, 0x0C, 0x00, 0x02, 0x10, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
			0x00, 0x00, 0x00, 0x54, 0x50, 0x4D, 0x20, 0x32, 0x2E, 0x30, 0x00, 0x49, 0x4E, 0x46,
			0x49, 0x4E, 0x45, 0x4F, 0x4E, 0x00, 0x00},
		Strings: []string{"0XFI", "INFINEON"},
	}
)

func TestParse(t *testing.T) {
	tpm, err := tpm.Parse(s)
	if assert.NoError(t, err) {
		t.Log(tpm)
	}
}
