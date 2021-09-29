package nftlabs

import "time"

const (
	txWaitTimeBetweenAttempts = time.Second * 1
	txMaxAttempts             = 20
)
