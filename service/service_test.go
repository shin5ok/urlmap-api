package service

import (
	"context"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	mock_urlmap "github.com/shin5ok/urlmap-api/mock_urlmap"
	pb "github.com/shin5ok/urlmap-api/pb"
)

func TestRedirection_GetInfoByUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := mock_urlmap.NewMockRedirectionClient(ctrl)

	// ctx := context.Background()

	redirects := []*pb.RedirectData{}
	redirects = append(redirects, &pb.RedirectData{Redirect: &pb.RedirectInfo{User: "tako"}})
	request := &pb.User{User: "tako", NotifyTo: "slack", SlackUrl: "test", Email: ""}
	mockClient.EXPECT().GetInfoByUser(
		gomock.Any(),
		request,
	).Return(&pb.ArrayRedirectData{Redirects: redirects}, nil)

	// u := &pb.User{User: "tako", NotifyTo: "slack"}
	testRedirection_GetInfoByUser(t, mockClient)

}

func testRedirection_GetOrgPath(t *testing.T, client *mock_urlmap.MockRedirectionClient) {
	t.Helper()
	ctx := context.Background()
	resp, err := client.GetOrgByPath(ctx, &pb.RedirectPath{})
	if err != nil {
		t.Error(err)
	}
	if resp.GetOrg() != "https://www.google.com/" {
		t.Errorf("Org looks like not match %v", resp.GetOrg())
	}
	if resp.GetEmail() != "user0@example.com" {
		t.Errorf("Email looks like not match %v", resp.GetEmail())
	}
	t.Log(resp.GetOrg())
}

func TestRedirection_GetOrgPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := mock_urlmap.NewMockRedirectionClient(ctrl)

	ctx := context.Background()

	path := &pb.RedirectPath{}
	org := &pb.OrgUrl{Org: "https://www.google.com/", Email: "user0@example.com"}
	mockClient.EXPECT().GetOrgByPath(
		ctx,
		path,
	).Return(org, nil)

	testRedirection_GetOrgPath(t, mockClient)

}

func testRedirection_GetInfoByUser(t *testing.T, client *mock_urlmap.MockRedirectionClient) {
	t.Helper()
	ctx := context.Background()
	user := &pb.User{User: "tako", SlackUrl: "test", NotifyTo: "slack"}
	resp, err := client.GetInfoByUser(ctx, user)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.GetRedirects())
}

func TestRedirection_PingPongMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := mock_urlmap.NewMockRedirectionClient(ctrl)

	ctx := context.Background()

	name := "foo"
	mockClient.EXPECT().
		PingPongMessage(
			ctx,
			&pb.Message{Name: name},
		).
		Return(&pb.Message{Name: name, ShowModeOneof: &pb.Message_Mode{Mode: "pong"}}, nil)

	func(t *testing.T, client *mock_urlmap.MockRedirectionClient) {
		t.Helper()
		ctx := context.Background()
		message := &pb.Message{Name: "foo"}
		resp, err := client.PingPongMessage(ctx, message)
		if err != nil {
			t.Error(err)
		}
		t.Log(resp.GetName())
	}(t, mockClient)

}

func TestRedirection_SetInfo(t *testing.T) {
	type args struct {
		ctx context.Context
		r   *pb.RedirectData
	}
	ctrl := gomock.NewController(t)
	mockClient := mock_urlmap.NewMockRedirectionClient(ctrl)
	ctx := context.Background()

	redirectInfo := &pb.RedirectInfo{User: "foo"}
	mockClient.EXPECT().SetInfo(
		ctx,
		&pb.RedirectData{Redirect: redirectInfo},
	).
		Return(&pb.OrgUrl{Org: "https://www.google.com/"}, nil)

	tests := []struct {
		name    string
		args    args
		want    *pb.OrgUrl
		wantErr bool
	}{
		{
			name: "test SetInfo()",
			want: &pb.OrgUrl{Org: "https://www.google.com/"},
			args: args{
				ctx: ctx,
				r: &pb.RedirectData{
					Redirect: &pb.RedirectInfo{
						User: "foo",
					}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mockClient.SetInfo(tt.args.ctx, tt.args.r)
			// t.Log(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("Redirection.SetInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Redirection.SetInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
