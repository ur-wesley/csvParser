package main

import "github.com/sqweek/dialog"

func ShowErrorDialog(e string) {
	dialog.Message("%v", e).Error()
}
