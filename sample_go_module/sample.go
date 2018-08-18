// Copyright 2018 The Nakama Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"database/sql"
	"github.com/heroiclabs/nakama/rtapi"
	"github.com/heroiclabs/nakama/runtime"
	"log"
)

func InitModule(ctx context.Context, logger *log.Logger, db *sql.DB, nk runtime.NakamaModule, initialiser runtime.Initialiser) {
	initialiser.RegisterRpc("go_echo_sample", rpcEcho)
	initialiser.RegisterBeforeRt("ChannelJoin", beforeChannelJoin)
}

func rpcEcho(ctx context.Context, logger *log.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error, int) {
	logger.Print("RUNNING IN GO")
	return payload, nil, 0
}

func beforeChannelJoin(ctx context.Context, logger *log.Logger, db *sql.DB, nk runtime.NakamaModule, envelope *rtapi.Envelope) (*rtapi.Envelope, error) {
	logger.Printf("Intercepted request to join channel '%v'", envelope.GetChannelJoin().Target)
	return envelope, nil
}
