package plaintext

import (
	"bufio"
	"io"
	"strings"

	"github.com/Tiiwoo/geoip/lib"
)

type textIn struct {
	Type        string
	Action      lib.Action
	Description string
	Name        string
	URI         string
	InputDir    string
	OnlyIPType  lib.IPType
}

func (t *textIn) scanFile(reader io.Reader, entry *lib.Entry) error {
	var err error
	switch t.Type {
	case typeTextIn:
		err = t.scanFileForTextIn(reader, entry)
	default:
		return lib.ErrNotSupportedFormat
	}

	return err
}

func (t *textIn) scanFileForTextIn(reader io.Reader, entry *lib.Entry) error {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		if err := entry.AddPrefix(line); err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
