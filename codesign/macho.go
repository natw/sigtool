package codesign

import "debug/macho"

type BuiltinMachoInspector struct{}

func (i *BuiltinMachoInspector) ImportedLibraries(fname string) ([]string, error) {
	var libs []string
	var err error

	f, err := macho.Open(fname)
	if err != nil {
		return libs, err
	}
	libs, err = f.ImportedLibraries()
	return libs, err
}
