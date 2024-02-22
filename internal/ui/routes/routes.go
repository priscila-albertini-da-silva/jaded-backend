package routes

import (
	"net/http"

	"github.com/priscila-albertini-silva/jaded-backend/internal/ui"
	"github.com/priscila-albertini-silva/jaded-backend/pkg/serverfx"
	"go.uber.org/fx"
)

type Routes struct {
	fx.Out
	T []serverfx.Route
}

func ProvideRoutes(controller ui.StockController) Routes {
	var routes = []serverfx.Route{
		{
			Method:  http.MethodGet,
			Path:    "/stock",
			Handler: controller.GetStock,
		},
	}

	return Routes{
		T: routes,
	}
}
