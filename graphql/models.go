package main

type Account struct {
	ID     string  `json:"id"`
	name   string  `json:"name"`
	orders []Order `json:"orders"`
}
