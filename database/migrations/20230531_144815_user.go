package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type User_20230531_144815 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &User_20230531_144815{}
	m.Created = "20230531_144815"

	migration.Register("User_20230531_144815", m)
}

// Run the migrations
func (m *User_20230531_144815) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL("CREATE TABLE user(`id` int(11) NOT NULL AUTO_INCREMENT,`name` varchar(128) NOT NULL,`account` varchar(128) NOT NULL,`password` varchar(128) NOT NULL,PRIMARY KEY (`id`))")
}

// Reverse the migrations
func (m *User_20230531_144815) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL("DROP TABLE `user`")
}
