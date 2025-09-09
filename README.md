## Parse ProxySQL Config Format.

```go
	b, _ := os.ReadFile("/etc/proxysql.cnf")
	c := Load(string(b))
	
	adminVariables, _ := c.GetGroup("admin_variables")
	mysqlIfaces, _ := adminVariables.GetScalar("mysql_ifaces")
```