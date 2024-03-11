package main

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/rm4n0s/wt-video-audio-transmission-example/app/containers"
)

func validateSettings(sets containers.Settings) error {
	if sets.Username == "" {
		return errors.New("username is empty")
	}
	_, err := url.ParseRequestURI(sets.Host)
	if err != nil {
		return fmt.Errorf("failed to parse host: %w", err)
	}
	return nil
}
