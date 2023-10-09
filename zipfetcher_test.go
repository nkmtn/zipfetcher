package zipfetcher

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type TestProvider struct {
	LastUpdateDate time.Time
	mock.Mock
}

func CreateTestProvider() *TestProvider {
	return &TestProvider{
		LastUpdateDate: getCurrentDate(),
	}
}

func (tp *TestProvider) GetLastModificationDate() (time.Time, error) {
	return getCurrentDate(), nil
}

func (tp *TestProvider) GetZips() ([]ZipCode, error) {
	args := tp.Called()
	return args.Get(0).([]ZipCode), args.Error(1)
}

func TestCreate(t *testing.T) {
	actual := Create()
	assert.IsType(t, &ZipFetcher{}, actual)
	assert.Implements(t, (*ZipProvider)(nil), actual.provider)

	actual1 := Create()
	expected1 := ZipFetcher{provider: CreateUspsProvider()}
	assert.Equal(t, actual1, &expected1)

	actual2 := Create(WithProvider(CreateUspsProvider()))
	expected2 := Create()
	assert.Equal(t, actual2, expected2)

	actual3 := Create(WithProvider(CreateTestProvider()))
	expected3 := Create()
	assert.NotEqual(t, actual3, expected3)
}

func TestWithProvider(t *testing.T) {
	zf := Create()
	assert.Equal(t, zf.provider, CreateUspsProvider())

	wp := WithProvider(CreateTestProvider())
	wp(zf)
	assert.Equal(t, zf.provider, CreateTestProvider())

	wp1 := WithProvider(CreateUspsProvider())
	wp1(zf)
	assert.Equal(t, zf.provider, CreateUspsProvider())
}

func TestCheckIfModifiedSince(t *testing.T) {
	zf := Create(WithProvider(CreateTestProvider()))
	actual, err := zf.CheckIfModifiedSince("1980-04-01")
	assert.NoError(t, err)
	assert.Equal(t, actual, true)

	date := getCurrentDate()
	actual1, err := zf.CheckIfModifiedSince(fmt.Sprintf("%d-%d-%d", date.Year(), date.Month(), date.Day()))
	assert.NoError(t, err)
	assert.Equal(t, actual1, false)

	actual2, err := zf.CheckIfModifiedSince(fmt.Sprintf("%d-%d-%d", date.Day(), date.Month(), date.Year()))
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "failed to parse")
	}
	assert.Equal(t, actual2, false)
}

func TestGetAllZips(t *testing.T) {
	testObj := new(TestProvider)
	testObj.On("GetZips").Return(nil, new(error))
	testObj.On("GetZips").Return([]ZipCode{}, nil)
	testObj.On("GetZips").Return([]ZipCode{{Code: "88888", State: "ST", City: "North", LocaleName: "North"}}, nil)
}
