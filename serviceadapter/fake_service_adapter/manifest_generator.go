// This file was generated by counterfeiter
package fake_service_adapter

import (
	"sync"

	"github.com/pivotal-cf/on-demand-service-broker-sdk/bosh"
	"github.com/pivotal-cf/on-demand-service-broker-sdk/serviceadapter"
)

type FakeManifestGenerator struct {
	GenerateManifestStub        func(serviceDeployment serviceadapter.ServiceDeployment, plan serviceadapter.Plan, requestParams serviceadapter.RequestParameters, previousManifest *bosh.BoshManifest, previousPlan *serviceadapter.Plan) (bosh.BoshManifest, error)
	generateManifestMutex       sync.RWMutex
	generateManifestArgsForCall []struct {
		serviceDeployment serviceadapter.ServiceDeployment
		plan              serviceadapter.Plan
		requestParams     serviceadapter.RequestParameters
		previousManifest  *bosh.BoshManifest
		previousPlan      *serviceadapter.Plan
	}
	generateManifestReturns struct {
		result1 bosh.BoshManifest
		result2 error
	}
}

func (fake *FakeManifestGenerator) GenerateManifest(serviceDeployment serviceadapter.ServiceDeployment, plan serviceadapter.Plan, requestParams serviceadapter.RequestParameters, previousManifest *bosh.BoshManifest, previousPlan *serviceadapter.Plan) (bosh.BoshManifest, error) {
	fake.generateManifestMutex.Lock()
	fake.generateManifestArgsForCall = append(fake.generateManifestArgsForCall, struct {
		serviceDeployment serviceadapter.ServiceDeployment
		plan              serviceadapter.Plan
		requestParams     serviceadapter.RequestParameters
		previousManifest  *bosh.BoshManifest
		previousPlan      *serviceadapter.Plan
	}{serviceDeployment, plan, requestParams, previousManifest, previousPlan})
	fake.generateManifestMutex.Unlock()
	if fake.GenerateManifestStub != nil {
		return fake.GenerateManifestStub(serviceDeployment, plan, requestParams, previousManifest, previousPlan)
	} else {
		return fake.generateManifestReturns.result1, fake.generateManifestReturns.result2
	}
}

func (fake *FakeManifestGenerator) GenerateManifestCallCount() int {
	fake.generateManifestMutex.RLock()
	defer fake.generateManifestMutex.RUnlock()
	return len(fake.generateManifestArgsForCall)
}

func (fake *FakeManifestGenerator) GenerateManifestArgsForCall(i int) (serviceadapter.ServiceDeployment, serviceadapter.Plan, serviceadapter.RequestParameters, *bosh.BoshManifest, *serviceadapter.Plan) {
	fake.generateManifestMutex.RLock()
	defer fake.generateManifestMutex.RUnlock()
	return fake.generateManifestArgsForCall[i].serviceDeployment, fake.generateManifestArgsForCall[i].plan, fake.generateManifestArgsForCall[i].requestParams, fake.generateManifestArgsForCall[i].previousManifest, fake.generateManifestArgsForCall[i].previousPlan
}

func (fake *FakeManifestGenerator) GenerateManifestReturns(result1 bosh.BoshManifest, result2 error) {
	fake.GenerateManifestStub = nil
	fake.generateManifestReturns = struct {
		result1 bosh.BoshManifest
		result2 error
	}{result1, result2}
}

var _ serviceadapter.ManifestGenerator = new(FakeManifestGenerator)
