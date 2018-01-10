package logspam2

import (
	"context"
	"fmt"

	"github.com/fnproject/fn/api/models"
	"github.com/fnproject/fn/api/server"
	"github.com/fnproject/fn/fnext"
	"github.com/treeder/yolib"
)

func init() {
	server.RegisterExtension(&LogSpam{})
}

type LogSpam struct {
}

func (e *LogSpam) Name() string {
	return "github.com/treeder/fn-ext-example2/logspam2"
}

func (e *LogSpam) Setup(s fnext.ExtServer) error {
	s.AddCallListener(&LogSpam{})
	return nil
}
func (l *LogSpam) BeforeCall(ctx context.Context, call *models.Call) error {
	fmt.Println("LOGSPAM2!!! LOGSPAM2!!! LOGSPAM2!!! LOGSPAM2!!!")
	yolib.Call()
	return nil
}

func (l *LogSpam) AfterCall(ctx context.Context, call *models.Call) error {
	fmt.Println("LOGSPAM2!!! MORE ANNOYING THAT LOGSPAM 1! LOGSPAM2!!!")
	return nil
}
