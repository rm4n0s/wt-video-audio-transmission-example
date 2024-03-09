package controllers

import (
	"context"
	"fmt"

	"github.com/rm4n0s/wt-video-audio-transmission-example/app/vpclient"
)

func connectToVoip(ctx context.Context, host, username string) (input, ouput *vpclient.VoipClient, err error) {
	ctx, cancel := context.WithCancel(ctx)
	inputClient := vpclient.NewVoipClient()
	outputClient := vpclient.NewVoipClient()

	err = inputClient.Connect(ctx, host+"/input?username="+username)
	if err != nil {
		cancel()
		err = fmt.Errorf("failed to connect for input: %w", err)
		return nil, nil, err
	}

	err = outputClient.Connect(ctx, host+"/output?username="+username)
	if err != nil {
		cancel()
		err = fmt.Errorf("failed to connect for output: %w", err)
		return nil, nil, err
	}
	return inputClient, outputClient, nil
}
