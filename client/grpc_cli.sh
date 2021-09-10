grpc_cli call localhost:8080 urlmap.Redirection.GetOrgByPath "path: 'tako'"
grpc_cli call localhost:8080 urlmap.Redirection.SetUser "user: 'tako',notify_to: 'slack'"
grpc_cli call localhost:8080 urlmap.Redirection.GetInfoByUser "user: 'tako'"
