package utils

import (
	"io/ioutil"
	"log"
	"os/exec"
	"time"
)

/**
 *
 * 备份MySql数据库
 * @param 	host: 			数据库地址: localhost
 * @param 	port:			端口: 3306
 * @param 	user:			用户名: root
 * @param 	password:		密码: root
 * @param 	databaseName:	需要被分的数据库名: test
 * @param 	tableName:		需要备份的表名: user
 * @param 	sqlPath:		备份SQL存储路径: D:/backup/test/
 * @return 	backupPath
 *
 */
func BackupMySqlDb(host, port, user, password, databaseName, tableName, sqlPath string) (error, string) {
	var cmd *exec.Cmd

	if tableName == "" {
		cmd = exec.Command("mysqldump", "--opt", "-h"+host, "-P"+port, "-u"+user, "-p"+password, databaseName)
	} else {
		cmd = exec.Command("mysqldump", "--opt", "-h"+host, "-P"+port, "-u"+user, "-p"+password, databaseName, tableName)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
		return err, ""
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
		return err, ""
	}

	bytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
		return err, ""
	}
	now := time.Now().Format("20060102150405")
	var backupPath string
	if tableName == "" {
		backupPath = sqlPath + databaseName + "_" + now + ".sql"
	} else {
		backupPath = sqlPath + databaseName + "_" + tableName + "_" + now + ".sql"
	}
	err = ioutil.WriteFile(backupPath, bytes, 0644)

	if err != nil {
		panic(err)
		return err, ""
	}
	return nil, backupPath
}
