package jazz

import (
	"context"
	"sync"
)

// QMCache will request every object only once
type QMCache[T QMObject] struct {
	project *QMProject
	m       sync.Map
}

// NewQMCache creates a new cache object for the given project
func NewQMCache[T QMObject](project *QMProject) *QMCache[T] {
	return &QMCache[T]{
		project: project,
	}
}

// Get element from cache or query from server
func (c *QMCache[T]) Get(ctx context.Context, ref QMRef) (T, error) {
	if entry, ok := c.m.Load(ref.Href); ok {
		return entry.(T), nil
	}

	entry, err := QMGet[T](ctx, c.project, ref.Href)
	if err != nil {
		var value T
		return value, err
	}

	c.m.Store(ref.Href, entry)
	return entry, nil
}
