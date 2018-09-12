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

package server

import (
	"github.com/gofrs/uuid"
	"github.com/heroiclabs/nakama/api"
	"go.opencensus.io/stats"
	"go.opencensus.io/tag"
	"go.opencensus.io/trace"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func (s *ApiServer) GetUsers(ctx context.Context, in *api.GetUsersRequest) (*api.Users, error) {
	if in.GetIds() == nil && in.GetUsernames() == nil && in.GetFacebookIds() == nil {
		return &api.Users{}, nil
	}

	ids := make([]string, 0)
	usernames := make([]string, 0)
	facebookIDs := make([]string, 0)

	if in.GetIds() != nil {
		for _, id := range in.GetIds() {
			if _, uuidErr := uuid.FromString(id); uuidErr != nil {
				return nil, status.Error(codes.InvalidArgument, "ID '"+id+"' is not a valid system ID.")
			}

			ids = append(ids, id)
		}
	}

	if in.GetUsernames() != nil {
		usernames = in.GetUsernames()
	}

	if in.GetFacebookIds() != nil {
		facebookIDs = in.GetFacebookIds()
	}

	// Before hook.
	if fn := s.runtime.beforeReqFunctions.beforeGetUsersFunction; fn != nil {
		// Stats measurement start boundary.
		name := "nakama.api-before.Nakama.GetUsers"
		statsCtx, _ := tag.New(context.Background(), tag.Upsert(MetricsFunction, name))
		startNanos := time.Now().UTC().UnixNano()
		span := trace.NewSpan(name, nil, trace.StartOptions{})

		// Extract request information and execute the hook.
		clientIP, clientPort := extractClientAddress(s.logger, ctx)
		result, err, code := fn(s.logger, ctx.Value(ctxUserIDKey{}).(uuid.UUID).String(), ctx.Value(ctxUsernameKey{}).(string), ctx.Value(ctxExpiryKey{}).(int64), clientIP, clientPort, in)
		if err != nil {
			return nil, status.Error(code, err.Error())
		}
		if result == nil {
			return nil, status.Error(codes.Internal, "Runtime Before hook returned no result.")
		}
		in = result

		// Stats measurement end boundary.
		span.End()
		stats.Record(statsCtx, MetricsApiTimeSpentMsec.M(float64(time.Now().UTC().UnixNano()-startNanos)/1000), MetricsApiCount.M(1))
	}

	users, err := GetUsers(s.logger, s.db, s.tracker, ids, usernames, facebookIDs)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error retrieving user accounts.")
	}

	// After hook.
	if fn := s.runtime.afterReqFunctions.afterGetUsersFunction; fn != nil {
		// Stats measurement start boundary.
		name := "nakama.api-after.Nakama.GetUsers"
		statsCtx, _ := tag.New(context.Background(), tag.Upsert(MetricsFunction, name))
		startNanos := time.Now().UTC().UnixNano()
		span := trace.NewSpan(name, nil, trace.StartOptions{})

		// Extract request information and execute the hook.
		clientIP, clientPort := extractClientAddress(s.logger, ctx)
		fn(s.logger, ctx.Value(ctxUserIDKey{}).(uuid.UUID).String(), ctx.Value(ctxUsernameKey{}).(string), ctx.Value(ctxExpiryKey{}).(int64), clientIP, clientPort, users)

		// Stats measurement end boundary.
		span.End()
		stats.Record(statsCtx, MetricsApiTimeSpentMsec.M(float64(time.Now().UTC().UnixNano()-startNanos)/1000), MetricsApiCount.M(1))
	}

	return users, nil
}
