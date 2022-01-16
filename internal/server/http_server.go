package server

import (
	orders "github.com/romaxa83/hra/proto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/jsonpb"
	_ "google.golang.org/grpc"
)

// RestServer реализует сервер REST для сервиса заказов
type RestServer struct {
	server       *http.Server
	orderService orders.OrderServiceServer // Тот же сервис заказов, что и в сервере gRPC
	errCh        chan error
}

// Функция NewRestServer отлично подходит для создания RestServer
func NewRestServer(orderService orders.OrderServiceServer, port string) RestServer {
	//logger.Infof("Create HTTP-server - [:%s]", port)
	router := gin.Default()

	rs := RestServer{
		server: &http.Server{
			Addr:    ":" + port,
			Handler: router,
		},
		orderService: orderService,
		errCh:        make(chan error),
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
	//logger.Infof("🚀 Start HTTP-server")
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
	var req orders.CreateTestRequest

	// Демаршализация запроса
	err := jsonpb.Unmarshal(c.Request.Body, &req)
	if err != nil {
		c.String(http.StatusInternalServerError, "error creating order request")
	}

	// Использует сервис заказов, чтобы создать заказ из запроса
	resp, err := r.orderService.CreateTest(c.Request.Context(), &req)
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
	resp, err := r.orderService.Create(c.Request.Context(), &req)
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
