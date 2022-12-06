package repo

import (
	"context"
	"errors"
	"testing"

	storageMock "github.com/gobox-preegnees/file_manager/pkg/mocks"

	"github.com/golang/mock/gomock"
)

func TestConnToDB(t *testing.T) {

	url := "user=docker password=docker host=localhost port=5431 dbname=docker"
	_, err := New(nil, url)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSaveBatch(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	storageOK := storageMock.NewMockIStorage(ctrl)
	dataMock1 := []byte("1")
	storageOK.EXPECT().SaveBatchOnDisk(gomock.Any(), gomock.Any()).Return(1, nil).AnyTimes()

	storageFali := storageMock.NewMockIStorage(ctrl)
	storageErr := errors.New("storage fail")
	storageFali.EXPECT().SaveBatchOnDisk(gomock.Any(), gomock.Any()).Return(1, storageErr).AnyTimes()

	type TestData struct {
		batch Batch
		mock *storageMock.MockIStorage
		err bool
		clear bool
		desc string
	}
	testData := []TestData{ 
		{
			desc: "тут не возникает никаких ошибок",
			batch: Batch{
				Username:   "usrename_",
				FolderID:   "folder_",
				ClientID:   "client_",
				Path:       "/folder/file.txt_",
				Hash:       "frekth46_",
				ModTime:    12347,
				Part:       2,
				CountParts: 10,
				PartSize:   len(dataMock1),
				Offset:     4096,
				SizeFile:   100000,
				Content:    dataMock1,
			},
			mock: storageOK,
			err: false,
			clear: true,
		},
		{
			desc: "тут возникает ошибка при сохранении файла на диск, но не возникает ошибка с базой",
			batch: Batch{
				Username:   "usrename_",
				FolderID:   "folder_",
				ClientID:   "client_",
				Path:       "/folder/file.txt_",
				Hash:       "frekth46_",
				ModTime:    12347,
				Part:       2,
				CountParts: 10,
				PartSize:   len(dataMock1),
				Offset:     4096,
				SizeFile:   100000,
				Content:    dataMock1,
			},
			mock: storageFali,
			err: true,
			clear: true,
		},
		{
			desc: "тут также не возникает никакх ошибок, при этом запись остается в базе",
			batch: Batch{
				Username:   "usrename_",
				FolderID:   "folder_",
				ClientID:   "client_",
				Path:       "/folder/file.txt_",
				Hash:       "frekth46_",
				ModTime:    12347,
				Part:       2,
				CountParts: 10,
				PartSize:   len(dataMock1),
				Offset:     4096,
				SizeFile:   100000,
				Content:    dataMock1,
			},
			mock: storageOK,
			err: false,
			clear: false,
		},
		{
			desc: "в прошлом тесте запись с таким именем была сохранена, значит сейчас возникнет ошибка базы",
			batch: Batch{
				Username:   "usrename_",
				FolderID:   "folder_",
				ClientID:   "client_",
				Path:       "/folder/file.txt_",
				Hash:       "frekth46_",
				ModTime:    12347,
				Part:       2,
				CountParts: 10,
				PartSize:   len(dataMock1),
				Offset:     4096,
				SizeFile:   100000,
				Content:    dataMock1,
			},
			mock: storageOK,
			err: true,
			clear: true,
		},
	}

	url := "user=docker password=docker host=localhost port=5431 dbname=docker"
	for _, d := range testData {
		repo, err := New(d.mock, url)
		if err != nil {
			t.Error(err)
		}
	
		id, err := repo.SaveBatch(context.TODO(), &d.batch)
		if d.err {
			if err == nil {
				t.Error(err)
			}
		}

		t.Log(id)

		if d.clear {
			repo.deleteTestData(`DELETE FROM batches WHERE username='usrename_'`)
		}
	}
	
}

func TestGetBatch(t *testing.T) {

	url := "user=docker password=docker host=localhost port=5431 dbname=docker"
	repo, err := New(nil, url)
	if err != nil {
		t.Fatal(err)
	}

	batch, err := repo.GetBatch(context.TODO(), 1)
	if err != nil {
		t.Fatal(err)
	}
	if batch != nil {
		t.Log(batch)
	} else {
		t.Fatal("batch is nil")
	}
}
