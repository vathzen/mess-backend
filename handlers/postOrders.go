package handlers

import (
    "fmt"
    "net/http"
    "encoding/json"
    "time"
    "reflect"
)

func PostOrders(w http.ResponseWriter, r *http.Request){
    date := time.Now()
    order := Order{}
    status := Confirm{
        Status: "error",
        Text: "",
    }
    var costs [6]int64
    var cost int64
    var codes [6]string
    err := json.NewDecoder(r.Body).Decode(&order)
    if(err!=nil){
        panic(err)
    }
    fmt.Println("Got these:",order)
    db:= GetDB();

    if errw != nil{
        fmt.Println(errw)
    }
    val := reflect.ValueOf(order)
    for i := 1; i < val.NumField(); i++ {
        cost = cost + val.Field(i).Int() * costs[i-1]
    }
    query := "INSERT INTO orders VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) ON CONFLICT (regno) DO UPDATE set bf1 = excluded.bf1,bf2 = excluded.bf2,lun1 = excluded.lun1, lun2 = excluded.lun2, din1 = excluded.din1, din2 = excluded.din2;"
    _,erre := db.Exec(query,date.AddDate(0,0,1).Format("2006-01-02"),order.Username,order.Bf1,order.Bf2,order.Lun1,order.Lun2,order.Din1,order.Din2,cost)
    if(erre == nil){
        status.Status = "true"
    }else{
        fmt.Println(erre)
    }
    for i := 1; i < val.NumField(); i++ {
        if(val.Field(i).Int() == 0){
            codes[i-1] = "null"
        }else{
            codes[i-1] = GenRand(5) + fmt.Sprintf("%02d",val.Field(i))
        }
    }
    _,errt := db.Exec("INSERT INTO order_codes_tomorrow values($1,$2,$3,$4,$5,$6,$7) ON CONFLICT (regno) DO UPDATE set bf1 = excluded.bf1,bf2 = excluded.bf2,lun1 = excluded.lun1,lun2 = excluded.lun2,din1 = excluded.din1,din2 = excluded.din2;",order.Username,codes[0],codes[1],codes[2],codes[3],codes[4],codes[5])
    if(errt == nil){
        status.Status = "true"
    }else{
        fmt.Println(errt)
    }
    statusJson,errj := json.Marshal(status)
    if errj!=nil{
        panic(errj)
    }
    w.Header().Set("Content-Type","application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(statusJson)
}
