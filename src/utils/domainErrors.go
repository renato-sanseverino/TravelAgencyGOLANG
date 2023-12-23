package utils

import "errors"


var ValidationError = errors.New("Invalid data")

var NotFoundError = errors.New("Item was not found")

var DuplicationError = errors.New("Item is already in the list")
