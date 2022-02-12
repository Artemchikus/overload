package stresstest

import (
	"github.com/go-resty/resty/v2"
	"overload/internal/models"
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
	errs := make(chan error)
	go func(resps chan<- *models.ResponseInfo, errs chan<- error) {
		for i := int64(0); i < config.TestingDuration.Milliseconds(); i += time.Second.Milliseconds() {
			now := time.Now()
			for j := int32(0); j <= config.RPS; j++ {
				go func() {
					switch config.ReqMethod {
					case methodPost:
						resp, err := b.client.R().EnableTrace().SetBody([]byte(config.ReqBody)).Post(config.ReqURL)
						if err != nil {
							errs <- err
						}
						resps <- &models.ResponseInfo{
							Status:    resp.StatusCode(),
							TotalTime: resp.Request.TraceInfo().TotalTime,
							Time:      time.Now(),
						}
					case methodGet:
						resp, err := b.client.R().EnableTrace().Get(config.ReqURL)
						if err != nil {
							errs <- err
							return
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
		close(resps)
		close(errs)
	}(resps, errs)
	go func() {
		for resp := range resps {
			log.Info(resp)
		}
	}()
	//log.Info("Логирую(удали меня)")

	return nil, nil
}
