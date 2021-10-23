grpcurl -plaintext -d '{"path": "foo"}' 127.0.0.1:8080 urlmap.Redirection.GetOrgByPath
grpcurl -plaintext -d '{"user": "tako", "notify_to":"slack"}' 127.0.0.1:8080 urlmap.Redirection.SetUser
grpcurl -plaintext -d '{"user": "tako"}' 127.0.0.1:8080 urlmap.Redirection.GetInfoByUser
grpcurl -plaintext -d '' 127.0.0.1:8080 urlmap.Redirection.ListUsers
