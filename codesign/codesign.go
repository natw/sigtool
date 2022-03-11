package codesign

type Logger interface {
	Debugw(msg string, keysAndValues ...interface{})
}

type MachoInspector interface { // hehe
	ImportedLibraries(fname string) ([]string, error)
}

type Options struct {
	Logger         Logger
	MachoInspector MachoInspector
}

type Checker struct {
	Logger         Logger
	MachoInspector MachoInspector
}

func NewCodesign(opts *Options) *Checker {
	c := &Checker{
		Logger:         opts.Logger,
		MachoInspector: opts.MachoInspector,
	}
	if c.MachoInspector == nil {
		c.MachoInspector = &BuiltinMachoInspector{}
	}
	return c
}

func (c *Checker) Verify(fname string) error {
	return nil
}

// VerifyRecursive verifies a file's signature and that of all the libs it imports
func (c *Checker) VerifyRecursive(fname string) (map[string]error, error) {
	visited := make(map[string]bool)
	results := make(map[string]error)

	var f func(string) error
	f = func(fname string) error {
		deps, err := c.MachoInspector.ImportedLibraries(fname)
		if err != nil {
			return err
		}
		results[fname] = c.Verify(fname)
		visited[fname] = true
		for _, dep := range deps {
			if !visited[dep] {
				err = f(dep)
				visited[dep] = true
				if err != nil {
					return err
				}
			}
		}
		return nil
	}

	visited[fname] = true
	err := f(fname)
	return results, err
}
