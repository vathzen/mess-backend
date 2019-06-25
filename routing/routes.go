package routing

import (
    "net/http"
    "backendSastraMess/handlers"
)

type Route struct{
    Name string
    Method string
    Pattern string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
    Route{
        "Get Date Time",
        "GET",
        "/",
        handlers.Index,
    },
    Route{
        "Return Order(s)",
        "GET",
        "/orders",
        handlers.GetOrders,
    },
    Route{
        "Add an Order",
        "POST",
        "/orders",
        handlers.PostOrders,
    },
    Route{
        "Return Menu",
        "GET",
        "/menu",
        handlers.GetMenu,
    },
    Route{
        "Add Menu",
        "POST",
        "/menu",
        handlers.PostMenu,
    },
    Route{
        "Add Menu Item",
        "PUT",
        "/menu",
        handlers.PutMenu,
    },
    Route{
        "Delete Menu Item",
        "DELETE",
        "/menu",
        handlers.DeleteMenu,
    },
    Route{
        "Login",
        "POST",
        "/users",
        handlers.PostUser,
    },
    Route{
        "Sign Up",
        "PUT",
        "/users",
        handlers.PutUser,
    },
    Route{
        "postCodes",
        "POST",
        "/codes",
        handlers.VerUser,
    },
}
