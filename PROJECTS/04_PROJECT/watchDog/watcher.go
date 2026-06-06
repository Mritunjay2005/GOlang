package watchdog

import (
    "sync"
    "time"

    "github.com/fsnotify/fsnotify"
)

// HandlerFunc is the function signature for all event handlers.
type HandlerFunc func(Event)

type handler struct {
    pattern string
    fn      HandlerFunc
}

// Watcher watches directories and files for changes,
// routing events to registered handlers.
type Watcher struct {
    cfg      config
    fsw      *fsnotify.Watcher
    handlers map[Op][]handler
    done     chan struct{}
    mu       sync.RWMutex
}

// New creates a Watcher with the given options.
// If no options are provided, sensible defaults are used.
func New(opts ...Option) *Watcher {
    cfg := config{
        debounce:  100 * time.Millisecond,
        recursive: true,
    }
    for _, opt := range opts {
        opt(&cfg)
    }
    return &Watcher{
        cfg:      cfg,
        handlers: make(map[Op][]handler),
        done:     make(chan struct{}),
    }
}

// On registers a handler function called when a file matching
// pattern is changed with the given Op.
func (w *Watcher) On(op Op, pattern string, fn HandlerFunc) error {
    w.mu.Lock()
    defer w.mu.Unlock()
    w.handlers[op] = append(w.handlers[op], handler{pattern: pattern, fn: fn})
    return nil
}

// Start begins watching the given paths.
// Watching runs in the background — Start returns immediately.
func (w *Watcher) Start(paths ...string) error { return nil }

// Stop shuts down the watcher and releases all resources.
func (w *Watcher) Stop() {}