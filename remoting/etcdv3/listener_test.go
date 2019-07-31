package etcdv3

import (
	"testing"

	"github.com/apache/dubbo-go/remoting"
	"github.com/stretchr/testify/assert"
)


var changedData = `
	dubbo.consumer.request_timeout=3s
	dubbo.consumer.connect_timeout=5s
	dubbo.application.organization=ikurento.com
	dubbo.application.name=BDTService
	dubbo.application.module=dubbogo user-info server
	dubbo.application.version=0.0.1
	dubbo.application.owner=ZX
	dubbo.application.environment=dev
	dubbo.registries.hangzhouzk.protocol=zookeeper
	dubbo.registries.hangzhouzk.timeout=3s
	dubbo.registries.hangzhouzk.address=127.0.0.1:2181
	dubbo.registries.shanghaizk.protocol=zookeeper
	dubbo.registries.shanghaizk.timeout=3s
	dubbo.registries.shanghaizk.address=127.0.0.1:2182
	dubbo.service.com.ikurento.user.UserProvider.protocol=dubbo
	dubbo.service.com.ikurento.user.UserProvider.interface=com.ikurento.user.UserProvider
	dubbo.service.com.ikurento.user.UserProvider.loadbalance=random
	dubbo.service.com.ikurento.user.UserProvider.warmup=100
	dubbo.service.com.ikurento.user.UserProvider.cluster=failover
`
func TestListener(t *testing.T) {

	var tests = []struct{
		input struct{
			k string
			v string
		}
	}{
		{input: struct {
			k string
			v string
		}{k: "/dubbo", v: changedData}},
	}

	c := initClient(t)
	defer c.Close()

	listener := NewEventListener(c)
	dataListener := &mockDataListener{client: c, changedData: changedData}
	listener.ListenServiceEvent("/dubbo", dataListener)


	for _, tc := range tests{

		k := tc.input.k
		v := tc.input.v
		if err := c.Create(k, v); err != nil{
			t.Fatal(err)
		}
	}
	assert.Equal(t, changedData, dataListener.eventList[0].Content)
}

type mockDataListener struct {
	eventList   []remoting.Event
	client      *Client
	changedData string
}

func (m *mockDataListener) DataChange(eventType remoting.Event) bool {
	m.eventList = append(m.eventList, eventType)
	if eventType.Content == m.changedData {
		//m.client.Close()
	}
	return true
}
