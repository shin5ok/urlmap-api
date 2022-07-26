package service

import (
	"context"
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

func TestRedirection_GetOrgPath(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := mock_urlmap.NewMockRedirectionClient(ctrl)

	ctx := context.Background()

	path := &pb.RedirectPath{}
	org := &pb.OrgUrl{}
	mockClient.EXPECT().GetOrgByPath(
		ctx,
		path,
	).Return(org, nil)

	testRedirection_GetOrgPath(t, mockClient)

}
func testRedirection_GetOrgPath(t *testing.T, client *mock_urlmap.MockRedirectionClient) {
	t.Helper()
	ctx := context.Background()
	resp, err := client.GetOrgByPath(ctx, &pb.RedirectPath{})
	if err != nil {
		t.Error(err)
	}
	t.Log(resp.GetOrg())
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
		Return(&pb.Message{Name: name, ShowModeOneof: &pb.Message_Mode{"pong"}}, nil)

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
