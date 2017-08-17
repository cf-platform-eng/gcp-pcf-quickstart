// Code generated by counterfeiter. DO NOT EDIT.
package googlefakes

/*
 * Copyright 2017 Google Inc.
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
 */

import (
	"omg-cli/google"
	"sync"
)

type FakeProjectService struct {
	QuotasStub        func() (map[string]google.Quota, error)
	quotasMutex       sync.RWMutex
	quotasArgsForCall []struct{}
	quotasReturns     struct {
		result1 map[string]google.Quota
		result2 error
	}
	quotasReturnsOnCall map[int]struct {
		result1 map[string]google.Quota
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProjectService) Quotas() (map[string]google.Quota, error) {
	fake.quotasMutex.Lock()
	ret, specificReturn := fake.quotasReturnsOnCall[len(fake.quotasArgsForCall)]
	fake.quotasArgsForCall = append(fake.quotasArgsForCall, struct{}{})
	fake.recordInvocation("Quotas", []interface{}{})
	fake.quotasMutex.Unlock()
	if fake.QuotasStub != nil {
		return fake.QuotasStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.quotasReturns.result1, fake.quotasReturns.result2
}

func (fake *FakeProjectService) QuotasCallCount() int {
	fake.quotasMutex.RLock()
	defer fake.quotasMutex.RUnlock()
	return len(fake.quotasArgsForCall)
}

func (fake *FakeProjectService) QuotasReturns(result1 map[string]google.Quota, result2 error) {
	fake.QuotasStub = nil
	fake.quotasReturns = struct {
		result1 map[string]google.Quota
		result2 error
	}{result1, result2}
}

func (fake *FakeProjectService) QuotasReturnsOnCall(i int, result1 map[string]google.Quota, result2 error) {
	fake.QuotasStub = nil
	if fake.quotasReturnsOnCall == nil {
		fake.quotasReturnsOnCall = make(map[int]struct {
			result1 map[string]google.Quota
			result2 error
		})
	}
	fake.quotasReturnsOnCall[i] = struct {
		result1 map[string]google.Quota
		result2 error
	}{result1, result2}
}

func (fake *FakeProjectService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.quotasMutex.RLock()
	defer fake.quotasMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProjectService) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ google.ProjectService = new(FakeProjectService)
