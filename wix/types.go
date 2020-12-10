package wix

import (
	"encoding/json"
	"errors"
)

// YesNoType implements Wix.YesNoType
type YesNoType int

// Supported YesNoType values
const (
	YesNoUnspecified = YesNoType(iota)
	Yes
	No
)

// InstallUninstallType implements Wix.InstallUninstallType
type InstallUninstallType int

// Supported InstallUninstallType values
const (
	InstallUninstallUnspecified = InstallUninstallType(iota) // unspecified
	InstallOnly                                              // install-only action
	UninstallOnly                                            // uninstall-only action
	InstallAndUninstall                                      // action that occurs both when installing and uninstalling
)

// ErrInvalidYesNoValue is an error that indicates invalid YesNoType value
var ErrInvalidYesNoValue = errors.New("invalid InstallUninstallType value")

// ToYesNo converts true->Yes and false->No
func ToYesNo(v bool) YesNoType {
	if v {
		return Yes
	}
	return No
}

func (v YesNoType) toString() (string, error) {
	switch v {
	case YesNoUnspecified:
		return "unspecified", nil
	case Yes:
		return "yes", nil
	case No:
		return "no", nil
	default:
		return "<invalid>", ErrInvalidInstallUninstall
	}
}

// MarshalJSON is used when writing YesNoType into a JSON stream
func (v *YesNoType) MarshalJSON() ([]byte, error) {
	s, err := v.toString()
	if err != nil {
		return nil, err
	}
	return json.Marshal(s)
}

// MarshalText implements TextMarshaler interface for YesNoType
func (v *YesNoType) MarshalText() ([]byte, error) {
	s, err := v.toString()
	return []byte(s), err
}

// ErrInvalidInstallUninstall is an error that indicates invalid InstallUninstallType value
var ErrInvalidInstallUninstall = errors.New("invalid InstallUninstallType value")

func (v InstallUninstallType) toString() (string, error) {
	switch v {
	case InstallUninstallUnspecified:
		return "unspecified", nil
	case InstallOnly:
		return "install", nil
	case UninstallOnly:
		return "uninstall", nil
	case InstallAndUninstall:
		return "both", nil
	default:
		return "<invalid>", ErrInvalidInstallUninstall
	}
}

// MarshalJSON is used when writing InstallUninstallType into a JSON stream
func (v *InstallUninstallType) MarshalJSON() ([]byte, error) {
	s, err := v.toString()
	if err != nil {
		return nil, err
	}
	return json.Marshal(s)
}

// MarshalText implements TextMarshaler interface for InstallUninstallType
func (v *InstallUninstallType) MarshalText() ([]byte, error) {
	s, err := v.toString()
	return []byte(s), err
}
