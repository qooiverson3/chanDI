package fake

import "chanLoader/pkg/domain"

var (
	FakeData1 domain.Info = domain.Info{
		Name: "fake-data-1",
		Data: []int{1, 2, 3},
	}

	FakeData2 domain.Info = domain.Info{
		Name: "fake-data-2",
		Data: []int{11, 22, 33},
	}

	FakeDataList []domain.Info = []domain.Info{
		FakeData1,
		FakeData2,
	}

	FakeAmount int = 2
)
