package http

import (
	"bst/app"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
)

type bstHandler struct {
	logger  *logrus.Logger
	storage app.Storage
}

func InitAndBindBSTHandler(
	_logger *logrus.Logger,
	rg *echo.Group,
	_storage app.Storage,
) *echo.Group {
	h := bstHandler{
		logger:  _logger,
		storage: _storage,
	}

	rg.GET("search/", h.Get)
	rg.POST("insert/", h.Insert)
	rg.DELETE("delete/", h.Delete)
	return rg
}

func (b *bstHandler) Get(c echo.Context) error {
	b.logger.Info("считывание параметра")
	valueStr := c.QueryParam("val")
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		b.logger.Errorf("ошибка конвертации str->int %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "value is not valid")
	}
	isExist := b.storage.Search(value)

	return c.JSON(http.StatusOK, app.ResponseDTO{Message: isExist})
}

func (b *bstHandler) Insert(c echo.Context) error {
	body := c.Request().Body
	bodyReq, err := ioutil.ReadAll(body)
	if err != nil {
		b.logger.Errorf("ошибка чтения тела запроса %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad body")
	}
	defer body.Close()
	var val app.ValueDTO
	err = json.Unmarshal(bodyReq, &val)
	if err != nil {
		b.logger.Errorf("ошибка парсинга тела запроса %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	b.storage.Insert(val.Value)

	return c.NoContent(http.StatusOK)
}

func (b *bstHandler) Delete(c echo.Context) error {
	b.logger.Info("считывание параметра")
	valueStr := c.QueryParam("val")
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		b.logger.Errorf("ошибка конвертации str->int %v", err)
		return echo.NewHTTPError(http.StatusBadRequest, "value is not valid")
	}
	b.storage.Remove(value)
	return c.NoContent(http.StatusOK)
}
