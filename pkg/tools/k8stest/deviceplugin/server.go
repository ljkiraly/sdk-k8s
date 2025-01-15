// Copyright (c) 2020-2022 Doc.ai and/or its affiliates.
//
// Copyright (c) 2023 Cisco and/or its affiliates.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !windows
// +build !windows

// Package deviceplugin provides device plugin GRPC API stubs for testing
package deviceplugin

import (
	"context"
	"path"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"

	"github.com/ljkiraly/sdk/pkg/tools/grpcutils"
	"github.com/ljkiraly/sdk/pkg/tools/log"
	"github.com/pkg/errors"

	"github.com/ljkiraly/sdk-k8s/pkg/tools/socketpath"
)

type registrationServer struct {
	devicePluginPath string
}

// StartRegistrationServer creates a new registrationServer and registers it on given GRPC server
func StartRegistrationServer(devicePluginPath string, server *grpc.Server) {
	pluginapi.RegisterRegistrationServer(server, &registrationServer{
		devicePluginPath: devicePluginPath,
	})
}

func (rs *registrationServer) Register(ctx context.Context, request *pluginapi.RegisterRequest) (*pluginapi.Empty, error) {
	logger := log.FromContext(ctx).WithField("registrationServer", "Register")

	socketPath := socketpath.SocketPath(path.Join(rs.devicePluginPath, request.Endpoint))
	socketURL := grpcutils.AddressToURL(socketPath)
	conn, err := grpc.DialContext(ctx, socketURL.String(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Errorf("failed to connect to %v", socketPath.String())
		return nil, errors.Wrapf(err, "failed to connect to %v", socketPath.String())
	}

	client := pluginapi.NewDevicePluginClient(conn)
	receiver, err := client.ListAndWatch(ctx, &pluginapi.Empty{})
	if err != nil {
		logger.Errorf("failed to ListAndWatch on %v", socketPath.String())
		return nil, errors.Wrapf(err, "failed to ListAndWatch on %v", socketPath.String())
	}

	go func() {
		logger.Info("client started")
		for {
			response, err := receiver.Recv()
			if err != nil {
				logger.Infof("client closed: %+v", err)
				return
			}
			logger.Infof("devices update -> %v", response.Devices)
		}
	}()

	return &pluginapi.Empty{}, nil
}
