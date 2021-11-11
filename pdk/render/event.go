package render

import treevents "github.com/negrel/paon/pdk/tree/events"

// Event define a render event that is triggered when an object need rendering.
type Event struct {
	treevents.Event

	Context Context
}
