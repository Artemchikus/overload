package stresstest

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"overload/internal/models"
	"sync"
	"time"

	"go.uber.org/zap"
)

const (
	methodPost = "POST"
	methodGet  = "GET"
)

type Tester struct {
	logger *zap.SugaredLogger
	client *resty.Client
}

func New() *Tester {
	log, _ := zap.NewProduction()
	defer log.Sync()
	sugar := log.Sugar()
	client := resty.New()

	return &Tester{
		logger: sugar,
		client: client,
	}
}

func (b *Tester) Test(config *models.TestingConfig) (*models.Metric, error) {
	log := b.logger.With("request_id", config.ID, "func", "test")
	resps := make(chan *models.ResponseInfo)
	wg := sync.WaitGroup{}
	wgChan := sync.WaitGroup{}
	go func(resps <-chan *models.ResponseInfo) {
		wgChan.Add(1)
		defer wgChan.Done()
		now := time.Now()
		medianTime := time.Duration(0)
		totalRequests := 0
		codesToAmount := make(map[int]uint32)
		for resp := range resps {
			totalRequests++
			medianTime += resp.TotalTime
			log.Info(resp)
			if _, ok := codesToAmount[resp.Status]; !ok {
				codesToAmount[resp.Status] = 1
				continue
			}
			codesToAmount[resp.Status]++
		}
		medianTime /= time.Duration(totalRequests)
		log.Info(fmt.Sprintf("Test ended in %d, with total %d requests\n"+
			"response codes: %v, median response time: %f m. / %f s. / %d ms.",
			time.Since(now),
			totalRequests,
			codesToAmount,
			medianTime.Minutes(),
			medianTime.Seconds(),
			medianTime.Milliseconds()))
	}(resps)
	for i := int64(0); i < config.DurationUNIX; i++ {
		now := time.Now()
		for j := int32(0); j <= config.RPS; j++ {
			go func() {
				wg.Add(1)
				defer wg.Done()
				switch config.ReqMethod {
				case methodPost:
					resp, err := b.client.SetTimeout(time.Second * 10).R().EnableTrace().SetBody([]byte(config.ReqBody)).Post(config.ReqURL)
					if err != nil {
						log.Errorf("request error: %v", err)
					}
					resps <- &models.ResponseInfo{
						Status:    resp.StatusCode(),
						TotalTime: resp.Request.TraceInfo().TotalTime,
						Time:      time.Now(),
					}
				case methodGet:
					resp, err := b.client.SetTimeout(time.Second * 10).R().EnableTrace().Get(config.ReqURL)
					if err != nil {
						log.Errorf("request error: %v", err)
					}
					resps <- &models.ResponseInfo{
						Status:    resp.StatusCode(),
						TotalTime: resp.Request.TraceInfo().TotalTime,
						Time:      time.Now(),
					}
				default:
					return
				}
			}()
		}
		time.Sleep(time.Second - time.Since(now))
	}
	wg.Wait()
	close(resps)
	//log.Info("Логирую(удали меня)")

	wgChan.Wait()
	return nil, nil
}
