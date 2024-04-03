package data

import (
	"testing"

	"youshare-api.anvo.dev/internal/assert"
)

func TestVideoModel_Insert(t *testing.T) {
	testCases := []struct {
		name          string
		video         *Video
		expectedError error
	}{
		{
			name: "valid video",
			video: &Video{
				Url:         "https://www.youtube.com/watch?v=KLuTLF3x9sA",
				Title:       "Norway 4K • Scenic Relaxation Film with Peaceful Relaxing Music and Nature Video Ultra HD",
				Description: "Norway 4K • Scenic Relaxation Film with Peaceful Relaxing Music and Nature Video Ultra HD",
				Author: User{
					Email: "alice_test@example.com",
				},
			},
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := newTestDB(t)

			// get a user for the video author
			user, err := UserModel{db}.GetByEmail(tc.video.Author.Email)
			if err != nil {
				t.Fatal(err)
			}
			tc.video.Author.ID = user.ID
			m := VideoModel{db}

			err = m.Insert(tc.video)
			assert.NilError(t, err)
		})
	}
}

func TestVideoModel_GetAll(t *testing.T) {
	testCases := []struct {
		name           string
		expectedError  error
		expectedVideos []*Video
	}{
		{
			name:          "list video",
			expectedError: nil,
			expectedVideos: []*Video{
				{
					ID:          1,
					Url:         "https://www.youtube.com/watch?v=KLuTLF3x9sA",
					Title:       "Norway 4K",
					Description: "Video description",
					Author: User{
						ID:    1,
						Name:  "Alice",
						Email: "alice_test@example.com",
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := newTestDB(t)

			m := VideoModel{db}
			videos, err := m.GetAll()
			assert.NilError(t, err)
			assert.Equal(t, len(tc.expectedVideos), len(videos))

			for i := 0; i < len(tc.expectedVideos); i++ {
				assert.Equal(t, tc.expectedVideos[i].ID, videos[i].ID)
				assert.Equal(t, tc.expectedVideos[i].Url, videos[i].Url)
				assert.Equal(t, tc.expectedVideos[i].Title, videos[i].Title)
				assert.Equal(t, tc.expectedVideos[i].Description, videos[i].Description)
				assert.Equal(t, tc.expectedVideos[i].Author.ID, videos[i].Author.ID)
				assert.Equal(t, tc.expectedVideos[i].Author.Name, videos[i].Author.Name)
				assert.Equal(t, tc.expectedVideos[i].Author.Email, videos[i].Author.Email)
			}
		})
	}
}
