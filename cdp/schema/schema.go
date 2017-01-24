// Package schema provides the Chrome Debugging Protocol
// commands, types, and events for the Chrome Schema domain.
//
// Provides information about the protocol schema.
//
// Generated by the chromedp-gen command.
package schema

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	. "github.com/knq/chromedp/cdp"
	"github.com/mailru/easyjson"
)

var (
	_ BackendNode
	_ BackendNodeID
	_ ComputedProperty
	_ ErrorType
	_ Frame
	_ FrameID
	_ LoaderID
	_ Message
	_ MessageError
	_ MethodType
	_ Node
	_ NodeID
	_ NodeType
	_ PseudoType
	_ RGBA
	_ ShadowRootType
	_ Timestamp
)

// GetDomainsParams returns supported domains.
type GetDomainsParams struct{}

// GetDomains returns supported domains.
func GetDomains() *GetDomainsParams {
	return &GetDomainsParams{}
}

// GetDomainsReturns return values.
type GetDomainsReturns struct {
	Domains []*Domain `json:"domains,omitempty"` // List of supported domains.
}

// Do executes Schema.getDomains.
//
// returns:
//   domains - List of supported domains.
func (p *GetDomainsParams) Do(ctxt context.Context, h FrameHandler) (domains []*Domain, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, CommandSchemaGetDomains, Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return nil, ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r GetDomainsReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return nil, ErrInvalidResult
			}

			return r.Domains, nil

		case error:
			return nil, v
		}

	case <-ctxt.Done():
		return nil, ErrContextDone
	}

	return nil, ErrUnknownResult
}