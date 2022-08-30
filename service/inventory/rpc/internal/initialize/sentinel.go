package initialize

import (
	"fmt"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/util"
	"go.uber.org/zap"
	"log"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/logging"
)

type stateChangeTestListener struct {
}

func (s *stateChangeTestListener) OnTransformToClosed(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	fmt.Printf("rule.steategy: %+v, From %s to Closed, time: %d\n", rule.Strategy, prev.String(), util.CurrentTimeMillis())
}

func (s *stateChangeTestListener) OnTransformToOpen(prev circuitbreaker.State, rule circuitbreaker.Rule, snapshot interface{}) {
	fmt.Printf("rule.steategy: %+v, From %s to Open, snapshot: %d, time: %d\n", rule.Strategy, prev.String(), snapshot, util.CurrentTimeMillis())
}

func (s *stateChangeTestListener) OnTransformToHalfOpen(prev circuitbreaker.State, rule circuitbreaker.Rule) {
	fmt.Printf("rule.steategy: %+v, From %s to Half-Open, time: %d\n", rule.Strategy, prev.String(), util.CurrentTimeMillis())
}

func InitSentinel() {
	conf := config.NewDefaultConfig()
	// for testing, logging output to console
	conf.Sentinel.Log.Logger = logging.NewConsoleLogger()
	err := sentinel.InitWithConfig(conf)
	if err != nil {
		log.Fatal(err)
	}

	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "limit",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject, //直接拒绝
			Threshold:              50,
			StatIntervalInMs:       5000,
		},
		{
			Resource:               "goods-update",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject, //直接拒绝
			Threshold:              5,
			StatIntervalInMs:       5000,
		},
		{
			Resource:               "goods-add",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject, //直接拒绝
			Threshold:              5,
			StatIntervalInMs:       5000,
		},
		{
			Resource:               "goods-detail",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject, //直接拒绝
			Threshold:              5,
			StatIntervalInMs:       5000,
		},
		{
			Resource:               "limit_WarmUp",
			TokenCalculateStrategy: flow.WarmUp,
			ControlBehavior:        flow.Reject, //直接拒绝
			Threshold:              100,         //并发数
			WarmUpPeriodSec:        30,          //30s达到1000的并发数
		},
	})

	// Register a state change listener so that we could observer the state change of the internal circuit breaker.
	circuitbreaker.RegisterStateChangeListeners(&stateChangeTestListener{})

	_, err = circuitbreaker.LoadRules([]*circuitbreaker.Rule{
		// Statistic time span=5s, recoveryTimeout=3s, maxErrorCount=50
		{
			Resource:         "breaker",
			Strategy:         circuitbreaker.ErrorCount,
			RetryTimeoutMs:   3000, //3秒回复
			MinRequestAmount: 10,   //静默数
			StatIntervalMs:   5000, //5s 统计一次
			//StatSlidingWindowBucketCount: 10,
			Threshold: 5, //50个请求限制
		},
	})
	if err != nil {
		zap.S().Fatalf("Unexpected error: %+v", err)
		return
	}
}
