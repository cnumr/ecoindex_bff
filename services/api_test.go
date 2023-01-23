package services

import (
	"reflect"
	"testing"

	"github.com/vvatelot/ecoindex-bff/models"
)

func Test_convertApIResult(t *testing.T) {
	type args struct {
		ecoindexes []models.Ecoindex
		host       string
		path       string
	}
	tests := []struct {
		name string
		args args
		want models.EcoindexSearchResults
	}{
		{
			name: "should return empty results",
			args: args{
				ecoindexes: []models.Ecoindex{},
				host:       "www.example.com",
				path:       "/",
			},
			want: models.EcoindexSearchResults{},
		},
		{
			name: "should return latest result",
			args: args{
				ecoindexes: []models.Ecoindex{
					{
						Url:      "https://www.example.com/",
						Grade:    "A",
						Score:    99,
						Date:     "2023-01-23T10:07:38",
						Requests: 10,
						Size:     100,
						Nodes:    10,
					},
				},
				host: "www.example.com",
				path: "/",
			},
			want: models.EcoindexSearchResults{
				Count: 1,
				LatestResult: models.Ecoindex{
					Url:      "https://www.example.com/",
					Grade:    "A",
					Color:    "#349A47",
					Score:    99,
					Date:     "2023-01-23T10:07:38",
					Requests: 10,
					Size:     100,
					Nodes:    10,
				},
			},
		},
		{
			name: "should return latest result and older results",
			args: args{
				ecoindexes: []models.Ecoindex{
					{
						Url:      "https://www.example.com/",
						Grade:    "A",
						Score:    99,
						Date:     "2023-01-23T10:08:00",
						Requests: 10,
						Size:     100,
						Nodes:    10,
					},
					{
						Url:      "https://www.example.com/",
						Grade:    "A",
						Score:    99,
						Date:     "2023-01-23T10:07:30",
						Requests: 10,
						Size:     100,
						Nodes:    10,
					},
					{
						Url:      "https://www.example.com/",
						Grade:    "A",
						Score:    99,
						Date:     "2023-01-23T10:07:00",
						Requests: 10,
						Size:     100,
						Nodes:    10,
					},
				},
				host: "www.example.com",
				path: "/",
			},
			want: models.EcoindexSearchResults{
				Count: 3,
				LatestResult: models.Ecoindex{
					Url:      "https://www.example.com/",
					Grade:    "A",
					Color:    "#349A47",
					Score:    99,
					Date:     "2023-01-23T10:08:00",
					Requests: 10,
					Size:     100,
					Nodes:    10,
				},
				OlderResults: []models.Ecoindex{
					{
						Url:      "https://www.example.com/",
						Grade:    "A",
						Color:    "#349A47",
						Score:    99,
						Date:     "2023-01-23T10:07:30",
						Requests: 10,
						Size:     100,
						Nodes:    10,
					},
					{
						Url:      "https://www.example.com/",
						Grade:    "A",
						Color:    "#349A47",
						Score:    99,
						Date:     "2023-01-23T10:07:00",
						Requests: 10,
						Size:     100,
						Nodes:    10,
					},
				},
			},
		},
		{
			name: "should return only other results",
			args: args{
				ecoindexes: []models.Ecoindex{
					{
						Url:      "https://www.example.com/other_page",
						Grade:    "B",
						Score:    87,
						Date:     "2023-01-23T10:07:00",
						Requests: 100,
						Size:     1000,
						Nodes:    10,
					},
					{
						Url:      "https://www.example.com/another_page",
						Grade:    "C",
						Score:    75,
						Date:     "2023-01-23T10:07:00",
						Requests: 100,
						Size:     3000,
						Nodes:    10,
					},
				},
				host: "www.example.com",
				path: "/",
			},
			want: models.EcoindexSearchResults{
				Count: 2,
				HostResults: []models.Ecoindex{
					{
						Url:      "https://www.example.com/other_page",
						Grade:    "B",
						Color:    "#51B84B",
						Score:    87,
						Date:     "2023-01-23T10:07:00",
						Requests: 100,
						Size:     1000,
						Nodes:    10,
					},
					{
						Url:      "https://www.example.com/another_page",
						Grade:    "C",
						Color:    "#CADB2A",
						Score:    75,
						Date:     "2023-01-23T10:07:00",
						Requests: 100,
						Size:     3000,
						Nodes:    10,
					},
				},
			},
		},
		{
			name: "should return latest result and other results",
			args: args{
				ecoindexes: []models.Ecoindex{
					{
						Grade:    "E",
						Score:    29,
						Date:     "2022-12-20T09:45:22",
						Requests: 73,
						Size:     1631.92,
						Nodes:    2052,
						Url:      "https://github.com/cnumr/GreenIT-Analysis-cli",
					},
					{
						Grade:    "E",
						Score:    36,
						Date:     "2023-01-20T16:37:57",
						Requests: 65,
						Size:     1083.48,
						Nodes:    1617,
						Url:      "https://github.com/EmmanuelDemey/eco-index-audit",
					},
					{
						Grade:    "E",
						Score:    27,
						Date:     "2023-01-18T15:42:04",
						Requests: 95,
						Size:     2514.77,
						Nodes:    1379,
						Url:      "https://github.com/",
					},
				},
				host: "github.com",
				path: "/cnumr/GreenIT-Analysis-cli",
			},
			want: models.EcoindexSearchResults{
				Count: 3,
				LatestResult: models.Ecoindex{
					Grade:    "E",
					Score:    29,
					Date:     "2022-12-20T09:45:22",
					Requests: 73,
					Size:     1631.92,
					Nodes:    2052,
					Url:      "https://github.com/cnumr/GreenIT-Analysis-cli",
					Color:    "#FECD06",
				},
				HostResults: []models.Ecoindex{
					{
						Grade:    "E",
						Score:    36,
						Date:     "2023-01-20T16:37:57",
						Requests: 65,
						Size:     1083.48,
						Nodes:    1617,
						Url:      "https://github.com/EmmanuelDemey/eco-index-audit",
						Color:    "#FECD06",
					},
					{
						Grade:    "E",
						Score:    27,
						Date:     "2023-01-18T15:42:04",
						Requests: 95,
						Size:     2514.77,
						Nodes:    1379,
						Url:      "https://github.com/",
						Color:    "#FECD06",
					},
				},
			},
		},
		{
			name: "should return latest result, older and other results",
			args: args{
				ecoindexes: []models.Ecoindex{
					{
						Grade:    "E",
						Score:    29,
						Date:     "2022-12-20T09:45:22",
						Requests: 73,
						Size:     1631.92,
						Nodes:    2052,
						Url:      "https://github.com/cnumr/GreenIT-Analysis-cli",
					},
					{
						Grade:    "E",
						Score:    36,
						Date:     "2023-01-20T16:37:57",
						Requests: 65,
						Size:     1083.48,
						Nodes:    1617,
						Url:      "https://github.com/EmmanuelDemey/eco-index-audit",
					},
					{
						Grade:    "E",
						Score:    27,
						Date:     "2023-01-18T15:42:04",
						Requests: 95,
						Size:     2514.77,
						Nodes:    1379,
						Url:      "https://github.com/",
					},
					{
						Grade:    "E",
						Score:    27,
						Date:     "2023-01-02T14:23:42",
						Requests: 96,
						Size:     2483.12,
						Nodes:    1380,
						Url:      "https://github.com",
					},
					{
						Grade:    "E",
						Score:    26,
						Date:     "2022-12-23T13:34:21",
						Requests: 96,
						Size:     2520.46,
						Nodes:    1380,
						Url:      "https://github.com/",
					},
				},
				host: "github.com",
				path: "/",
			},
			want: models.EcoindexSearchResults{
				Count: 5,
				LatestResult: models.Ecoindex{
					Grade:    "E",
					Score:    27,
					Date:     "2023-01-18T15:42:04",
					Requests: 95,
					Size:     2514.77,
					Nodes:    1379,
					Url:      "https://github.com/",
					Color:    "#FECD06",
				},
				OlderResults: []models.Ecoindex{
					{
						Grade:    "E",
						Score:    27,
						Date:     "2023-01-02T14:23:42",
						Requests: 96,
						Size:     2483.12,
						Nodes:    1380,
						Url:      "https://github.com",
						Color:    "#FECD06",
					},
					{
						Grade:    "E",
						Score:    26,
						Date:     "2022-12-23T13:34:21",
						Requests: 96,
						Size:     2520.46,
						Nodes:    1380,
						Url:      "https://github.com/",
						Color:    "#FECD06",
					},
				},
				HostResults: []models.Ecoindex{
					{
						Grade:    "E",
						Score:    29,
						Date:     "2022-12-20T09:45:22",
						Requests: 73,
						Size:     1631.92,
						Nodes:    2052,
						Url:      "https://github.com/cnumr/GreenIT-Analysis-cli",
						Color:    "#FECD06",
					},
					{
						Grade:    "E",
						Score:    36,
						Date:     "2023-01-20T16:37:57",
						Requests: 65,
						Size:     1083.48,
						Nodes:    1617,
						Url:      "https://github.com/EmmanuelDemey/eco-index-audit",
						Color:    "#FECD06",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertApIResult(tt.args.ecoindexes, tt.args.host, tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertApIResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
