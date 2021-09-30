package nftlabs

import "time"

// TODO: Make configurable at sdk level/env variables

const (
	txWaitTimeBetweenAttempts = time.Second * 1
	txMaxAttempts             = 20
)
