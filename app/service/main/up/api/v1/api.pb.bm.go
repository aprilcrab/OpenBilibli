// Package v1 Code generated by go-common/app/tool/protoc-gen-bm. DO NOT EDIT.
package v1

import (
	"bytes"
	"context"
	"encoding/json"

	"go-common/app/tool/protoc-gen-bm/jsonpb"

	bm "go-common/library/net/http/blademaster"
)

// BMUpServer interface as same as gGRPC server define
type BMUpServer interface {
	UpArcs(context.Context, *UpArcsReq) (*UpArcsReply, error)
	UpsArcs(context.Context, *UpsArcsReq) (*UpsArcsReply, error)
	UpCount(context.Context, *UpCountReq) (*UpCountReply, error)
	UpsCount(context.Context, *UpsCountReq) (*UpsCountReply, error)
}

// _BMServerUpserver
type _BMServerUp struct {
	BMUpServer
}

func (b *_BMServerUp) bmUpUpArcsHandler(c *bm.Context) {
	req := new(UpArcsReq)
	if err := c.Bind(req); err != nil {
		return
	}
	reply, err := b.UpArcs(c.Context, req)
	if err != nil {
		c.JSON(nil, err)
		return
	}
	body := &bytes.Buffer{}
	marshaler := jsonpb.Marshaler{EmitDefaults: true, OrigName: true}
	err = marshaler.Marshal(body, reply)
	c.JSON(json.RawMessage(body.Bytes()), err)
}

func (b *_BMServerUp) bmUpUpsArcsHandler(c *bm.Context) {
	req := new(UpsArcsReq)
	if err := c.Bind(req); err != nil {
		return
	}
	reply, err := b.UpsArcs(c.Context, req)
	if err != nil {
		c.JSON(nil, err)
		return
	}
	body := &bytes.Buffer{}
	marshaler := jsonpb.Marshaler{EmitDefaults: true, OrigName: true}
	err = marshaler.Marshal(body, reply)
	c.JSON(json.RawMessage(body.Bytes()), err)
}

func (b *_BMServerUp) bmUpUpCountHandler(c *bm.Context) {
	req := new(UpCountReq)
	if err := c.Bind(req); err != nil {
		return
	}
	reply, err := b.UpCount(c.Context, req)
	if err != nil {
		c.JSON(nil, err)
		return
	}
	body := &bytes.Buffer{}
	marshaler := jsonpb.Marshaler{EmitDefaults: true, OrigName: true}
	err = marshaler.Marshal(body, reply)
	c.JSON(json.RawMessage(body.Bytes()), err)
}

func (b *_BMServerUp) bmUpUpsCountHandler(c *bm.Context) {
	req := new(UpsCountReq)
	if err := c.Bind(req); err != nil {
		return
	}
	reply, err := b.UpsCount(c.Context, req)
	if err != nil {
		c.JSON(nil, err)
		return
	}
	body := &bytes.Buffer{}
	marshaler := jsonpb.Marshaler{EmitDefaults: true, OrigName: true}
	err = marshaler.Marshal(body, reply)
	c.JSON(json.RawMessage(body.Bytes()), err)
}

// RegisterUpBMServer register bm server
func RegisterUpBMServer(e *bm.Engine, s BMUpServer) {
	bs := &_BMServerUp{s}
	e.GET("/x/internal/uper/archive/up/passed", bs.bmUpUpArcsHandler)
	e.GET("/x/internal/uper/archive/ups/passed", bs.bmUpUpsArcsHandler)
	e.GET("/x/internal/uper/archive/up/count", bs.bmUpUpCountHandler)
	e.GET("/x/internal/uper/archive/ups/count", bs.bmUpUpsCountHandler)
}