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
	"github.com/hazelcast/hazelcast-go-client/internal/proto"
	"github.com/hazelcast/hazelcast-go-client/internal/proto/codec/internal"
)

const (
	// hex: 0x090300
	AtomicLongAddAndGetCodecRequestMessageType = int32(590592)
	// hex: 0x090301
	AtomicLongAddAndGetCodecResponseMessageType = int32(590593)

	AtomicLongAddAndGetCodecRequestDeltaOffset      = proto.PartitionIDOffset + proto.IntSizeInBytes
	AtomicLongAddAndGetCodecRequestInitialFrameSize = AtomicLongAddAndGetCodecRequestDeltaOffset + proto.LongSizeInBytes

	AtomicLongAddAndGetResponseResponseOffset = proto.ResponseBackupAcksOffset + proto.ByteSizeInBytes
)

// Atomically adds the given value to the current value.
type atomiclongAddAndGetCodec struct{}

var AtomicLongAddAndGetCodec atomiclongAddAndGetCodec

func (atomiclongAddAndGetCodec) EncodeRequest(groupId proto.RaftGroupId, name string, delta int64) *proto.ClientMessage {
	clientMessage := proto.NewClientMessageForEncode()
	clientMessage.SetRetryable(false)

	initialFrame := proto.NewFrame(make([]byte, AtomicLongAddAndGetCodecRequestInitialFrameSize))
	internal.FixSizedTypesCodec.EncodeLong(initialFrame.Content, AtomicLongAddAndGetCodecRequestDeltaOffset, delta)
	clientMessage.AddFrame(initialFrame)
	clientMessage.SetMessageType(AtomicLongAddAndGetCodecRequestMessageType)
	clientMessage.SetPartitionId(-1)

	internal.RaftGroupIdCodec.Encode(clientMessage, groupId)
	internal.StringCodec.Encode(clientMessage, name)

	return clientMessage
}

func (atomiclongAddAndGetCodec) DecodeResponse(clientMessage *proto.ClientMessage) int64 {
	frameIterator := clientMessage.FrameIterator()
	initialFrame := frameIterator.Next()

	return internal.FixSizedTypesCodec.DecodeLong(initialFrame.Content, AtomicLongAddAndGetResponseResponseOffset)
}