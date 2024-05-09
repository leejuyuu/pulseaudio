package pulseaudio

type SubscriptionMask uint32

const (
	SubscriptionMaskNull         SubscriptionMask = 0x0000
	SubscriptionMaskSink         SubscriptionMask = 0x0001
	SubscriptionMaskSource       SubscriptionMask = 0x0002
	SubscriptionMaskSinkInput    SubscriptionMask = 0x0004
	SubscriptionMaskSourceOutput SubscriptionMask = 0x0008
	SubscriptionMaskModule       SubscriptionMask = 0x0010
	SubscriptionMaskClient       SubscriptionMask = 0x0020
	SubscriptionMaskSampleCache  SubscriptionMask = 0x0040
	SubscriptionMaskServer       SubscriptionMask = 0x0080
	SubscriptionMaskAutoload     SubscriptionMask = 0x0100
	SubscriptionMaskCard         SubscriptionMask = 0x0200
	SubscriptionMaskAll          SubscriptionMask = 0x02ff
)

const (
	subscriptionEventFacilityMask = 0x000F
	subscriptionEventTypeMask     = 0x0030
)

type SubscriptionEventType uint32

const (
	SubscriptionEventNew    SubscriptionEventType = 0x0000
	SubscriptionEventChange SubscriptionEventType = 0x0010
	SubscriptionEventRemove SubscriptionEventType = 0x0020
)

type SubscriptionEventFacility uint32

const (
	SubscriptionEventSink         SubscriptionEventFacility = 0x0000
	SubscriptionEventSource       SubscriptionEventFacility = 0x0001
	SubscriptionEventSinkInput    SubscriptionEventFacility = 0x0002
	SubscriptionEventSourceOutput SubscriptionEventFacility = 0x0003
	SubscriptionEventModule       SubscriptionEventFacility = 0x0004
	SubscriptionEventClient       SubscriptionEventFacility = 0x0005
	SubscriptionEventSampleCache  SubscriptionEventFacility = 0x0006
	SubscriptionEventServer       SubscriptionEventFacility = 0x0007
	SubscriptionEventAutoload     SubscriptionEventFacility = 0x0008
	SubscriptionEventCard         SubscriptionEventFacility = 0x0009
)

type SubscriptionEvent struct {
	Type     SubscriptionEventType
	Facility SubscriptionEventFacility
	Idx      uint32
}

// Subscribe subscribes the client to the pulseaudio events provided by mask
func (c *Client) Subscribe(mask SubscriptionMask) (<-chan SubscriptionEvent, error) {
	// TODO: Can this be called twice?
	_, err := c.request(commandSubscribe, uint32Tag, mask)
	if err != nil {
		return nil, err
	}
	return c.events, nil
}
