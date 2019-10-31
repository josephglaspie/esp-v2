// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package components

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/GoogleCloudPlatform/api-proxy/tests/env/platform"
)

type GrpcInteropGrpcServer struct {
	*Cmd
}

func NewGrpcInteropGrpcServer(port uint16) (*GrpcInteropGrpcServer, error) {
	cmd := exec.Command(platform.GetFilePath(platform.GrpcInteropServer), fmt.Sprintf("--port=%v", port))
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return &GrpcInteropGrpcServer{
		Cmd: &Cmd{
			name: "grpcInteropGRPCServer",
			Cmd:  cmd,
		},
	}, nil
}
