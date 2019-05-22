package suggests

import "github.com/c-bata/go-prompt"

var MysqlOptions = []prompt.Suggest{
	{Text: "-B", Description: "Print Bytes received from/send to MySQL(Bytes_received,Bytes_sent)."},
	{Text: "-H", Description: "Mysql连接主机，默认127.0.0.1 (default \"127.0.0.1\")"},
	{Text: "-L", Description: "Print to Logfile. (default \"none\")"},
	{Text: "-P", Description: "Mysql连接端口,默认3306 (default \"3306\")"},
	{Text: "-u", Description: "Mysql 用户名,默认: root"},
	{Text: "-p", Description: "Mysql 密码"},
	{Text: "-db", Description: "Mysql 指定databases,默认：mysql"},
	{Text: "-S", Description: "mysql socket连接文件地址 (default \"/tmp/mysql.sock\")"},
	{Text: "-T", Description: "Print Threads Status(Threads_running,Threads_connected,Threads_created,Threads_cached)."},
	{Text: "-com", Description: "Print MySQL Status(Com_select,Com_insert,Com_update,Com_delete)."},
	{Text: "-hit", Description: "Print Innodb Hit%."},
	{Text: "-nocolor", Description: "不显示颜"},
	{Text: "-t", Description: "打印当前时间"},
	{Text: "-innodb", Description: "Print InnodbInfo(include -t,-innodb_pages,-innodb_data,-innodb_log,-innodb_status)"},
	{Text: "-innodb_rows", Description: "Print Innodb Rows Status(Innodb_rows_inserted/updated/deleted/read)."},
	{Text: "-innodb_pages", Description: "Print Innodb Buffer Pool Pages Status(Innodb_buffer_pool_pages_data/free/dirty/flushed)"},
	{Text: "-innodb_data", Description: "Print Innodb Data Status(Innodb_data_reads/writes/read/written)"},
	{Text: "-innodb_log", Description: "Print Innodb Log  Status(Innodb_os_log_fsyncs/written)"},
	{Text: "-innodb_status", Description: "Print Innodb Status from Command: 'Show Engine Innodb Status'"},
	{Text: "-T", Description: "Print Threads Status(Threads_running,Threads_connected,Threads_created,Threads_cached)."},
	{Text: "-mysql", Description: "Print MySQLInfo (include -t,-com,-hit,-T,-B)."},
	{Text: "-lazy", Description: "Print Info  (include -t,-l,-c,-s,-com,-hit)."},
	{Text: "-semi", Description: "半同步监控"},
	{Text: "-slave", Description: "打印Slave info"},
}
