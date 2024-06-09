// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/loft-sh/api/v4/pkg/apis/management/v1"
	scheme "github.com/loft-sh/api/v4/pkg/client/clientset_generated/clientset/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AgentAuditEventsGetter has a method to return a AgentAuditEventInterface.
// A group's client should implement this interface.
type AgentAuditEventsGetter interface {
	AgentAuditEvents() AgentAuditEventInterface
}

// AgentAuditEventInterface has methods to work with AgentAuditEvent resources.
type AgentAuditEventInterface interface {
	Create(ctx context.Context, agentAuditEvent *v1.AgentAuditEvent, opts metav1.CreateOptions) (*v1.AgentAuditEvent, error)
	Update(ctx context.Context, agentAuditEvent *v1.AgentAuditEvent, opts metav1.UpdateOptions) (*v1.AgentAuditEvent, error)
	UpdateStatus(ctx context.Context, agentAuditEvent *v1.AgentAuditEvent, opts metav1.UpdateOptions) (*v1.AgentAuditEvent, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.AgentAuditEvent, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.AgentAuditEventList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AgentAuditEvent, err error)
	AgentAuditEventExpansion
}

// agentAuditEvents implements AgentAuditEventInterface
type agentAuditEvents struct {
	client rest.Interface
}

// newAgentAuditEvents returns a AgentAuditEvents
func newAgentAuditEvents(c *ManagementV1Client) *agentAuditEvents {
	return &agentAuditEvents{
		client: c.RESTClient(),
	}
}

// Get takes name of the agentAuditEvent, and returns the corresponding agentAuditEvent object, and an error if there is any.
func (c *agentAuditEvents) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.AgentAuditEvent, err error) {
	result = &v1.AgentAuditEvent{}
	err = c.client.Get().
		Resource("agentauditevents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AgentAuditEvents that match those selectors.
func (c *agentAuditEvents) List(ctx context.Context, opts metav1.ListOptions) (result *v1.AgentAuditEventList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.AgentAuditEventList{}
	err = c.client.Get().
		Resource("agentauditevents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested agentAuditEvents.
func (c *agentAuditEvents) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("agentauditevents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a agentAuditEvent and creates it.  Returns the server's representation of the agentAuditEvent, and an error, if there is any.
func (c *agentAuditEvents) Create(ctx context.Context, agentAuditEvent *v1.AgentAuditEvent, opts metav1.CreateOptions) (result *v1.AgentAuditEvent, err error) {
	result = &v1.AgentAuditEvent{}
	err = c.client.Post().
		Resource("agentauditevents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(agentAuditEvent).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a agentAuditEvent and updates it. Returns the server's representation of the agentAuditEvent, and an error, if there is any.
func (c *agentAuditEvents) Update(ctx context.Context, agentAuditEvent *v1.AgentAuditEvent, opts metav1.UpdateOptions) (result *v1.AgentAuditEvent, err error) {
	result = &v1.AgentAuditEvent{}
	err = c.client.Put().
		Resource("agentauditevents").
		Name(agentAuditEvent.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(agentAuditEvent).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *agentAuditEvents) UpdateStatus(ctx context.Context, agentAuditEvent *v1.AgentAuditEvent, opts metav1.UpdateOptions) (result *v1.AgentAuditEvent, err error) {
	result = &v1.AgentAuditEvent{}
	err = c.client.Put().
		Resource("agentauditevents").
		Name(agentAuditEvent.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(agentAuditEvent).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the agentAuditEvent and deletes it. Returns an error if one occurs.
func (c *agentAuditEvents) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("agentauditevents").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *agentAuditEvents) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("agentauditevents").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched agentAuditEvent.
func (c *agentAuditEvents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AgentAuditEvent, err error) {
	result = &v1.AgentAuditEvent{}
	err = c.client.Patch(pt).
		Resource("agentauditevents").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}