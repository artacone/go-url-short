package services

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/artacone/go-url-short/internal/repository.Repository -o ./repository_mock.go -n RepositoryMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/artacone/go-url-short/internal/models"
	"github.com/gojuno/minimock/v3"
)

// RepositoryMock implements repository.Repository
type RepositoryMock struct {
	t minimock.Tester

	funcCreate          func(ctx context.Context, url models.URL, code models.ShortURL) (s1 models.ShortURL, err error)
	inspectFuncCreate   func(ctx context.Context, url models.URL, code models.ShortURL)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mRepositoryMockCreate

	funcGetCode          func(ctx context.Context, url models.URL) (s1 models.ShortURL, err error)
	inspectFuncGetCode   func(ctx context.Context, url models.URL)
	afterGetCodeCounter  uint64
	beforeGetCodeCounter uint64
	GetCodeMock          mRepositoryMockGetCode

	funcGetURL          func(ctx context.Context, code models.ShortURL) (u1 models.URL, err error)
	inspectFuncGetURL   func(ctx context.Context, code models.ShortURL)
	afterGetURLCounter  uint64
	beforeGetURLCounter uint64
	GetURLMock          mRepositoryMockGetURL
}

// NewRepositoryMock returns a mock for repository.Repository
func NewRepositoryMock(t minimock.Tester) *RepositoryMock {
	m := &RepositoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mRepositoryMockCreate{mock: m}
	m.CreateMock.callArgs = []*RepositoryMockCreateParams{}

	m.GetCodeMock = mRepositoryMockGetCode{mock: m}
	m.GetCodeMock.callArgs = []*RepositoryMockGetCodeParams{}

	m.GetURLMock = mRepositoryMockGetURL{mock: m}
	m.GetURLMock.callArgs = []*RepositoryMockGetURLParams{}

	return m
}

type mRepositoryMockCreate struct {
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockCreateExpectation
	expectations       []*RepositoryMockCreateExpectation

	callArgs []*RepositoryMockCreateParams
	mutex    sync.RWMutex
}

// RepositoryMockCreateExpectation specifies expectation struct of the Repository.Create
type RepositoryMockCreateExpectation struct {
	mock    *RepositoryMock
	params  *RepositoryMockCreateParams
	results *RepositoryMockCreateResults
	Counter uint64
}

// RepositoryMockCreateParams contains parameters of the Repository.Create
type RepositoryMockCreateParams struct {
	ctx  context.Context
	url  models.URL
	code models.ShortURL
}

// RepositoryMockCreateResults contains results of the Repository.Create
type RepositoryMockCreateResults struct {
	s1  models.ShortURL
	err error
}

// Expect sets up expected params for Repository.Create
func (mmCreate *mRepositoryMockCreate) Expect(ctx context.Context, url models.URL, code models.ShortURL) *mRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("RepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &RepositoryMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &RepositoryMockCreateParams{ctx, url, code}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the Repository.Create
func (mmCreate *mRepositoryMockCreate) Inspect(f func(ctx context.Context, url models.URL, code models.ShortURL)) *mRepositoryMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for RepositoryMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by Repository.Create
func (mmCreate *mRepositoryMockCreate) Return(s1 models.ShortURL, err error) *RepositoryMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("RepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &RepositoryMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &RepositoryMockCreateResults{s1, err}
	return mmCreate.mock
}

//Set uses given function f to mock the Repository.Create method
func (mmCreate *mRepositoryMockCreate) Set(f func(ctx context.Context, url models.URL, code models.ShortURL) (s1 models.ShortURL, err error)) *RepositoryMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the Repository.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the Repository.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the Repository.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mRepositoryMockCreate) When(ctx context.Context, url models.URL, code models.ShortURL) *RepositoryMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("RepositoryMock.Create mock is already set by Set")
	}

	expectation := &RepositoryMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &RepositoryMockCreateParams{ctx, url, code},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up Repository.Create return parameters for the expectation previously defined by the When method
func (e *RepositoryMockCreateExpectation) Then(s1 models.ShortURL, err error) *RepositoryMock {
	e.results = &RepositoryMockCreateResults{s1, err}
	return e.mock
}

// Create implements repository.Repository
func (mmCreate *RepositoryMock) Create(ctx context.Context, url models.URL, code models.ShortURL) (s1 models.ShortURL, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx, url, code)
	}

	mm_params := &RepositoryMockCreateParams{ctx, url, code}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := RepositoryMockCreateParams{ctx, url, code}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("RepositoryMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the RepositoryMock.Create")
		}
		return (*mm_results).s1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx, url, code)
	}
	mmCreate.t.Fatalf("Unexpected call to RepositoryMock.Create. %v %v %v", ctx, url, code)
	return
}

// CreateAfterCounter returns a count of finished RepositoryMock.Create invocations
func (mmCreate *RepositoryMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of RepositoryMock.Create invocations
func (mmCreate *RepositoryMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mRepositoryMockCreate) Calls() []*RepositoryMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*RepositoryMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *RepositoryMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.Create")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to RepositoryMock.Create")
	}
}

type mRepositoryMockGetCode struct {
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockGetCodeExpectation
	expectations       []*RepositoryMockGetCodeExpectation

	callArgs []*RepositoryMockGetCodeParams
	mutex    sync.RWMutex
}

// RepositoryMockGetCodeExpectation specifies expectation struct of the Repository.GetCode
type RepositoryMockGetCodeExpectation struct {
	mock    *RepositoryMock
	params  *RepositoryMockGetCodeParams
	results *RepositoryMockGetCodeResults
	Counter uint64
}

// RepositoryMockGetCodeParams contains parameters of the Repository.GetCode
type RepositoryMockGetCodeParams struct {
	ctx context.Context
	url models.URL
}

// RepositoryMockGetCodeResults contains results of the Repository.GetCode
type RepositoryMockGetCodeResults struct {
	s1  models.ShortURL
	err error
}

// Expect sets up expected params for Repository.GetCode
func (mmGetCode *mRepositoryMockGetCode) Expect(ctx context.Context, url models.URL) *mRepositoryMockGetCode {
	if mmGetCode.mock.funcGetCode != nil {
		mmGetCode.mock.t.Fatalf("RepositoryMock.GetCode mock is already set by Set")
	}

	if mmGetCode.defaultExpectation == nil {
		mmGetCode.defaultExpectation = &RepositoryMockGetCodeExpectation{}
	}

	mmGetCode.defaultExpectation.params = &RepositoryMockGetCodeParams{ctx, url}
	for _, e := range mmGetCode.expectations {
		if minimock.Equal(e.params, mmGetCode.defaultExpectation.params) {
			mmGetCode.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetCode.defaultExpectation.params)
		}
	}

	return mmGetCode
}

// Inspect accepts an inspector function that has same arguments as the Repository.GetCode
func (mmGetCode *mRepositoryMockGetCode) Inspect(f func(ctx context.Context, url models.URL)) *mRepositoryMockGetCode {
	if mmGetCode.mock.inspectFuncGetCode != nil {
		mmGetCode.mock.t.Fatalf("Inspect function is already set for RepositoryMock.GetCode")
	}

	mmGetCode.mock.inspectFuncGetCode = f

	return mmGetCode
}

// Return sets up results that will be returned by Repository.GetCode
func (mmGetCode *mRepositoryMockGetCode) Return(s1 models.ShortURL, err error) *RepositoryMock {
	if mmGetCode.mock.funcGetCode != nil {
		mmGetCode.mock.t.Fatalf("RepositoryMock.GetCode mock is already set by Set")
	}

	if mmGetCode.defaultExpectation == nil {
		mmGetCode.defaultExpectation = &RepositoryMockGetCodeExpectation{mock: mmGetCode.mock}
	}
	mmGetCode.defaultExpectation.results = &RepositoryMockGetCodeResults{s1, err}
	return mmGetCode.mock
}

//Set uses given function f to mock the Repository.GetCode method
func (mmGetCode *mRepositoryMockGetCode) Set(f func(ctx context.Context, url models.URL) (s1 models.ShortURL, err error)) *RepositoryMock {
	if mmGetCode.defaultExpectation != nil {
		mmGetCode.mock.t.Fatalf("Default expectation is already set for the Repository.GetCode method")
	}

	if len(mmGetCode.expectations) > 0 {
		mmGetCode.mock.t.Fatalf("Some expectations are already set for the Repository.GetCode method")
	}

	mmGetCode.mock.funcGetCode = f
	return mmGetCode.mock
}

// When sets expectation for the Repository.GetCode which will trigger the result defined by the following
// Then helper
func (mmGetCode *mRepositoryMockGetCode) When(ctx context.Context, url models.URL) *RepositoryMockGetCodeExpectation {
	if mmGetCode.mock.funcGetCode != nil {
		mmGetCode.mock.t.Fatalf("RepositoryMock.GetCode mock is already set by Set")
	}

	expectation := &RepositoryMockGetCodeExpectation{
		mock:   mmGetCode.mock,
		params: &RepositoryMockGetCodeParams{ctx, url},
	}
	mmGetCode.expectations = append(mmGetCode.expectations, expectation)
	return expectation
}

// Then sets up Repository.GetCode return parameters for the expectation previously defined by the When method
func (e *RepositoryMockGetCodeExpectation) Then(s1 models.ShortURL, err error) *RepositoryMock {
	e.results = &RepositoryMockGetCodeResults{s1, err}
	return e.mock
}

// GetCode implements repository.Repository
func (mmGetCode *RepositoryMock) GetCode(ctx context.Context, url models.URL) (s1 models.ShortURL, err error) {
	mm_atomic.AddUint64(&mmGetCode.beforeGetCodeCounter, 1)
	defer mm_atomic.AddUint64(&mmGetCode.afterGetCodeCounter, 1)

	if mmGetCode.inspectFuncGetCode != nil {
		mmGetCode.inspectFuncGetCode(ctx, url)
	}

	mm_params := &RepositoryMockGetCodeParams{ctx, url}

	// Record call args
	mmGetCode.GetCodeMock.mutex.Lock()
	mmGetCode.GetCodeMock.callArgs = append(mmGetCode.GetCodeMock.callArgs, mm_params)
	mmGetCode.GetCodeMock.mutex.Unlock()

	for _, e := range mmGetCode.GetCodeMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.s1, e.results.err
		}
	}

	if mmGetCode.GetCodeMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetCode.GetCodeMock.defaultExpectation.Counter, 1)
		mm_want := mmGetCode.GetCodeMock.defaultExpectation.params
		mm_got := RepositoryMockGetCodeParams{ctx, url}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetCode.t.Errorf("RepositoryMock.GetCode got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetCode.GetCodeMock.defaultExpectation.results
		if mm_results == nil {
			mmGetCode.t.Fatal("No results are set for the RepositoryMock.GetCode")
		}
		return (*mm_results).s1, (*mm_results).err
	}
	if mmGetCode.funcGetCode != nil {
		return mmGetCode.funcGetCode(ctx, url)
	}
	mmGetCode.t.Fatalf("Unexpected call to RepositoryMock.GetCode. %v %v", ctx, url)
	return
}

// GetCodeAfterCounter returns a count of finished RepositoryMock.GetCode invocations
func (mmGetCode *RepositoryMock) GetCodeAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetCode.afterGetCodeCounter)
}

// GetCodeBeforeCounter returns a count of RepositoryMock.GetCode invocations
func (mmGetCode *RepositoryMock) GetCodeBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetCode.beforeGetCodeCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.GetCode.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetCode *mRepositoryMockGetCode) Calls() []*RepositoryMockGetCodeParams {
	mmGetCode.mutex.RLock()

	argCopy := make([]*RepositoryMockGetCodeParams, len(mmGetCode.callArgs))
	copy(argCopy, mmGetCode.callArgs)

	mmGetCode.mutex.RUnlock()

	return argCopy
}

// MinimockGetCodeDone returns true if the count of the GetCode invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockGetCodeDone() bool {
	for _, e := range m.GetCodeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetCodeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCodeCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetCode != nil && mm_atomic.LoadUint64(&m.afterGetCodeCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetCodeInspect logs each unmet expectation
func (m *RepositoryMock) MinimockGetCodeInspect() {
	for _, e := range m.GetCodeMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.GetCode with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetCodeMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCodeCounter) < 1 {
		if m.GetCodeMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.GetCode")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.GetCode with params: %#v", *m.GetCodeMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetCode != nil && mm_atomic.LoadUint64(&m.afterGetCodeCounter) < 1 {
		m.t.Error("Expected call to RepositoryMock.GetCode")
	}
}

type mRepositoryMockGetURL struct {
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockGetURLExpectation
	expectations       []*RepositoryMockGetURLExpectation

	callArgs []*RepositoryMockGetURLParams
	mutex    sync.RWMutex
}

// RepositoryMockGetURLExpectation specifies expectation struct of the Repository.GetURL
type RepositoryMockGetURLExpectation struct {
	mock    *RepositoryMock
	params  *RepositoryMockGetURLParams
	results *RepositoryMockGetURLResults
	Counter uint64
}

// RepositoryMockGetURLParams contains parameters of the Repository.GetURL
type RepositoryMockGetURLParams struct {
	ctx  context.Context
	code models.ShortURL
}

// RepositoryMockGetURLResults contains results of the Repository.GetURL
type RepositoryMockGetURLResults struct {
	u1  models.URL
	err error
}

// Expect sets up expected params for Repository.GetURL
func (mmGetURL *mRepositoryMockGetURL) Expect(ctx context.Context, code models.ShortURL) *mRepositoryMockGetURL {
	if mmGetURL.mock.funcGetURL != nil {
		mmGetURL.mock.t.Fatalf("RepositoryMock.GetURL mock is already set by Set")
	}

	if mmGetURL.defaultExpectation == nil {
		mmGetURL.defaultExpectation = &RepositoryMockGetURLExpectation{}
	}

	mmGetURL.defaultExpectation.params = &RepositoryMockGetURLParams{ctx, code}
	for _, e := range mmGetURL.expectations {
		if minimock.Equal(e.params, mmGetURL.defaultExpectation.params) {
			mmGetURL.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetURL.defaultExpectation.params)
		}
	}

	return mmGetURL
}

// Inspect accepts an inspector function that has same arguments as the Repository.GetURL
func (mmGetURL *mRepositoryMockGetURL) Inspect(f func(ctx context.Context, code models.ShortURL)) *mRepositoryMockGetURL {
	if mmGetURL.mock.inspectFuncGetURL != nil {
		mmGetURL.mock.t.Fatalf("Inspect function is already set for RepositoryMock.GetURL")
	}

	mmGetURL.mock.inspectFuncGetURL = f

	return mmGetURL
}

// Return sets up results that will be returned by Repository.GetURL
func (mmGetURL *mRepositoryMockGetURL) Return(u1 models.URL, err error) *RepositoryMock {
	if mmGetURL.mock.funcGetURL != nil {
		mmGetURL.mock.t.Fatalf("RepositoryMock.GetURL mock is already set by Set")
	}

	if mmGetURL.defaultExpectation == nil {
		mmGetURL.defaultExpectation = &RepositoryMockGetURLExpectation{mock: mmGetURL.mock}
	}
	mmGetURL.defaultExpectation.results = &RepositoryMockGetURLResults{u1, err}
	return mmGetURL.mock
}

//Set uses given function f to mock the Repository.GetURL method
func (mmGetURL *mRepositoryMockGetURL) Set(f func(ctx context.Context, code models.ShortURL) (u1 models.URL, err error)) *RepositoryMock {
	if mmGetURL.defaultExpectation != nil {
		mmGetURL.mock.t.Fatalf("Default expectation is already set for the Repository.GetURL method")
	}

	if len(mmGetURL.expectations) > 0 {
		mmGetURL.mock.t.Fatalf("Some expectations are already set for the Repository.GetURL method")
	}

	mmGetURL.mock.funcGetURL = f
	return mmGetURL.mock
}

// When sets expectation for the Repository.GetURL which will trigger the result defined by the following
// Then helper
func (mmGetURL *mRepositoryMockGetURL) When(ctx context.Context, code models.ShortURL) *RepositoryMockGetURLExpectation {
	if mmGetURL.mock.funcGetURL != nil {
		mmGetURL.mock.t.Fatalf("RepositoryMock.GetURL mock is already set by Set")
	}

	expectation := &RepositoryMockGetURLExpectation{
		mock:   mmGetURL.mock,
		params: &RepositoryMockGetURLParams{ctx, code},
	}
	mmGetURL.expectations = append(mmGetURL.expectations, expectation)
	return expectation
}

// Then sets up Repository.GetURL return parameters for the expectation previously defined by the When method
func (e *RepositoryMockGetURLExpectation) Then(u1 models.URL, err error) *RepositoryMock {
	e.results = &RepositoryMockGetURLResults{u1, err}
	return e.mock
}

// GetURL implements repository.Repository
func (mmGetURL *RepositoryMock) GetURL(ctx context.Context, code models.ShortURL) (u1 models.URL, err error) {
	mm_atomic.AddUint64(&mmGetURL.beforeGetURLCounter, 1)
	defer mm_atomic.AddUint64(&mmGetURL.afterGetURLCounter, 1)

	if mmGetURL.inspectFuncGetURL != nil {
		mmGetURL.inspectFuncGetURL(ctx, code)
	}

	mm_params := &RepositoryMockGetURLParams{ctx, code}

	// Record call args
	mmGetURL.GetURLMock.mutex.Lock()
	mmGetURL.GetURLMock.callArgs = append(mmGetURL.GetURLMock.callArgs, mm_params)
	mmGetURL.GetURLMock.mutex.Unlock()

	for _, e := range mmGetURL.GetURLMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.u1, e.results.err
		}
	}

	if mmGetURL.GetURLMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetURL.GetURLMock.defaultExpectation.Counter, 1)
		mm_want := mmGetURL.GetURLMock.defaultExpectation.params
		mm_got := RepositoryMockGetURLParams{ctx, code}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetURL.t.Errorf("RepositoryMock.GetURL got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetURL.GetURLMock.defaultExpectation.results
		if mm_results == nil {
			mmGetURL.t.Fatal("No results are set for the RepositoryMock.GetURL")
		}
		return (*mm_results).u1, (*mm_results).err
	}
	if mmGetURL.funcGetURL != nil {
		return mmGetURL.funcGetURL(ctx, code)
	}
	mmGetURL.t.Fatalf("Unexpected call to RepositoryMock.GetURL. %v %v", ctx, code)
	return
}

// GetURLAfterCounter returns a count of finished RepositoryMock.GetURL invocations
func (mmGetURL *RepositoryMock) GetURLAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetURL.afterGetURLCounter)
}

// GetURLBeforeCounter returns a count of RepositoryMock.GetURL invocations
func (mmGetURL *RepositoryMock) GetURLBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetURL.beforeGetURLCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.GetURL.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetURL *mRepositoryMockGetURL) Calls() []*RepositoryMockGetURLParams {
	mmGetURL.mutex.RLock()

	argCopy := make([]*RepositoryMockGetURLParams, len(mmGetURL.callArgs))
	copy(argCopy, mmGetURL.callArgs)

	mmGetURL.mutex.RUnlock()

	return argCopy
}

// MinimockGetURLDone returns true if the count of the GetURL invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockGetURLDone() bool {
	for _, e := range m.GetURLMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetURLMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetURLCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetURL != nil && mm_atomic.LoadUint64(&m.afterGetURLCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetURLInspect logs each unmet expectation
func (m *RepositoryMock) MinimockGetURLInspect() {
	for _, e := range m.GetURLMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.GetURL with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetURLMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetURLCounter) < 1 {
		if m.GetURLMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.GetURL")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.GetURL with params: %#v", *m.GetURLMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetURL != nil && mm_atomic.LoadUint64(&m.afterGetURLCounter) < 1 {
		m.t.Error("Expected call to RepositoryMock.GetURL")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RepositoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateInspect()

		m.MinimockGetCodeInspect()

		m.MinimockGetURLInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RepositoryMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *RepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockGetCodeDone() &&
		m.MinimockGetURLDone()
}
