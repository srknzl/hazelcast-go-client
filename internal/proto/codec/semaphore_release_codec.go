// Copyright (c) 2008-2020, Hazelcast, Inc. All Rights Reserved.
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
package codec

import (
	"github.com/hazelcast/hazelcast-go-client/core"
	"github.com/hazelcast/hazelcast-go-client/internal/proto"
	"github.com/hazelcast/hazelcast-go-client/internal/proto/codec/internal"
)

const (
	// hex: 0x0C0300
	SemaphoreReleaseCodecRequestMessageType = int32(787200)
	// hex: 0x0C0301
	SemaphoreReleaseCodecResponseMessageType = int32(787201)

	SemaphoreReleaseCodecRequestSessionIdOffset     = proto.PartitionIDOffset + proto.IntSizeInBytes
	SemaphoreReleaseCodecRequestThreadIdOffset      = SemaphoreReleaseCodecRequestSessionIdOffset + proto.LongSizeInBytes
	SemaphoreReleaseCodecRequestInvocationUidOffset = SemaphoreReleaseCodecRequestThreadIdOffset + proto.LongSizeInBytes
	SemaphoreReleaseCodecRequestPermitsOffset       = SemaphoreReleaseCodecRequestInvocationUidOffset + proto.UuidSizeInBytes
	SemaphoreReleaseCodecRequestInitialFrameSize    = SemaphoreReleaseCodecRequestPermitsOffset + proto.IntSizeInBytes

	SemaphoreReleaseResponseResponseOffset = proto.ResponseBackupAcksOffset + proto.ByteSizeInBytes
)

// Releases the given number of permits and increases the number of
// available permits by that amount.
type semaphoreReleaseCodec struct{}

var SemaphoreReleaseCodec semaphoreReleaseCodec

func (semaphoreReleaseCodec) EncodeRequest(groupId proto.RaftGroupId, name string, sessionId int64, threadId int64, invocationUid core.UUID, permits int32) *proto.ClientMessage {
	clientMessage := proto.NewClientMessageForEncode()
	clientMessage.SetRetryable(true)

	initialFrame := proto.NewFrame(make([]byte, SemaphoreReleaseCodecRequestInitialFrameSize))
	internal.FixSizedTypesCodec.EncodeLong(initialFrame.Content, SemaphoreReleaseCodecRequestSessionIdOffset, sessionId)
	internal.FixSizedTypesCodec.EncodeLong(initialFrame.Content, SemaphoreReleaseCodecRequestThreadIdOffset, threadId)
	internal.FixSizedTypesCodec.EncodeUUID(initialFrame.Content, SemaphoreReleaseCodecRequestInvocationUidOffset, invocationUid)
	internal.FixSizedTypesCodec.EncodeInt(initialFrame.Content, SemaphoreReleaseCodecRequestPermitsOffset, permits)
	clientMessage.AddFrame(initialFrame)
	clientMessage.SetMessageType(SemaphoreReleaseCodecRequestMessageType)
	clientMessage.SetPartitionId(-1)

	internal.RaftGroupIdCodec.Encode(clientMessage, groupId)
	internal.StringCodec.Encode(clientMessage, name)

	return clientMessage
}

func (semaphoreReleaseCodec) DecodeResponse(clientMessage *proto.ClientMessage) bool {
	frameIterator := clientMessage.FrameIterator()
	initialFrame := frameIterator.Next()

	return internal.FixSizedTypesCodec.DecodeBoolean(initialFrame.Content, SemaphoreReleaseResponseResponseOffset)
}