package server

import (
	"github.com/opentracing/opentracing-go"
	"github.com/romaxa83/hra/config"
	"github.com/romaxa83/hra/pkg/logger"
	"github.com/romaxa83/hra/pkg/tracing"
	orders "github.com/romaxa83/hra/proto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	_ "google.golang.org/grpc"
)

// RestServer реализует сервер REST для сервиса заказов
type RestServer struct {
	server           *http.Server
	grpsOrderService orders.OrderServiceServer // Тот же сервис заказов, что и в сервере gRPC
	errCh            chan error
	logger           logger.Logger
	cfg              config.Config
}

// Функция NewRestServer отлично подходит для создания RestServer
func NewRestServer(
	orderService orders.OrderServiceServer,
	port string,
	logger logger.Logger,
	cfg config.Config,
) RestServer {
	logger.Infof("Create HTTP-server - [:%s]", port)
	router := gin.Default()

	if cfg.Jaeger.Enable {
		tracer, closer, err := tracing.NewJaegerTracer(cfg.Jaeger)
		if err != nil {
			logger.Error("Problem with Jaeger", err)
		}
		defer closer.Close()
		logger.Warn("HTTP Tracer", tracer)
		opentracing.SetGlobalTracer(tracer)
	}

	rs := RestServer{
		server: &http.Server{
			Addr:    ":" + port,
			Handler: router,
		},
		grpsOrderService: orderService,
		errCh:            make(chan error),
		logger:           logger,
	}

	// Регистрация маршрутов
	router.POST("/test", rs.test)
	router.POST("/order", rs.create)
	router.GET("/order/:id", rs.retrieve)
	router.PUT("/order", rs.update)
	router.DELETE("/order", rs.delete)
	router.GET("/order", rs.list)

	return rs
}

// Start запускает сервер REST в фоновом режиме, отправляя ошибку в канал ошибок
func (r RestServer) Start() {
	r.logger.Infof("🚀 Start HTTP-server")
	go func() {
		r.errCh <- r.server.ListenAndServe()
	}()
}

// Stop останавливает сервер
func (r RestServer) Stop() error {
	return r.server.Close()
}

// Error возвращает канал ошибок сервера
func (r RestServer) Error() chan error {
	return r.errCh
}

// Функция-обработчик create создает заказ из запроса (тело JSON)
func (r RestServer) test(c *gin.Context) {
	ctx, span := tracing.StartHttpServerTracerSpan(*c, "ordersHandlers.CreatOrder")
	//defer span.Finish()

	r.logger.Warn("Test HEADER", c.Request.Header)
	r.logger.Warn("Test CREATE ", ctx)
	r.logger.Warn("Test CREATE SPAN ", span)

	var req orders.CreateTestRequest

	// Демаршализация запроса
	err := jsonpb.Unmarshal(c.Request.Body, &req)
	if err != nil {
		c.String(http.StatusInternalServerError, "error creating order request")
	}

	// Использует сервис заказов, чтобы создать заказ из запроса
	resp, err := r.grpsOrderService.CreateTest(c.Request.Context(), &req)
	if err != nil {
		c.String(http.StatusInternalServerError, "error creating order")
	}
	//logger.Warnf("RESPONSE %+v", resp)
	//logger.Warnf("REQUEST %+v", req)
	m := &jsonpb.Marshaler{}
	if err := m.Marshal(c.Writer, resp); err != nil {
		c.String(http.StatusInternalServerError, "error sending order response")
	}
}

// Функция-обработчик create создает заказ из запроса (тело JSON)
func (r RestServer) create(c *gin.Context) {
	var req orders.CreateOrderRequest

	// Демаршализация запроса
	err := jsonpb.Unmarshal(c.Request.Body, &req)
	if err != nil {
		c.String(http.StatusInternalServerError, "error creating order request")
	}

	// Использует сервис заказов, чтобы создать заказ из запроса
	resp, err := r.grpsOrderService.Create(c.Request.Context(), &req)
	if err != nil {
		c.String(http.StatusInternalServerError, "error creating order")
	}
	m := &jsonpb.Marshaler{}
	if err := m.Marshal(c.Writer, resp); err != nil {
		c.String(http.StatusInternalServerError, "error sending order response")
	}
}

func (r RestServer) retrieve(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func (r RestServer) update(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func (r RestServer) delete(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}

func (r RestServer) list(c *gin.Context) {
	c.String(http.StatusNotImplemented, "not implemented yet")
}
