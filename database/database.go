package database

import harperdb "github.com/HarperDB-Add-Ons/sdk-go"

type Database interface {
	GetDb() *harperdb.Client
}


