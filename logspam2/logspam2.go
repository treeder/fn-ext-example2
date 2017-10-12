package logspam2

import (
	"context"
	"fmt"

	"github.com/fnproject/fn/api/models"
)

// This implements the extender.CallListener interface
type LogSpam struct {
}

func (l *LogSpam) BeforeCall(ctx context.Context, call *models.Call) error {
	fmt.Println("LOGSPAM2!!! LOGSPAM2!!! LOGSPAM2!!! LOGSPAM2!!!")
	return nil
}

func (l *LogSpam) AfterCall(ctx context.Context, call *models.Call) error {
	fmt.Println("LOGSPAM2!!! MORE ANNOYING THAT LOGSPAM 1! LOGSPAM2!!!")
	return nil
}
