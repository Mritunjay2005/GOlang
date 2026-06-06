package watchdog

import "time"

type config struct {
    debounce  time.Duration
    recursive bool
    ignore    []string
}

type Option func(*config)

func WithDebounce(d time.Duration) Option {
    return func(c *config) { c.debounce = d }
}

func WithRecursive(r bool) Option {
    return func(c *config) { c.recursive = r }
}

func WithIgnore(patterns ...string) Option {
    return func(c *config) { c.ignore = append(c.ignore, patterns...) }
}