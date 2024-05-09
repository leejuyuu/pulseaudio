package pulseaudio

import "fmt"

// Updates returns a channel with PulseAudio updates.
func (c *Client) Updates() (updates <-chan struct{}, err error) {
	events, err := c.Subscribe(SubscriptionMaskAll)
	if err != nil {
		return nil, err
	}

	u := make(chan struct{})
	go func() {
		defer close(u)
		for e := range events {
			fmt.Println(e)
			u <- struct{}{}
		}
	}()

	return u, nil
}
