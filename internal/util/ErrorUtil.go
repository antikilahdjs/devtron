/*
 * Copyright (c) 2020 Devtron Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package util

import (
	"fmt"
	"github.com/go-pg/pg"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ApiError struct {
	HttpStatusCode    int         `json:"-"`
	Code              string      `json:"code,omitempty"`
	InternalMessage   string      `json:"internalMessage,omitempty"`
	UserMessage       interface{} `json:"userMessage,omitempty"`
	UserDetailMessage string      `json:"userDetailMessage,omitempty"`
}

func (e *ApiError) Error() string {
	return e.InternalMessage
}

// default internal will be set
func (e *ApiError) ErrorfInternal(format string, a ...interface{}) error {
	return &ApiError{InternalMessage: fmt.Sprintf(format, a...)}
}

// default user message will be set
func (e ApiError) ErrorfUser(format string, a ...interface{}) error {
	return &ApiError{InternalMessage: fmt.Sprintf(format, a...)}
}

func IsErrNoRows(err error) bool {
	return pg.ErrNoRows == err
}

type grpcCodeExtended struct {
	Code codes.Code
}

func (r grpcCodeExtended) IsInvalidArgumentCode() bool {
	return r.Code == codes.InvalidArgument
}

func (r grpcCodeExtended) IsNotFoundCode() bool {
	return r.Code == codes.NotFound
}

func (r grpcCodeExtended) IsFailedPreconditionCode() bool {
	return r.Code == codes.FailedPrecondition
}

func GetGRPCErrorDetailedMessage(err error) string {
	if errStatus, ok := status.FromError(err); ok {
		return errStatus.Message()
	}
	return err.Error()
}

func GetGRPCDetailedError(err error) (grpcCodeExtended, string) {
	grpcCode := grpcCodeExtended{Code: codes.Unknown}
	if errStatus, ok := status.FromError(err); ok {
		grpcCode.Code = errStatus.Code()
		return grpcCode, errStatus.Message()
	}
	return grpcCode, err.Error()
}
